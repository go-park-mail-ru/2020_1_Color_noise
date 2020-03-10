package http

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"

	//"fmt"

	//"fmt"
	"io/ioutil"

	//"io/ioutil"
	"net/http/httptest"
	"pinterest/internal/mock"
	"pinterest/internal/models"

	"strings"

	gomock "github.com/golang/mock/gomock"
	"testing"
	//"time"
)

type TestCase struct {
	IsAuth	   bool
	ID         uint
	GetID      string
	Login      string
	Password   string
	Email      string
	ErrAdd     error
	ErrGet     error
	ErrUpdate  error
	CookieName string
	Cookie     string
	TokenName  string
	User       models.User
	Token      string
	ErrSession error
	ContextID  uint
	Response   string
	StatusCode int
}

func TestAddUser(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockUserUsecase := mock.NewMockIUserUsecase(ctl)
	mockSessionUsecase := mock.NewMockISessionUsecase(ctl)

	userDelivery := NewUserDelivery(mockUserUsecase, mockSessionUsecase)

	cases := []TestCase{
		TestCase{
			IsAuth:     true,
			ID:         2,
			Login:		"login1",
			Password:	"password1",
			StatusCode: 200,
			Response:   `{"status":"200","body":{"id":2}}
`,
		},
		TestCase{
			IsAuth:     false,
			ID:         1,
			ErrAdd:     nil,
			Login:		"login1",
			Password:	"password1",
			CookieName: "session_id",
			Cookie:     "cookie",
			TokenName:  "csrf_token",
			Token:      "token",
			ErrSession: nil,
			StatusCode: 200,
			Response:   `{"status":"200","body":{"id":1}}
`,
		},
		TestCase{
			IsAuth:     false,
			ID:         1,
			ErrAdd:     fmt.Errorf("some error"),
			Login:		"login1",
			Password:	"password1",
			CookieName: "session_id",
			Cookie:     "cookie",
			TokenName:  "csrf_token",
			Token:      "token",
			ErrSession: nil,
			StatusCode: 200,
			Response:   `{"status":"500","body":{"error":"some error"}}
`,
		},
		TestCase{
			IsAuth:     false,
			ID:         1,
			ErrAdd:     nil,
			Login:		"login1",
			Password:	"password1",
			CookieName: "session_id",
			Cookie:     "cookie",
			TokenName:  "csrf_token",
			Token:      "token",
			ErrSession: fmt.Errorf("some error"),
			StatusCode: 200,
			Response:   `{"status":"500","body":{"error":"some error"}}
`,
		},
	}

	for caseNum, item := range cases {
		r := httptest.NewRequest("POST", "/signup",
			strings.NewReader("login=" + item.Login +"&password=" + item.Password))
		r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
		w := httptest.NewRecorder()
		ctx := r.Context()
		ctx = context.WithValue(ctx, "isAuth", item.IsAuth)
		ctx = context.WithValue(ctx, "Id", item.ID)
		r = r.WithContext(ctx)

		if !item.IsAuth {
			user := &models.User{
				Login:    item.Login,
				Password: item.Password,
			}
			gomock.InOrder(
				mockUserUsecase.EXPECT().Add(user).Return(item.ID, item.ErrAdd),
			)

			if item.ErrAdd == nil {
				session := &models.Session{
					Id:     item.ID,
					Cookie: item.Cookie,
					Token:  item.Token,
				}

				gomock.InOrder(
					mockSessionUsecase.EXPECT().CreateSession(item.ID).Return(session, item.ErrSession),
				)
			}
		}

		userDelivery.AddUser(w, r)

		if w.Code != item.StatusCode {
			t.Errorf("[%d] wrong StatusCode: got %d, expected %d",
				caseNum, w.Code, item.StatusCode)
		}
		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		bodyStr := string(body)
		if bodyStr != item.Response {
			t.Errorf("[%d] wrong Response: got %+v, expected %+v",
				caseNum, bodyStr, item.Response)
		}
		if !item.IsAuth && item.ErrSession == nil && item.ErrAdd == nil {
			cookies := w.Result().Cookies()
			if len(cookies) != 2 {
				t.Errorf("[%d] No Cookie", caseNum)
			}
			if cookies[0].Name != item.CookieName {
				t.Errorf("[%d] wrong Cookie: got %+v, expected %+v",
					caseNum, cookies[0].Name, item.CookieName)
			}
			if cookies[0].Value != item.Cookie {
				t.Errorf("[%d] wrong Cookie: got %+v, expected %+v",
					caseNum, cookies[0].Value, item.Cookie)
			}

			if cookies[1].Name != item.TokenName {
				t.Errorf("[%d] wrong Cookie: got %+v, expected %+v",
					caseNum, cookies[1].Name, item.TokenName)
			}
			if cookies[1].Value != item.Token {
				t.Errorf("[%d] wrong Cookie: got %+v, expected %+v",
					caseNum, cookies[1].Value, item.Token)
			}
		}
	}
}

