package http

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"

	//"github.com/gorilla/mux"
	"net/http"

	"2020_1_Color_noise/internal/models"
	"2020_1_Color_noise/internal/pkg/comment/mock"
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
	IsAuth    bool
	UserId    uint
	Comment   *models.Comment
	Response  string
	IdErr     bool
	InputErr  bool
	ValidErr  bool
	CreateErr bool
}

func TestHandler_Create(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	zap := logger.Sugar().With(
		zap.String("mode", "[access_log]"),
		zap.String("logger", "ZAP"),
	)

	mockCommentUsecase := mock.NewMockIUsecase(ctl)
	commentDelivery := NewHandler(mockCommentUsecase, zap)

	cases := []TestCaseCreate{
		TestCaseCreate{
			Comment:  &models.Comment{},
			UserId:   1,
			IsAuth:   false,
			Response: `{"status":401,"body":{"error":"User is unauthorized"}}`,
		},
		TestCaseCreate{
			IsAuth:   true,
			UserId:   1,
			IdErr:    true,
			Comment:  &models.Comment{},
			Response: `{"status":500,"body":{"error":"Internal server error"}}`,
		},
		TestCaseCreate{
			IsAuth:   true,
			InputErr: true,
			UserId:   1,
			Comment:  &models.Comment{},
			Response: `{"status":400,"body":{"error":"Bad request"}}`,
		},
		TestCaseCreate{
			IsAuth:   true,
			ValidErr: true,
			UserId:   1,
			Comment: &models.Comment{
				Id:     2,
				UserId: 1,
				PinId:  1,
				Text:   "",
			},
			Response: `{"status":400,"body":{"error":"Bad request"}}`,
		},
		TestCaseCreate{
			IsAuth:    true,
			CreateErr: true,
			UserId:    1,
			Comment: &models.Comment{
				Id:     2,
				UserId: 1,
				PinId:  1,
				Text:   "comment",
			},
		},
		TestCaseCreate{
			IsAuth: true,
			UserId: 1,
			Comment: &models.Comment{
				Id:    2,
				PinId: 1,
				Text:  "comment",
			},
			Response: `{"status":201,"body":{"id":2}}`,
		},
	}

	for caseNum, item := range cases {
		var r *http.Request
		if item.InputErr == false {
			r = httptest.NewRequest("POST", "/api/comment",
				strings.NewReader(fmt.Sprintf(`{"pin_id":%d,"comment":"%s"}`, item.Comment.PinId, item.Comment.Text)))
		} else {
			r = httptest.NewRequest("POST", "/api/pin",
				strings.NewReader(fmt.Sprintf(`{"comment:"%s"}`, item.Comment.Text)))
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

			input := &models.InputComment{
				Text:  item.Comment.Text,
				PinId: item.Comment.PinId,
			}

			var err error = nil
			if item.CreateErr {
				err = NoType.New("")
			}

			gomock.InOrder(
				mockCommentUsecase.EXPECT().Create(input, item.UserId).Return(item.Comment.Id, err),
			)
		}

		commentDelivery.Create(w, r)

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

type TestCaseGetComment struct {
	IsAuth   bool
	UserId   uint
	Comment  *models.Comment
	Response string
	IdErr    bool
	GetErr   bool
}

func TestHandler_GetComment(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	zap := logger.Sugar().With(
		zap.String("mode", "[access_log]"),
		zap.String("logger", "ZAP"),
	)

	mockCommentUsecase := mock.NewMockIUsecase(ctl)
	commentDelivery := NewHandler(mockCommentUsecase, zap)

	cases := []TestCaseGetComment{
		TestCaseGetComment{
			IsAuth:   true,
			IdErr:    true,
			Comment:  &models.Comment{},
			Response: `{"status":400,"body":{"error":"Bad request"}}`,
		},
		TestCaseGetComment{
			IsAuth:  true,
			GetErr:  true,
			Comment: &models.Comment{},
		},
		TestCaseGetComment{
			IsAuth: true,
			Comment: &models.Comment{
				Id:     2,
				UserId: 1,
				PinId:  1,
				Text:   "comment",
			},
			Response: `{"status":200,"body":{"id":2,"user_id":1,"pin_id":1,"created_at":"0001-01-01T00:00:00Z","comment":"comment"}}`,
		},
	}

	for caseNum, item := range cases {
		r := httptest.NewRequest("GET", "/api/comment/", strings.NewReader(""))
		if !item.IdErr {
			r = mux.SetURLVars(r, map[string]string{"id": fmt.Sprintf("%d", item.Comment.Id)})
		} else {
			r = mux.SetURLVars(r, map[string]string{"id": "j"})
		}

		w := httptest.NewRecorder()

		ctx := r.Context()
		ctx = context.WithValue(ctx, "IsAuth", item.IsAuth)
		r = r.WithContext(ctx)

		if !item.IdErr {

			var err error = nil
			if item.GetErr {
				err = NoType.New("")
			}

			gomock.InOrder(
				mockCommentUsecase.EXPECT().GetById(item.Comment.Id).Return(item.Comment, err),
			)
		}

		commentDelivery.GetComment(w, r)

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
	IsAuth   bool
	PinId    uint
	Comments []*models.Comment
	Response string
	IdErr    bool
	GetErr   bool
	Start    int
	Limit    int
}

func TestHandler_Fetch(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	zap := logger.Sugar().With(
		zap.String("mode", "[access_log]"),
		zap.String("logger", "ZAP"),
	)

	mockCommentUsecase := mock.NewMockIUsecase(ctl)
	commentDelivery := NewHandler(mockCommentUsecase, zap)

	cases := []TestCaseFetch{
		TestCaseFetch{
			IsAuth:   true,
			IdErr:    true,
			PinId:    1,
			Start:    1,
			Limit:    15,
			Response: `{"status":400,"body":{"error":"Bad request"}}`,
		},
		TestCaseFetch{
			IsAuth:   true,
			GetErr:   true,
			PinId:    1,
			Start:    1,
			Limit:    15,
			Comments: nil,
		},
		TestCaseFetch{
			IsAuth: true,
			PinId:  1,
			Start:  1,
			Limit:  15,
			Comments: []*models.Comment{
				&models.Comment{
					Id:     2,
					UserId: 1,
					PinId:  1,
					Text:   "comment",
				},
				&models.Comment{
					Id:     3,
					UserId: 4,
					PinId:  1,
					Text:   "comment",
				},
			},
			Response: `{"status":200,"body":[{"id":2,"user_id":1,"pin_id":1,"created_at":"0001-01-01T00:00:00Z","comment":"comment"},` +
				`{"id":3,"user_id":4,"pin_id":1,"created_at":"0001-01-01T00:00:00Z","comment":"comment"}]}`,
		},
	}

	for caseNum, item := range cases {
		r := httptest.NewRequest("GET", fmt.Sprintf("/api/comment/pin/%d/?start=%d&limit=%d", item.PinId, item.Start, item.Limit), strings.NewReader(""))
		if !item.IdErr {
			r = mux.SetURLVars(r, map[string]string{"id": fmt.Sprintf("%d", item.PinId)})
		} else {
			r = mux.SetURLVars(r, map[string]string{"id": "j"})
		}

		w := httptest.NewRecorder()

		ctx := r.Context()
		ctx = context.WithValue(ctx, "IsAuth", item.IsAuth)
		r = r.WithContext(ctx)

		if !item.IdErr {

			var err error = nil
			if item.GetErr {
				err = NoType.New("")
			}

			gomock.InOrder(
				mockCommentUsecase.EXPECT().GetByPinId(uint(item.PinId), item.Start, item.Limit).Return(item.Comments, err),
			)
		}

		commentDelivery.Fetch(w, r)

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
