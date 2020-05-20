package http

import (
	"2020_1_Color_noise/internal/models"
	boardMock "2020_1_Color_noise/internal/pkg/board/mock"
	. "2020_1_Color_noise/internal/pkg/error"
	authServ "2020_1_Color_noise/internal/pkg/proto/session"
	authServMock "2020_1_Color_noise/internal/pkg/proto/session/mock"
	userServ "2020_1_Color_noise/internal/pkg/proto/user"
	userServMock "2020_1_Color_noise/internal/pkg/proto/user/mock"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	//"io/ioutil"
	"net/http/httptest"

	"strings"

	gomock "github.com/golang/mock/gomock"
	"testing"
	//"time"
)

type TestCaseCreate struct {
	IsAuth     bool
	Login      string
	Password   string
	Email      string
	CookieName string
	Cookie     string
	TokenName  string
	Token      string
	Response   string
	InputErr   bool
	ValidErr   bool
	CreateErr  bool
	SessErr    bool
}

func TestHandler_Create(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockBoard := boardMock.NewMockIUsecase(ctl)
	mockUserService := userServMock.NewMockUserServiceClient(ctl)
	mockAuthService := authServMock.NewMockAuthSeviceClient(ctl)

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	zap := logger.Sugar().With(
		zap.String("mode", "[access_log]"),
		zap.String("logger", "ZAP"),
	)

	userDelivery := NewHandler(mockUserService, mockAuthService, mockBoard, zap)

	cases := []TestCaseCreate{
		TestCaseCreate{
			IsAuth:   true,
			Response: `{"status":200,"body":{"message":"Ok"}}`,
		},
		TestCaseCreate{
			InputErr: true,
			Response: `{"status":400,"body":{"error":"Wrong body of request"}}`,
		},
		/*TestCaseCreate{
			ValidErr: true,
			Response: `{"status":400,"body":{"error":"Password should be longer than 6 characters and shorter 100. Login should be letters and numbers, and shorter than 20 characters Email should be like hello@example.com and shorter than 50 characters."}}`,
			Email:    "helloexam.com",
			Login:    "Login1",
			Password: "Password1",
		},
		TestCaseCreate{
			ValidErr: true,
			Response: `{"status":400,"body":{"error":"Password should be longer than 6 characters and shorter 100. Login should be letters and numbers, and shorter than 20 characters Email should be like hello@example.com and shorter than 50 characters."}}`,
			Email:    "hello@exam.com",
			Login:    "",
			Password: "Password1",
		},
		TestCaseCreate{
			ValidErr: true,
			Response: `{"status":400,"body":{"error":"Password should be longer than 6 characters and shorter 100. Login should be letters and numbers, and shorter than 20 characters Email should be like hello@example.com and shorter than 50 characters."}}`,
			Email:    "hello@exam.com",
			Login:    "Login1",
			Password: "Passw",
		},*/
		TestCaseCreate{
			CreateErr: true,
			Email:     "hello@exam.com",
			Login:     "Login1",
			Password:  "Password",
		},
		TestCaseCreate{
			SessErr:  true,
			Email:    "hello@exam.com",
			Login:    "Login1",
			Password: "Password",
		},
		TestCaseCreate{
			Email:      "hello@exam.com",
			Login:      "Login1",
			Password:   "Password",
			Response:   `{"status":201,"body":{"id":1,"login":"","subscriptions":0,"subscribers":0}}`,
			CookieName: "session_id",
			Cookie:     "cookie",
			//TokenName:  "csrf_token",
			//Token:      "token",
		},
	}

	for caseNum, item := range cases {
		var r *http.Request
		if item.InputErr == false {
			r = httptest.NewRequest("POST", "/api/user",
				strings.NewReader(fmt.Sprintf(`{"login":"%s", "email":"%s", "password":"%s"}`, item.Login, item.Email, item.Password)))
		} else {
			r = httptest.NewRequest("POST", "/api/user",
				strings.NewReader(fmt.Sprintf(`{"login:"%s", "email":"%s", "password":"%s"}`, item.Login, item.Email, item.Password)))
		}

		r.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		ctx := r.Context()
		ctx = context.WithValue(ctx, "IsAuth", item.IsAuth)
		r = r.WithContext(ctx)

		if !item.IsAuth && !item.InputErr && !item.ValidErr {

			input := &userServ.SignUp{
				Email:    item.Email,
				Login:    item.Login,
				Password: item.Password,
			}

			var err error = nil
			if item.CreateErr {
				err = NoType.New("")
			}

			gomock.InOrder(
				mockUserService.EXPECT().Create(r.Context(), input).Return(&userServ.User{Id: int64(1)}, err),
			)

			if !item.CreateErr {
				session := &authServ.Session{
					Id:     1,
					Cookie: item.Cookie,
					Token:  item.Token,
				}

				if item.SessErr {
					err = NoType.New("")
				}

				gomock.InOrder(
					mockAuthService.EXPECT().Create(r.Context(),
						&authServ.UserID{Id: int64(1)}).Return(session, err),
				)
			}
		}

		userDelivery.Create(w, r)

		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		bodyStr := string(body)

		if item.CreateErr || item.SessErr {
			var output map[string]interface{}

			err := json.NewDecoder(strings.NewReader(bodyStr)).Decode(&output)
			if err != nil {
				t.Fatalf("[%d] wrong decoding Response: got %+v, err: %v",
					caseNum, bodyStr, err.Error())
			}

			status, ok := output["status"]
			if !ok {
				t.Fatalf("[%d] wrong Response: got %+v - no status",
					caseNum, bodyStr)
			}

			if status == 200 || status == 201 {
				t.Errorf("[%d] wrong status Response: got %+v, expected not success status",
					caseNum, status)
			}
		} else {
			if bodyStr != item.Response {
				t.Errorf("[%d] wrong Response: got %+v, expected %+v",
					caseNum, bodyStr, item.Response)
			}

			if !item.IsAuth && !item.InputErr && !item.ValidErr {
				cookies := w.Result().Cookies()
				if len(cookies) < 1 {
					t.Fatalf("[%d] No Cookie", caseNum)
				}

				if cookies[0].Name != item.CookieName {
					t.Errorf("[%d] wrong Cookie: got %+v, expected %+v",
						caseNum, cookies[0].Name, item.CookieName)
				}

				if cookies[0].Value != item.Cookie {
					t.Errorf("[%d] wrong Cookie: got %+v, expected %+v",
						caseNum, cookies[0].Value, item.Cookie)
				}

				/*
					if cookies[1].Name != item.TokenName {
						t.Errorf("[%d] wrong Cookie: got %+v, expected %+v",
							caseNum, cookies[1].Name, item.TokenName)
					}

					if cookies[1].Value != item.Token {
						t.Errorf("[%d] wrong Cookie: got %+v, expected %+v",
							caseNum, cookies[1].Value, item.Token)
					}
				*/
			}
		}
	}
}

