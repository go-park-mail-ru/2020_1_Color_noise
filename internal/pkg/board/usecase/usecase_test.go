package usecase

import (
	"2020_1_Color_noise/internal/models"
	"2020_1_Color_noise/internal/pkg/board/mock"
	. "2020_1_Color_noise/internal/pkg/error"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

type TestCaseCreate struct {
	ErrFunc error
	Board   *models.InputBoard
	BoardId uint
	UserId  uint
}

func TestHandler_Create(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockBoardRepository := mock.NewMockIRepository(ctl)
	boardUsecase := NewUsecase(mockBoardRepository)

	cases := []TestCaseCreate{
		TestCaseCreate{
			BoardId: 0,
			UserId:  2,
			Board: &models.InputBoard{
				Name:        "name",
				Description: "desc",
			},
			ErrFunc: NoType.New("error"),
		},
		TestCaseCreate{
			BoardId: 1,
			UserId:  2,
			Board: &models.InputBoard{
				Name:        "name",
				Description: "desc",
			},
			ErrFunc: nil,
		},
	}

	for caseNum, item := range cases {

		board := &models.Board{
			UserId:      item.UserId,
			Name:        item.Board.Name,
			Description: item.Board.Description,
		}

		gomock.InOrder(
			mockBoardRepository.EXPECT().Create(board).Return(item.BoardId, item.ErrFunc),
		)

		id, err := boardUsecase.Create(item.Board, item.UserId)
		if id != item.BoardId {
			t.Errorf("[%d] wrong Id: got %+v, expected %+v",
				caseNum, id, item.BoardId)
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

type TestCaseGet struct {
	ErrFunc   error
	BoardFunc *models.Board
	BoardExp  *models.Board
	BoardId   uint
}

func TestHandler_GetById(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockBoardRepository := mock.NewMockIRepository(ctl)
	boardUsecase := NewUsecase(mockBoardRepository)

	cases := []TestCaseGet{
		TestCaseGet{
			BoardId:   1,
			ErrFunc:   NoType.New("error"),
			BoardFunc: nil,
			BoardExp:  nil,
		},
		TestCaseGet{
			BoardId: 1,
			ErrFunc: nil,
			BoardFunc: &models.Board{
				Id: 1,
			},
			BoardExp: &models.Board{
				Id: 1,
			},
		},
	}

	for caseNum, item := range cases {

		gomock.InOrder(
			mockBoardRepository.EXPECT().GetByID(item.BoardId).Return(item.BoardFunc, item.ErrFunc),
		)

		board, err := boardUsecase.GetById(item.BoardId)
		if !reflect.DeepEqual(board, item.BoardExp) {
			t.Errorf("[%d] wrong Board: got %+v, expected %+v",
				caseNum, board, item.BoardExp)
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

func TestHandler_GetByNameId(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockBoardRepository := mock.NewMockIRepository(ctl)
	boardUsecase := NewUsecase(mockBoardRepository)

	cases := []TestCaseGet{
		TestCaseGet{
			BoardId:   1,
			ErrFunc:   NoType.New("error"),
			BoardFunc: nil,
			BoardExp:  nil,
		},
		TestCaseGet{
			BoardId: 1,
			ErrFunc: nil,
			BoardFunc: &models.Board{
				Id: 1,
			},
			BoardExp: &models.Board{
				Id: 1,
			},
		},
	}

	for caseNum, item := range cases {

		gomock.InOrder(
			mockBoardRepository.EXPECT().GetByNameID(item.BoardId).Return(item.BoardFunc, item.ErrFunc),
		)

		board, err := boardUsecase.GetByNameId(item.BoardId)
		if !reflect.DeepEqual(board, item.BoardExp) {
			t.Errorf("[%d] wrong Board: got %+v, expected %+v",
				caseNum, board, item.BoardExp)
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
	BoardFunc []*models.Board
	BoardExp  []*models.Board
	UserId    uint
	Name      string
	Start     int
	Limit     int
}

func TestHandler_GetByName(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockBoardRepository := mock.NewMockIRepository(ctl)
	boardUsecase := NewUsecase(mockBoardRepository)

	cases := []TestCaseFetch{
		TestCaseFetch{
			UserId:    1,
			ErrFunc:   NoType.New("error"),
			BoardFunc: nil,
			Name:      "pin",
			BoardExp:  nil,
			Start:     5,
			Limit:     15,
		},
		TestCaseFetch{
			UserId:  1,
			Name:    "pin",
			ErrFunc: nil,
			BoardFunc: []*models.Board{
				&models.Board{
					Id: 1,
				},
				&models.Board{
					Id: 2,
				},
			},
			BoardExp: []*models.Board{
				&models.Board{
					Id: 1,
				},
				&models.Board{
					Id: 2,
				},
			},
			Start: 5,
			Limit: 15,
		},
	}

	for caseNum, item := range cases {

		gomock.InOrder(
			mockBoardRepository.EXPECT().GetByName(item.Name, item.Start, item.Limit).Return(item.BoardFunc, item.ErrFunc),
		)

		boards, err := boardUsecase.GetByName(item.Name, item.Start, item.Limit)
		if !reflect.DeepEqual(boards, item.BoardExp) {
			t.Errorf("[%d] wrong Boards: got %+v, expected %+v",
				caseNum, boards, item.BoardExp)
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

	mockBoardRepository := mock.NewMockIRepository(ctl)
	boardUsecase := NewUsecase(mockBoardRepository)

	cases := []TestCaseFetch{
		TestCaseFetch{
			UserId:    1,
			ErrFunc:   NoType.New("error"),
			BoardFunc: nil,
			Name:      "pin",
			BoardExp:  nil,
			Start:     5,
			Limit:     15,
		},
		TestCaseFetch{
			UserId:  1,
			Name:    "pin",
			ErrFunc: nil,
			BoardFunc: []*models.Board{
				&models.Board{
					Id: 1,
				},
				&models.Board{
					Id: 2,
				},
			},
			BoardExp: []*models.Board{
				&models.Board{
					Id: 1,
				},
				&models.Board{
					Id: 2,
				},
			},
			Start: 5,
			Limit: 15,
		},
	}

	for caseNum, item := range cases {

		gomock.InOrder(
			mockBoardRepository.EXPECT().GetByUserID(item.UserId, item.Start, item.Limit).Return(item.BoardFunc, item.ErrFunc),
		)

		boards, err := boardUsecase.GetByUserId(item.UserId, item.Start, item.Limit)
		if !reflect.DeepEqual(boards, item.BoardExp) {
			t.Errorf("[%d] wrong Boards: got %+v, expected %+v",
				caseNum, boards, item.BoardExp)
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
	ErrFunc error
	Board   *models.InputBoard
	BoardId uint
	UserId  uint
}

func TestHandler_Update(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockBoardRepository := mock.NewMockIRepository(ctl)
	boardUsecase := NewUsecase(mockBoardRepository)

	cases := []TestCaseUpdate{
		TestCaseUpdate{
			BoardId: 1,
			UserId:  2,
			Board: &models.InputBoard{
				Name:        "name",
				Description: "desc",
			},
			ErrFunc: NoType.New("error"),
		},
		TestCaseUpdate{
			BoardId: 1,
			UserId:  2,
			Board: &models.InputBoard{
				Name:        "name",
				Description: "desc",
			},
			ErrFunc: nil,
		},
	}

	for caseNum, item := range cases {

		board := &models.Board{
			Id:          item.BoardId,
			UserId:      item.UserId,
			Name:        item.Board.Name,
			Description: item.Board.Description,
		}

		gomock.InOrder(
			mockBoardRepository.EXPECT().Update(board).Return(item.ErrFunc),
		)

		err := boardUsecase.Update(item.Board, item.BoardId, item.UserId)
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
	ErrFunc error
	BoardId uint
	UserId  uint
}

func TestHandler_Delete(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockBoardRepository := mock.NewMockIRepository(ctl)
	boardUsecase := NewUsecase(mockBoardRepository)

	cases := []TestCaseDelete{
		TestCaseDelete{
			UserId:  1,
			BoardId: 2,
			ErrFunc: NoType.New("error"),
		},
		TestCaseDelete{
			UserId:  1,
			BoardId: 2,
			ErrFunc: nil,
		},
	}

	for caseNum, item := range cases {

		gomock.InOrder(
			mockBoardRepository.EXPECT().Delete(item.BoardId, item.UserId).Return(item.ErrFunc),
		)

		err := boardUsecase.Delete(item.BoardId, item.UserId)
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