func TestGetUser(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockUserUsecase := mock.NewMockIUserUsecase(ctl)
	mockSessionUsecase := mock.NewMockISessionUsecase(ctl)

	userDelivery := NewUserDelivery(mockUserUsecase, mockSessionUsecase)

	cases := []TestCase{
		TestCase{
			IsAuth:     false,
			GetID:      "2",
			StatusCode: 200,
			Response:   `{"status":"403","body":{"error":"User not found"}}
`,
		},
		TestCase{
			IsAuth:     true,
			GetID:      "h",
			ErrGet:     nil,
			StatusCode: 200,
			Response:   `{"status":"500","body":{"error":"Invalid id"}}
`,
		},
		TestCase{
			IsAuth:     true,
			GetID:      "1",
			ErrGet:     fmt.Errorf("some error"),
			StatusCode: 200,
			Response:   `{"status":"500","body":{"error":"some error"}}
`,
		},
		TestCase{
			IsAuth:     true,
			GetID:      "1",
			User:		models.User{
						Id: 1,
						Login: "login1",
			},
			ErrGet:     nil,
			StatusCode: 200,
			Response:   `{"status":"200","body":{"user":{"id":1,"login":"login1"}}}
`,
		},
	}

	for caseNum, item := range cases {
		r := httptest.NewRequest("GET", "/profile/" + item.GetID,
			strings.NewReader(""))
		r = mux.SetURLVars(r, map[string]string{"id": item.GetID})
		w := httptest.NewRecorder()
		ctx := r.Context()
		ctx = context.WithValue(ctx, "isAuth", item.IsAuth)
		r = r.WithContext(ctx)

		id, err := strconv.Atoi(item.GetID)
		if item.IsAuth && err == nil {
			gomock.InOrder(
				mockUserUsecase.EXPECT().GetById(uint(id)).Return(&item.User, item.ErrGet),
			)
		}

		userDelivery.GetUser(w, r)

		if w.Code != item.StatusCode {
			t.Errorf("[%d] wrong StatusCode: got %d, expected %d",
				caseNum, w.Code, item.StatusCode)
		}
		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		bodyStr := string(body)
		if bodyStr != item.Response {
			t.Errorf("[%d] wrong Response: got %+v, expected %+v",
				caseNum, bodyStr, item.Response)
		}
	}
}