type TestCaseGetUser struct {
	IsAuth   bool
	User     *models.User
	Response string
	IdErr    bool
	GetErr   bool
}

func TestHandler_GetUser(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockBoard := boardMock.NewMockIUsecase(ctl)
	mockUserService := userServMock.NewMockUserServiceClient(ctl)
	mockAuthService := authServMock.NewMockAuthSeviceClient(ctl)

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	zap := logger.Sugar().With(
		zap.String("mode", "[access_log]"),
		zap.String("logger", "ZAP"),
	)

	userDelivery := NewHandler(mockUserService, mockAuthService, mockBoard, zap)

	cases := []TestCaseGetUser{
		TestCaseGetUser{
			IsAuth:   false,
			User:     &models.User{},
			Response: `{"status":401,"body":{"error":"User is unauthorized"}}`,
		},
		TestCaseGetUser{
			IsAuth:   true,
			IdErr:    true,
			User:     &models.User{},
			Response: `{"status":500,"body":{"error":"Internal server error"}}`,
		},
		TestCaseGetUser{
			IsAuth: true,
			GetErr: true,
			User:   &models.User{},
		},
		TestCaseGetUser{
			IsAuth: true,
			User: &models.User{
				Id:            1,
				Email:         "a@b.com",
				Login:         "login",
				About:         "about me",
				Avatar:        "avatar.jpg",
				Subscribers:   11000,
				Subscriptions: 100,
			},
			Response: `{"status":200,"body":{"id":1,"email":"a@b.com","login":"login","about":"about me","avatar":"avatar.jpg","subscriptions":100,"subscribers":11000}}`,
		},
	}

	for caseNum, item := range cases {
		r := httptest.NewRequest("GET", "/api/user", strings.NewReader(""))

		w := httptest.NewRecorder()

		ctx := r.Context()
		ctx = context.WithValue(ctx, "IsAuth", item.IsAuth)
		if !item.IdErr {
			ctx = context.WithValue(ctx, "Id", item.User.Id)
		}
		r = r.WithContext(ctx)

		if item.IsAuth && !item.IdErr {

			var err error = nil
			if item.GetErr {
				err = NoType.New("")
			}

			gomock.InOrder(
				mockUserService.EXPECT().GetById(r.Context(),
					&userServ.UserID{Id: int64(item.User.Id)}).Return(
					&userServ.User{Id: int64(item.User.Id),
						Email:         item.User.Email,
						Login:         item.User.Login,
						Avatar:        item.User.Avatar,
						About:         item.User.About,
						Subscribers:   int64(item.User.Subscribers),
						Subscriptions: int64(item.User.Subscriptions),
					},
					err,
				),
			)
		}

		userDelivery.GetUser(w, r)

		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		bodyStr := string(body)

		if item.GetErr {
			var output map[string]interface{}

			err := json.NewDecoder(strings.NewReader(bodyStr)).Decode(&output)
			if err != nil {
				t.Fatalf("[%d] wrong decoding Response: got %+v, err: %v",
					caseNum, bodyStr, err.Error())
			}

			status, ok := output["status"]
			if !ok {
				t.Fatalf("[%d] wrong Response: got %+v - no status",
					caseNum, bodyStr)
			}

			if status == 200 {
				t.Errorf("[%d] wrong status Response: got %+v, expected not success status",
					caseNum, status)
			}
		} else {
			if bodyStr != item.Response {
				t.Errorf("[%d] wrong Response: got %+v, expected %+v",
					caseNum, bodyStr, item.Response)
			}
		}
	}
}

