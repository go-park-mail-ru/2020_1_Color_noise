package http

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"

	"2020_1_Color_noise/internal/models"
	. "2020_1_Color_noise/internal/pkg/error"
	"2020_1_Color_noise/internal/pkg/list/mock"
	"context"
	"io/ioutil"
	//"io/ioutil"
	"net/http/httptest"

	"strings"

	gomock "github.com/golang/mock/gomock"
	"testing"
	//"time"
)

type TestCaseFetch struct {
	IsAuth     bool
	UserId     uint
	Pins	   []*models.Pin
	Response   string
	IdErr      bool
	GetErr     bool
	Start      int
	Limit 	   int
}

func TestHandler_GetSubList(t *testing.T) {
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

	mockListUsecase := mock.NewMockIUsecase(ctl)
	listDelivery := NewHandler(mockListUsecase, zap)

	cases := []TestCaseFetch{
		TestCaseFetch{
			IsAuth:     false,
			UserId:     1,
			Start:		1,
			Limit:		15,
			Response:	`{"status":401,"body":{"error":"User is unauthorized"}}`,
		},
		TestCaseFetch{
			IsAuth:     true,
			IdErr:		true,
			UserId:     1,
			Start:		1,
			Limit:		15,
		},
		TestCaseFetch{
			IsAuth:     true,
			GetErr:     true,
			UserId:     1,
			Start:		1,
			Limit:		15,
			Pins:       nil,
		},
		TestCaseFetch{
			IsAuth:     true,
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
				`{"id":2,"user_id":1,"board_id":5,"name":"name2","description":"desc2","image":"image.jpg"}]}`,
		},
	}

	for caseNum, item := range cases {
		r := httptest.NewRequest("GET", fmt.Sprintf("/api/list/sub?start=%d&limit=%d", item.Start, item.Limit), strings.NewReader(""))

		w := httptest.NewRecorder()

		ctx := r.Context()
		ctx = context.WithValue(ctx, "IsAuth", item.IsAuth)
		if !item.IdErr {
			ctx = context.WithValue(ctx, "Id", item.UserId)
		}
		r = r.WithContext(ctx)

		if item.IsAuth && !item.IdErr {

			var err error = nil
			if item.GetErr {
				err = NoType.New("")
			}

			gomock.InOrder(
				mockListUsecase.EXPECT().GetSubList(uint(item.UserId), item.Start, item.Limit).Return(item.Pins, err),
			)
		}

		listDelivery.GetSubList(w, r)

		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		bodyStr := string(body)

		if item.GetErr || item.IdErr {
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

func TestHandler_GetRecommendationList(t *testing.T) {
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

	mockListUsecase := mock.NewMockIUsecase(ctl)
	listDelivery := NewHandler(mockListUsecase, zap)

	cases := []TestCaseFetch{
		TestCaseFetch{
			IsAuth:     false,
			UserId:     1,
			Start:		1,
			Limit:		15,
			Response:	`{"status":401,"body":{"error":"User is unauthorized"}}`,
		},
		TestCaseFetch{
			IsAuth:     true,
			IdErr:		true,
			UserId:     1,
			Start:		1,
			Limit:		15,
		},
		TestCaseFetch{
			IsAuth:     true,
			GetErr:     true,
			UserId:     1,
			Start:		1,
			Limit:		15,
			Pins:       nil,
		},
		TestCaseFetch{
			IsAuth:     true,
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
				`{"id":2,"user_id":1,"board_id":5,"name":"name2","description":"desc2","image":"image.jpg"}]}`,
		},
	}

	for caseNum, item := range cases {
		r := httptest.NewRequest("GET", fmt.Sprintf("/api/list/recommendation?start=%d&limit=%d", item.Start, item.Limit), strings.NewReader(""))

		w := httptest.NewRecorder()

		ctx := r.Context()
		ctx = context.WithValue(ctx, "IsAuth", item.IsAuth)
		if !item.IdErr {
			ctx = context.WithValue(ctx, "Id", item.UserId)
		}
		r = r.WithContext(ctx)

		if item.IsAuth && !item.IdErr {

			var err error = nil
			if item.GetErr {
				err = NoType.New("")
			}

			gomock.InOrder(
				mockListUsecase.EXPECT().GetRecommendationList(uint(item.UserId), item.Start, item.Limit).Return(item.Pins, err),
			)
		}

		listDelivery.GetRecommendationList(w, r)

		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		bodyStr := string(body)

		if item.GetErr || item.IdErr {
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

func TestHandler_GetMainList(t *testing.T) {
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

	mockListUsecase := mock.NewMockIUsecase(ctl)
	listDelivery := NewHandler(mockListUsecase, zap)

	cases := []TestCaseFetch{
		TestCaseFetch{
			GetErr:     true,
			UserId:     1,
			Start:		1,
			Limit:		15,
			Pins:       nil,
		},
		TestCaseFetch{
			IsAuth:     true,
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
				`{"id":2,"user_id":1,"board_id":5,"name":"name2","description":"desc2","image":"image.jpg"}]}`,
		},
	}

	for caseNum, item := range cases {
		r := httptest.NewRequest("GET", fmt.Sprintf("/api/list?start=%d&limit=%d", item.Start, item.Limit), strings.NewReader(""))

		w := httptest.NewRecorder()

		var err error = nil
		if item.GetErr {
			err = NoType.New("")
		}

		gomock.InOrder(
			mockListUsecase.EXPECT().GetMainList(item.Start, item.Limit).Return(item.Pins, err),
		)

		listDelivery.GetMainList(w, r)

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

