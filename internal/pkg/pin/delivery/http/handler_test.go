package http

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"

	//"github.com/gorilla/mux"
	"net/http"

	"2020_1_Color_noise/internal/models"
	. "2020_1_Color_noise/internal/pkg/error"
	"2020_1_Color_noise/internal/pkg/pin/mock"
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
	Pin       *models.Pin
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

	mockPinUsecase := mock.NewMockIUsecase(ctl)
	pinDelivery := NewHandler(mockPinUsecase, zap)

	cases := []TestCaseCreate{
		TestCaseCreate{
			Pin:      &models.Pin{},
			IsAuth:   false,
			Response: `{"status":401,"body":{"error":"User is unauthorized"}}`,
		},
		TestCaseCreate{
			IsAuth:   true,
			IdErr:    true,
			Pin:      &models.Pin{},
			Response: `{"status":500,"body":{"error":"Internal server error"}}`,
		},
		TestCaseCreate{
			IsAuth:   true,
			InputErr: true,
			Pin:      &models.Pin{},
			Response: `{"status":400,"body":{"error":"Wrong body of request"}}`,
		},
		TestCaseCreate{
			IsAuth:   true,
			ValidErr: true,
			Pin: &models.Pin{
				Name:        "",
				Description: "desc",
				Image: "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAoAAAAKCAYAAACNMs+9AAAABmJLR0QA/wD/AP+gvaeTAAAAB3RJ" +
					"TUUH1ggDCwMADQ4NnwAAAFVJREFUGJWNkMEJADEIBEcbSDkXUnfSg" +
					"nBVeZ8LSAjiwjyEQXSFEIcHGP9oAi+H0Bymgx9MhxbFdZE2a0s9kT" +
					"Zdw01ZhhYkABSwgmf1Z6r1SNyfFf4BZ+ZUExcNUQUAAAAASUVORK5CYII=",
			},
			Response: `{"status":400,"body":{"error":"Name shouldn't be empty and longer 60 characters. ` +
				`Description shouldn't be empty and longer 1000 characters. ` +
				`Image should be base64"}}`,
		},
		TestCaseCreate{
			IsAuth:   true,
			ValidErr: true,
			Pin: &models.Pin{
				Name:        "nameddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddd",
				Description: "desc",
				Image: "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAoAAAAKCAYAAACNMs+9AAAABmJLR0QA/wD/AP+gvaeTAAAAB3RJ" +
					"TUUH1ggDCwMADQ4NnwAAAFVJREFUGJWNkMEJADEIBEcbSDkXUnfSg" +
					"nBVeZ8LSAjiwjyEQXSFEIcHGP9oAi+H0Bymgx9MhxbFdZE2a0s9kT" +
					"Zdw01ZhhYkABSwgmf1Z6r1SNyfFf4BZ+ZUExcNUQUAAAAASUVORK5CYII=",
			},
			Response: `{"status":400,"body":{"error":"Name shouldn't be empty and longer 60 characters. ` +
				`Description shouldn't be empty and longer 1000 characters. ` +
				`Image should be base64"}}`,
		},
		/*TestCaseCreate{
			IsAuth:     true,
			ValidErr:   true,
			Pin:		&models.Pin{
				Name:	     "name",
				Description: "",
				Image:		 "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAoAAAAKCAYAAACNMs+9AAAABmJLR0QA/wD/AP+gvaeTAAAAB3RJ" +
					"TUUH1ggDCwMADQ4NnwAAAFVJREFUGJWNkMEJADEIBEcbSDkXUnfSg" +
					"nBVeZ8LSAjiwjyEQXSFEIcHGP9oAi+H0Bymgx9MhxbFdZE2a0s9kT" +
					"Zdw01ZhhYkABSwgmf1Z6r1SNyfFf4BZ+ZUExcNUQUAAAAASUVORK5CYII=",
			},
			Response:	`{"status":400,"body":{"error":"Name shouldn't be empty and longer 60 characters. `+
				`Description shouldn't be empty and longer 1000 characters. `+
				`Image should be base64"}}`,
		},*/
		TestCaseCreate{
			IsAuth:   true,
			ValidErr: true,
			Pin: &models.Pin{
				Name:        "name",
				Description: "desc",
				Image: "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAoAAAAKCAYAAACNMs+9AAAABmJLR0QA/wD/AP+gvaeTAAAAB3RJ" +
					"TUUH1ggDCwMADQ4NnwAAAFVJREFUGJWNkMEJADEIBEcbSDkXUnfSg" +
					"nBVeZ8LSAjiwjyEQXSFEIcHGP9oAi+H0Bymgx9MhxbFdZE2a0s9kT" +
					"Zdw01ZhhYkABSwgmf1Z6r1SNyfFf4BZ+ZUExcNUQUAAAAASUVORK5CYII",
			},
			Response: `{"status":400,"body":{"error":"Name shouldn't be empty and longer 60 characters. ` +
				`Description shouldn't be empty and longer 1000 characters. ` +
				`Image should be base64"}}`,
		},
		TestCaseCreate{
			IsAuth:    true,
			CreateErr: true,
			Pin: &models.Pin{
				Id:          1,
				Name:        "name",
				Description: "desc",
				Image: "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAoAAAAKCAYAAACNMs+9AAAABmJLR0QA/wD/AP+gvaeTAAAAB3RJ" +
					"TUUH1ggDCwMADQ4NnwAAAFVJREFUGJWNkMEJADEIBEcbSDkXUnfSg" +
					"nBVeZ8LSAjiwjyEQXSFEIcHGP9oAi+H0Bymgx9MhxbFdZE2a0s9kT" +
					"Zdw01ZhhYkABSwgmf1Z6r1SNyfFf4BZ+ZUExcNUQUAAAAASUVORK5CYII=",
			},
		},
		TestCaseCreate{
			IsAuth: true,
			Pin: &models.Pin{
				Id:          1,
				Name:        "name",
				Description: "desc",
				Image: "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAoAAAAKCAYAAACNMs+9AAAABmJLR0QA/wD/AP+gvaeTAAAAB3RJ" +
					"TUUH1ggDCwMADQ4NnwAAAFVJREFUGJWNkMEJADEIBEcbSDkXUnfSg" +
					"nBVeZ8LSAjiwjyEQXSFEIcHGP9oAi+H0Bymgx9MhxbFdZE2a0s9kT" +
					"Zdw01ZhhYkABSwgmf1Z6r1SNyfFf4BZ+ZUExcNUQUAAAAASUVORK5CYII=",
			},
			Response: `{"status":201,"body":{"id":1}}`,
		},
	}

	for caseNum, item := range cases {
		var r *http.Request
		if item.InputErr == false {
			r = httptest.NewRequest("POST", "/api/pin",
				strings.NewReader(fmt.Sprintf(`{"name":"%s", "description":"%s", "image":"%s"}`, item.Pin.Name, item.Pin.Description, item.Pin.Image)))
		} else {
			r = httptest.NewRequest("POST", "/api/pin",
				strings.NewReader(fmt.Sprintf(`{"name:"%s", "description":"%s", "image":"%s"}`, item.Pin.Name, item.Pin.Description, item.Pin.Image)))
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

			input := &models.InputPin{
				Name:        item.Pin.Name,
				Description: item.Pin.Description,
				Image:       item.Pin.Image,
			}

			var err error = nil
			if item.CreateErr {
				err = NoType.New("")
			}

			gomock.InOrder(
				mockPinUsecase.EXPECT().Create(input, item.UserId).Return(uint(item.Pin.Id), err),
			)
		}

		pinDelivery.Create(w, r)

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

type TestCaseGetPin struct {
	IsAuth   bool
	UserId   uint
	Pin      *models.Pin
	Response string
	IdErr    bool
	GetErr   bool
}

func TestHandler_GetPin(t *testing.T) {
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

	mockPinUsecase := mock.NewMockIUsecase(ctl)
	pinDelivery := NewHandler(mockPinUsecase, zap)

	cases := []TestCaseGetPin{
		TestCaseGetPin{
			IsAuth:   false,
			Pin:      &models.Pin{},
			Response: `{"status":401,"body":{"error":"User is unauthorized"}}`,
		},
		TestCaseGetPin{
			IsAuth:   true,
			IdErr:    true,
			Pin:      &models.Pin{},
			Response: `{"status":400,"body":{"error":"Bad id"}}`,
		},
		TestCaseGetPin{
			IsAuth: true,
			GetErr: true,
			Pin:    &models.Pin{},
		},
		TestCaseGetPin{
			IsAuth: true,
			Pin: &models.Pin{
				Id:          1,
				BoardId:     2,
				UserId:      3,
				Name:        "name",
				Description: "desc",
				Image:       "image.jpg",
			},
			Response: `{"status":200,"body":{"id":1,"user_id":3,"board_id":2,"name":"name","description":"desc","image":"image.jpg"}}`,
		},
	}

	for caseNum, item := range cases {
		r := httptest.NewRequest("GET", "/api/pin/", strings.NewReader(""))
		if !item.IdErr {
			r = mux.SetURLVars(r, map[string]string{"id": fmt.Sprintf("%d", item.Pin.Id)})
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
				mockPinUsecase.EXPECT().GetById(item.Pin.Id).Return(item.Pin, err),
			)
		}

		pinDelivery.GetPin(w, r)

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
	UserId   uint
	Pins     []*models.Pin
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

	mockPinUsecase := mock.NewMockIUsecase(ctl)
	pinDelivery := NewHandler(mockPinUsecase, zap)

	cases := []TestCaseFetch{
		TestCaseFetch{
			IsAuth:   false,
			UserId:   1,
			Start:    1,
			Limit:    15,
			Response: `{"status":401,"body":{"error":"User is unauthorized"}}`,
		},
		TestCaseFetch{
			IsAuth:   true,
			IdErr:    true,
			UserId:   1,
			Start:    1,
			Limit:    15,
			Response: `{"status":400,"body":{"error":"Bad id"}}`,
		},
		TestCaseFetch{
			IsAuth: true,
			GetErr: true,
			UserId: 1,
			Start:  1,
			Limit:  15,
			Pins:   nil,
		},
		TestCaseFetch{
			IsAuth: true,
			UserId: 1,
			Start:  1,
			Limit:  15,
			Pins: []*models.Pin{
				&models.Pin{
					Id:          1,
					BoardId:     2,
					UserId:      1,
					Name:        "name1",
					Description: "desc1",
					Image:       "image.jpg",
				},
				&models.Pin{
					Id:          2,
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
		r := httptest.NewRequest("GET", fmt.Sprintf("/api/pin/user/%d/?start=%d&limit=%d", item.UserId, item.Start, item.Limit), strings.NewReader(""))
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
				mockPinUsecase.EXPECT().GetByUserId(uint(item.UserId), item.Start, item.Limit).Return(item.Pins, err),
			)
		}

		pinDelivery.Fetch(w, r)

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
	IsAuth    bool
	UserId    uint
	Pin       *models.Pin
	Response  string
	UserIdErr bool
	IdErr     bool
	InputErr  bool
	ValidErr  bool
	UpdateErr bool
}

func TestHandler_Update(t *testing.T) {
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

	mockPinUsecase := mock.NewMockIUsecase(ctl)
	pinDelivery := NewHandler(mockPinUsecase, zap)

	cases := []TestCaseUpdate{
		TestCaseUpdate{
			Pin:      &models.Pin{},
			IsAuth:   false,
			Response: `{"status":401,"body":{"error":"User is unauthorized"}}`,
		},
		TestCaseUpdate{
			IsAuth:    true,
			UserIdErr: true,
			Pin:       &models.Pin{},
		},
		TestCaseUpdate{
			IsAuth:   true,
			IdErr:    true,
			UserId:   1,
			Pin:      &models.Pin{},
			Response: `{"status":400,"body":{"error":"Bad id"}}`,
		},
		TestCaseUpdate{
			IsAuth:   true,
			InputErr: true,
			Pin:      &models.Pin{},
			Response: `{"status":400,"body":{"error":"Wrong body of request"}}`,
		},
		TestCaseUpdate{
			IsAuth:   true,
			ValidErr: true,
			Pin: &models.Pin{
				Name:        "",
				Description: "desc",
			},
			Response: `{"status":400,"body":{"error":"Name shouldn't be empty and longer 60 characters. ` +
				`Description shouldn't be empty and longer 1000 characters."}}`,
		},
		TestCaseUpdate{
			IsAuth:   true,
			ValidErr: true,
			Pin: &models.Pin{
				Name:        "nameddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddd",
				Description: "desc",
			},
			Response: `{"status":400,"body":{"error":"Name shouldn't be empty and longer 60 characters. ` +
				`Description shouldn't be empty and longer 1000 characters."}}`,
		},
		/*TestCaseUpdate{
			IsAuth:     true,
			ValidErr:   true,
			Pin:		&models.Pin{
				Name:	     "name",
				Description: "",
			},
			Response:	`{"status":400,"body":{"error":"Name shouldn't be empty and longer 60 characters. `+
				`Description shouldn't be empty and longer 1000 characters."}}`,
		},*/
		TestCaseUpdate{
			IsAuth:    true,
			UpdateErr: true,
			Pin: &models.Pin{
				Id:          1,
				Name:        "name",
				Description: "desc",
			},
		},
		TestCaseUpdate{
			IsAuth: true,
			Pin: &models.Pin{
				Id:          1,
				Name:        "name",
				Description: "desc",
			},
			Response: `{"status":200,"body":{"message":"Ok"}}`,
		},
	}

	for caseNum, item := range cases {
		var r *http.Request
		if item.InputErr == false {
			r = httptest.NewRequest("PUT", "/api/pin",
				strings.NewReader(fmt.Sprintf(`{"name":"%s", "description":"%s"}`, item.Pin.Name, item.Pin.Description)))
		} else {
			r = httptest.NewRequest("PUT", "/api/pin",
				strings.NewReader(fmt.Sprintf(`{"name:"%s", "description":"%s"}`, item.Pin.Name, item.Pin.Description)))
		}

		if !item.IdErr {
			r = mux.SetURLVars(r, map[string]string{"id": fmt.Sprintf("%d", item.Pin.Id)})
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

			input := &models.UpdatePin{
				Name:        item.Pin.Name,
				Description: item.Pin.Description,
			}

			var err error = nil
			if item.UpdateErr {
				err = NoType.New("")
			}

			gomock.InOrder(
				mockPinUsecase.EXPECT().Update(input, item.Pin.Id, item.UserId).Return(err),
			)
		}

		pinDelivery.Update(w, r)

		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		bodyStr := string(body)

		if item.UpdateErr || item.UserIdErr {
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
	IsAuth    bool
	UserId    uint
	PinId     uint
	Response  string
	UserIdErr bool
	IdErr     bool
	DeleteErr bool
}

func TestHandler_DeletePin(t *testing.T) {
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

	mockPinUsecase := mock.NewMockIUsecase(ctl)
	pinDelivery := NewHandler(mockPinUsecase, zap)

	cases := []TestCaseDelete{
		TestCaseDelete{
			IsAuth:   false,
			Response: `{"status":401,"body":{"error":"User is unauthorized"}}`,
		},
		TestCaseDelete{
			IsAuth:    true,
			UserIdErr: true,
			PinId:     1,
			UserId:    1,
		},
		TestCaseDelete{
			IsAuth:   true,
			IdErr:    true,
			PinId:    1,
			UserId:   1,
			Response: `{"status":400,"body":{"error":"Bad id"}}`,
		},
		TestCaseDelete{
			IsAuth:    true,
			DeleteErr: true,
			PinId:     1,
			UserId:    1,
		},
		TestCaseDelete{
			IsAuth:   true,
			PinId:    1,
			UserId:   1,
			Response: `{"status":200,"body":{"message":"Ok"}}`,
		},
	}

	for caseNum, item := range cases {
		r := httptest.NewRequest("DELETE", fmt.Sprintf("/api/pin/%d", item.PinId), strings.NewReader(""))

		if !item.IdErr {
			r = mux.SetURLVars(r, map[string]string{"id": fmt.Sprintf("%d", item.PinId)})
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
				mockPinUsecase.EXPECT().Delete(item.PinId, item.UserId).Return(err),
			)
		}

		pinDelivery.DeletePin(w, r)

		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		bodyStr := string(body)

		if item.DeleteErr || item.UserIdErr {
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