func TestHandler_GetOtherUser(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockBoard := boardMock.NewMockIUsecase(ctl)
	mockUserService := userServMock.NewMockUserServiceClient(ctl)
	mockAuthService := authServMock.NewMockAuthSeviceClient(ctl)

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	zap := logger.Sugar().With(
		zap.String("mode", "[access_log]"),
		zap.String("logger", "ZAP"),
	)

	userDelivery := NewHandler(mockUserService, mockAuthService, mockBoard, zap)

	cases := []TestCaseGetUser{
		TestCaseGetUser{
			IsAuth:   false,
			User:     &models.User{},
			Response: `{"status":401,"body":{"error":"User is unauthorized"}}`,
		},
		TestCaseGetUser{
			IsAuth: true,
			GetErr: true,
			User:   &models.User{},
		},
		TestCaseGetUser{
			IsAuth: true,
			User: &models.User{
				Id:            1,
				Email:         "a@b.com",
				Login:         "login",
				About:         "about me",
				Avatar:        "avatar.jpg",
				Subscribers:   11000,
				Subscriptions: 100,
			},
			Response: `{"status":200,"body":{"id":1,"login":"login","about":"about me","avatar":"avatar.jpg","subscriptions":100,"subscribers":11000}}`,
		},
	}

	for caseNum, item := range cases {
		r := httptest.NewRequest("GET", "/api/user/", strings.NewReader(""))
		r = mux.SetURLVars(r, map[string]string{"id": fmt.Sprintf("%d", item.User.Id)})

		w := httptest.NewRecorder()

		ctx := r.Context()
		ctx = context.WithValue(ctx, "IsAuth", item.IsAuth)
		if !item.IdErr {
			ctx = context.WithValue(ctx, "Id", item.User.Id)
		}
		r = r.WithContext(ctx)

		if item.IsAuth && !item.IdErr {

			var err error = nil
			if item.GetErr {
				err = NoType.New("")
			}

			gomock.InOrder(
				mockUserService.EXPECT().GetById(r.Context(),
					&userServ.UserID{Id: int64(item.User.Id)}).Return(
					&userServ.User{Id: int64(item.User.Id),
						Email:         item.User.Email,
						Login:         item.User.Login,
						Avatar:        item.User.Avatar,
						About:         item.User.About,
						Subscribers:   int64(item.User.Subscribers),
						Subscriptions: int64(item.User.Subscriptions),
					},
					err,
				),
			)
		}

		userDelivery.GetOtherUser(w, r)

		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		bodyStr := string(body)

		if item.GetErr {
			var output map[string]interface{}

			err := json.NewDecoder(strings.NewReader(bodyStr)).Decode(&output)
			if err != nil {
				t.Fatalf("[%d] wrong decoding Response: got %+v, err: %v",
					caseNum, bodyStr, err.Error())
			}

			status, ok := output["status"]
			if !ok {
				t.Fatalf("[%d] wrong Response: got %+v - no status",
					caseNum, bodyStr)
			}

			if status == 200 {
				t.Errorf("[%d] wrong status Response: got %+v, expected not success status",
					caseNum, status)
			}

		} else {
			if bodyStr != item.Response {
				t.Errorf("[%d] wrong Response: got %+v, expected %+v",
					caseNum, bodyStr, item.Response)
			}
		}

	}
}

type TestCaseUpdateProfile struct {
	IsAuth    bool
	User      *models.User
	Response  string
	IdErr     bool
	InputErr  bool
	ValidErr  bool
	UpdateErr bool
}

