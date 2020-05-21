package repository

import (
	"2020_1_Color_noise/internal/models"
	"2020_1_Color_noise/internal/pkg/config"
	"2020_1_Color_noise/internal/pkg/database"
	. "2020_1_Color_noise/internal/pkg/error"
	"fmt"
	"log"
	"testing"
	"time"
)

var db = database.NewPgxDB()
var repo = NewRepository(db)

type Case struct {
	c      models.Message
	answer error
}

func TestRepository_AddMessage(t *testing.T) {
	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	db.Open(c)

	id, _ := db.CreateUser(models.DataBaseUser{
		Login: fmt.Sprint(time.Now()) + "add",
	})

	log.Print(id)

	cases := []Case{
		{
			models.Message{
				SendUser:  &models.User{Id: id},
				RecUser:   &models.User{Id: id},
				Message:   "",
				CreatedAt: time.Time{},
			},
			nil,
		},
		{
			models.Message{
				SendUser:  &models.User{Id: 0},
				RecUser:   &models.User{Id: 0},
				Message:   "",
				CreatedAt: time.Time{},
			},
			UserNotFound.Newf("User not found, id: %d", 0),
		},
	}

	for i, item := range cases {
		_, answer := repo.AddMessage(item.c.SendUser.Id, item.c.SendUser.Id, item.c.Message, "")
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

func TestRepository_GetMessages(t *testing.T) {

	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	db.Open(c)

	id, _ := db.CreateUser(models.DataBaseUser{
		Login: fmt.Sprint(time.Now()) + "get",
	})

	cases := []Case{
		{
			models.Message{
				SendUser:  &models.User{Id: id},
				RecUser:   &models.User{Id: id},
				Message:   "",
				CreatedAt: time.Time{},
			},
			nil,
		},
		{
			models.Message{
				SendUser:  &models.User{Id: id},
				RecUser:   &models.User{Id: 0},
				Message:   "",
				CreatedAt: time.Time{},
			},
			UserNotFound.Newf("User not found, id: %d", id),
		},
	}

	for i, item := range cases {
		_, _ = repo.AddMessage(item.c.SendUser.Id, item.c.RecUser.Id, item.c.Message, "")

		_, answer := repo.GetMessages(item.c.SendUser.Id, item.c.RecUser.Id, 0, 5)
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

func TestRepository_GetUsers(t *testing.T) {

	c, err := config.GetTestConfing()

	if err != nil {
		t.SkipNow()
	}
	db.Open(c)

	id, _ := db.CreateUser(models.DataBaseUser{
		Login: fmt.Sprint(time.Now()) + "getusers",
	})

	cases := []Case{
		{
			models.Message{
				SendUser:  &models.User{Id: id},
				RecUser:   &models.User{Id: id},
				Message:   "",
				CreatedAt: time.Time{},
			},
			nil,
		},
		{
			models.Message{
				SendUser:  &models.User{Id: id},
				RecUser:   &models.User{Id: 0},
				Message:   "",
				CreatedAt: time.Time{},
			},
			nil,
		},
		{
			models.Message{
				SendUser:  &models.User{Id: 0},
				RecUser:   &models.User{Id: id},
				Message:   "",
				CreatedAt: time.Time{},
			},
			UserNotFound.Newf("User not found, id: %d", 0),
		},
	}

	for i, item := range cases {
		_, _ = repo.AddMessage(item.c.SendUser.Id, item.c.RecUser.Id, item.c.Message, "")

		_, answer := repo.GetUsers(item.c.SendUser.Id, 0, 5)
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
