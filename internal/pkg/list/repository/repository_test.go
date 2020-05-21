package repository

import (
	"2020_1_Color_noise/internal/models"
	"2020_1_Color_noise/internal/pkg/config"
	"2020_1_Color_noise/internal/pkg/database"
	"fmt"
	"testing"
	"time"
)

var db = database.NewPgxDB()
var repo = NewRepo(db)

type FeedCase struct {
	u      models.User
	answer error
}

func TestRepository_GetMainList(t *testing.T) {
	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	db.Open(c)

	id, _ := db.CreateUser(models.DataBaseUser{
		Login: fmt.Sprint(time.Now()),
	})

	cases := []FeedCase{
		{
			u: models.User{
				Id: id,
			},
			answer: nil,
		},
		{
			u:      models.User{},
			answer: nil,
		},
	}

	for i, item := range cases {
		_, answer := repo.GetMainList(0, 2)
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

func TestRepository_GetSubList(t *testing.T) {
	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	db.Open(c)

	id, _ := db.CreateUser(models.DataBaseUser{
		Login: fmt.Sprint(time.Now()),
	})

	cases := []FeedCase{
		{
			u: models.User{
				Id: id,
			},
			answer: nil,
		},
		{
			u:      models.User{},
			answer: nil,
		},
	}

	for i, item := range cases {
		_, answer := repo.GetSubList(item.u.Id, 0, 2)
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

func TestRepository_GetSRecList(t *testing.T) {
	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	db.Open(c)

	id, _ := db.CreateUser(models.DataBaseUser{
		Login: fmt.Sprint(time.Now()),
	})

	cases := []FeedCase{
		{
			u: models.User{
				Id: id,
			},
			answer: nil,
		},
		{
			u:      models.User{},
			answer: fmt.Errorf("User not found, user id = 0"),
		},
	}

	for i, item := range cases {
		_, answer := repo.GetRecommendationList(item.u.Id, 0, 2)
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