func TestHandler_UpdateProfile(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockBoard := boardMock.NewMockIUsecase(ctl)
	mockUserService := userServMock.NewMockUserServiceClient(ctl)
	mockAuthService := authServMock.NewMockAuthSeviceClient(ctl)

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	zap := logger.Sugar().With(
		zap.String("mode", "[access_log]"),
		zap.String("logger", "ZAP"),
	)

	userDelivery := NewHandler(mockUserService, mockAuthService, mockBoard, zap)

	cases := []TestCaseUpdateProfile{
		TestCaseUpdateProfile{
			IsAuth:   false,
			User:     &models.User{},
			Response: `{"status":401,"body":{"error":"User is unauthorized"}}`,
		},
		TestCaseUpdateProfile{
			IsAuth:   true,
			IdErr:    true,
			User:     &models.User{},
			Response: `{"status":500,"body":{"error":"Internal server error"}}`,
		},
		TestCaseUpdateProfile{
			IsAuth:   true,
			InputErr: true,
			User:     &models.User{},
			Response: `{"status":400,"body":{"error":"Wrong body of request"}}`,
		},
		/*TestCaseUpdateProfile{
			IsAuth:   true,
			ValidErr: true,
			Response: `{"status":400,"body":{"error":"Login should be letters and numbers, shorter than 20 characters Email should be like hello@example.com and shorter than 50 characters"}}`,
			User: &models.User{
				Email: "helloexam.com",
				Login: "Login1",
			},
		},
		TestCaseUpdateProfile{
			IsAuth:   true,
			ValidErr: true,
			Response: `{"status":400,"body":{"error":"Login should be letters and numbers, shorter than 20 characters Email should be like hello@example.com and shorter than 50 characters"}}`,
			User: &models.User{
				Email: "helloexam.com",
				Login: "",
			},
		},*/
		TestCaseUpdateProfile{
			IsAuth:    true,
			UpdateErr: true,
			User: &models.User{
				Id:    1,
				Email: "hello@exam.com",
				Login: "login",
			},
		},
		TestCaseUpdateProfile{
			IsAuth: true,
			User: &models.User{
				Id:    1,
				Email: "hello@exam.com",
				Login: "login",
			},
			Response: `{"status":200,"body":{"message":"Ok"}}`,
		},
	}

	for caseNum, item := range cases {
		var r *http.Request
		if item.InputErr == false {
			r = httptest.NewRequest("PUT", "/api/user/settings/profile",
				strings.NewReader(fmt.Sprintf(`{"login":"%s", "email":"%s"}`, item.User.Login, item.User.Email)))
		} else {
			r = httptest.NewRequest("PUT", "/api/user/settings/profile",
				strings.NewReader(fmt.Sprintf(`{"login:"%s", "email":"%s"}`, item.User.Login, item.User.Email)))
		}

		r.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		ctx := r.Context()
		ctx = context.WithValue(ctx, "IsAuth", item.IsAuth)
		if !item.IdErr {
			ctx = context.WithValue(ctx, "Id", item.User.Id)
		}
		r = r.WithContext(ctx)

		if item.IsAuth && !item.InputErr && !item.ValidErr && !item.IdErr {

			input := &userServ.Profile{
				Id: &userServ.UserID{Id: int64(item.User.Id)},
				Input: &userServ.UpdateProfileInput{
					Email: item.User.Email,
					Login: item.User.Login,
				},
			}

			var err error = nil
			if item.UpdateErr {
				err = NoType.New("")
			}

			gomock.InOrder(
				mockUserService.EXPECT().UpdateProfile(r.Context(), input).Return(
					&userServ.Nothing{}, err),
			)

		}

		userDelivery.UpdateProfile(w, r)

		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		bodyStr := string(body)

		if item.UpdateErr {
			var output map[string]interface{}

			err := json.NewDecoder(strings.NewReader(bodyStr)).Decode(&output)
			if err != nil {
				t.Fatalf("[%d] wrong decoding Response: got %+v, err: %v",
					caseNum, bodyStr, err.Error())
			}

			status, ok := output["status"]
			if !ok {
				t.Fatalf("[%d] wrong Response: got %+v - no status",
					caseNum, bodyStr)
			}

			if status == 200 {
				t.Errorf("[%d] wrong status Response: got %+v, expected not success status",
					caseNum, status)
			}
		} else {
			if bodyStr != item.Response {
				t.Errorf("[%d] wrong Response: got %+v, expected %+v",
					caseNum, bodyStr, item.Response)
			}
		}
	}
}

type TestCaseUpdateDescription struct {
	IsAuth    bool
	User      *models.User
	Response  string
	IdErr     bool
	InputErr  bool
	ValidErr  bool
	UpdateErr bool
}

func TestHandler_UpdateDescription(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockBoard := boardMock.NewMockIUsecase(ctl)
	mockUserService := userServMock.NewMockUserServiceClient(ctl)
	mockAuthService := authServMock.NewMockAuthSeviceClient(ctl)

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	zap := logger.Sugar().With(
		zap.String("mode", "[access_log]"),
		zap.String("logger", "ZAP"),
	)

	userDelivery := NewHandler(mockUserService, mockAuthService, mockBoard, zap)

	cases := []TestCaseUpdateDescription{
		TestCaseUpdateDescription{
			IsAuth:   false,
			User:     &models.User{},
			Response: `{"status":401,"body":{"error":"User is unauthorized"}}`,
		},
		TestCaseUpdateDescription{
			IsAuth:   true,
			IdErr:    true,
			User:     &models.User{},
			Response: `{"status":500,"body":{"error":"Internal server error"}}`,
		},
		TestCaseUpdateDescription{
			IsAuth:   true,
			InputErr: true,
			User:     &models.User{},
			Response: `{"status":400,"body":{"error":"Wrong body of request"}}`,
		},
		TestCaseUpdateDescription{
			IsAuth:    true,
			UpdateErr: true,
			User: &models.User{
				Id:    1,
				About: "about me",
			},
		},
		TestCaseUpdateDescription{
			IsAuth: true,
			User: &models.User{
				Id:    1,
				About: "about me",
			},
			Response: `{"status":200,"body":{"message":"Ok"}}`,
		},
	}

	for caseNum, item := range cases {
		var r *http.Request
		if item.InputErr == false {
			r = httptest.NewRequest("PUT", "/api/user/settings/description",
				strings.NewReader(fmt.Sprintf(`{"description":"%s"}`, item.User.About)))
		} else {
			r = httptest.NewRequest("PUT", "/api/user/settings/description",
				strings.NewReader(fmt.Sprintf(`{"description:"%s"}`, item.User.About)))
		}

		r.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		ctx := r.Context()
		ctx = context.WithValue(ctx, "IsAuth", item.IsAuth)
		if !item.IdErr {
			ctx = context.WithValue(ctx, "Id", item.User.Id)
		}
		r = r.WithContext(ctx)

		if item.IsAuth && !item.InputErr && !item.ValidErr && !item.IdErr {

			input := &userServ.Description{
				Id:          &userServ.UserID{Id: int64(item.User.Id)},
				Description: item.User.About,
			}

			var err error = nil
			if item.UpdateErr {
				err = NoType.New("")
			}

			gomock.InOrder(
				mockUserService.EXPECT().UpdateDescription(r.Context(), input).Return(
					&userServ.Nothing{}, err),
			)

		}

		userDelivery.UpdateDescription(w, r)

		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		bodyStr := string(body)

		if item.UpdateErr {
			var output map[string]interface{}

			err := json.NewDecoder(strings.NewReader(bodyStr)).Decode(&output)
			if err != nil {
				t.Fatalf("[%d] wrong decoding Response: got %+v, err: %v",
					caseNum, bodyStr, err.Error())
			}

			status, ok := output["status"]
			if !ok {
				t.Fatalf("[%d] wrong Response: got %+v - no status",
					caseNum, bodyStr)
			}

			if status == 200 {
				t.Errorf("[%d] wrong status Response: got %+v, expected not success status",
					caseNum, status)
			}
		} else {
			if bodyStr != item.Response {
				t.Errorf("[%d] wrong Response: got %+v, expected %+v",
					caseNum, bodyStr, item.Response)
			}
		}

	}
}

