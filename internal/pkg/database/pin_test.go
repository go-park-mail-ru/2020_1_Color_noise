package database

/*
import (
	"2020_1_Color_noise/internal/models"
	"2020_1_Color_noise/internal/pkg/config"
	"fmt"
	"testing"
	"time"
)

var PDB = NewPgxDB()

type PinCase struct {
	pin    models.DataBasePin
	answer error
}

func TestPgxDB_CreatePin(t *testing.T) {
	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	PDB.Open(c)

	id, _ := PDB.CreateUser(models.DataBaseUser{
		Login: fmt.Sprint(time.Now()),
	})

	bid, _ := PDB.CreateBoard(models.DataBaseBoard{
		UserId: id,
	})

	cases := []PinCase{{
		pin: models.DataBasePin{
			UserId:  id,
			BoardId: bid,
		},
		answer: nil,
	},{
		pin: models.DataBasePin{
		},
		answer: fmt.Errorf("pin creation error"),
	},
	}

	for i, item := range cases {
		_, answer := PDB.CreatePin(item.pin)
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

func TestPgxDB_DeletePin(t *testing.T) {

	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	PDB.Open(c)

	id, _ := PDB.CreateUser(models.DataBaseUser{
		Login: fmt.Sprint(time.Now()),
	})

	bid, _ := PDB.CreateBoard(models.DataBaseBoard{
		UserId: id,
	})

	pid, _ := PDB.CreatePin(models.DataBasePin{
		UserId:  id,
		BoardId: bid,
	})


	cases := []PinCase{{
		pin: models.DataBasePin{
			Id:pid,
			UserId:  id,
			BoardId: bid,
		},
		answer: nil,
	},
	{pin: models.DataBasePin{
		UserId:  id,
		BoardId: bid,
	},
		answer: nil,
	},
		{pin: models.DataBasePin{
			Id:pid,
			UserId:  id,
		},
			answer: nil,
		},
	}

	for i, item := range cases {
		PDB.CreatePin(item.pin)
		answer := PDB.DeletePin(item.pin)
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

func TestPgxDB_UpdatePin(t *testing.T) {

	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	PDB.Open(c)

	id, _ := PDB.CreateUser(models.DataBaseUser{
		Login: fmt.Sprint(time.Now()),
	})

	bid, _ := PDB.CreateBoard(models.DataBaseBoard{
		UserId: id,
	})

	pid, _ := PDB.CreatePin(models.DataBasePin{
		UserId:  id,
		BoardId: bid,
	})


	cases := []PinCase{{
		pin: models.DataBasePin{
			Id:pid,
			UserId:  id,
			BoardId: bid,
		},
		answer: nil,
	},
		{pin: models.DataBasePin{
			UserId:  id,
			BoardId: bid,
		},
			answer: nil,
		},
		{pin: models.DataBasePin{
			Id:pid,
			UserId:  id,
		},
			answer: fmt.Errorf("board not found"),
		},
	}

	for i, item := range cases {
		PDB.CreatePin(item.pin)
		answer := PDB.UpdatePin(item.pin)
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

func TestPgxDB_GetPinById(t *testing.T) {

	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	PDB.Open(c)

	id, _ := PDB.CreateUser(models.DataBaseUser{
		Login: fmt.Sprint(time.Now()),
	})

	bid, _ := PDB.CreateBoard(models.DataBaseBoard{
		UserId: id,
	})

	cases := []PinCase{{
		pin: models.DataBasePin{
			UserId:  id,
			BoardId: bid,
		},
		answer: nil},

		{pin: models.DataBasePin{},
			answer: fmt.Errorf("pin not found")},
	}

	for i, item := range cases {
		id, err := PDB.CreatePin(item.pin)
		if err == nil {
			item.pin.Id = id
		}
		_, answer := PDB.GetPinById(item.pin)
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

func TestPgxDB_GetPinsByUserId(t *testing.T) {

	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	PDB.Open(c)

	id, _ := PDB.CreateUser(models.DataBaseUser{
		Login: fmt.Sprint(time.Now()),
	})

	bid, _ := PDB.CreateBoard(models.DataBaseBoard{
		UserId: id,
	})

	cases := []PinCase{{
		pin: models.DataBasePin{
			UserId:  id,
			BoardId: bid,
		},
		answer: nil},
		{pin: models.DataBasePin{},
			answer: nil},
	}

	for i, item := range cases {
		PDB.CreatePin(item.pin)
		_, answer := PDB.GetPinsByUserId(item.pin)
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

func TestPgxDB_GetPinsByName(t *testing.T) {

	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	PDB.Open(c)

	id, _ := PDB.CreateUser(models.DataBaseUser{
		Login: fmt.Sprint(time.Now()),
	})

	bid, _ := PDB.CreateBoard(models.DataBaseBoard{
		UserId: id,
	})
	cases := []PinCase{{
		pin: models.DataBasePin{
			UserId:  id,
			BoardId: bid,
		},
		answer: nil},
		{pin: models.DataBasePin{},
			answer: nil},
	}

	for i, item := range cases {
		PDB.CreatePin(item.pin)
		_, answer := PDB.GetPinsByName(item.pin)
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
*/