func TestUpdateUser(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockUserUsecase := mock.NewMockIUserUsecase(ctl)
	mockSessionUsecase := mock.NewMockISessionUsecase(ctl)

	userDelivery := NewUserDelivery(mockUserUsecase, mockSessionUsecase)

	cases := []TestCase{
		TestCase{
			IsAuth:     false,
			ID:         1,
			Login:		"login1",
			Email:	    "email1",
			StatusCode: 200,
			Response:   `{"status":"403","body":{"error":"User not found"}}
`,
		},
		TestCase{
			IsAuth:     true,
			ID:         1,
			GetID:      "h",
			ErrAdd:     nil,
			Login:		"login1",
			Email:	    "email1",
			ErrSession: nil,
			StatusCode: 200,
			Response:   `{"status":"500","body":{"error":"Invalid id"}}
`,
		},
		TestCase{
			IsAuth:     true,
			ID:         1,
			GetID:      "1",
			ErrGet:     fmt.Errorf("some error"),
			Login:		"login1",
			Email:	    "email1",
			User:		models.User{
				Id: 1,
				Login: "login1",
				Email: "email1",
			},
			ErrSession: nil,
			StatusCode: 200,
			Response:   `{"status":"500","body":{"error":"Internal error"}}
`,
		},
		TestCase{
			IsAuth:     true,
			ID:         1,
			GetID:      "1",
			ErrGet:     nil,
			Login:		"login1",
			Email:	    "email1",
			User:		models.User{
				Id: 1,
				Login: "login1",
				Email: "email1",
			},
			ErrSession: nil,
			StatusCode: 200,
			Response:   `{"status":"500","body":{"error":"Internal error"}}
`,
		},
		TestCase{
			IsAuth:     true,
			ID:         1,
			GetID:      "1",
			ErrGet:     nil,
			ContextID:  2,
			Login:		"login1",
			Email:	    "email1",
			User:		models.User{
				Id: 1,
				Login: "login1",
				Email: "email1",
			},
			ErrSession: nil,
			StatusCode: 200,
			Response:   `{"status":"500","body":{"error":"Internal error"}}
`,
		},
		TestCase{
			IsAuth:     true,
			ID:         1,
			GetID:      "1",
			ErrGet:     nil,
			ContextID:  1,
			ErrUpdate:  fmt.Errorf("some error"),
			Login:		"login1",
			Email:	    "email1",
			User:		models.User{
				Id: 1,
				Login: "login1",
				Email: "email1",
			},
			ErrSession: nil,
			StatusCode: 200,
			Response:   `{"status":"500","body":{"error":"some error"}}
`,
		},
		TestCase{
			IsAuth:     true,
			ID:         1,
			GetID:      "1",
			ErrGet:     nil,
			ContextID:  1,
			ErrUpdate:  nil,
			Login:		"login1",
			Email:	    "email1",
			User:		models.User{
				Id: 1,
				Login: "login1",
				Email: "email1",
			},
			ErrSession: nil,
			StatusCode: 200,
			Response:   `{"status":"200"}
`,
		},
	}

	for caseNum, item := range cases {
		r := httptest.NewRequest("POST", "/profile/" + item.GetID,
			strings.NewReader("login=" + item.Login +"&email=" + item.Email))
		r = mux.SetURLVars(r, map[string]string{"id": item.GetID})
		r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
		w := httptest.NewRecorder()
		ctx := r.Context()
		ctx = context.WithValue(ctx, "isAuth", item.IsAuth)
		ctx = context.WithValue(ctx, "Id", item.ContextID)
		r = r.WithContext(ctx)

		id, err := strconv.Atoi(item.GetID)
		if item.IsAuth && err == nil {
			gomock.InOrder(
				mockUserUsecase.EXPECT().GetById(uint(id)).Return(&item.User, item.ErrGet),
			)
			if item.ErrGet == nil && item.ContextID == item.ID && uint(id) == item.ID {
				gomock.InOrder(
					mockUserUsecase.EXPECT().Update(&item.User).Return(item.ErrUpdate),
				)
			}
		}

		userDelivery.UpdateUser(w, r)

		if w.Code != item.StatusCode {
			t.Errorf("[%d] wrong StatusCode: got %d, expected %d",
				caseNum, w.Code, item.StatusCode)
		}
		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		bodyStr := string(body)
		if bodyStr != item.Response {
			t.Errorf("[%d] wrong Response: got %+v, expected %+v",
				caseNum, bodyStr, item.Response)
		}
	}
}