type TestCaseUpdatePassword struct {
	IsAuth    bool
	UserId    uint
	Password  string
	Response  string
	IdErr     bool
	InputErr  bool
	ValidErr  bool
	UpdateErr bool
}

func TestHandler_UpdatePassword(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockBoard := boardMock.NewMockIUsecase(ctl)
	mockUserService := userServMock.NewMockUserServiceClient(ctl)
	mockAuthService := authServMock.NewMockAuthSeviceClient(ctl)

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	zap := logger.Sugar().With(
		zap.String("mode", "[access_log]"),
		zap.String("logger", "ZAP"),
	)

	userDelivery := NewHandler(mockUserService, mockAuthService, mockBoard, zap)

	cases := []TestCaseUpdatePassword{
		TestCaseUpdatePassword{
			IsAuth:   false,
			UserId:   1,
			Password: "password",
			Response: `{"status":401,"body":{"error":"User is unauthorized"}}`,
		},
		TestCaseUpdatePassword{
			IsAuth:   true,
			IdErr:    true,
			UserId:   1,
			Password: "password",
			Response: `{"status":500,"body":{"error":"Internal server error"}}`,
		},
		TestCaseUpdatePassword{
			IsAuth:   true,
			InputErr: true,
			UserId:   1,
			Password: "password",
			Response: `{"status":400,"body":{"error":"Wrong body of request"}}`,
		},
		/*TestCaseUpdatePassword{
					IsAuth:   true,
					ValidErr: true,
					Response: `{"status":400,"body":{"error":"Password should be longer than 6 characters and shorter 100."}}
		`,
					UserId:   1,
					Password: "pas",
				},*/
		TestCaseUpdatePassword{
			IsAuth:    true,
			UpdateErr: true,
			UserId:    1,
			Password:  "password",
		},
		TestCaseUpdatePassword{
			IsAuth:   true,
			UserId:   1,
			Password: "password",
			Response: `{"status":200,"body":{"message":"Ok"}}`,
		},
	}

	for caseNum, item := range cases {
		var r *http.Request
		if item.InputErr == false {
			r = httptest.NewRequest("PUT", "/api/user/settings/password",
				strings.NewReader(fmt.Sprintf(`{"password":"%s"}`, item.Password)))
		} else {
			r = httptest.NewRequest("PUT", "/api/user/settings/password",
				strings.NewReader(fmt.Sprintf(`{"password:"%s"}`, item.Password)))
		}

		r.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		ctx := r.Context()
		ctx = context.WithValue(ctx, "IsAuth", item.IsAuth)
		if !item.IdErr {
			ctx = context.WithValue(ctx, "Id", item.UserId)
		}
		r = r.WithContext(ctx)

		if item.IsAuth && !item.InputErr && !item.ValidErr && !item.IdErr {

			input := &userServ.Password{
				Id:       &userServ.UserID{Id: int64(item.UserId)},
				Password: item.Password,
			}

			var err error = nil
			if item.UpdateErr {
				err = NoType.New("")
			}

			gomock.InOrder(
				mockUserService.EXPECT().UpdatePassword(r.Context(), input).Return(
					&userServ.Nothing{}, err))

		}

		userDelivery.UpdatePassword(w, r)

		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		bodyStr := string(body)

		if item.UpdateErr {
			var output map[string]interface{}

			err := json.NewDecoder(strings.NewReader(bodyStr)).Decode(&output)
			if err != nil {
				t.Fatalf("[%d] wrong decoding Response: got %+v, err: %v",
					caseNum, bodyStr, err.Error())
			}

			status, ok := output["status"]
			if !ok {
				t.Fatalf("[%d] wrong Response: got %+v - no status",
					caseNum, bodyStr)
			}

			if status == 200 {
				t.Errorf("[%d] wrong status Response: got %+v, expected not success status",
					caseNum, status)
			}
		} else {
			if bodyStr != item.Response {
				t.Errorf("[%d] wrong Response: got %+v, expected %+v",
					caseNum, bodyStr, item.Response)
			}
		}

	}
}

