package usecase

import (
	"2020_1_Color_noise/internal/models"
	. "2020_1_Color_noise/internal/pkg/error"
	"2020_1_Color_noise/internal/pkg/session/mock"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

type TestCaseGet struct {
	ErrFunc   error
	SessFunc  *models.Session
	SessExp   *models.Session
	Cookie	  string
}

func TestHandler_GetByCookie(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockSessionRepository := mock.NewMockIRepository(ctl)
	sessionUsecase := NewUsecase(mockSessionRepository)

	cases := []TestCaseGet{
		TestCaseGet{
			Cookie:   "cookie1",
			ErrFunc:  NoType.New("error"),
			SessFunc: nil,
			SessExp:  nil,
		},
		TestCaseGet{
			Cookie:  "cookie1",
			ErrFunc: nil,
			SessFunc: &models.Session{
				Id:     1,
				Cookie: "cookie1",
			},
			SessExp: &models.Session{
				Id:     1,
				Cookie: "cookie1",
			},
		},
	}

	for caseNum, item := range cases {

		gomock.InOrder(
			mockSessionRepository.EXPECT().GetByCookie(item.Cookie).Return(item.SessFunc, item.ErrFunc),
		)

		sess, err := sessionUsecase.GetByCookie(item.Cookie)
		if !reflect.DeepEqual(sess, item.SessExp) {
			t.Errorf("[%d] wrong Session: got %+v, expected %+v",
				caseNum, sess, item.SessExp)
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
	ErrFunc   error
	Sess      *models.Session
	Token     string
}

func TestHandler_UpdateToken(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockSessionRepository := mock.NewMockIRepository(ctl)
	sessionUsecase := NewUsecase(mockSessionRepository)

	cases := []TestCaseUpdate{
		TestCaseUpdate{
			Token:    "token1",
			Sess: 	   &models.Session{
				Cookie: "token",
				Token:  "cookie",
			},
			ErrFunc:   NoType.New("error"),
		},
		TestCaseUpdate{
			Token:    "token1",
			Sess: 	   &models.Session{
				Cookie: "token",
				Token:  "cookie",
			},
			ErrFunc:   nil,
		},
	}

	for caseNum, item := range cases {

		gomock.InOrder(
			mockSessionRepository.EXPECT().Update(item.Sess).Return(item.ErrFunc),
		)

		err := sessionUsecase.Update(item.Sess)
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

func TestHandler_Delete(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockSessionRepository := mock.NewMockIRepository(ctl)
	sessionUsecase := NewUsecase(mockSessionRepository)

	cases := []TestCaseGet{
		TestCaseGet{
			Cookie:    "cookie",
			ErrFunc:   NoType.New("error"),
		},
		TestCaseGet{
			Cookie:    "cookie",
			ErrFunc:   nil,
		},
	}

	for caseNum, item := range cases {

		gomock.InOrder(
			mockSessionRepository.EXPECT().Delete(item.Cookie).Return(item.ErrFunc),
		)

		err := sessionUsecase.Delete(item.Cookie)
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
