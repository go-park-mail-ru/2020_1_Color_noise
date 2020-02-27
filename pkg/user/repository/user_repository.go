package repository

import (
	"pinterest/pkg/models"
)

type UserRepository struct {
	data []*models.User
}

func NewUserRepo() *UserRepository {
	return &UserRepository{
		data: make([]*models.User, 0),
	}
}

func (ur *UserRepository) Add(user *models.User) (uint, error) {
	user.Id = uint(len(ur.data) + 1)
	ur.data = append(ur.data, user)
	return user.Id, nil
}

func (ur *UserRepository) GetByID(id uint) ([]*models.User, error) {
	for _, user := range ur.data {
		if user.Id == id {
			return []*models.User{user}, nil
		}
	}
	return []*models.User{}, nil
}

func (ur *UserRepository) GetByLogin(login string) (*models.User, error) {
	for _, user := range ur.data {
		if user.Login == login {
			//fmt.Println(login, user.Login)
			return user, nil
		}
	}
	return nil, nil
}

func (ur *UserRepository) GetByEmail(email string) (*models.User, error) {
	for _, user := range ur.data {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, nil
}

func (ur *UserRepository) Update(id uint, user *models.User) (bool, error) {
	for i, user := range ur.data {
		if user.Id == id {
			ur.data[i] = user
			return true, nil
		}
	}
	return false, nil
}

func (ur *UserRepository) Delete(id uint) (bool, error) {
	for i, user := range ur.data {
		if user.Id == id {
			newData := ur.data[:i]
			for j := i + 1; j < len(ur.data); j++ {
				newData = append(newData, ur.data[j])
			}
			ur.data = newData
			return true, nil
		}
	}
	return false, nil
}