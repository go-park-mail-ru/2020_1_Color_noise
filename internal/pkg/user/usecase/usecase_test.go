package usecase

import (
	"2020_1_Color_noise/internal/models"
	. "2020_1_Color_noise/internal/pkg/error"
	"2020_1_Color_noise/internal/pkg/user/mock"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

type TestCaseGet struct {
	ErrFunc   error
	UserFunc  *models.User
	UserExp   *models.User
	UserId    uint
	Login	  string
}

func TestHandler_GetById(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockUserRepository := mock.NewMockIRepository(ctl)
	userUsecase := NewUsecase(mockUserRepository)

	cases := []TestCaseGet{
		TestCaseGet{
			UserId:    1,
			ErrFunc:   NoType.New("error"),
			UserFunc:  nil,
			UserExp:   nil,
		},
		TestCaseGet{
			UserId:    1,
			ErrFunc:   nil,
			UserFunc:  &models.User{
				Id: 1,
			},
			UserExp:   &models.User{
				Id: 1,
			},
		},
	}

	for caseNum, item := range cases {

		gomock.InOrder(
			mockUserRepository.EXPECT().GetByID(item.UserId).Return(item.UserFunc, item.ErrFunc),
		)
		user, err := userUsecase.GetById(item.UserId)
		if !reflect.DeepEqual(user, item.UserExp) {
			t.Errorf("[%d] wrong User: got %+v, expected %+v",
				caseNum, user, item.UserExp)
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

func TestHandler_GetByLogin(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockUserRepository := mock.NewMockIRepository(ctl)
	userUsecase := NewUsecase(mockUserRepository)

	cases := []TestCaseGet{
		TestCaseGet{
			UserId:    1,
			Login: 	   "login1",
			ErrFunc:   NoType.New("error"),
			UserFunc:  nil,
			UserExp:   nil,
		},
		TestCaseGet{
			UserId:    1,
			Login: 	   "login1",
			ErrFunc:   nil,
			UserFunc:  &models.User{
				Id: 1,
				Login: "login1",
			},
			UserExp:   &models.User{
				Id: 1,
				Login: "login1",
			},
		},
	}

	for caseNum, item := range cases {

		gomock.InOrder(
			mockUserRepository.EXPECT().GetByLogin(item.Login).Return(item.UserFunc, item.ErrFunc),
		)

		user, err := userUsecase.GetByLogin(item.Login)
		if !reflect.DeepEqual(user, item.UserExp) {
			t.Errorf("[%d] wrong User: got %+v, expected %+v",
				caseNum, user, item.UserExp)
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

type TestCaseSearch struct {
	ErrFunc   error
	UserFunc  []*models.User
	UserExp   []*models.User
	UserId    uint
	Login	  string
	Start	  int
	Limit     int
}

func TestHandler_Search(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockUserRepository := mock.NewMockIRepository(ctl)
	userUsecase := NewUsecase(mockUserRepository)

	cases := []TestCaseSearch{
		TestCaseSearch{
			UserId:    1,
			Start:     5,
			Limit:     15,
			Login: 	   "login1",
			ErrFunc:   NoType.New("error"),
			UserFunc:  nil,
			UserExp:   nil,
		},
		TestCaseSearch{
			UserId:    1,
			Start:     5,
			Limit:     15,
			Login: 	   "login1",
			ErrFunc:   nil,
			UserFunc:  []*models.User{
				&models.User{
				Id:    1,
				Login: "login1",
				},
				&models.User{
					Id:    2,
					Login: "login2",
				},
			},
			UserExp:   []*models.User{
				&models.User{
				Id:    1,
				Login: "login1",
				},
				&models.User{
					Id:    2,
					Login: "login2",
				},
			},
		},
	}

	for caseNum, item := range cases {

		gomock.InOrder(
			mockUserRepository.EXPECT().Search(item.Login, item.Start, item.Limit).Return(item.UserFunc, item.ErrFunc),
		)

		user, err := userUsecase.Search(item.Login, item.Start, item.Limit)
		if !reflect.DeepEqual(user, item.UserExp) {
			t.Errorf("[%d] wrong Users: got %+v, expected %+v",
				caseNum, user, item.UserExp)
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
	ErrValid  error
	ErrFunc   error
	User      *models.UpdateProfileInput
	UserId    uint
}

func TestHandler_UpdateProfile(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockUserRepository := mock.NewMockIRepository(ctl)
	userUsecase := NewUsecase(mockUserRepository)

	cases := []TestCaseUpdate{
		TestCaseUpdate{
			ErrValid:  nil,
			UserId:    1,
			User: 	   &models.UpdateProfileInput{
				Login: "login1",
				Email: "email@ad.com",
			},
			ErrFunc:   NoType.New("error"),
		},
		TestCaseUpdate{
			ErrValid:  nil,
			UserId:    1,
			User: 	   &models.UpdateProfileInput{
				Login: "login1",
				Email: "email@ad.com",
			},
			ErrFunc:   nil,
		},
	}

	for caseNum, item := range cases {

			gomock.InOrder(
				mockUserRepository.EXPECT().UpdateProfile(item.UserId, item.User.Email, item.User.Login).Return(item.ErrFunc),
			)

		err := userUsecase.UpdateProfile(item.UserId, item.User)
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

type TestCaseDescription struct {
	ErrFunc   error
	Desc      *models.UpdateDescriptionInput
	UserId    uint
}

func TestHandler_UpdateDescription(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockUserRepository := mock.NewMockIRepository(ctl)
	userUsecase := NewUsecase(mockUserRepository)

	cases := []TestCaseDescription{
		TestCaseDescription{
			UserId:    1,
			Desc: 	   &models.UpdateDescriptionInput{
				Description: "desc1",
			},
			ErrFunc:   NoType.New("error"),
		},
		TestCaseDescription{
			UserId:    1,
			Desc: 	   &models.UpdateDescriptionInput{
				Description: "desc1",
			},
			ErrFunc:   nil,
		},
	}

	for caseNum, item := range cases {

		gomock.InOrder(
			mockUserRepository.EXPECT().UpdateDescription(item.UserId, &item.Desc.Description).Return(item.ErrFunc),
		)

		err := userUsecase.UpdateDescription(item.UserId, item.Desc)
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

	mockUserRepository := mock.NewMockIRepository(ctl)
	userUsecase := NewUsecase(mockUserRepository)

	cases := []TestCaseGet{
		TestCaseGet{
			UserId:    1,
			ErrFunc:   NoType.New("error"),
		},
		TestCaseGet{
			UserId:    1,
			ErrFunc:   nil,
		},
	}

	for caseNum, item := range cases {

		gomock.InOrder(
			mockUserRepository.EXPECT().Delete(item.UserId).Return(item.ErrFunc),
		)

		err := userUsecase.Delete(item.UserId)
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

type TestCaseFollow struct {
	ErrFunc   error
	UserId    uint
	SubId	  uint
}

func TestHandler_Follow(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockUserRepository := mock.NewMockIRepository(ctl)
	userUsecase := NewUsecase(mockUserRepository)

	cases := []TestCaseFollow{
		TestCaseFollow{
			UserId:    1,
			SubId:	   2,
			ErrFunc:   NoType.New("error"),
		},
		TestCaseFollow{
			UserId:    1,
			SubId:	   2,
			ErrFunc:   nil,
		},
	}

	for caseNum, item := range cases {

		gomock.InOrder(
			mockUserRepository.EXPECT().Follow(item.UserId, item.SubId).Return(item.ErrFunc),
		)

		err := userUsecase.Follow(item.UserId, item.SubId)
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

func TestHandler_Unfollow(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockUserRepository := mock.NewMockIRepository(ctl)
	userUsecase := NewUsecase(mockUserRepository)

	cases := []TestCaseFollow{
		TestCaseFollow{
			UserId:    1,
			SubId:	   2,
			ErrFunc:   NoType.New("error"),
		},
		TestCaseFollow{
			UserId:    1,
			SubId:	   2,
			ErrFunc:   nil,
		},
	}

	for caseNum, item := range cases {

		gomock.InOrder(
			mockUserRepository.EXPECT().Unfollow(item.UserId, item.SubId).Return(item.ErrFunc),
		)

		err := userUsecase.Unfollow(item.UserId, item.SubId)
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

type TestCaseSub struct {
	ErrFunc   error
	UserFunc  []*models.User
	UserExp   []*models.User
	UserId    uint
	Start	  int
	Limit     int
}

func TestHandler_GetSubscribers(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockUserRepository := mock.NewMockIRepository(ctl)
	userUsecase := NewUsecase(mockUserRepository)

	cases := []TestCaseSub{
		TestCaseSub{
			UserId:    3,
			Start:     5,
			Limit:     15,
			ErrFunc:   NoType.New("error"),
			UserFunc:  nil,
			UserExp:   nil,
		},
		TestCaseSub{
			UserId:    3,
			Start:     5,
			Limit:     15,
			ErrFunc:   nil,
			UserFunc:  []*models.User{
				&models.User{
					Id:    1,
					Login: "login1",
				},
				&models.User{
					Id:    2,
					Login: "login2",
				},
			},
			UserExp:   []*models.User{
				&models.User{
					Id:    1,
					Login: "login1",
				},
				&models.User{
					Id:    2,
					Login: "login2",
				},
			},
		},
	}

	for caseNum, item := range cases {

		gomock.InOrder(
			mockUserRepository.EXPECT().GetSubscribers(item.UserId, item.Start, item.Limit).Return(item.UserFunc, item.ErrFunc),
		)

		users, err := userUsecase.GetSubscribers(item.UserId, item.Start, item.Limit)
		if !reflect.DeepEqual(users, item.UserExp) {
			t.Errorf("[%d] wrong Users: got %+v, expected %+v",
				caseNum, users, item.UserExp)
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

func TestHandler_GetSubscriptions(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockUserRepository := mock.NewMockIRepository(ctl)
	userUsecase := NewUsecase(mockUserRepository)

	cases := []TestCaseSub{
		TestCaseSub{
			UserId:    3,
			Start:     5,
			Limit:     15,
			ErrFunc:   NoType.New("error"),
			UserFunc:  nil,
			UserExp:   nil,
		},
		TestCaseSub{
			UserId:    3,
			Start:     5,
			Limit:     15,
			ErrFunc:   nil,
			UserFunc:  []*models.User{
				&models.User{
					Id:    1,
					Login: "login1",
				},
				&models.User{
					Id:    2,
					Login: "login2",
				},
			},
			UserExp:   []*models.User{
				&models.User{
					Id:    1,
					Login: "login1",
				},
				&models.User{
					Id:    2,
					Login: "login2",
				},
			},
		},
	}

	for caseNum, item := range cases {

		gomock.InOrder(
			mockUserRepository.EXPECT().GetSubscriptions(item.UserId, item.Start, item.Limit).Return(item.UserFunc, item.ErrFunc),
		)

		users, err := userUsecase.GetSubscriptions(item.UserId, item.Start, item.Limit)
		if !reflect.DeepEqual(users, item.UserExp) {
			t.Errorf("[%d] wrong Users: got %+v, expected %+v",
				caseNum, users, item.UserExp)
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

type TestCasePassword struct {
	Err  error
	Pass string
	User *models.User
}
/*
func TestHandler_ComparePassword(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockUserRepository := mock.NewMockIRepository(ctl)
	userUsecase := NewUsecase(mockUserRepository)

	cases := []TestCasePassword{
		TestCasePassword{
			Err:   BadPassword.Newf("Password is incorrect"),
			Pass:  "password",
			User:  &models.User{},
		},
		TestCasePassword{
			Err:   nil,
			Pass:  "password",
			User:  &models.User{},
		},
	}

	for caseNum, item := range cases {

		if item.Err != nil {
			item.User.EncryptedPassword, _ = encryptPassword("somepass")
		} else {
			item.User.EncryptedPassword, _ = encryptPassword(item.Pass)
		}

		err := .ComparePassword(item.User, item.Pass)
		if item.Err == nil && err != nil {
			t.Errorf("[%d] wrong Error: got %+v, expected %+v",
				caseNum, err, item.Err)
		}
		if item.Err != nil && err == nil {
			t.Errorf("[%d] wrong Error: got %+v, expected %+v",
				caseNum, err, item.Err)
		}
	}
}

 */

