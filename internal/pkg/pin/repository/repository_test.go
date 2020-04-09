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

type pinCase struct {
	p models.Pin
	answer error
}

func TestRepository_Create(t *testing.T) {
	db.Open()

	id, _ := db.CreateUser(models.DataBaseUser{
		Login: fmt.Sprint(time.Now()),
	})

	bid, _ := db.CreateBoard(models.DataBaseBoard{
		UserId: id,
	})

	cases := []pinCase{
		{
			p:      models.Pin{
				UserId:      id,
				BoardId:     bid,
			},
			answer: nil,
		},
		{
			p:      models.Pin{
				UserId:      0,
				BoardId:     bid,
			},
			answer: fmt.Errorf("Pin can not be created, err: pin creation error"),
		},
	}

	for i, item := range cases {
		_, answer := repo.Create(&item.p)
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

func TestRepository_Delete(t *testing.T) {
	db.Open()

	id, _ := db.CreateUser(models.DataBaseUser{
		Login: fmt.Sprint(time.Now()),
	})

	bid, _ := db.CreateBoard(models.DataBaseBoard{
		UserId: id,
	})

	cases := []pinCase{
		{
			p:      models.Pin{
				UserId:      id,
				BoardId:     bid,
			},
			answer: nil,
		},
		{
			p:      models.Pin{
				UserId:      0,
				BoardId:     bid,
			},
			answer: nil,
		},
	}

	for i, item := range cases {
		id, _ := repo.Create(&item.p)
		answer := repo.Delete(id, item.p.UserId)
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

func TestRepository_Update(t *testing.T) {
	db.Open()

	id, _ := db.CreateUser(models.DataBaseUser{
		Login: fmt.Sprint(time.Now()),
	})

	bid, _ := db.CreateBoard(models.DataBaseBoard{
		UserId: id,
	})

	cases := []pinCase{
		{
			p:      models.Pin{
				UserId:      id,
				BoardId:     bid,
			},
			answer: nil,
		},
		{
			p:      models.Pin{
				UserId:      0,
				BoardId:     bid,
			},
			answer: nil,
		},
	}

	for i, item := range cases {
		repo.Create(&item.p)
		answer := repo.Update(&item.p)
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

func TestRepository_GetByID(t *testing.T) {

	db.Open()

	id, _ := db.CreateUser(models.DataBaseUser{
		Login: fmt.Sprint(time.Now()),
	})

	bid, _ := db.CreateBoard(models.DataBaseBoard{
		UserId: id,
	})

	cases := []pinCase{
		{
			p:      models.Pin{
				UserId:      id,
				BoardId:     bid,
			},
			answer: nil,
		},
		{
			p:      models.Pin{
				UserId:      0,
				BoardId:     bid,
			},
			answer: fmt.Errorf("Pin not found, id: 0"),
		},
	}

	for i, item := range cases {
		id, _ := repo.Create(&item.p)
		item.p.Id = id
		_, answer := repo.GetByID(item.p.Id)
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

func TestRepository_GetByName(t *testing.T) {

	db.Open()

	id, _ := db.CreateUser(models.DataBaseUser{
		Login: fmt.Sprint(time.Now()),
	})

	bid, _ := db.CreateBoard(models.DataBaseBoard{
		UserId: id,
	})

	cases := []pinCase{
		{
			p:      models.Pin{
				Name: "name",
				UserId:      id,
				BoardId:     bid,
			},
			answer: nil,
		},
		{
			p:      models.Pin{
				UserId:      0,
				BoardId:     bid,
			},
			answer: nil,
		},
	}

	for i, item := range cases {
		id, _ := repo.Create(&item.p)
		item.p.Id = id
		_, answer := repo.GetByName(item.p.Name, 0, 2)
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

func TestRepository_GetByUserID(t *testing.T) {

	db.Open()

	id, _ := db.CreateUser(models.DataBaseUser{
		Login: fmt.Sprint(time.Now()),
	})

	bid, _ := db.CreateBoard(models.DataBaseBoard{
		UserId: id,
	})

	cases := []pinCase{
		{
			p:      models.Pin{
				UserId:      id,
				BoardId:     bid,
			},
			answer: nil,
		},
		{
			p:      models.Pin{
				UserId:      0,
				BoardId:     bid,
			},
			answer: nil,
		},
	}

	for i, item := range cases {
		id, _ := repo.Create(&item.p)
		item.p.Id = id
		_, answer := repo.GetByUserID(item.p.UserId, 0, 2)
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