type TestCaseFollow struct {
	IsAuth    bool
	UserId    uint
	SubId     uint
	Response  string
	IdErr     bool
	BadIdErr  bool
	SubIdErr  bool
	FollowErr bool
}

func TestHandler_Follow(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockBoard := boardMock.NewMockIUsecase(ctl)
	mockUserService := userServMock.NewMockUserServiceClient(ctl)
	mockAuthService := authServMock.NewMockAuthSeviceClient(ctl)

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	zap := logger.Sugar().With(
		zap.String("mode", "[access_log]"),
		zap.String("logger", "ZAP"),
	)

	userDelivery := NewHandler(mockUserService, mockAuthService, mockBoard, zap)

	cases := []TestCaseFollow{
		TestCaseFollow{
			IsAuth:   false,
			UserId:   1,
			SubId:    2,
			Response: `{"status":401,"body":{"error":"User is unauthorized"}}`,
		},
		TestCaseFollow{
			IsAuth:   true,
			IdErr:    true,
			UserId:   1,
			SubId:    2,
			Response: `{"status":500,"body":{"error":"Internal server error"}}`,
		},
		TestCaseFollow{
			IsAuth:   true,
			BadIdErr: true,
			UserId:   1,
			SubId:    2,
			Response: `{"status":400,"body":{"error":"Bad id"}}`,
		},
		TestCaseFollow{
			IsAuth:   true,
			SubIdErr: true,
			UserId:   1,
			SubId:    1,
			Response: `{"status":400,"body":{"error":"Your id and following id shoudn't match"}}`,
		},
		TestCaseFollow{
			IsAuth:    true,
			FollowErr: true,
			UserId:    1,
			SubId:     2,
		},
		TestCaseFollow{
			IsAuth:   true,
			UserId:   1,
			SubId:    2,
			Response: `{"status":201,"body":{"message":"Ok"}}`,
		},
	}

	for caseNum, item := range cases {
		r := httptest.NewRequest("POST", "/api/user/following/",
			strings.NewReader(""))
		if !item.BadIdErr {
			r = mux.SetURLVars(r, map[string]string{"id": fmt.Sprintf("%d", item.SubId)})
		} else {
			r = mux.SetURLVars(r, map[string]string{"id": "h-6"})
		}

		//r.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		ctx := r.Context()
		ctx = context.WithValue(ctx, "IsAuth", item.IsAuth)
		if !item.IdErr {
			ctx = context.WithValue(ctx, "Id", item.UserId)
		}
		r = r.WithContext(ctx)

		if item.IsAuth && !item.IdErr && !item.SubIdErr && !item.BadIdErr {

			var err error = nil
			if item.FollowErr {
				err = NoType.New("")
			}

			gomock.InOrder(
				mockUserService.EXPECT().Follow(r.Context(),
					&userServ.Following{
						Id:    &userServ.UserID{Id: int64(item.UserId)},
						SubId: &userServ.UserID{Id: int64(item.SubId)},
					}).Return(&userServ.Nothing{}, err),
			)

		}

		userDelivery.Follow(w, r)

		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		bodyStr := string(body)

		if item.FollowErr {
			var output map[string]interface{}

			err := json.NewDecoder(strings.NewReader(bodyStr)).Decode(&output)
			if err != nil {
				t.Fatalf("[%d] wrong decoding Response: got %+v, err: %v",
					caseNum, bodyStr, err.Error())
			}

			status, ok := output["status"]
			if !ok {
				t.Fatalf("[%d] wrong Response: got %+v - no status",
					caseNum, bodyStr)
			}

			if status == 200 {
				t.Errorf("[%d] wrong status Response: got %+v, expected not success status",
					caseNum, status)
			}

		} else {
			if bodyStr != item.Response {
				t.Errorf("[%d] wrong Response: got %+v, expected %+v",
					caseNum, bodyStr, item.Response)
			}
		}

	}
}

