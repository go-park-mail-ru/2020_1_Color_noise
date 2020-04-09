package usecase

import (
	"2020_1_Color_noise/internal/models"
	"2020_1_Color_noise/internal/pkg/comment/mock"
	. "2020_1_Color_noise/internal/pkg/error"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

type TestCaseCreate struct {
	ErrFunc   error
	Comment   *models.InputComment
	CommentId uint
	UserId    uint
}

func TestHandler_Create(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockCommentRepository := mock.NewMockIRepository(ctl)
	commentUsecase := NewUsecase(mockCommentRepository)

	cases := []TestCaseCreate{
		TestCaseCreate{
			CommentId:  0,
			UserId:     2,
			Comment: 	&models.InputComment{
				Text:   "comment",
				PinId:	1,
			},
			ErrFunc:   NoType.New("error"),
		},
		TestCaseCreate{
			CommentId:  1,
			UserId:     2,
			Comment: 	&models.InputComment{
				Text:   "comment",
				PinId:	 1,
			},
			ErrFunc:   nil,
		},
	}

	for caseNum, item := range cases {

		comment := &models.Comment{
			UserId:      item.UserId,
			PinId: 		 item.Comment.PinId,
			Text:        item.Comment.Text,
		}

		gomock.InOrder(
			mockCommentRepository.EXPECT().Create(comment).Return(item.CommentId, item.ErrFunc),
		)

		id, err := commentUsecase.Create(item.Comment, item.UserId)
		if id != item.CommentId {
			t.Errorf("[%d] wrong Id: got %+v, expected %+v",
				caseNum, id, item.CommentId)
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
	ErrFunc       error
	CommentFunc   *models.Comment
	CommentExp    *models.Comment
	CommentId        uint
}

func TestHandler_GetById(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockCommentRepository := mock.NewMockIRepository(ctl)
	commentUsecase := NewUsecase(mockCommentRepository)

	cases := []TestCaseGet{
		TestCaseGet{
			CommentId:    1,
			ErrFunc:      NoType.New("error"),
			CommentFunc:  nil,
			CommentExp:   nil,
		},
		TestCaseGet{
			CommentId:    1,
			ErrFunc:  nil,
			CommentFunc:  &models.Comment{
				Text:   "comment",
				PinId:	 1,
			},
			CommentExp:   &models.Comment{
				Text:   "comment",
				PinId:	 1,
			},
		},
	}

	for caseNum, item := range cases {

		gomock.InOrder(
			mockCommentRepository.EXPECT().GetByID(item.CommentId).Return(item.CommentFunc, item.ErrFunc),
		)

		comment, err := commentUsecase.GetById(item.CommentId)
		if !reflect.DeepEqual(comment, item.CommentExp) {
			t.Errorf("[%d] wrong Comment: got %+v, expected %+v",
				caseNum, comment, item.CommentExp)
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
	CommentFunc []*models.Comment
	CommentExp  []*models.Comment
	PinId	    uint
	Text	    string
	Start       int
	Limit 	    int
}

func TestHandler_GetByName(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockCommentRepository := mock.NewMockIRepository(ctl)
	commentUsecase := NewUsecase(mockCommentRepository)

	cases := []TestCaseFetch{
		TestCaseFetch{
			PinId:     1,
			ErrFunc:    NoType.New("error"),
			CommentFunc:  nil,
			CommentExp:   nil,
			Start:      5,
			Limit:      15,
		},
		TestCaseFetch{
			PinId:   1,
			ErrFunc:  nil,
			CommentFunc:  []*models.Comment{
				&models.Comment{
					Id:    1,
				},
				&models.Comment{
					Id:    2,
				},
			},
			CommentExp:  []*models.Comment{
				&models.Comment{
					Id:    1,
				},
				&models.Comment{
					Id:    2,
				},
			},
			Start:    5,
			Limit:    15,
		},
	}

	for caseNum, item := range cases {

		gomock.InOrder(
			mockCommentRepository.EXPECT().GetByPinID(item.PinId, item.Start, item.Limit).Return(item.CommentFunc, item.ErrFunc),
		)

		comments, err := commentUsecase.GetByPinId(item.PinId, item.Start, item.Limit)
		if !reflect.DeepEqual(comments, item.CommentExp) {
			t.Errorf("[%d] wrong Comments: got %+v, expected %+v",
				caseNum, comments, item.CommentExp)
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

func TestHandler_GetByText(t *testing.T) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockCommentRepository := mock.NewMockIRepository(ctl)
	commentUsecase := NewUsecase(mockCommentRepository)

	cases := []TestCaseFetch{
		TestCaseFetch{
			Text:      "comment",
			ErrFunc:    NoType.New("error"),
			CommentFunc:  nil,
			CommentExp:   nil,
			Start:      5,
			Limit:      15,
		},
		TestCaseFetch{
			Text:      "comment",
			ErrFunc:  nil,
			CommentFunc:  []*models.Comment{
				&models.Comment{
					Id:    1,
				},
				&models.Comment{
					Id:    2,
				},
			},
			CommentExp:  []*models.Comment{
				&models.Comment{
					Id:    1,
				},
				&models.Comment{
					Id:    2,
				},
			},
			Start:    5,
			Limit:    15,
		},
	}

	for caseNum, item := range cases {

		gomock.InOrder(
			mockCommentRepository.EXPECT().GetByText(item.Text, item.Start, item.Limit).Return(item.CommentFunc, item.ErrFunc),
		)

		comments, err := commentUsecase.GetByText(item.Text, item.Start, item.Limit)
		if !reflect.DeepEqual(comments, item.CommentExp) {
			t.Errorf("[%d] wrong Comments: got %+v, expected %+v",
				caseNum, comments, item.CommentExp)
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

