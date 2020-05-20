package usecase

import (
	_ "2020_1_Color_noise/internal/models"
	_ "github.com/golang/mock/gomock"
	_ "testing"
)

/*
import (
	"2020_1_Color_noise/internal/models"
	. "2020_1_Color_noise/internal/pkg/error"
	"2020_1_Color_noise/internal/pkg/pin/mock"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

type TestCaseGet struct {
	ErrFunc   error
	PinFunc   *models.Pin
	PinExp    *models.Pin
	PinId     uint
}

func TestHandler_GetById(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockPinRepository := mock.NewMockIRepository(ctl)
	pinUsecase := NewUsecase(mockPinRepository)

	cases := []TestCaseGet{
		TestCaseGet{
			PinId:    1,
			ErrFunc:  NoType.New("error"),
			PinFunc:  nil,
			PinExp:   nil,
		},
		TestCaseGet{
			PinId:    1,
			ErrFunc:  nil,
			PinFunc:  &models.Pin{
				Id: 1,
			},
			PinExp:   &models.Pin{
				Id: 1,
			},
		},
	}

	for caseNum, item := range cases {

		gomock.InOrder(
			mockPinRepository.EXPECT().GetByID(item.PinId).Return(item.PinFunc, item.ErrFunc),
		)

		pin, err := pinUsecase.GetById(item.PinId)
		if !reflect.DeepEqual(pin, item.PinExp) {
			t.Errorf("[%d] wrong Pin: got %+v, expected %+v",
				caseNum, pin, item.PinExp)
		}

		if item.ErrFunc == nil && err != nil {
			t.Errorf("[%d] wrong Error: got %+v, expected %+v",
				caseNum, err, item.ErrFunc)
		}
		if item.ErrFunc != nil && err == nil {
			t.Errorf("[%d] wrong Error: got %+v, expected %+v",
				caseNum, err, item.ErrFunc)
		}
	}
}

type TestCaseFetch struct {
	ErrFunc   error
	PinFunc   []*models.Pin
	PinExp    []*models.Pin
	UserId	  uint
	Name	  string
	Start     int
	Limit 	  int
}

func TestHandler_GetByName(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockPinRepository := mock.NewMockIRepository(ctl)
	pinUsecase := NewUsecase(mockPinRepository)

	cases := []TestCaseFetch{
		TestCaseFetch{
			UserId:    1,
			ErrFunc:   NoType.New("error"),
			PinFunc:  nil,
			Name:	  "pin",
			PinExp:   nil,
			Start:    5,
			Limit:    15,
		},
		TestCaseFetch{
			UserId:   1,
			Name:	  "pin",
			ErrFunc:  nil,
			PinFunc:  []*models.Pin{
				&models.Pin{
					Id:    1,
				},
				&models.Pin{
					Id:    2,
				},
			},
			PinExp:  []*models.Pin{
				&models.Pin{
					Id:    1,
				},
				&models.Pin{
					Id:    2,
				},
			},
			Start:    5,
			Limit:    15,
		},
	}

	for caseNum, item := range cases {

		gomock.InOrder(
			mockPinRepository.EXPECT().GetByName(item.Name, item.Start, item.Limit).Return(item.PinFunc, item.ErrFunc),
		)

		pins, err := pinUsecase.GetByName(item.Name, item.Start, item.Limit)
		if !reflect.DeepEqual(pins, item.PinExp) {
			t.Errorf("[%d] wrong Pin: got %+v, expected %+v",
				caseNum, pins, item.PinExp)
		}

		if item.ErrFunc == nil && err != nil {
			t.Errorf("[%d] wrong Error: got %+v, expected %+v",
				caseNum, err, item.ErrFunc)
		}
		if item.ErrFunc != nil && err == nil {
			t.Errorf("[%d] wrong Error: got %+v, expected %+v",
				caseNum, err, item.ErrFunc)
		}
	}
}

func TestHandler_GetByUserId(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockPinRepository := mock.NewMockIRepository(ctl)
	pinUsecase := NewUsecase(mockPinRepository)

	cases := []TestCaseFetch{
		TestCaseFetch{
			UserId:    1,
			ErrFunc:   NoType.New("error"),
			PinFunc:  nil,
			Name:	  "pin",
			PinExp:   nil,
			Start:    5,
			Limit:    15,
		},
		TestCaseFetch{
			UserId:   1,
			Name:	  "pin",
			ErrFunc:  nil,
			PinFunc:  []*models.Pin{
				&models.Pin{
					Id:    1,
				},
				&models.Pin{
					Id:    2,
				},
			},
			PinExp:  []*models.Pin{
				&models.Pin{
					Id:    1,
				},
				&models.Pin{
					Id:    2,
				},
			},
			Start:    5,
			Limit:    15,
		},
	}

	for caseNum, item := range cases {

		gomock.InOrder(
			mockPinRepository.EXPECT().GetByUserID(item.UserId, item.Start, item.Limit).Return(item.PinFunc, item.ErrFunc),
		)

		pins, err := pinUsecase.GetByUserId(item.UserId, item.Start, item.Limit)
		if !reflect.DeepEqual(pins, item.PinExp) {
			t.Errorf("[%d] wrong Pin: got %+v, expected %+v",
				caseNum, pins, item.PinExp)
		}

		if item.ErrFunc == nil && err != nil {
			t.Errorf("[%d] wrong Error: got %+v, expected %+v",
				caseNum, err, item.ErrFunc)
		}
		if item.ErrFunc != nil && err == nil {
			t.Errorf("[%d] wrong Error: got %+v, expected %+v",
				caseNum, err, item.ErrFunc)
		}
	}
}

type TestCaseUpdate struct {
	ErrFunc  error
	Pin      *models.UpdatePin
	PinId    uint
	UserId   uint
}

func TestHandler_Update(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockPinRepository := mock.NewMockIRepository(ctl)
	pinUsecase := NewUsecase(mockPinRepository)

	cases := []TestCaseUpdate{
		TestCaseUpdate{
			PinId:    1,
			UserId:   2,
			Pin: 	   &models.UpdatePin{
				Name:        "name",
				Description: "desc",
			},
			ErrFunc:   NoType.New("error"),
		},
		TestCaseUpdate{
			PinId:    1,
			UserId:   2,
			Pin: 	   &models.UpdatePin{
				Name:        "name",
				Description: "desc",
			},
			ErrFunc:   nil,
		},
	}

	for caseNum, item := range cases {

		pin := &models.Pin{
			Id:          item.PinId,
			UserId:      item.UserId,
			Name:        item.Pin.Name,
			Description: item.Pin.Description,
		}

		gomock.InOrder(
			mockPinRepository.EXPECT().Update(pin).Return(item.ErrFunc),
		)

		err := pinUsecase.Update(item.Pin, item.PinId, item.UserId)
		if item.ErrFunc == nil && err != nil {
			t.Errorf("[%d] wrong Error: got %+v, expected %+v",
				caseNum, err, item.ErrFunc)
		}
		if item.ErrFunc != nil && err == nil {
			t.Errorf("[%d] wrong Error: got %+v, expected %+v",
				caseNum, err, item.ErrFunc)
		}
	}
}

type TestCaseDelete struct {
	ErrFunc   error
	PinId     uint
	UserId    uint
}

func TestHandler_Delete(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockPinRepository := mock.NewMockIRepository(ctl)
	pinUsecase := NewUsecase(mockPinRepository)

	cases := []TestCaseDelete{
		TestCaseDelete{
			UserId:    1,
			PinId:     2,
			ErrFunc:   NoType.New("error"),
		},
		TestCaseDelete{
			UserId:    1,
			PinId:     2,
			ErrFunc:   nil,
		},
	}

	for caseNum, item := range cases {

		gomock.InOrder(
			mockPinRepository.EXPECT().Delete(item.PinId, item.UserId).Return(item.ErrFunc),
		)

		err := pinUsecase.Delete(item.PinId, item.UserId)
		if item.ErrFunc == nil && err != nil {
			t.Errorf("[%d] wrong Error: got %+v, expected %+v",
				caseNum, err, item.ErrFunc)
		}
		if item.ErrFunc != nil && err == nil {
			t.Errorf("[%d] wrong Error: got %+v, expected %+v",
				caseNum, err, item.ErrFunc)
		}
	}
}

*/
