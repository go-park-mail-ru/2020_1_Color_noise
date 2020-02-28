package usecase

import (
	"pinterest/pkg/models"
)

type IUserUsecase interface {
	Add(user *models.User) (uint, error)
	GetById(id uint) (*models.User, error)
	GetByLogin(login string) (*models.User, error)
	Update(user *models.User) error
	ChangePassword(id uint, newPassword string) error
	Delete(id uint) error
	ComparePassword(user *models.User, password string) bool
	CheckLogin(user *models.User) error
	CheckEmail(user *models.User) error
	SaveAvatar(user *models.User) (string, error)
}



