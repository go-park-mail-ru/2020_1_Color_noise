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

type UserCase struct {
	u      models.User
	answer error
}

type FollowCase struct {
	u      models.User
	s      models.User
	answer error
}

func TestRepository_Create(t *testing.T) {
	db.Open()

	login := fmt.Sprint(time.Now())
	cases := []UserCase{
		{
			u: models.User{
			},
			answer: fmt.Errorf("Repo: Error in during creating"),
		},{
			u: models.User{
				Login: fmt.Sprint(time.Now()),
				Email: "email",
			},
			answer: fmt.Errorf("Repo: Error in during creating"),
		},
		{
			u: models.User{
				Login: login,
				Email: login + "@mail.com",
			},
			answer: nil,
		},
	}

	for i, item := range cases {
		_, answer := repo.Create(&item.u)
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
	id, _ := repo.Create(&models.User{
		Login: fmt.Sprint(time.Now()),
		Email: "email@mail.com",
	},)

	cases := []UserCase{
		{
			u: models.User{
				Id:id,
			},
			answer:nil,
		},{
			u: models.User{
			},
			answer:nil,
		},
	}

	for i, item := range cases {
		answer := repo.Delete(item.u.Id)
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

func TestRepository_UpdateProfile(t *testing.T) {

	db.Open()

	login := fmt.Sprint(time.Now())
	cases := []UserCase{
		{
			u: models.User{
			},
			answer: fmt.Errorf("Repo: Error in during updating profile"),
		},{
			u: models.User{
				Login: fmt.Sprint(time.Now()),
				Email: "email@email.com",
			},
			answer:fmt.Errorf("Repo: Error in during updating profile"),
		},
		{
			u: models.User{
				Login: login,
				Email: login + "@mail.com",
			},
			answer: nil,
		},
	}

	for i, item := range cases {
		id, _ := repo.Create(&item.u)
		item.u.Id = id
		answer := repo.UpdateProfile(item.u.Id, item.u.Email, item.u.Login)
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

func TestRepository_UpdateAvatar(t *testing.T) {
	db.Open()

	db.Open()

	login := fmt.Sprint(time.Now())
	cases := []UserCase{
		{
			u: models.User{
			},
			answer: fmt.Errorf("User to update not found, id: 0"),
		},{
			u: models.User{
				Login: fmt.Sprint(time.Now()),
				Email: "email@email.com",
			},
			answer: fmt.Errorf("User to update not found, id: 0"),
		},
		{
			u: models.User{
				Login: login,
				Email: login + "@mail.com",
			},
			answer: nil,
		},
	}

	for i, item := range cases {
		id, _ := repo.Create(&item.u)
		item.u.Id = id
		answer := repo.UpdateAvatar(item.u.Id,  item.u.Avatar)
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

func TestRepository_UpdateDescription(t *testing.T) {

	db.Open()

	db.Open()

	login := fmt.Sprint(time.Now())
	cases := []UserCase{
		{
			u: models.User{
			},
			answer: fmt.Errorf("User to update not found, id: 0"),
		},{
			u: models.User{
				Login: fmt.Sprint(time.Now()),
				Email: "email@email.com",
			},
			answer: fmt.Errorf("User to update not found, id: 0"),
		},
		{
			u: models.User{
				Login: login,
				Email: login + "@mail.com",
			},
			answer: nil,
		},
	}

	for i, item := range cases {
		id, _ := repo.Create(&item.u)
		item.u.Id = id
		answer := repo.UpdateDescription(item.u.Id,  &item.u.About)
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

func TestRepository_UpdatePassword(t *testing.T) {

	db.Open()

	db.Open()

	login := fmt.Sprint(time.Now())
	cases := []UserCase{
		{
			u: models.User{
			},
			answer: fmt.Errorf("User to update not found, id: 0"),
		},{
			u: models.User{
				Login: fmt.Sprint(time.Now()),
				Email: "email@email.com",
			},
			answer: fmt.Errorf("User to update not found, id: 0"),
		},
		{
			u: models.User{
				Login: login,
				Email: login + "@mail.com",
			},
			answer: nil,
		},
	}

	for i, item := range cases {
		id, _ := repo.Create(&item.u)
		item.u.Id = id
		answer := repo.UpdatePassword(item.u.Id,  item.u.EncryptedPassword)
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



func TestRepository_Follow(t *testing.T) {

	db.Open()

	login := fmt.Sprint(time.Now())
	cases := []FollowCase{
		{
			u: models.User{
				Login:login,
				Email: login + "@mail.com",
			},
			s: models.User{
				Login:login + "2",
				Email: login + "2@mail.com",
			},
			answer: fmt.Errorf("User to get not found, id: 0"),
		},{
			u: models.User{
			},
			s: models.User{},
			answer: fmt.Errorf("User to get not found, id: 0"),
		},
	}

	for i, item := range cases {
		_, _ = repo.Create(&item.u)
		_, _ = repo.Create(&item.s)
		answer := repo.Follow(item.u.Id, item.s.Id)
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

func TestRepository_Unfollow(t *testing.T) {


	db.Open()


	login := fmt.Sprint(time.Now())
	cases := []FollowCase{
		{
			u: models.User{
				Login:login,
				Email: login + "@mail.com",
			},
			s: models.User{
				Login:login + "2",
				Email: login + "2@mail.com",
			},
			answer: fmt.Errorf("User to get not found, id: 0"),
		},{
			u: models.User{
			},
			s: models.User{},
			answer: fmt.Errorf("User to get not found, id: 0"),
		},
	}

	for i, item := range cases {
		id, _ := repo.Create(&item.u)
		sid, _ := repo.Create(&item.u)
		_ = repo.Follow(id, sid)
		answer := repo.Unfollow(id, sid)
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

	login := fmt.Sprint(time.Now())
	cases := []UserCase{
		{
			u: models.User{
			},
			answer: fmt.Errorf("User to get not found, id: 0"),
		},{
			u: models.User{
				Login: fmt.Sprint(time.Now()),
				Email: "email",
			},
			answer: fmt.Errorf("User to get not found, id: 0"),
		},
		{
			u: models.User{
				Login: login,
				Email: "email@mail.com",
			},
			answer: fmt.Errorf("User to get not found, id: 0"),
		},
		{
			u: models.User{
				Login: login,
				Email: "email@mail.com",
			},
			answer: fmt.Errorf("User to get not found, id: 0"),
		},
	}

	for i, item := range cases {
		id, _ := repo.Create(&item.u)
		item.u.Id = id
		_, answer := repo.GetByID(item.u.Id)
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

func TestRepository_GetByLogin(t *testing.T) {

	db.Open()

	login := fmt.Sprint(time.Now())
	cases := []UserCase{
		{
			u: models.User{
			},
			answer: fmt.Errorf("User is not found"),
		},{
			u: models.User{
				Login: fmt.Sprint(time.Now()),
				Email: "email",
			},
			answer: fmt.Errorf("User is not found"),
		},
		{
			u: models.User{
				Login: login,
				Email: "email@mail.com",
			},
			answer: fmt.Errorf("User is not found"),
		},
		{
			u: models.User{
				Login: login,
				Email: "email@mail.com",
			},
			answer: fmt.Errorf("User is not found"),
		},
	}

	for i, item := range cases {
		_, _ = repo.Create(&item.u)
		_, answer := repo.GetByLogin(item.u.Login)
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

func TestRepository_GetSubscribers(t *testing.T) {

}

func TestRepository_GetSubscriptions(t *testing.T) {

}

func TestRepository_Search(t *testing.T) {

	db.Open()

	login := fmt.Sprint(time.Now())
	cases := []UserCase{
		{
			u: models.User{
			},
		},{
			u: models.User{
				Login: fmt.Sprint(time.Now()),
				Email: "email@email.com",
			},
		},
		{
			u: models.User{
				Login: login,
				Email: login + "@mail.com",
			},
		},
	}

	for i, item := range cases {
		id, _ := repo.Create(&item.u)
		item.u.Id = id
		_, answer := repo.Search(item.u.Login, 0, 2)
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
