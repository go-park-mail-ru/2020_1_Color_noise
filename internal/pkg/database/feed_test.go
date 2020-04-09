package database

import (
	"2020_1_Color_noise/internal/models"
	"fmt"
	"testing"
	"time"
)

var FBD = NewPgxDB()

type FeedCase struct {
	user   models.DataBaseUser
	answer error
}

func TestPgxDB_GetMainFeed(t *testing.T) {
	FBD.Open()
	id, _ := FBD.CreateUser(models.DataBaseUser{
		Login:             fmt.Sprint(time.Now()),
	})

	bid, _ := CommentDB.CreateBoard(models.DataBaseBoard{
		UserId: id,
	})

	_, _ = CommentDB.CreatePin(models.DataBasePin{
		UserId:  id,
		BoardId: bid,
	})

	cases := []FeedCase{
		{models.DataBaseUser{},
			nil,
		},
		{models.DataBaseUser{
			Id: id,},
			nil,
		},
	}

	for i, item := range cases {
		_, answer := FBD.GetMainFeed(item.user, 0, 2)
		if answer != nil && item.answer  != nil{
			if answer.Error() != item.answer.Error() {
				t.Errorf("error in test case №[%d], expected: [%v], got [%v]", i, item.answer, answer)
			}
		} else {
			if item.answer != nil || answer != nil{
				t.Errorf("error in test case №[%d], expected: [%v], got [%v]", i, item.answer, answer)
			}
		}
	}
}

func TestPgxDB_GetRecFeed(t *testing.T) {
	FBD.Open()
	id, _ := FBD.CreateUser(models.DataBaseUser{
		Email:             "update",
		Login:             fmt.Sprint(time.Now()),
		EncryptedPassword: "password",
		Subscriptions:     0,
		Subscribers:       0,
		CreatedAt:         time.Time{},
	})

	cases := []FeedCase{
		{models.DataBaseUser{},
			nil,
		},
		{models.DataBaseUser{
			Id: id,
		},
			nil,
		},
	}

	for i, item := range cases {
		_, answer := FBD.GetRecFeed(item.user, 0, 1)
		if answer != nil && item.answer  != nil{
			if answer.Error() != item.answer.Error() {
				t.Errorf("error in test case №[%d], expected: [%v], got [%v]", i, item.answer, answer)
			}
		} else {
			if item.answer != nil || answer != nil{
				t.Errorf("error in test case №[%d], expected: [%v], got [%v]", i, item.answer, answer)
			}
		}
	}
}

func TestPgxDB_GetSubFeed(t *testing.T) {
	FBD.Open()
	id, _ := FBD.CreateUser(models.DataBaseUser{
		Login:             fmt.Sprint(time.Now()),
	})

	sid, _ := FBD.CreateUser(models.DataBaseUser{
		Login:             fmt.Sprint(time.Now()),
	})

	FBD.Follow(id, sid)

	bid, _ := CommentDB.CreateBoard(models.DataBaseBoard{
		UserId: sid,
	})

	_, _ = CommentDB.CreatePin(models.DataBasePin{
		UserId:  sid,
		BoardId: bid,
	})
	_, _ = CommentDB.CreatePin(models.DataBasePin{
		UserId:  sid,
		BoardId: bid,
	})

	cases := []FeedCase{
		{models.DataBaseUser{},
			nil,
		},
		{models.DataBaseUser{
			Id: id,
		},
			nil,
		},
		{models.DataBaseUser{
			Id: sid,
		},
			nil,
		},
	}

	for i, item := range cases {
		_, answer := FBD.GetSubFeed(item.user, 0, 1)
		if answer != nil && item.answer  != nil{
			if answer.Error() != item.answer.Error() {
				t.Errorf("error in test case №[%d], expected: [%v], got [%v]", i, item.answer, answer)
			}
		} else {
			if item.answer != nil || answer != nil{
				t.Errorf("error in test case №[%d], expected: [%v], got [%v]", i, item.answer, answer)
			}
		}
	}
}