func TestHandler_Unfollow(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockBoard := boardMock.NewMockIUsecase(ctl)
	mockUserService := userServMock.NewMockUserServiceClient(ctl)
	mockAuthService := authServMock.NewMockAuthSeviceClient(ctl)

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	zap := logger.Sugar().With(
		zap.String("mode", "[access_log]"),
		zap.String("logger", "ZAP"),
	)

	userDelivery := NewHandler(mockUserService, mockAuthService, mockBoard, zap)

	cases := []TestCaseFollow{
		TestCaseFollow{
			IsAuth:   false,
			UserId:   1,
			SubId:    2,
			Response: `{"status":401,"body":{"error":"User is unauthorized"}}`,
		},
		TestCaseFollow{
			IsAuth:   true,
			IdErr:    true,
			UserId:   1,
			SubId:    2,
			Response: `{"status":500,"body":{"error":"Internal server error"}}`,
		},
		TestCaseFollow{
			IsAuth:   true,
			BadIdErr: true,
			UserId:   1,
			SubId:    2,
			Response: `{"status":400,"body":{"error":"Bad id"}}`,
		},
		TestCaseFollow{
			IsAuth:   true,
			SubIdErr: true,
			UserId:   1,
			SubId:    1,
			Response: `{"status":400,"body":{"error":"Your id and unfollowing id shoudn't match"}}`,
		},
		TestCaseFollow{
			IsAuth:    true,
			FollowErr: true,
			UserId:    1,
			SubId:     2,
		},
		TestCaseFollow{
			IsAuth:   true,
			UserId:   1,
			SubId:    2,
			Response: `{"status":200,"body":{"message":"Ok"}}`,
		},
	}

	for caseNum, item := range cases {
		r := httptest.NewRequest("POST", "/api/user/unfollowing/",
			strings.NewReader(""))
		if !item.BadIdErr {
			r = mux.SetURLVars(r, map[string]string{"id": fmt.Sprintf("%d", item.SubId)})
		} else {
			r = mux.SetURLVars(r, map[string]string{"id": "h-6"})
		}

		w := httptest.NewRecorder()

		ctx := r.Context()
		ctx = context.WithValue(ctx, "IsAuth", item.IsAuth)
		if !item.IdErr {
			ctx = context.WithValue(ctx, "Id", item.UserId)
		}
		r = r.WithContext(ctx)

		if item.IsAuth && !item.IdErr && !item.SubIdErr && !item.BadIdErr {

			var err error = nil
			if item.FollowErr {
				err = NoType.New("")
			}

			gomock.InOrder(
				mockUserService.EXPECT().Unfollow(r.Context(),
					&userServ.Following{
						Id:    &userServ.UserID{Id: int64(item.UserId)},
						SubId: &userServ.UserID{Id: int64(item.SubId)},
					}).Return(&userServ.Nothing{}, err),
			)

		}

		userDelivery.Unfollow(w, r)

		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		bodyStr := string(body)

		if item.FollowErr {
			var output map[string]interface{}

			err := json.NewDecoder(strings.NewReader(bodyStr)).Decode(&output)
			if err != nil {
				t.Fatalf("[%d] wrong decoding Response: got %+v, err: %v",
					caseNum, bodyStr, err.Error())
			}

			status, ok := output["status"]
			if !ok {
				t.Fatalf("[%d] wrong Response: got %+v - no status",
					caseNum, bodyStr)
			}

			if status == 200 || status == 201 {
				t.Errorf("[%d] wrong status Response: got %+v, expected not success status",
					caseNum, status)
			}
		} else {
			if bodyStr != item.Response {
				t.Errorf("[%d] wrong Response: got %+v, expected %+v",
					caseNum, bodyStr, item.Response)
			}
		}

	}
}

type TestCaseGet struct {
	IsAuth   bool
	UserId   uint
	Response string
	GetErr   bool
	Users    []*models.User
}

func TestHandler_GetSubscribers(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockBoard := boardMock.NewMockIUsecase(ctl)
	mockUserService := userServMock.NewMockUserServiceClient(ctl)
	mockAuthService := authServMock.NewMockAuthSeviceClient(ctl)

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	zap := logger.Sugar().With(
		zap.String("mode", "[access_log]"),
		zap.String("logger", "ZAP"),
	)

	userDelivery := NewHandler(mockUserService, mockAuthService, mockBoard, zap)

	cases := []TestCaseGet{
		TestCaseGet{
			IsAuth:   false,
			UserId:   1,
			Response: `{"status":401,"body":{"error":"User is unauthorized"}}`,
		},
		TestCaseGet{
			IsAuth: true,
			GetErr: true,
			UserId: 1,
			Users:  nil,
		},
		TestCaseGet{
			IsAuth: true,
			UserId: 1,
			Users: []*models.User{
				&models.User{
					Id:            2,
					Login:         "login1",
					About:         "about me",
					Avatar:        "avatar.jpg",
					Subscribers:   5,
					Subscriptions: 1,
				},
				&models.User{
					Id:            3,
					Login:         "login2",
					About:         "about me",
					Avatar:        "avatar.jpg",
					Subscribers:   6,
					Subscriptions: 2,
				},
			},
			Response: `{"status":200,"body":[{"id":2,"login":"login1","about":"about me","avatar":"avatar.jpg","subscriptions":1,"subscribers":5},` +
				`{"id":3,"login":"login2","about":"about me","avatar":"avatar.jpg","subscriptions":2,"subscribers":6}]}`,
		},
	}

	for caseNum, item := range cases {
		r := httptest.NewRequest("GET", fmt.Sprintf("/api/user/subscribers/%d/?start=1&limit=15", item.UserId),
			strings.NewReader(""))
		r = mux.SetURLVars(r, map[string]string{"id": fmt.Sprintf("%d", item.UserId)})

		//r.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		ctx := r.Context()
		ctx = context.WithValue(ctx, "IsAuth", item.IsAuth)
		r = r.WithContext(ctx)

		if item.IsAuth {

			var err error = nil
			if item.GetErr {
				err = NoType.New("")
			}

			u := &userServ.Users{}

			for _, us := range item.Users {
				u.Users = append(u.Users,
					&userServ.User{
						Id:            int64(us.Id),
						Login:         us.Login,
						About:         us.About,
						Avatar:        us.Avatar,
						Subscribers:   int64(us.Subscribers),
						Subscriptions: int64(us.Subscriptions),
					})
			}

			gomock.InOrder(
				mockUserService.EXPECT().GetSubscribers(r.Context(),
					&userServ.Sub{Id: &userServ.UserID{Id: int64(item.UserId)},
						Start: 1,
						Limit: 15,
					}).Return(
					u, err),
			)

		}

		userDelivery.GetSubscribers(w, r)

		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		bodyStr := string(body)

		if item.GetErr {
			var output map[string]interface{}

			err := json.NewDecoder(strings.NewReader(bodyStr)).Decode(&output)
			if err != nil {
				t.Fatalf("[%d] wrong decoding Response: got %+v, err: %v",
					caseNum, bodyStr, err.Error())
			}

			status, ok := output["status"]
			if !ok {
				t.Fatalf("[%d] wrong Response: got %+v - no status",
					caseNum, bodyStr)
			}

			if status == 200 {
				t.Errorf("[%d] wrong status Response: got %+v, expected not success status",
					caseNum, status)
			}

		} else {
			if bodyStr != item.Response {
				t.Errorf("[%d] wrong Response: got %+v, expected %+v",
					caseNum, bodyStr, item.Response)
			}
		}

	}
}

