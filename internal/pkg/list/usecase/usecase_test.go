package usecase

import (
	"2020_1_Color_noise/internal/models"
	. "2020_1_Color_noise/internal/pkg/error"
	"2020_1_Color_noise/internal/pkg/list/mock"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)


type TestCaseFetch struct {
	ErrFunc   error
	PinFunc   []*models.Pin
	PinExp    []*models.Pin
	UserId	  uint
	Start     int
	Limit 	  int
}

func TestHandler_GetSubList(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockListRepository := mock.NewMockIRepository(ctl)
	listUsecase := NewUsecase(mockListRepository)

	cases := []TestCaseFetch{
		TestCaseFetch{
			UserId:    1,
			ErrFunc:   NoType.New("error"),
			PinFunc:  nil,
			PinExp:   nil,
			Start:    5,
			Limit:    15,
		},
		TestCaseFetch{
			UserId:   1,
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
			mockListRepository.EXPECT().GetSubList(item.UserId, item.Start, item.Limit).Return(item.PinFunc, item.ErrFunc),
		)

		pins, err := listUsecase.GetSubList(item.UserId, item.Start, item.Limit)
		if !reflect.DeepEqual(pins, item.PinExp) {
			t.Errorf("[%d] wrong User: got %+v, expected %+v",
				caseNum, pins, item.PinExp)
		}

		if item.ErrFunc == nil && err != nil {
			t.Errorf("[%d] wrong User: got %+v, expected %+v",
				caseNum, err, item.ErrFunc)
		}
		if item.ErrFunc != nil && err == nil {
			t.Errorf("[%d] wrong User: got %+v, expected %+v",
				caseNum, err, item.ErrFunc)
		}
	}
}


func TestHandler_GetRecommendationList(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockListRepository := mock.NewMockIRepository(ctl)
	listUsecase := NewUsecase(mockListRepository)

	cases := []TestCaseFetch{
		TestCaseFetch{
			UserId:    1,
			ErrFunc:   NoType.New("error"),
			PinFunc:  nil,
			PinExp:   nil,
			Start:    5,
			Limit:    15,
		},
		TestCaseFetch{
			UserId:   1,
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
			mockListRepository.EXPECT().GetRecommendationList(item.UserId, item.Start, item.Limit).Return(item.PinFunc, item.ErrFunc),
		)

		pins, err := listUsecase.GetRecommendationList(item.UserId, item.Start, item.Limit)
		if !reflect.DeepEqual(pins, item.PinExp) {
			t.Errorf("[%d] wrong User: got %+v, expected %+v",
				caseNum, pins, item.PinExp)
		}

		if item.ErrFunc == nil && err != nil {
			t.Errorf("[%d] wrong User: got %+v, expected %+v",
				caseNum, err, item.ErrFunc)
		}
		if item.ErrFunc != nil && err == nil {
			t.Errorf("[%d] wrong User: got %+v, expected %+v",
				caseNum, err, item.ErrFunc)
		}
	}
}

func TestHandler_GetMainList(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockListRepository := mock.NewMockIRepository(ctl)
	listUsecase := NewUsecase(mockListRepository)

	cases := []TestCaseFetch{
		TestCaseFetch{
			UserId:    1,
			ErrFunc:   NoType.New("error"),
			PinFunc:  nil,
			PinExp:   nil,
			Start:    5,
			Limit:    15,
		},
		TestCaseFetch{
			UserId:   1,
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
			mockListRepository.EXPECT().GetMainList(item.Start, item.Limit).Return(item.PinFunc, item.ErrFunc),
		)

		pins, err := listUsecase.GetMainList(item.Start, item.Limit)
		if !reflect.DeepEqual(pins, item.PinExp) {
			t.Errorf("[%d] wrong Pins: got %+v, expected %+v",
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