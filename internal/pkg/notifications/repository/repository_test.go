package repository

import (
	"2020_1_Color_noise/internal/models"
	"2020_1_Color_noise/internal/pkg/database"
	"fmt"
	"testing"
	"time"
)

var db = database.NewPgxDB()
var repo = NewRepo(db)

type NotiCase struct {
	n models.Notification
	answer error
}

func TestRepository_GetNotifications(t *testing.T) {
	db.Open()

	id, _ := db.CreateUser(models.DataBaseUser{
		Login: fmt.Sprint(time.Now()),
	})

	bid, _ := db.CreateBoard(models.DataBaseBoard{
		UserId: id,
	})

	pid, _ := db.CreatePin(models.DataBasePin{
		UserId:  id,
		BoardId: bid,
	})

	db.PutNotifications(models.DataBaseComment{
		UserId:    id,
		PinId:     pid,
		Text:      "",
	})


	cases := []NotiCase{
		{
			n:      models.Notification{
				User:    models.User{
					Id:id,
				},
				Message: "hello",
			},
			answer: nil,
		},
		{
			n:      models.Notification{
				User:    models.User{
				},
				Message: "hello",
			},
			answer: nil,
		},
	}

	for i, item := range cases {
		_, answer := repo.GetNotifications(item.n.User.Id, 0, 2)
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
