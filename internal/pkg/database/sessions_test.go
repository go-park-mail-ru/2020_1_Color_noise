package database

import (
	"2020_1_Color_noise/internal/models"
	"2020_1_Color_noise/internal/pkg/config"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

var SDB = NewPgxDB()

type SessionCase struct {
	s      models.DataBaseSession
	answer error
}

func TestPgxDB_CreateSession(t *testing.T) {
	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	SDB.Open(c)
	id, _ := SDB.CreateUser(models.DataBaseUser{
		Email:             "create_board",
		Login:             fmt.Sprint(time.Now()),
		EncryptedPassword: "create_board",
		Subscriptions:     0,
		Subscribers:       0,
		CreatedAt:         time.Time{},
	})

	cases := []SessionCase{
		{
			s: models.DataBaseSession{
				Id:     id,
				Token:  "token",
				Cookie: "cookie",
			},
			answer: nil,
		},
	}

	for i, item := range cases {
		answer := SDB.CreateSession(item.s)
		if answer != item.answer {
			t.Errorf("error in test case №[%d], expected: [%v], got [%v]", i, item.answer, answer)
		}
	}
}

func TestPgxDB_DeleteSession(t *testing.T) {
	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	SDB.Open(c)

	id, _ := SDB.CreateUser(models.DataBaseUser{
		Email:             "create_board",
		Login:             fmt.Sprint(time.Now()),
		EncryptedPassword: "create_board",
		Subscriptions:     0,
		Subscribers:       0,
		CreatedAt:         time.Time{},
	})

	cases := []SessionCase{
		{
			s: models.DataBaseSession{
				Id:     id,
				Token:  "token",
				Cookie: "cookie",
			},
			answer: nil,
		},
	}

	for i, item := range cases {
		SDB.CreateSession(item.s)
		answer := SDB.DeleteSession(item.s)
		if answer != item.answer {
			t.Errorf("error in test case №[%d], expected: [%v], got [%v]", i, item.answer, answer)
		}
	}
}

func TestPgxDB_UpdateSession(t *testing.T) {
	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	SDB.Open(c)

	id, _ := SDB.CreateUser(models.DataBaseUser{
		Email:             "create_board",
		Login:             fmt.Sprint(time.Now()),
		EncryptedPassword: "create_board",
		Subscriptions:     0,
		Subscribers:       0,
		CreatedAt:         time.Time{},
	})

	SDB.CreateSession(models.DataBaseSession{
		Id:         id,
		Cookie:     "cookie",
		Token:      "token",
		CreatedAt:  time.Time{},
		DeletingAt: sql.NullTime{},
	})

	cases := []SessionCase{
		{
			s: models.DataBaseSession{
				Id:     id,
				Token:  "new token",
				Cookie: "cookie",
			},
			answer: nil,
		},
	}

	for i, item := range cases {
		answer := SDB.UpdateSession(item.s)
		if answer != item.answer {
			t.Errorf("error in test case №[%d], expected: [%v], got [%v]", i, item.answer, answer)
		}
	}
}

func TestPgxDB_GetSessionByCookie(t *testing.T) {
	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	SDB.Open(c)

	id, _ := SDB.CreateUser(models.DataBaseUser{
		Email:             "create_board",
		Login:             fmt.Sprint(time.Now()),
		EncryptedPassword: "create_board",
		Subscriptions:     0,
		Subscribers:       0,
		CreatedAt:         time.Time{},
	})

	SDB.CreateSession(models.DataBaseSession{
		Id:         id,
		Cookie:     "cookie",
		Token:      "token",
		CreatedAt:  time.Time{},
		DeletingAt: sql.NullTime{},
	})

	cases := []SessionCase{
		{
			s: models.DataBaseSession{
				Cookie: "cookie",
			},
			answer: nil,
		},
	}

	for i, item := range cases {
		_, answer := SDB.GetSessionByCookie(item.s)
		if answer != item.answer {
			t.Errorf("error in test case №[%d], expected: [%v], got [%v]", i, item.answer, answer)
		}
	}
}
