package database

import (
	"2020_1_Color_noise/internal/models"
	"fmt"
	"testing"
	"time"
)

var BDB = NewPgxDB()

type BoardCase struct {
	answer error
	b      models.DataBaseBoard
}

func TestPgxDB_CreateBoard(t *testing.T) {
	BDB.Open()

	id, _ := BDB.CreateUser(models.DataBaseUser{
		Login: fmt.Sprint(time.Now()),
	})

	cases := []BoardCase{
		{answer: nil,
			b: models.DataBaseBoard{
				UserId:    id,
				Name:      fmt.Sprint(time.Now()),
				CreatedAt: time.Time{},
			},
		},
		{answer: fmt.Errorf("user not found"),
			b: models.DataBaseBoard{
				Name:      fmt.Sprint(time.Now()),
			},
		},
	}

	for i, item := range cases {
		_, answer := BDB.CreateBoard(item.b)
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

func TestPgxDB_DeleteBoard(t *testing.T) {
	BDB.Open()

	id, _ := BDB.CreateUser(models.DataBaseUser{
		Login:             fmt.Sprint(time.Now()),
	})

	bid, _ := BDB.CreateBoard(models.DataBaseBoard{
		UserId: id,
	})

	cases := []BoardCase{
		{answer: nil,
			b: models.DataBaseBoard{
				Id: bid,
			},
		},
		{answer: nil,
			b: models.DataBaseBoard{},
		},
	}

	for i, item := range cases {
		answer := BDB.DeleteBoard(item.b)
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

func TestPgxDB_GetBoardById(t *testing.T) {
	BDB.Open()

	id, _ := BDB.CreateUser(models.DataBaseUser{
		Login:             fmt.Sprint(time.Now()),
	})

	bid, _ := BDB.CreateBoard(models.DataBaseBoard{
		UserId: id,
	})

	cases := []BoardCase{
		{answer: nil,
			b: models.DataBaseBoard{
				Id: bid,
			},
		},
		{answer: fmt.Errorf("board not found"),
			b: models.DataBaseBoard{
			},
		},
	}

	for i, item := range cases {
		_, answer := BDB.GetBoardById(item.b)
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

func TestPgxDB_GetBoardLastPin(t *testing.T) {
	BDB.Open()

	id, _ := BDB.CreateUser(models.DataBaseUser{
		Login:             fmt.Sprint(time.Now()),
	})

	bid, _ := BDB.CreateBoard(models.DataBaseBoard{
		UserId: id,
	})

	_, _ = BDB.CreatePin(models.DataBasePin{
		BoardId:bid,
		UserId:id,
	})

	cases := []BoardCase{
		{answer: nil,
			b: models.DataBaseBoard{
				Id: bid,
			},
		},
		{answer: fmt.Errorf("board not found"),
			b: models.DataBaseBoard{
			},
		},
	}

	for i, item := range cases {
		_, answer := BDB.GetBoardLastPin(item.b)
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

func TestPgxDB_UpdateBoard(t *testing.T) {
	BDB.Open()

	id, _ := BDB.CreateUser(models.DataBaseUser{
		Login:             fmt.Sprint(time.Now()),
	})

	bid, _ := BDB.CreateBoard(models.DataBaseBoard{
		UserId: id,
	})

	cases := []BoardCase{
		{answer: nil,
			b: models.DataBaseBoard{
				Name: "new name",
				Id:   bid,
			},
		},
		{answer: nil,
			b: models.DataBaseBoard{
			},
		},
	}

	for i, item := range cases {
		answer := BDB.UpdateBoard(item.b)
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

func TestPgxDB_GetBoardsByName(t *testing.T) {
	BDB.Open()

	id, _ := BDB.CreateUser(models.DataBaseUser{
		Login:             fmt.Sprint(time.Now()),
	})

	_, _ = BDB.CreateBoard(models.DataBaseBoard{
		UserId: id,
		Name: "board",
	})

	cases := []BoardCase{
		{answer: nil,
			b: models.DataBaseBoard{
				Name: "board",
			},
		},
		{answer: nil,
			b: models.DataBaseBoard{
				Name: fmt.Sprint(time.Now()),
			},
		},
	}

	for i, item := range cases {
		_, answer := BDB.GetBoardsByName(item.b, 0, 5)
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

func TestPgxDB_GetBoardsByUserId(t *testing.T) {
	BDB.Open()

	id, _ := BDB.CreateUser(models.DataBaseUser{
		Login:             fmt.Sprint(time.Now()),
	})

	_, _ = BDB.CreateBoard(models.DataBaseBoard{
		UserId: id,
	})

	cases := []BoardCase{
		{answer: nil,
			b: models.DataBaseBoard{
				UserId: id,
			},
		},
		{answer: nil,
			b: models.DataBaseBoard{
				UserId: id,
			},
		},
		{answer: nil,
			b: models.DataBaseBoard{},
		},
	}

	for i, item := range cases {
		_, answer := BDB.GetBoardsByUserId(item.b, 0, i)
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

func TestPgxDB_GetPinsByBoardID(t *testing.T) {
	BDB.Open()

	id, _ := BDB.CreateUser(models.DataBaseUser{
		Login:             fmt.Sprint(time.Now()),
	})

	bid, _ := BDB.CreateBoard(models.DataBaseBoard{
		UserId: id,
	})

	_, _ = BDB.CreatePin(models.DataBasePin{
		UserId:      id,
		BoardId:     bid,
	})

	cases := []BoardCase{
		{answer: nil,
			b: models.DataBaseBoard{
				Id: bid,
			},
		},
		{answer: nil,
			b: models.DataBaseBoard{},
		},
	}

	for i, item := range cases {
		_, answer := BDB.GetPinsByBoardID(item.b)
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
