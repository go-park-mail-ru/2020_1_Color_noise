package repository

import (
	"pinterest/internal/models"
	. "pinterest/internal/pkg/error"
)

type Repository struct {
	data []*models.User
}

func NewRepo() *Repository {
	return &Repository{
		data: make([]*models.User, 0),
	}
}

func (ur *Repository) Add(user *models.User) (uint, error) {
	user.Id = uint(len(ur.data) + 1)
	ur.data = append(ur.data, user)

	return user.Id, nil
}

func (ur *Repository) GetByID(id uint) (*models.User, error) {
	for _, user := range ur.data {
		if user.Id == id {
			return user, nil
		}
	}

	return nil, NotFound.Newf("User to get not found, id: %d", id)
}

func (ur *Repository) GetByLogin(login string) (*models.User, error) {
	for _, user := range ur.data {
		if user.Login == login {
			return user, nil
		}
	}

	return nil, NotFound.Newf("User to get not found, login: %s", login)
}

func (ur *Repository) GetByEmail(email string) (*models.User, error) {
	for _, user := range ur.data {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, NotFound.Newf("User to get not found, id: %s", email)
}

func (ur *Repository) Update(newUser *models.User) error {
	for i, user := range ur.data {
		if user.Id == newUser.Id {
			ur.data[i] = newUser
			return nil
		}
	}

	return NotFound.Newf("User to update not found, id: %d", newUser.Id)
}

func (ur *Repository) Delete(id uint) error {
	for i, user := range ur.data {
		if user.Id == id {
			newData := ur.data[:i]
			for j := i + 1; j < len(ur.data); j++ {
				newData = append(newData, ur.data[j])
			}
			ur.data = newData
			return nil
		}
	}

	return NotFound.Newf("User to delete not found, id: %d", id)
}