func TestHandler_GetSubscriptions(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockBoard := boardMock.NewMockIUsecase(ctl)
	mockUserService := userServMock.NewMockUserServiceClient(ctl)
	mockAuthService := authServMock.NewMockAuthSeviceClient(ctl)

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	zap := logger.Sugar().With(
		zap.String("mode", "[access_log]"),
		zap.String("logger", "ZAP"),
	)

	userDelivery := NewHandler(mockUserService, mockAuthService, mockBoard, zap)

	cases := []TestCaseGet{
		TestCaseGet{
			IsAuth:   false,
			UserId:   1,
			Response: `{"status":401,"body":{"error":"User is unauthorized"}}`,
		},
		TestCaseGet{
			IsAuth: true,
			GetErr: true,
			UserId: 1,
			Users:  nil,
		},
		TestCaseGet{
			IsAuth: true,
			UserId: 1,
			Users: []*models.User{
				&models.User{
					Id:            2,
					Login:         "login1",
					About:         "about me",
					Avatar:        "avatar.jpg",
					Subscribers:   5,
					Subscriptions: 1,
				},
				&models.User{
					Id:            3,
					Login:         "login2",
					About:         "about me",
					Avatar:        "avatar.jpg",
					Subscribers:   6,
					Subscriptions: 2,
				},
			},
			Response: `{"status":200,"body":[{"id":2,"login":"login1","about":"about me","avatar":"avatar.jpg","subscriptions":1,"subscribers":5},` +
				`{"id":3,"login":"login2","about":"about me","avatar":"avatar.jpg","subscriptions":2,"subscribers":6}]}`,
		},
	}

	for caseNum, item := range cases {
		r := httptest.NewRequest("GET", fmt.Sprintf("/api/user/subscriptions/%d/?start=1&limit=15", item.UserId),
			strings.NewReader(""))
		r = mux.SetURLVars(r, map[string]string{"id": fmt.Sprintf("%d", item.UserId)})

		//r.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		ctx := r.Context()
		ctx = context.WithValue(ctx, "IsAuth", item.IsAuth)
		r = r.WithContext(ctx)

		if item.IsAuth {

			var err error = nil
			if item.GetErr {
				err = NoType.New("")
			}

			u := &userServ.Users{}

			for _, us := range item.Users {
				u.Users = append(u.Users,
					&userServ.User{
						Id:            int64(us.Id),
						Login:         us.Login,
						About:         us.About,
						Avatar:        us.Avatar,
						Subscribers:   int64(us.Subscribers),
						Subscriptions: int64(us.Subscriptions),
					})
			}

			gomock.InOrder(
				mockUserService.EXPECT().GetSubscriptions(r.Context(),
					&userServ.Sub{Id: &userServ.UserID{Id: int64(item.UserId)},
						Start: 1,
						Limit: 15,
					}).Return(
					u, err),
			)

		}

		userDelivery.GetSubscribtions(w, r)

		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		bodyStr := string(body)

		if item.GetErr {
			var output map[string]interface{}

			err := json.NewDecoder(strings.NewReader(bodyStr)).Decode(&output)
			if err != nil {
				t.Fatalf("[%d] wrong decoding Response: got %+v, err: %v",
					caseNum, bodyStr, err.Error())
			}

			status, ok := output["status"]
			if !ok {
				t.Fatalf("[%d] wrong Response: got %+v - no status",
					caseNum, bodyStr)
			}

			if status == 200 {
				t.Errorf("[%d] wrong status Response: got %+v, expected not success status",
					caseNum, status)
			}

		} else {
			if bodyStr != item.Response {
				t.Errorf("[%d] wrong Response: got %+v, expected %+v",
					caseNum, bodyStr, item.Response)
			}
		}

	}
}
