package database

import (
	"2020_1_Color_noise/internal/models"
	"2020_1_Color_noise/internal/pkg/config"
	"fmt"
	"testing"
	"time"
)

type UserCreateTestCase struct {
	user   models.DataBaseUser
	id     uint
	answer error
}

var dbTest = NewPgxDB()

func TestUserCreate(t *testing.T) {
	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	dbTest.Open(c)

	login := fmt.Sprint(time.Now())

	cases := []UserCreateTestCase{
		{
			user: models.DataBaseUser{
				Login: login,
			},
			answer: nil,
		},
		{
			//корректные, отправлены второй раз, нарушают
			user: models.DataBaseUser{
				Login: login,
			},
			answer: fmt.Errorf("user is not unique"),
		},
	}

	for i, item := range cases {
		_, answer := dbTest.CreateUser(item.user)
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

type UserDeleteTestCase struct {
	user   models.DataBaseUser
	answer error
}

func TestUserDelete(t *testing.T) {
	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	dbTest.Open(c)

	id, _ := dbTest.CreateUser(models.DataBaseUser{
		Login: fmt.Sprint(time.Now()),
	})

	cases := []UserDeleteTestCase{
		{
			//данные есть в БД
			user: models.DataBaseUser{
				Id: id,
			},
			answer: nil,
		},
		{
			//данных неи в БД
			user: models.DataBaseUser{
				Id: id,
			},
			answer: nil,
		},
	}

	for i, item := range cases {
		answer := dbTest.DeleteUser(item.user)
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

type UserUpdateTestCase struct {
	user   models.DataBaseUser
	answer error
}

func TestUserUpdate(t *testing.T) {
	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	dbTest.Open(c)

	login := fmt.Sprint(time.Now())
	id, _ := dbTest.CreateUser(models.DataBaseUser{
		Login: login,
	})
	id2, _ := dbTest.CreateUser(models.DataBaseUser{
		Login: fmt.Sprint(time.Now()),
	})

	cases := []UserDeleteTestCase{
		{
			//данные есть
			user: models.DataBaseUser{
				Id:    id,
				Email: "new_email@mail.ru",
				Login: login,
			},
			answer: nil,
		},
		{
			//данных нет
			user: models.DataBaseUser{
				Email: "email@mail.ru",
				Login: "brand_new_login",
			},
			answer: fmt.Errorf("user can not be updated"),
		},
		{
			//данные нарушают уникальность
			user: models.DataBaseUser{
				Id:    id2,
				Email: "new_email@mail.ru",
				Login: login,
			},
			answer: fmt.Errorf("user can not be updated"),
		},
	}

	for i, item := range cases {
		answer := dbTest.UpdateUser(item.user)
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

func TestUpdateUserPassword(t *testing.T) {
	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	dbTest.Open(c)

	id, _ := dbTest.CreateUser(models.DataBaseUser{
		Login: fmt.Sprint(time.Now()),
	})

	cases := []UserUpdateTestCase{
		{
			user: models.DataBaseUser{
				EncryptedPassword: "password",
				Id:                id,
			},
			answer: nil,
		},
		{
			user: models.DataBaseUser{
				EncryptedPassword: "password",
			},
			answer: fmt.Errorf("user not found"),
		},
	}

	for i, item := range cases {
		answer := dbTest.UpdateUserPassword(item.user)
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

func TestUpdateUserDescription(t *testing.T) {
	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	dbTest.Open(c)

	id, _ := dbTest.CreateUser(models.DataBaseUser{
		Login: fmt.Sprint(time.Now()),
	})
	cases := []UserUpdateTestCase{
		{
			user: models.DataBaseUser{
				Id: id,
				About: struct {
					String string
					Valid  bool
				}{String: "about", Valid: true},
			},
			answer: nil,
		},
		{
			user: models.DataBaseUser{
				About: struct {
					String string
					Valid  bool
				}{String: "about", Valid: true},
			},
			answer: fmt.Errorf("user not found"),
		},
	}

	for i, item := range cases {
		answer := dbTest.UpdateUserDescription(item.user)
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

func TestUpdateUserAvatar(t *testing.T) {
	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	dbTest.Open(c)

	id, _ := dbTest.CreateUser(models.DataBaseUser{
		Login: fmt.Sprint(time.Now()),
	})
	cases := []UserUpdateTestCase{
		{
			user: models.DataBaseUser{
				Id: id,
				Avatar: struct {
					String string
					Valid  bool
				}{String: "picture", Valid: true},
			},
			answer: nil,
		},
		{
			user: models.DataBaseUser{
				Avatar: struct {
					String string
					Valid  bool
				}{String: "picture", Valid: true},
			},
			answer: fmt.Errorf("user not found"),
		},
	}

	for i, item := range cases {
		answer := dbTest.UpdateUserAvatar(item.user)
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

type UserSimpleSelectTestCase struct {
	user   models.DataBaseUser
	output models.User
	answer error
}

func TestGetUserById(t *testing.T) {
	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	dbTest.Open(c)

	id, _ := dbTest.CreateUser(models.DataBaseUser{
		Login: fmt.Sprint(time.Now()),
	})

	cases := []UserSimpleSelectTestCase{
		{
			user: models.DataBaseUser{
				Id: id,
			},
			output: models.User{},
			answer: nil,
		},
		{
			user:   models.DataBaseUser{},
			output: models.User{},
			answer: fmt.Errorf("user not found"),
		},
	}

	for i, item := range cases {
		_, answer := dbTest.GetUserById(item.user)
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

func TestGetUserByName(t *testing.T) {
	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	dbTest.Open(c)

	_, _ = dbTest.CreateUser(models.DataBaseUser{
		Login: "testing",
	})

	cases := []UserSimpleSelectTestCase{
		{
			user: models.DataBaseUser{
				Login: "testing",
			},
			output: models.User{},
			answer: nil,
		},
		{
			user:   models.DataBaseUser{},
			output: models.User{},
			answer: fmt.Errorf("user not found"),
		},
	}

	for i, item := range cases {
		_, answer := dbTest.GetUserByName(item.user)
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

func TestGetUserByEmail(t *testing.T) {
	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	dbTest.Open(c)

	_, _ = dbTest.CreateUser(models.DataBaseUser{
		Email: "email",
		Login: fmt.Sprint(time.Now()),
	})
	cases := []UserSimpleSelectTestCase{
		{
			user:   models.DataBaseUser{},
			output: models.User{},
			answer: nil,
		},
		{
			user: models.DataBaseUser{
				Email: "email",
			},
			output: models.User{},
			answer: nil,
		},
	}

	for i, item := range cases {
		_, answer := dbTest.GetUserByEmail(item.user)
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

func TestGetUserSubscriptions(t *testing.T) {
	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	dbTest.Open(c)

	cases := []UserSimpleSelectTestCase{
		{
			user:   models.DataBaseUser{},
			output: models.User{},
			answer: fmt.Errorf("user not found"),
		},
	}

	for i, item := range cases {
		_, answer := dbTest.GetUserSubscriptions(item.user)
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

func TestGetUserSubscribers(t *testing.T) {

	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	dbTest.Open(c)

	cases := []UserSimpleSelectTestCase{
		{
			user:   models.DataBaseUser{},
			output: models.User{},
			answer: fmt.Errorf("user not found"),
		},
	}

	for i, item := range cases {
		_, answer := dbTest.GetUserSubscribers(item.user)
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

type UserSelectTestCase struct {
	user   models.DataBaseUser
	output []*models.User
	answer error
}

func TestGetUserSubUsers(t *testing.T) {
	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	dbTest.Open(c)

	cases := []UserSelectTestCase{
		{
			user:   models.DataBaseUser{},
			output: nil,
			answer: nil,
		},
	}

	for i, item := range cases {
		_, answer := dbTest.GetUserSubUsers(item.user)
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

func TestGetUserSupUsers(t *testing.T) {

	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	dbTest.Open(c)

	cases := []UserSelectTestCase{
		{
			user:   models.DataBaseUser{},
			output: nil,
			answer: nil,
		},
	}

	for i, item := range cases {
		_, answer := dbTest.GetUserSupUsers(item.user)
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

func TestGetUserByLogin(t *testing.T) {
	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	dbTest.Open(c)

	_, _ = dbTest.CreateUser(models.DataBaseUser{
		Login: "loginsearch",
	})

	cases := []UserSelectTestCase{
		{
			user:   models.DataBaseUser{},
			output: nil,
			answer: nil,
		},
		{
			user: models.DataBaseUser{
				Login: "loginsearch",
			},
			output: nil,
			answer: nil,
		},
	}

	for i, item := range cases {
		_, answer := dbTest.GetUserByLogin(item.user, 0, 1)
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

func TestPgxDB_Follow(t *testing.T) {
	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	dbTest.Open(c)

	id, _ := dbTest.CreateUser(models.DataBaseUser{
		Login: fmt.Sprint(time.Now()),
	})

	sid, _ := dbTest.CreateUser(models.DataBaseUser{
		Login: fmt.Sprint(time.Now()),
	})

	dbTest.Follow(id, sid)
}

func TestPgxDB_Unfollow(t *testing.T) {
	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	dbTest.Open(c)

	id, _ := dbTest.CreateUser(models.DataBaseUser{
		Login: fmt.Sprint(time.Now()),
	})

	sid, _ := dbTest.CreateUser(models.DataBaseUser{
		Login: fmt.Sprint(time.Now()),
	})

	dbTest.Unfollow(id, sid)
}
