package user

import (
	"pinterest/internal/models"
)

type IRepository interface {
	Add(user *models.User) (uint, error)
	GetByID(id uint) (*models.User, error)
	GetByLogin(login string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	Update(newUser *models.User) error
	Delete(id uint) error
}
