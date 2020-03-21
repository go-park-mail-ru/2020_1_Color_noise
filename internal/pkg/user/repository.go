package user

import (
	"2020_1_Color_noise/internal/models"
)

type IRepository interface {
	Create(user *models.User) (uint, error)
	GetByID(id uint) (*models.User, error)
	GetByLogin(login string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	UpdateProfile(id uint, email string, login string) error
	UpdateDescription(id uint, description *string) error
	UpdatePassword(id uint, encryptredPassword string) error
	UpdateAvatar(id uint, path string) error
	Delete(id uint) error
}
