package http

import (
	"2020_1_Color_noise/internal/models"
	. "2020_1_Color_noise/internal/pkg/error"
	sessionMock "2020_1_Color_noise/internal/pkg/session/mock"
	userMock "2020_1_Color_noise/internal/pkg/user/mock"
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type TestCaseLogin struct {
	IsAuth     bool
	Login      string
	Password   string
	User       *models.User
	CookieName string
	Cookie     string
	TokenName  string
	Token      string
	Response   string
	InputErr   bool
	GetErr     bool
	CompareErr bool
	CreateErr  bool
}

func TestHandler_Login(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockUserUsecase := userMock.NewMockIUsecase(ctl)
	mockSessionUsecase := sessionMock.NewMockIUsecase(ctl)

	sessionDelivery := NewHandler(mockSessionUsecase, mockUserUsecase)

	cases := []TestCaseLogin{
		TestCaseLogin{
			IsAuth: true,
			Response: `{"status":200,"body":{"message":"Ok"}}
`,
		},
		TestCaseLogin{
			InputErr: true,
			Login:    "Login1",
			Password: "Password1",
			Response: `{"status":400,"body":{"error":"Wrong body of request"}}
`,
		},
		TestCaseLogin{
			GetErr: true,
			User: &models.User{
				Id: 1,
			},
			Login:    "Login1",
			Password: "Password1",
		},
		TestCaseLogin{
			CompareErr: true,
			User: &models.User{
				Id: 1,
			},
			Login:    "Login1",
			Password: "Password1",
		},
		TestCaseLogin{
			CreateErr: true,
			User: &models.User{
				Id: 1,
			},
			Login:    "Login1",
			Password: "Password1",
		},
		TestCaseLogin{
			User: &models.User{
				Id: 1,
			},
			Login:    "Login1",
			Password: "Password",
			Response: `{"status":200,"body":{"message":"Ok"}}
`,
			CookieName: "session_id",
			Cookie:     "cookie",
			//TokenName:  "csrf_token",
			//Token:      "token",
		},
	}

	for caseNum, item := range cases {
		var r *http.Request
		if item.InputErr == false {
			r = httptest.NewRequest("POST", "/api/auth",
				strings.NewReader(fmt.Sprintf(`{"login":"%s", "password":"%s"}`, item.Login, item.Password)))
		} else {
			r = httptest.NewRequest("POST", "/api/auth",
				strings.NewReader(fmt.Sprintf(`{"login:"%s", "password":"%s"}`, item.Login, item.Password)))
		}

		r.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		ctx := r.Context()
		ctx = context.WithValue(ctx, "IsAuth", item.IsAuth)
		r = r.WithContext(ctx)

		if !item.IsAuth && !item.InputErr {

			input := &models.SignUpInput{
				Login:    item.Login,
				Password: item.Password,
			}

			var err error = nil
			if item.GetErr {
				err = NoType.New("")
			}

			gomock.InOrder(
				mockUserUsecase.EXPECT().GetByLogin(input.Login).Return(item.User, err),
			)

			if !item.GetErr {
				if item.CompareErr {
					err = NoType.New("")
				}

				gomock.InOrder(
					mockUserUsecase.EXPECT().ComparePassword(item.User, item.Password).Return(err),
				)

				if !item.CompareErr {
					session := &models.Session{
						Id:     item.User.Id,
						Cookie: item.Cookie,
						//Token:  item.Token,
					}

					if item.CreateErr {
						err = NoType.New("")
					}

					gomock.InOrder(
						mockSessionUsecase.EXPECT().CreateSession(item.User.Id).Return(session, err),
					)

				}
			}
		}

		sessionDelivery.Login(w, r)

		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		bodyStr := string(body)

		if item.CreateErr || item.CompareErr || item.GetErr {
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

			if !item.IsAuth && !item.InputErr {
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

type TestCaseLogout struct {
	IsAuth     bool
	CookieName string
	Cookie     string
	TokenName  string
	Token      string
	Response   string
	CookieErr  bool
	DeleteErr  bool
}

func TestHandler_Logout(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockUserUsecase := userMock.NewMockIUsecase(ctl)
	mockSessionUsecase := sessionMock.NewMockIUsecase(ctl)

	sessionDelivery := NewHandler(mockSessionUsecase, mockUserUsecase)

	cases := []TestCaseLogout{
		TestCaseLogout{
			IsAuth: false,
			Response: `{"status":401,"body":{"error":"User is unauthorized"}}
`,
		},
		TestCaseLogout{
			IsAuth:    true,
			CookieErr: true,
		},
		TestCaseLogout{
			DeleteErr: true,
		},
		TestCaseLogout{
			IsAuth: true,
			Response: `{"status":200,"body":{"message":"Ok"}}
`,
			CookieName: "session_id",
			Cookie:     "cookie",
			//TokenName:  "csrf_token",
			//Token:      "token",
		},
	}

	for caseNum, item := range cases {
		r := httptest.NewRequest("DELETE", "/api/auth", strings.NewReader(""))

		r.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		ctx := r.Context()
		ctx = context.WithValue(ctx, "IsAuth", item.IsAuth)
		r = r.WithContext(ctx)

		if !item.CookieErr {
			cookie := &http.Cookie{
				Name:     item.CookieName,
				Value:    item.Cookie,
				Expires:  time.Now().Add(5),
				HttpOnly: true,
				//Domain:   r.Host,
			}
			r.AddCookie(cookie)
		}

		if item.IsAuth && !item.CookieErr {
			var err error = nil
			if item.DeleteErr {
				err = NoType.New("")
			}

			gomock.InOrder(
				mockSessionUsecase.EXPECT().Delete(item.Cookie).Return(err),
			)
		}

		sessionDelivery.Logout(w, r)

		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		bodyStr := string(body)

		if item.CookieErr || item.DeleteErr {
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

			if item.IsAuth {
				cookies := w.Result().Cookies()
				fmt.Println(cookies)
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
