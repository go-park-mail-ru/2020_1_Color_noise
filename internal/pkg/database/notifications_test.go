package database

/*
import (
	"2020_1_Color_noise/internal/models"
	"2020_1_Color_noise/internal/pkg/config"
	"fmt"
	"testing"
	"time"
)

var NBD = NewPgxDB()

type NotiCase struct {
	user   models.DataBaseUser
	cm     models.DataBaseComment
	answer error
}

func TestPgxDB_PutNotifications(t *testing.T) {
	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	NBD.Open(c)

	id, _ := NBD.CreateUser(models.DataBaseUser{
		Login:             fmt.Sprint(time.Now()),
	})

	bid, _ := NBD.CreateBoard(models.DataBaseBoard{
		UserId: id,
	})

	pid, _ := NBD.CreatePin(models.DataBasePin{
		UserId:  id,
		BoardId: bid,
	})

	cases := []NotiCase{
		{
			cm: models.DataBaseComment{
				UserId:    id,
				PinId:     pid,
				Text:      "new comment",
			},
			answer: nil,
		},
		{
			cm: models.DataBaseComment{
				PinId:     pid,
				Text:      "new comment",
			},
			answer: fmt.Errorf("no notifications"),
		},
	}

	for i, item := range cases {
		_, answer := NBD.PutNotifications(item.cm)
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

func TestPgxDB_GetNotifications(t *testing.T) {
	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	NBD.Open(c)

	id, _ := NBD.CreateUser(models.DataBaseUser{
		Login:             fmt.Sprint(time.Now()),
	})

	bid, _ := NBD.CreateBoard(models.DataBaseBoard{
		UserId: id,
	})

	pid, _ := NBD.CreatePin(models.DataBasePin{
		UserId:  id,
		BoardId: bid,
	})

	_, _ = NBD.PutNotifications(models.DataBaseComment{
		Id:id,
		UserId:id,
		PinId:pid,
	})

	cases := []NotiCase{
		{
			models.DataBaseUser{
				Id: id,
			},
			models.DataBaseComment{},
			nil,
		},
	}

	for i, item := range cases {
		_, answer := NBD.GetNotifications(item.user)
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


*/