func TestDeleteUser(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockUserUsecase := mock.NewMockIUserUsecase(ctl)
	mockSessionUsecase := mock.NewMockISessionUsecase(ctl)

	userDelivery := NewUserDelivery(mockUserUsecase, mockSessionUsecase)

	cases := []TestCase{
		TestCase{
			IsAuth:     false,
			ID:         1,
			Cookie:     "cookie",
			StatusCode: 200,
			Response:   `{"status":"403","body":{"error":"User not found"}}
`,
		},
		TestCase{
			IsAuth:     true,
			ID:         1,
			GetID:      "h",
			ErrAdd:     nil,
			Login:		"login1",
			Password:	"password1",
			CookieName: "session_id",
			Cookie:     "cookie",
			TokenName:  "token",
			Token:      "token",
			ErrSession: nil,
			StatusCode: 200,
			Response:   `{"status":"500","body":{"error":"Invalid id"}}
`,
		},
		TestCase{
			IsAuth:     true,
			ID:         1,
			GetID:      "1",
			ErrGet:     fmt.Errorf("some error"),
			CookieName: "session_id",
			Cookie:     "cookie",
			TokenName:  "token",
			Token:      "token",
			User:		models.User{
				Id: 1,
				Login: "login1",
				Email: "email1",
			},
			ErrSession: nil,
			StatusCode: 200,
			Response:   `{"status":"500","body":{"error":"some error"}}
`,
		},
		TestCase{
			IsAuth:     true,
			ID:         1,
			GetID:      "1",
			ErrGet:     nil,
			CookieName: "session_id",
			Cookie:     "cookie",
			TokenName:  "token",
			Token:      "token",
			User:		models.User{
				Id: 1,
				Login: "login1",
				Email: "email1",
			},
			ErrSession: nil,
			StatusCode: 200,
			Response:   `{"status":"500","body":{"error":"Internal error"}}
`,
		},
		TestCase{
			IsAuth:     true,
			ID:         1,
			GetID:      "1",
			ErrGet:     nil,
			ContextID:  2,
			CookieName: "session_id",
			Cookie:     "cookie",
			TokenName:  "token",
			Token:      "token",
			User:		models.User{
				Id: 1,
				Login: "login1",
				Email: "email1",
			},
			ErrSession: nil,
			StatusCode: 200,
			Response:   `{"status":"500","body":{"error":"Internal error"}}
`,
		},
		TestCase{
			IsAuth:     true,
			ID:         1,
			GetID:      "1",
			ErrGet:     nil,
			CookieName: "session_id",
			Cookie:     "cookie",
			TokenName:  "token",
			Token:      "token",
			ContextID:  1,
			ErrUpdate:  fmt.Errorf("some error"),
			User:		models.User{
				Id: 1,
				Login: "login1",
				Email: "email1",
			},
			ErrSession: nil,
			StatusCode: 200,
			Response:   `{"status":"500","body":{"error":"some error"}}
`,
		},
		TestCase{
			IsAuth:     true,
			ID:         1,
			GetID:      "1",
			ErrGet:     nil,
			CookieName: "session_id",
			Cookie:     "cookie",
			TokenName:  "token",
			Token:      "token",
			ContextID:  1,
			ErrUpdate:  nil,
			User:		models.User{
				Id: 1,
				Login: "login1",
				Email: "email1",
			},
			ErrSession: nil,
			StatusCode: 200,
			Response:   `{"status":"200"}
`,
		},
	}

	for caseNum, item := range cases {
		r := httptest.NewRequest("DELETE", "/profile/" + item.GetID,
			strings.NewReader(""))
		r = mux.SetURLVars(r, map[string]string{"id": item.GetID})
		w := httptest.NewRecorder()
		ctx := r.Context()
		ctx = context.WithValue(ctx, "isAuth", item.IsAuth)
		ctx = context.WithValue(ctx, "Id", item.ContextID)
		r = r.WithContext(ctx)
		cookie := &http.Cookie{
			Name:    item.CookieName,
			Value:   item.Cookie,
			Expires: time.Now().Add(10 * time.Hour),
			HttpOnly: true,
			Domain: r.Host,
		}
		token := &http.Cookie{
			Name:    item.TokenName,
			Value:   item.Token,
			Expires: time.Now().Add(10 * time.Hour),
			HttpOnly: true,
			Domain: r.Host,
		}
		r.AddCookie(cookie)
		r.AddCookie(token)
		//http.Cookie(w, cookie)

		id, err := strconv.Atoi(item.GetID)
		if item.IsAuth && err == nil {
			gomock.InOrder(
				mockUserUsecase.EXPECT().GetById(uint(id)).Return(&item.User, item.ErrGet),
			)
			if item.ErrGet == nil && item.ContextID == item.ID && uint(id) == item.ID {
				gomock.InOrder(
					mockUserUsecase.EXPECT().Delete(uint(id)).Return(item.ErrUpdate),
				)
				if item.ErrUpdate == nil {
					gomock.InOrder(
						mockSessionUsecase.EXPECT().Delete("cookie").Return(item.ErrUpdate),
					)
				}
			}
		}

		userDelivery.DeleteUser(w, r)

		if w.Code != item.StatusCode {
			t.Errorf("[%d] wrong StatusCode: got %d, expected %d",
				caseNum, w.Code, item.StatusCode)
		}
		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		bodyStr := string(body)
		if bodyStr != item.Response {
			t.Errorf("[%d] wrong Response: got %+v, expected %+v",
				caseNum, bodyStr, item.Response)
		}
	}
}




