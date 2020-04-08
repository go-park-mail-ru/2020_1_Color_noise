package search

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"2020_1_Color_noise/internal/models"
	commentMock "2020_1_Color_noise/internal/pkg/comment/mock"
	. "2020_1_Color_noise/internal/pkg/error"
	pinMock "2020_1_Color_noise/internal/pkg/pin/mock"
	userMock "2020_1_Color_noise/internal/pkg/user/mock"
	"context"
	"io/ioutil"
	//"io/ioutil"
	"net/http/httptest"

	"strings"

	gomock "github.com/golang/mock/gomock"
	"testing"
	//"time"
)

type TestCaseSearch struct {
	IsAuth     bool
	UserId     uint
	IsComment  bool
	IsPin	   bool
	IsUser	   bool
	Pins	   []*models.Pin
	Users	   []*models.User
	Comments   []*models.Comment
	Response   string
	Err        bool
	Line       string
	Start      int
	Limit 	   int
}

func TestHandler_Search(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockUserUsecase := userMock.NewMockIUsecase(ctl)
	mockPinUsecase := pinMock.NewMockIUsecase(ctl)
	mockCommentUsecase := commentMock.NewMockIUsecase(ctl)

	searchDelivery := NewHandler(mockCommentUsecase, mockPinUsecase, mockUserUsecase)

	cases := []TestCaseSearch{
		TestCaseSearch{
			IsAuth:     false,
			UserId:     1,
			Start:		1,
			Limit:		15,
			Response:	`{"status":401,"body":{"error":"User is unauthorized"}}
`,
		},
		TestCaseSearch{
			IsAuth:     true,
			UserId:     1,
			IsComment:	true,
			Err:        true,
			Start:		1,
			Limit:		15,
		},
		TestCaseSearch{
			IsAuth:     true,
			IsComment:	true,
			UserId:     1,
			Start:		1,
			Limit:		15,
			Comments:       []*models.Comment {
				&models.Comment{
					Id:		2,
					UserId: 1,
					PinId:	1,
					Text:   "comment",
				},
				&models.Comment{
					Id:		3,
					UserId: 4,
					PinId:	1,
					Text:   "comment",
				},
			},
			Response: `{"status":200,"body":[{"id":2,"user_id":1,"pin_id":1,"created_at":"0001-01-01T00:00:00Z","comment":"comment"},` +
				`{"id":3,"user_id":4,"pin_id":1,"created_at":"0001-01-01T00:00:00Z","comment":"comment"}]}
`,
		},
		TestCaseSearch{
			IsAuth:     true,
			UserId:     1,
			IsUser: 	true,
			Err:        true,
			Start:		1,
			Limit:		15,
		},
		TestCaseSearch{
			IsAuth:     true,
			IsUser:     true,
			UserId:     1,
			Start:		1,
			Limit:		15,
			Users:      []*models.User {
				&models.User{
					Id:           1,
					Login:        "login",
					About:         "about me",
					Avatar:        "avatar.jpg",
					Subscribers:   11000,
					Subscriptions: 100,
				},
				&models.User{
					Id:           1,
					Login:        "login",
					About:         "about me",
					Avatar:        "avatar.jpg",
					Subscribers:   11000,
					Subscriptions: 100,
				},
			},
			Response: `{"status":200,"body":[{"id":1,"login":"login","about":"about me","avatar":"avatar.jpg","subscriptions":100,"subscribers":11000},` +
				`{"id":1,"login":"login","about":"about me","avatar":"avatar.jpg","subscriptions":100,"subscribers":11000}]}
`,
		},
		TestCaseSearch{
			IsAuth:     true,
			UserId:     1,
			IsPin:   	true,
			Err:        true,
			Start:		1,
			Limit:		15,
		},
		TestCaseSearch{
			IsAuth:     true,
			IsPin:      true,
			UserId:     1,
			Start:		1,
			Limit:		15,
			Pins:       []*models.Pin {
				&models.Pin {
					Id: 		 1,
					BoardId:     2,
					UserId:      1,
					Name:        "name1",
					Description: "desc1",
					Image:       "image.jpg",
				},
				&models.Pin {
					Id: 		 2,
					BoardId:     5,
					UserId:      1,
					Name:        "name2",
					Description: "desc2",
					Image:       "image.jpg",
				},
			},
			Response: `{"status":200,"body":[{"id":1,"user_id":1,"board_id":2,"name":"name1","description":"desc1","image":"image.jpg"},` +
				`{"id":2,"user_id":1,"board_id":5,"name":"name2","description":"desc2","image":"image.jpg"}]}
`,
		},
		TestCaseSearch{
			IsAuth:     true,
			UserId:     1,
			Start:		1,
			Limit:		15,
			Response:	`{"status":404,"body":{"error":"Not found"}}
`,
		},
	}

	for caseNum, item := range cases {
		var r *http.Request
		switch {
		case item.IsComment:
			r = httptest.NewRequest("GET", fmt.Sprintf("/api/search?what=%s&description=%s&start=%d&limit=%d", "comment", item.Line, item.Start, item.Limit), strings.NewReader(""))
		case item.IsPin:
			r = httptest.NewRequest("GET", fmt.Sprintf("/api/search?what=%s&description=%s&start=%d&limit=%d", "pin", item.Line, item.Start, item.Limit), strings.NewReader(""))
		case item.IsUser:
			r = httptest.NewRequest("GET", fmt.Sprintf("/api/search?what=%s&description=%s&start=%d&limit=%d", "user", item.Line, item.Start, item.Limit), strings.NewReader(""))
		default:
			r = httptest.NewRequest("GET", fmt.Sprintf("/api/search"), strings.NewReader(""))
		}

		r = mux.SetURLVars(r, map[string]string{"description": item.Line})

		w := httptest.NewRecorder()

		ctx := r.Context()
		ctx = context.WithValue(ctx, "IsAuth", item.IsAuth)
		/*if !item.IdErr {
			ctx = context.WithValue(ctx, "Id", item.UserId)
		}*/
		r = r.WithContext(ctx)

		if item.IsAuth && item.IsComment {

			var err error = nil
			if item.Err {
				err = NoType.New("")
			}

			gomock.InOrder(
				mockCommentUsecase.EXPECT().GetByText(item.Line, item.Start, item.Limit).Return(item.Comments, err),
			)
		}

		if item.IsAuth && item.IsPin {

			var err error = nil
			if item.Err {
				err = NoType.New("")
				log.Println(err.Error())
			}

			gomock.InOrder(
				mockPinUsecase.EXPECT().GetByName(item.Line, item.Start, item.Limit).Return(item.Pins, err),
			)
		}

		if item.IsAuth && item.IsUser {

			var err error = nil
			if item.Err {
				err = NoType.New("")
			}

			gomock.InOrder(
				mockUserUsecase.EXPECT().Search(item.Line, item.Start, item.Limit).Return(item.Users, err),
			)
		}

		searchDelivery.Search(w, r)

		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		bodyStr := string(body)

		if item.Err {
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
