package user

import (
	"2020_1_Color_noise/internal/models"
	"bytes"
)

type IUsecase interface {
	Create(input *models.SignUpInput) (uint, error)
	GetById(id uint) (*models.User, error)
	GetByLogin(login string) (*models.User, error)
	Update(id uint, input *models.UpdateProfileInput) error
	UpdatePassword(id uint, password string) error
	Delete(id uint) error
	ComparePassword(user *models.User, password string) error
	UpdateAvatar(id uint, buffer *bytes.Buffer) (string, error)
}



