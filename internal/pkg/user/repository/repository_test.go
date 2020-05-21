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
	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	db.Open(c)

	login := fmt.Sprint(time.Now())
	cases := []UserCase{
		{
			u:      models.User{},
			answer: fmt.Errorf("Repo: Error in during creating"),
		}, {
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

	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	db.Open(c)
	user, err := repo.Create(&models.User{
		Login: fmt.Sprint(time.Now()),
		Email: fmt.Sprint(time.Now()),
	})

	id := user.Id
	cases := []UserCase{
		{
			u: models.User{
				Id: id,
			},
			answer: nil,
		}, {
			u:      models.User{},
			answer: nil,
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

	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	db.Open(c)

	login := fmt.Sprint(time.Now())
	cases := []UserCase{
		{
			u: models.User{
				Login: login,
				Email: login + "@mail.com",
			},
			answer: nil,
		},
	}

	for i, item := range cases {
		user, _ := repo.Create(&item.u)
		item.u.Id = user.Id
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
	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	db.Open(c)

	login := fmt.Sprint(time.Now())
	cases := []UserCase{
		{
			u: models.User{
				Login: login,
				Email: login + "@mail.com",
			},
			answer: nil,
		},
	}

	for i, item := range cases {
		user, _ := repo.Create(&item.u)
		item.u.Id = user.Id
		answer := repo.UpdateAvatar(item.u.Id, item.u.Avatar)
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
	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	db.Open(c)

	login := fmt.Sprint(time.Now())
	cases := []UserCase{
		{
			u: models.User{
				Login: login,
				Email: login + "@mail.com",
			},
			answer: nil,
		},
	}

	for i, item := range cases {
		user, _ := repo.Create(&item.u)
		item.u.Id = user.Id
		answer := repo.UpdateDescription(item.u.Id, &item.u.About)
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

	c, err := config.GetTestConfing()
	if err != nil {
		t.SkipNow()
	}
	db.Open(c)

	login := fmt.Sprint(time.Now())
	cases := []UserCase{
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
		item.u.Id = id.Id
		answer := repo.UpdatePassword(item.u.Id, item.u.EncryptedPassword)
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

		c, err := config.GetTestConfing()
		if err != nil {
			t.SkipNow()
		}
		db.Open(c)

		cases := []FollowCase{
			{
				u: models.User{
				},
				s: models.User{},
				answer: fmt.Errorf("User not found, id: 0"),
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
		/*
		c, err := config.GetTestConfing()
		if err != nil {
			t.SkipNow()
		}
		db.Open(c)


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
				answer: fmt.Errorf("User not found, id: 0"),
			},
		}

		for i, item := range cases {
			id, _ := repo.Create(&item.u)
			sid, _ := repo.Create(&item.u)
			_ = repo.Follow(id.Id, sid.Id)
			answer := repo.Unfollow(id.Id, sid.Id)
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

		 */



}

func TestRepository_GetByID(t *testing.T) {
	/*
		c, err := config.GetTestConfing()
		if err != nil {
			t.SkipNow()
		}
		db.Open(c)

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
		}

		for i, item := range cases {
			id, _ := repo.Create(&item.u)
			item.u.Id = id.Id
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
	
	 */

}

func TestRepository_GetByLogin(t *testing.T) {


		c, err := config.GetTestConfing()
		if err != nil {
			t.SkipNow()
		}
		db.Open(c)

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


		c, err := config.GetTestConfing()
		if err != nil {
			t.SkipNow()
		}
		db.Open(c)

		login := fmt.Sprint(time.Now())
		cases := []UserCase{
			{
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
			item.u.Id = id.Id
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
