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
	c    models.Comment
	answer error
}

func TestRepository_Create(t *testing.T) {
	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	db.Open(c)
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

	cases := []Case{
		{
			c:      models.Comment{
				UserId:    id,
				PinId:     pid,
				Text:      "text",
			},
			answer: nil,
		},
	}

	for i, item := range cases {
		_, answer := repo.Create(&item.c)
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

	bid, _ := db.CreateBoard(models.DataBaseBoard{
		UserId: id,
	})

	pid, _ := db.CreatePin(models.DataBasePin{
		UserId:  id,
		BoardId: bid,
	})

	cases := []Case{
		{
			c:      models.Comment{
				UserId:    id,
				PinId:     pid,
				Text:      "text",
			},
			answer: nil,
		},
		{
			c:      models.Comment{
				PinId:     pid,
				Text:      "text",
			},
			answer: nil,
		},
	}

	for i, item := range cases {
		id , _ = repo.Create(&item.c)
		item.c.Id = id
		answer := repo.Update(&item.c)
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

	cases := []Case{
		{
			c:      models.Comment{
				UserId:    id,
				PinId:     pid,
				Text:      "text",
			},
			answer: nil,
		},

		{
			c:      models.Comment{
				PinId:     pid,
				Text:      "text",
			},
			answer: nil,
		},
	}

	for i, item := range cases {
		id , _ = repo.Create(&item.c)
		item.c.Id = id
		answer := repo.Delete(item.c.Id)
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


	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	db.Open(c)
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

	cases := []Case{
		{
			c:      models.Comment{
				UserId:    id,
				PinId:     pid,
				Text:      "text",
			},
			answer: nil,
		},

		{
			c:      models.Comment{
			},
			answer: fmt.Errorf("Repo: Getting by id comment error, id: 0"),
		},
	}

	for i, item := range cases {
		id , _ = repo.Create(&item.c)
		item.c.Id = id
		_, answer := repo.GetByID(item.c.Id)
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

func TestRepository_GetByPinID(t *testing.T) {


	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	db.Open(c)
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

	cases := []Case{
		{
			c:      models.Comment{
				UserId:    id,
				PinId:     pid,
				Text:      "text",
			},
			answer: nil,
		},
		{
			c:      models.Comment{
			},
			answer: nil,
		},
	}

	for i, item := range cases {
		_ , _ = repo.Create(&item.c)
		_, answer := repo.GetByPinID(item.c.PinId, 0, 2)
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

func TestRepository_GetByText(t *testing.T) {


	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	db.Open(c)
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

	cases := []Case{
		{
			c:      models.Comment{
				UserId:    id,
				PinId:     pid,
				Text:      "text",
			},
			answer: nil,
		},
		{
			c:      models.Comment{
			},
			answer: nil,
		},
	}

	for i, item := range cases {
		id , _ = repo.Create(&item.c)
		item.c.Id = id
		_, answer := repo.GetByText(item.c.Text, 0, 2)
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