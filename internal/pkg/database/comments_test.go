package database

import (
	"2020_1_Color_noise/internal/models"
	"2020_1_Color_noise/internal/pkg/config"
	"fmt"
	"testing"
	"time"
)

var CommentDB = NewPgxDB()

type CommentCase struct {
	answer error
	cm     models.DataBaseComment
}

func TestPgxDB_CreateComment(t *testing.T) {
	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	CommentDB.Open(c)

	id, _ := CommentDB.CreateUser(models.DataBaseUser{
		Login: fmt.Sprint(time.Now()) + "TestPgxDB_GetPinsByBoardID",
	})

	bid, _ := CommentDB.CreateBoard(models.DataBaseBoard{
		UserId: id,
	})

	pid, _ := CommentDB.CreatePin(models.DataBasePin{
		UserId:  id,
		BoardId: bid,
	})

	cases := []CommentCase{
		{
			answer: nil,
			cm: models.DataBaseComment{
				Text:   "my new comment",
				UserId: id,
				PinId:  pid,
			},
		},
		{
			answer: fmt.Errorf("comment error"),
			cm: models.DataBaseComment{
				Text:   "new comment",
				UserId: id,
				PinId:  pid + 1,
			},
		},
	}

	for i, item := range cases {
		_, answer := CommentDB.CreateComment(item.cm)
		if answer != nil && item.answer != nil {
			if answer.Error() != item.answer.Error() {
				t.Errorf("error in test case №[%d], expected: [%v], got [%v]", i, item.answer, answer)
			}
		} else {
			if item.answer != nil || answer != nil {
				t.Errorf("error in test case №[%d], expected: [%v], got [%v]", i, item.answer, answer)
			}
		}
	}
}

func TestPgxDB_UpdateComment(t *testing.T) {
	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	CommentDB.Open(c)

	id, _ := CommentDB.CreateUser(models.DataBaseUser{
		Login: fmt.Sprint(time.Now()) + "TestPgxDB_UpdateComment",
	})

	bid, _ := CommentDB.CreateBoard(models.DataBaseBoard{
		UserId: id,
	})

	pid, _ := CommentDB.CreatePin(models.DataBasePin{
		UserId:  id,
		BoardId: bid,
	})

	cid, _ := CommentDB.CreateComment(models.DataBaseComment{
		UserId: id,
		PinId:  pid,
		Text:   "text",
	})

	cases := []CommentCase{
		{
			answer: nil,
			cm: models.DataBaseComment{
				Text: "new new comment",
				Id:   cid,
			},
		},
		{
			answer: nil,
			cm: models.DataBaseComment{
				Text: "new new comment",
				Id:   1000,
			},
		},
	}

	for i, item := range cases {
		answer := CommentDB.UpdateComment(item.cm)
		if answer != nil && item.answer != nil {
			if answer.Error() != item.answer.Error() {
				t.Errorf("error in test case №[%d], expected: [%v], got [%v]", i, item.answer, answer)
			}
		} else {
			if item.answer != nil || answer != nil {
				t.Errorf("error in test case №[%d], expected: [%v], got [%v]", i, item.answer, answer)
			}
		}
	}
}

func TestPgxDB_DeleteComment(t *testing.T) {

	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	CommentDB.Open(c)

	id, _ := CommentDB.CreateUser(models.DataBaseUser{
		Login: fmt.Sprint(time.Now()) + "TestPgxDB_DeleteComment",
	})

	bid, _ := CommentDB.CreateBoard(models.DataBaseBoard{
		UserId: id,
	})

	pid, _ := CommentDB.CreatePin(models.DataBasePin{
		UserId:  id,
		BoardId: bid,
	})

	cid, _ := CommentDB.CreateComment(models.DataBaseComment{
		UserId: id,
		PinId:  pid,
		Text:   "text",
	})

	cases := []CommentCase{
		{
			answer: nil,
			cm: models.DataBaseComment{
				Id: cid,
			},
		},
		{
			answer: nil,
			cm: models.DataBaseComment{
				Id: cid,
			},
		},
	}

	for i, item := range cases {
		answer := CommentDB.DeleteComment(item.cm)
		if answer != nil && item.answer != nil {
			if answer.Error() != item.answer.Error() {
				t.Errorf("error in test case №[%d], expected: [%v], got [%v]", i, item.answer, answer)
			}
		} else {
			if item.answer != nil || answer != nil {
				t.Errorf("error in test case №[%d], expected: [%v], got [%v]", i, item.answer, answer)
			}
		}
	}

}

