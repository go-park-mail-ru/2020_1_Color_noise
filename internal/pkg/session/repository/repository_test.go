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

type Case struct {
	s     models.Session
	answer error
}

func TestRepository_Add(t *testing.T) {

	db.Open()

	cases := []Case{
		{
			s:      models.Session{
				Cookie:     "",
				Token:      "",
				CreatedAt:  time.Time{},
				DeletingAt: time.Time{},
			},
			answer: nil,
		},
		{
			s:      models.Session{
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

	db.Open()

	cases := []Case{
		{
			s:      models.Session{
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

	db.Open()

	cases := []Case{
		{
			s:      models.Session{
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

	db.Open()

	cases := []Case{
		{
			s:      models.Session{
				Cookie:     "cookie",
				Token:      "token",
				CreatedAt:  time.Time{},
				DeletingAt: time.Time{},
			},
			answer: nil,
		},
		{
		s:      models.Session{
		},
		answer: fmt.Errorf("Session is not found, cookie: "),
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