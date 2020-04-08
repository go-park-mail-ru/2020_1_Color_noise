package http

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	//"github.com/gorilla/mux"
	"net/http"

	"2020_1_Color_noise/internal/models"
	"2020_1_Color_noise/internal/pkg/board/mock"
	. "2020_1_Color_noise/internal/pkg/error"
	"context"
	"io/ioutil"
	//"io/ioutil"
	"net/http/httptest"

	"strings"

	gomock "github.com/golang/mock/gomock"
	"testing"
	//"time"
)

type TestCaseCreate struct {
	IsAuth     bool
	UserId     uint
	Board	   *models.Board
	Response   string
	IdErr      bool
	InputErr   bool
	ValidErr   bool
	CreateErr  bool
}

func TestHandler_Create(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockBoardUsecase := mock.NewMockIUsecase(ctl)
	boardDelivery := NewHandler(mockBoardUsecase)

	cases := []TestCaseCreate{
		TestCaseCreate{
			Board:        &models.Board{},
			IsAuth:     false,
			UserId:     1,
			Response:	`{"status":401,"body":{"error":"User is unauthorized"}}
`,
		},
		TestCaseCreate{
			IsAuth:     true,
			IdErr:      true,
			Board:      &models.Board{},
			UserId:     1,
			Response:	`{"status":500,"body":{"error":"Internal server error"}}
`,
		},
		TestCaseCreate{
			IsAuth:     true,
			InputErr:   true,
			UserId:     1,
			Board:      &models.Board{},
			Response:	`{"status":400,"body":{"error":"Wrong body of request"}}
`,
		},
		TestCaseCreate{
			IsAuth:     true,
			ValidErr:   true,
			UserId:     1,
			Board:		&models.Board{
				Name:	     "",
				Description: "desc",
			},
			Response:	`{"status":400,"body":{"error":"Name shouldn't be empty and longer 60 characters. `+
				`Description shouldn't be empty and longer 1000 characters."}}
`,
		},
		TestCaseCreate{
			IsAuth:     true,
			ValidErr:   true,
			UserId:     1,
			Board:		&models.Board{
				Name:	     "nameddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddd",
				Description: "desc",
			},
			Response:	`{"status":400,"body":{"error":"Name shouldn't be empty and longer 60 characters. `+
				`Description shouldn't be empty and longer 1000 characters."}}
`,
		},
		TestCaseCreate{
			IsAuth:     true,
			ValidErr:   true,
			UserId:     1,
			Board:		&models.Board{
				Name:	     "name",
				Description: "",
			},
			Response:	`{"status":400,"body":{"error":"Name shouldn't be empty and longer 60 characters. `+
				`Description shouldn't be empty and longer 1000 characters."}}
`,
		},
		TestCaseCreate{
			IsAuth:     true,
			CreateErr:   true,
			UserId:     1,
			Board:		&models.Board{
				Id:			 1,
				Name:	     "name",
				Description: "desc",
			},
		},
		TestCaseCreate{
			IsAuth:     true,
			UserId:     1,
			Board:		&models.Board{
				Id:			 1,
				Name:	     "name",
				Description: "desc",
			},
			Response:	`{"status":201,"body":{"id":1}}
`,
		},
	}

	for caseNum, item := range cases {
		var r *http.Request
		if item.InputErr == false {
			r = httptest.NewRequest("POST", "/api/board",
				strings.NewReader(fmt.Sprintf(`{"name":"%s", "description":"%s"}`, item.Board.Name, item.Board.Description)))
		} else {
			r = httptest.NewRequest("POST", "/api/pin",
				strings.NewReader(fmt.Sprintf(`{"name:"%s", "description":"%s"}`, item.Board.Name, item.Board.Description)))
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

			input := &models.InputBoard{
				Name:        	item.Board.Name,
				Description:    item.Board.Description,
			}

			var err error = nil
			if item.CreateErr {
				err = NoType.New("")
			}

			gomock.InOrder(
				mockBoardUsecase.EXPECT().Create(input, item.UserId).Return(uint(item.Board.Id), err),
			)
		}

		boardDelivery.Create(w, r)

		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		bodyStr := string(body)

		if item.CreateErr {
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

type TestCaseGetBoard struct {
	IsAuth     bool
	UserId     uint
	Board      *models.Board
	Response   string
	IdErr      bool
	GetErr     bool
}

func TestHandler_GetBoard(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockBoardUsecase := mock.NewMockIUsecase(ctl)
	boardDelivery := NewHandler(mockBoardUsecase)

	cases := []TestCaseGetBoard{
		TestCaseGetBoard{
			IsAuth:     false,
			Board:       &models.Board{},
			Response:	`{"status":401,"body":{"error":"User is unauthorized"}}
`,
		},
		TestCaseGetBoard{
			IsAuth:     true,
			IdErr:      true,
			Board:       &models.Board{},
			Response:	`{"status":400,"body":{"error":"Bad id"}}
`,
		},
		TestCaseGetBoard{
			IsAuth:     true,
			GetErr:     true,
			Board:       &models.Board{},
		},
		TestCaseGetBoard{
			IsAuth:     true,
			GetErr:     true,
			Board:      &models.Board{
				Id:          2,
				UserId:      3,
				Name:        "name",
				Description: "desc",
				Pins:        []*models.Pin{
					&models.Pin{
						Id:          1,
						BoardId:     2,
						UserId:      3,
						Name:        "name1",
						Description: "desc1",
						Image:       "image.jpg",
					},
					&models.Pin{
						Id:          2,
						BoardId:     2,
						UserId:      3,
						Name:        "name2",
						Description: "desc2",
						Image:       "image.jpg",
					},
				},
			},
			Response: `{"status":200,"body":{"id":2,"user_id":3,"name":"name","description":"desc","pins":[{"id":1,"user_id":3,"board_id":2,"name":"name1","description":"desc1","image":"image.jpg"},` +
				`{"id":2,"user_id":3,"board_id":2,"name":"name2","description":"desc2","image":"image.jpg"}]}}
`,
		},
	}

	for caseNum, item := range cases {
		r := httptest.NewRequest("GET", "/api/board/", strings.NewReader(""))
		if !item.IdErr {
			r = mux.SetURLVars(r, map[string]string{"id": fmt.Sprintf("%d", item.Board.Id)})
		} else {
			r = mux.SetURLVars(r, map[string]string{"id": "j"})
		}

		w := httptest.NewRecorder()

		ctx := r.Context()
		ctx = context.WithValue(ctx, "IsAuth", item.IsAuth)
		r = r.WithContext(ctx)

		if item.IsAuth && !item.IdErr {

			var err error = nil
			if item.GetErr {
				err = NoType.New("")
			}

			gomock.InOrder(
				mockBoardUsecase.EXPECT().GetById(item.Board.Id).Return(item.Board, err),
			)
		}

		boardDelivery.GetBoard(w, r)

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

func TestHandler_GetNameBoard(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockBoardUsecase := mock.NewMockIUsecase(ctl)
	boardDelivery := NewHandler(mockBoardUsecase)

	cases := []TestCaseGetBoard{
		TestCaseGetBoard{
			IsAuth:     false,
			Board:       &models.Board{},
			Response:	`{"status":401,"body":{"error":"User is unauthorized"}}
`,
		},
		TestCaseGetBoard{
			IsAuth:     true,
			IdErr:      true,
			Board:       &models.Board{},
			Response:	`{"status":400,"body":{"error":"Bad id"}}
`,
		},
		TestCaseGetBoard{
			IsAuth:     true,
			GetErr:     true,
			Board:       &models.Board{},
		},
		TestCaseGetBoard{
			IsAuth:     true,
			GetErr:     true,
			Board:      &models.Board{
				Id:          2,
				UserId:      3,
				Name:        "name",
				Description: "desc",
				LastPin:      models.Pin{
						Id:          1,
						BoardId:     2,
						UserId:      3,
						Name:        "name1",
						Description: "desc1",
						Image:       "image.jpg",
				},
			},
			Response: `{"status":200,"body":{"id":2,"user_id":3,"name":"name","description":"desc","last_pin":{"id":1,"user_id":3,"board_id":2,"name":"name1","description":"desc1","image":"image.jpg"}}
`,
		},
	}

	for caseNum, item := range cases {
		r := httptest.NewRequest("GET", "/api/board/name/user/", strings.NewReader(""))
		if !item.IdErr {
			r = mux.SetURLVars(r, map[string]string{"id": fmt.Sprintf("%d", item.Board.Id)})
		} else {
			r = mux.SetURLVars(r, map[string]string{"id": "j"})
		}

		w := httptest.NewRecorder()

		ctx := r.Context()
		ctx = context.WithValue(ctx, "IsAuth", item.IsAuth)
		r = r.WithContext(ctx)

		if item.IsAuth && !item.IdErr {

			var err error = nil
			if item.GetErr {
				err = NoType.New("")
			}

			gomock.InOrder(
				mockBoardUsecase.EXPECT().GetById(item.Board.Id).Return(item.Board, err),
			)
		}

		boardDelivery.GetNameBoard(w, r)

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

type TestCaseFetch struct {
	IsAuth     bool
	UserId     uint
	Boards	   []*models.Board
	Response   string
	IdErr      bool
	GetErr     bool
	Start      int
	Limit 	   int
}


func TestHandler_Fetch(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockBoardUsecase := mock.NewMockIUsecase(ctl)
	boardDelivery := NewHandler(mockBoardUsecase)

	cases := []TestCaseFetch{
		TestCaseFetch{
			IsAuth:     false,
			UserId:     3,
			Start:		1,
			Limit:		15,
			Response:	`{"status":401,"body":{"error":"User is unauthorized"}}
`,
		},
		TestCaseFetch{
			IsAuth:     true,
			IdErr:		true,
			UserId:     3,
			Start:		1,
			Limit:		15,
			Response:	`{"status":400,"body":{"error":"Bad id"}}
`,
		},
		TestCaseFetch{
			IsAuth:     true,
			GetErr:     true,
			UserId:     3,
			Start:		1,
			Limit:		15,
			Boards:     nil,
		},
		TestCaseFetch{
			IsAuth:     true,
			GetErr:     true,
			UserId:     3,
			Start:		1,
			Limit:		15,
			Boards:     []*models.Board {
				&models.Board{
					Id:          2,
					UserId:      3,
					Name:        "name1",
					Description: "desc1",
					LastPin:      models.Pin{
						Id:          1,
						BoardId:     2,
						UserId:      3,
						Name:        "name1",
						Description: "desc1",
						Image:       "image.jpg",
					},
				},
				&models.Board{
					Id:          4,
					UserId:      3,
					Name:        "name2",
					Description: "desc2",
					LastPin:      models.Pin{
						Id:          6,
						BoardId:     4,
						UserId:      3,
						Name:        "name2",
						Description: "desc2",
						Image:       "image.jpg",
					},
				},
			},
			Response: `{"status":200,"body":[{{"id":2,"user_id":3,"name":"name1","description":"desc1","last_pin":{"id":1,"user_id":3,"board_id":2,"name":"name1","description":"desc1","image":"image.jpg"},` +
				`{"id":4,"user_id":3,"name":"name2","description":"desc2","last_pin":{"id":6,"user_id":3,"board_id":4,"name":"name2","description":"desc2","image":"image.jpg"}]}
`,
		},
	}

	for caseNum, item := range cases {
		r := httptest.NewRequest("GET", fmt.Sprintf("/api/board/user/%d/?start=%d&limit=%d",item.UserId, item.Start, item.Limit), strings.NewReader(""))
		if !item.IdErr {
			r = mux.SetURLVars(r, map[string]string{"id": fmt.Sprintf("%d", item.UserId)})
		} else {
			r = mux.SetURLVars(r, map[string]string{"id": "j"})
		}

		w := httptest.NewRecorder()

		ctx := r.Context()
		ctx = context.WithValue(ctx, "IsAuth", item.IsAuth)
		r = r.WithContext(ctx)

		if item.IsAuth && !item.IdErr {

			var err error = nil
			if item.GetErr {
				err = NoType.New("")
			}

			gomock.InOrder(
				mockBoardUsecase.EXPECT().GetByUserId(uint(item.UserId), item.Start, item.Limit).Return(item.Boards, err),
			)
		}

		boardDelivery.Fetch(w, r)

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

type TestCaseUpdate struct {
	IsAuth     bool
	UserId     uint
	Board	   *models.Board
	Response   string
	UserIdErr  bool
	IdErr      bool
	InputErr   bool
	ValidErr   bool
	UpdateErr  bool
}

func TestHandler_Update(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockBoardUsecase := mock.NewMockIUsecase(ctl)
	boardDelivery := NewHandler(mockBoardUsecase)

	cases := []TestCaseUpdate{
		TestCaseUpdate{
			Board:      &models.Board{},
			IsAuth:     false,
			UserId:     1,
			Response:	`{"status":401,"body":{"error":"User is unauthorized"}}
`,
		},
		TestCaseUpdate{
			IsAuth:     true,
			UserIdErr:  true,
			UserId:     1,
			Board:      &models.Board{},
		},
		TestCaseUpdate{
			IsAuth:     true,
			IdErr:		true,
			UserId:     1,
			Board:      &models.Board{},
			Response:	`{"status":400,"body":{"error":"Bad id"}}
`,
		},
		TestCaseUpdate{
			IsAuth:     true,
			InputErr:   true,
			UserId:     1,
			Board:      &models.Board{},
			Response:	`{"status":400,"body":{"error":"Wrong body of request"}}
`,
		},
		TestCaseUpdate{
			IsAuth:     true,
			ValidErr:   true,
			UserId:     1,
			Board:      &models.Board{
				Name:	     "",
				Description: "desc",
			},
			Response:	`{"status":400,"body":{"error":"Name shouldn't be empty and longer 60 characters. `+
				`Description shouldn't be empty and longer 1000 characters."}}
`,
		},
		TestCaseUpdate{
			IsAuth:     true,
			ValidErr:   true,
			UserId:     1,
			Board:      &models.Board{
				Name:	     "nameddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddd",
				Description: "desc",
			},
			Response:	`{"status":400,"body":{"error":"Name shouldn't be empty and longer 60 characters. `+
				`Description shouldn't be empty and longer 1000 characters."}}
`,
		},
		TestCaseUpdate{
			IsAuth:     true,
			ValidErr:   true,
			UserId:     1,
			Board:      &models.Board{
				Name:	     "name",
				Description: "",
			},
			Response:	`{"status":400,"body":{"error":"Name shouldn't be empty and longer 60 characters. `+
				`Description shouldn't be empty and longer 1000 characters."}}
`,
		},
		TestCaseUpdate{
			IsAuth:     true,
			UpdateErr:  true,
			UserId:     1,
			Board:      &models.Board{
				Id:			 1,
				Name:	     "name",
				Description: "desc",
			},
		},
		TestCaseUpdate{
			IsAuth:     true,
			UserId:     1,
			Board:      &models.Board{
				Id:			 1,
				Name:	     "name",
				Description: "desc",
			},
			Response:	`{"status":200,"body":{"message":"Ok"}}
`,
		},
	}

	for caseNum, item := range cases {
		var r *http.Request
		if item.InputErr == false {
			r = httptest.NewRequest("PUT", "/api/board",
				strings.NewReader(fmt.Sprintf(`{"name":"%s", "description":"%s"}`, item.Board.Name, item.Board.Description)))
		} else {
			r = httptest.NewRequest("PUT", "/api/board",
				strings.NewReader(fmt.Sprintf(`{"name:"%s", "description":"%s"}`, item.Board.Name, item.Board.Description)))
		}

		if !item.IdErr {
			r = mux.SetURLVars(r, map[string]string{"id": fmt.Sprintf("%d", item.Board.Id)})
		} else {
			r = mux.SetURLVars(r, map[string]string{"id": "j"})
		}

		r.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		ctx := r.Context()
		ctx = context.WithValue(ctx, "IsAuth", item.IsAuth)
		if !item.UserIdErr {
			ctx = context.WithValue(ctx, "Id", item.UserId)
		}
		r = r.WithContext(ctx)

		if item.IsAuth && !item.InputErr && !item.ValidErr && !item.IdErr && !item.UserIdErr {

			input := &models.InputBoard{
				Name:        	item.Board.Name,
				Description:    item.Board.Description,
			}

			var err error = nil
			if item.UpdateErr {
				err = NoType.New("")
			}

			gomock.InOrder(
				mockBoardUsecase.EXPECT().Update(input, item.Board.Id, item.UserId).Return(err),
			)
		}

		boardDelivery.Update(w, r)

		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		bodyStr := string(body)

		if item.UpdateErr || item.UserIdErr  {
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

type TestCaseDelete struct {
	IsAuth     bool
	UserId     uint
	BoardId	   uint
	Response   string
	UserIdErr  bool
	IdErr      bool
	DeleteErr  bool
}

func TestHandler_DeletePin(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockBoardUsecase := mock.NewMockIUsecase(ctl)
	boardDelivery := NewHandler(mockBoardUsecase)

	cases := []TestCaseDelete{
		TestCaseDelete{
			IsAuth:     false,
			Response:	`{"status":401,"body":{"error":"User is unauthorized"}}
`,
		},
		TestCaseDelete{
			IsAuth:     true,
			UserIdErr:  true,
			BoardId:    1,
			UserId:		1,
		},
		TestCaseDelete{
			IsAuth:     true,
			IdErr:		true,
			BoardId:    1,
			UserId:		1,
			Response:	`{"status":400,"body":{"error":"Bad id"}}
`,
		},
		TestCaseDelete{
			IsAuth:     true,
			DeleteErr:  true,
			BoardId:    1,
			UserId:		1,
		},
		TestCaseDelete{
			IsAuth:     true,
			BoardId:    1,
			UserId:		1,
			Response:	`{"status":200,"body":{"message":"Ok"}}
`,
		},
	}

	for caseNum, item := range cases {
		r := httptest.NewRequest("DELETE", fmt.Sprintf("/api/board/%d", item.BoardId), strings.NewReader(""))

		if !item.IdErr {
			r = mux.SetURLVars(r, map[string]string{"id": fmt.Sprintf("%d", item.BoardId)})
		} else {
			r = mux.SetURLVars(r, map[string]string{"id": "j"})
		}

		r.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		ctx := r.Context()
		ctx = context.WithValue(ctx, "IsAuth", item.IsAuth)
		if !item.UserIdErr {
			ctx = context.WithValue(ctx, "Id", item.UserId)
		}
		r = r.WithContext(ctx)

		if item.IsAuth && !item.IdErr && !item.UserIdErr {

			var err error = nil
			if item.DeleteErr {
				err = NoType.New("")
			}

			gomock.InOrder(
				mockBoardUsecase.EXPECT().Delete(item.BoardId, item.UserId).Return(err),
			)
		}

		boardDelivery.Delete(w, r)

		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		bodyStr := string(body)

		if item.DeleteErr || item.UserIdErr  {
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