type SimpleCase struct {
	in     models.DataBaseComment
	out    models.Comment
	answer error
}

func TestPgxDB_GetCommentById(t *testing.T) {

	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	CommentDB.Open(c)

	id, _ := CommentDB.CreateUser(models.DataBaseUser{
		Login: fmt.Sprint(time.Now()) + "TestPgxDB_GetCommentById",
	})

	bid, _ := CommentDB.CreateBoard(models.DataBaseBoard{
		UserId: id,
	})

	pid, _ := CommentDB.CreatePin(models.DataBasePin{
		UserId:  id,
		BoardId: bid,
	})

	cid, _ := CommentDB.CreateComment(models.DataBaseComment{
		UserId: id,
		PinId:  pid,
		Text:   "text",
	})

	cases := []SimpleCase{
		{
			in: models.DataBaseComment{
				Id: cid,
			},
			out:    models.Comment{},
			answer: nil,
		},
		{
			in:     models.DataBaseComment{},
			out:    models.Comment{},
			answer: fmt.Errorf("comment not found"),
		},
	}

	for i, item := range cases {
		_, answer := CommentDB.GetCommentById(item.in)
		if answer != nil && item.answer != nil {
			if answer.Error() != item.answer.Error() {
				t.Errorf("error in test case №[%d], expected: [%v], got [%v]", i, item.answer, answer)
			}
		} else {
			if item.answer != nil || answer != nil {
				t.Errorf("error in test case №[%d], expected: [%v], got [%v]", i, item.answer, answer)
			}
		}
	}
}

type ComplicatedCase struct {
	in     models.DataBaseComment
	out    []*models.Comment
	answer error
}

func TestPgxDB_GetCommentsByPinId(t *testing.T) {
	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	CommentDB.Open(c)

	id, _ := CommentDB.CreateUser(models.DataBaseUser{
		Login: fmt.Sprint(time.Now()) + "TestPgxDB_GetCommentByPinId",
	})

	bid, _ := CommentDB.CreateBoard(models.DataBaseBoard{
		UserId: id,
	})

	pid, _ := CommentDB.CreatePin(models.DataBasePin{
		UserId:  id,
		BoardId: bid,
	})

	_, _ = CommentDB.CreateComment(models.DataBaseComment{
		UserId: id,
		PinId:  pid,
		Text:   "text",
	})

	cases := []ComplicatedCase{
		{
			in: models.DataBaseComment{
				PinId: pid,
			},
			out:    nil,
			answer: nil,
		},
		{
			in:     models.DataBaseComment{},
			out:    nil,
			answer: nil,
		},
	}

	for i, item := range cases {
		_, answer := CommentDB.GetCommentsByPinId(item.in, 0, 2)
		if answer != nil && item.answer != nil {
			if answer.Error() != item.answer.Error() {
				t.Errorf("error in test case №[%d], expected: [%v], got [%v]", i, item.answer, answer)
			}
		} else {
			if item.answer != nil || answer != nil {
				t.Errorf("error in test case №[%d], expected: [%v], got [%v]", i, item.answer, answer)
			}
		}
	}
}

func TestPgxDB_GetCommentsByText(t *testing.T) {
	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	CommentDB.Open(c)

	cases := []ComplicatedCase{
		{
			in: models.DataBaseComment{
				Text: "text",
			},
			out:    nil,
			answer: nil,
		},
		{
			in: models.DataBaseComment{
				Text: "no such text",
			},
			out:    nil,
			answer: nil,
		},
	}

	for i, item := range cases {
		_, answer := CommentDB.GetCommentsByText(item.in, 0, 2)
		if answer != nil && item.answer != nil {
			if answer.Error() != item.answer.Error() {
				t.Errorf("error in test case №[%d], expected: [%v], got [%v]", i, item.answer, answer)
			}
		} else {
			if item.answer != nil || answer != nil {
				t.Errorf("error in test case №[%d], expected: [%v], got [%v]", i, item.answer, answer)
			}
		}
	}
}
