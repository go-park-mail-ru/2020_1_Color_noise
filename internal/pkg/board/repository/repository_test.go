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
	b     models.Board
	answer error
}

func TestRepository_Create(t *testing.T) {
	db.Open()
	id, _ := db.CreateUser(models.DataBaseUser{
		Login: fmt.Sprint(time.Now()),
	})

	cases := []Case{
		{
			b:      models.Board{
				UserId:    id,
			},
			answer: nil,
		},
		{
			b:      models.Board{
			},
			answer: fmt.Errorf("board can not be created, err: user not found"),
		},
	}

	for i, item := range cases {
		_, answer := repo.Create(&item.b)
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
	id, _ := db.CreateUser(models.DataBaseUser{
		Login: fmt.Sprint(time.Now()),
	})

	cases := []Case{
		{
			b:      models.Board{
				UserId:    id,
			},
			answer: nil,
		},
		{
			b:      models.Board{
			},
			answer: nil,
		},
	}

	for i, item := range cases {
		id, _ := repo.Create(&item.b)
		answer := repo.Delete(id, item.b.UserId)
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
	id, _ := db.CreateUser(models.DataBaseUser{
		Login: fmt.Sprint(time.Now()),
	})

	cases := []Case{
		{
			b:      models.Board{
				UserId:    id,
			},
			answer: nil,
		},
		{
			b:      models.Board{
			},
			answer:nil,
		},
	}

	for i, item := range cases {
		id, _ := repo.Create(&item.b)
		item.b.Id = id
		answer := repo.Update(&item.b)
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

func TestRepository_GetByID(t *testing.T) {
	db.Open()
	id, _ := db.CreateUser(models.DataBaseUser{
		Login: fmt.Sprint(time.Now()),
	})

	cases := []Case{
		{
			b:      models.Board{
				UserId:    id,
			},
			answer: nil,
		},
		{
			b:      models.Board{
			},
			answer: fmt.Errorf("Board not found, board id: 0"),
		},
	}

	for i, item := range cases {
		id, _ := repo.Create(&item.b)
		item.b.Id = id
		_, answer := repo.GetByID(item.b.Id)
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

func TestRepository_GetByUserID(t *testing.T) {

	db.Open()
	id, _ := db.CreateUser(models.DataBaseUser{
		Login: fmt.Sprint(time.Now()),
	})

	cases := []Case{
		{
			b:      models.Board{
				UserId:    id,
			},
			answer: nil,
		},
	}

	for i, item := range cases {
		_, _ = repo.Create(&item.b)
		_, answer := repo.GetByUserID(item.b.UserId, 0, 2)
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

func TestRepository_GetByName(t *testing.T) {
	db.Open()
	id, _ := db.CreateUser(models.DataBaseUser{
		Login: fmt.Sprint(time.Now()),
	})

	cases := []Case{
		{
			b:      models.Board{
				UserId:    id,
			},
			answer: nil,
		},
	}

	for i, item := range cases {
		id, _ := repo.Create(&item.b)
		item.b.Id = id
		_, answer := repo.GetByName(item.b.Name, 0, 2)
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

func TestRepository_GetByNameID(t *testing.T) {
	db.Open()
	id, _ := db.CreateUser(models.DataBaseUser{
		Login: fmt.Sprint(time.Now()),
	})

	cases := []Case{
		{
			b:      models.Board{
				UserId:    id,
			},
			answer: fmt.Errorf("board not found"),
		},
	}

	for i, item := range cases {
		id, _ := repo.Create(&item.b)
		item.b.Id = id
		_, answer := repo.GetByNameID(item.b.Id)
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