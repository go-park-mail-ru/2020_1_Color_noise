package usecase

import (
	"2020_1_Color_noise/internal/models"
	. "2020_1_Color_noise/internal/pkg/error"
	"2020_1_Color_noise/internal/pkg/notifications/mock"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

type TestCaseGet struct {
	ErrFunc    error
	NotFunc    []*models.Notification
	NotExp     []*models.Notification
	UserId     uint
	Start      int
	Limit      int
}

func TestHandler_GetByName(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockNotRepository := mock.NewMockIRepository(ctl)
	notUsecase := NewUsecase(mockNotRepository)

	cases := []TestCaseGet{
		TestCaseGet{
			UserId:    1,
			ErrFunc:   NoType.New("error"),
			NotFunc:  nil,
			NotExp:   nil,
			Start:    5,
			Limit:    15,
		},
		TestCaseGet{
			UserId:   1,
			ErrFunc:  nil,
			NotFunc:  []*models.Notification{
				&models.Notification{
					User: models.User{
						Id: 3,
					},
					Message: "comment",
				},
				&models.Notification{
					User: models.User{
						Id: 2,
					},
					Message: "comment",
				},
			},
			NotExp:  []*models.Notification{
				&models.Notification{
					User: models.User{
						Id: 3,
					},
					Message: "comment",
				},
				&models.Notification{
					User: models.User{
						Id: 2,
					},
					Message: "comment",
				},
			},
			Start:    5,
			Limit:    15,
		},
	}

	for caseNum, item := range cases {

		gomock.InOrder(
			mockNotRepository.EXPECT().GetNotifications(item.UserId, item.Start, item.Limit).Return(item.NotFunc, item.ErrFunc),
		)

		pins, err := notUsecase.GetNotifications(item.UserId, item.Start, item.Limit)
		if !reflect.DeepEqual(pins, item.NotExp) {
			t.Errorf("[%d] wrong Pins: got %+v, expected %+v",
				caseNum, pins, item.NotExp)
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
