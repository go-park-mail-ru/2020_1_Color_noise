package database

import (
	"2020_1_Color_noise/internal/pkg/config"
	"errors"
	"testing"
)

var dbServiceTest = NewPgxDB()

type ServiceCase struct {
	answer error
}

func TestPgxDB_Open(t *testing.T) {
	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	cases := []ServiceCase{
		{answer: nil},
		{answer: errors.New("pool was created already")},
	}

	for i, item := range cases {
		answer := dbServiceTest.Open(c)
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

func TestPgxDB_Close(t *testing.T) {
	cases := []ServiceCase{
		{answer: nil},
	}

	for i, item := range cases {
		answer := dbServiceTest.Close()
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

func TestPgxDB_Ping(t *testing.T) {
	cases := []ServiceCase{
		{answer: nil},
	}

	for i, item := range cases {
		answer := dbServiceTest.Ping()
		if answer != item.answer {
			t.Errorf("error in test case №[%d], expected: [%v], got [%v]", i, item.answer, answer)
		}
	}
}
