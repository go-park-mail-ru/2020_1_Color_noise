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

type Case struct {
	s      models.Session
	answer error
}

func TestRepository_Add(t *testing.T) {


	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	db.Open(c)

	id, _ := db.CreateUser(models.DataBaseUser{
		Login: fmt.Sprint(time.Now()),
	})

	cases := []Case{
		{
			s: models.Session{
				Id:id,
				Cookie:     "",
				Token:      "",
				CreatedAt:  time.Time{},
				DeletingAt: time.Time{},
			},
			answer: nil,
		},
		{
			s: models.Session{
				Id:id,
				Cookie:     "cookie",
				Token:      "token",
				CreatedAt:  time.Time{},
				DeletingAt: time.Time{},
			},
			answer: nil,
		},
	}

	for i, item := range cases {
		answer := repo.Add(&item.s)
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

func TestRepository_Delete(t *testing.T) {

	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	db.Open(c)

	cases := []Case{
		{
			s: models.Session{
				Cookie:     "",
				Token:      "",
				CreatedAt:  time.Time{},
				DeletingAt: time.Time{},
			},
			answer: nil,
		},
	}

	for i, item := range cases {
		repo.Add(&item.s)
		answer := repo.Delete(item.s.Cookie)
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

func TestRepository_Update(t *testing.T) {

	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	db.Open(c)

	id, _ := db.CreateUser(models.DataBaseUser{
		Login: fmt.Sprint(time.Now()),
	})


	cases := []Case{
		{
			s: models.Session{
				Id: id,
				Cookie:     "",
				Token:      "",
				CreatedAt:  time.Time{},
				DeletingAt: time.Time{},
			},
			answer: nil,
		},
	}

	for i, item := range cases {
		repo.Add(&item.s)
		answer := repo.Update(&item.s)
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

func TestRepository_GetByCookie(t *testing.T) {

	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	db.Open(c)

	cases := []Case{
		{
			s: models.Session{
				Cookie:     "cookie",
				Token:      "token",
				CreatedAt:  time.Time{},
				DeletingAt: time.Time{},
			},
			answer: nil,
		},
		{
			s:      models.Session{},
			answer: nil,
		},
	}

	for i, item := range cases {
		repo.Add(&item.s)
		_, answer := repo.GetByCookie(item.s.Cookie)
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
