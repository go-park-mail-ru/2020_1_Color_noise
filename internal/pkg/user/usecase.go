package user

import (
	"2020_1_Color_noise/internal/models"
	"bytes"
)

type IUsecase interface {
	Create(input *models.SignUpInput) (uint, error)
	GetById(id uint) (*models.User, error)
	GetByLogin(login string) (*models.User, error)
	UpdateProfile(id uint, input *models.UpdateProfileInput) error
	UpdateDescription(id uint, input *models.UpdateDescriptionInput) error
	UpdatePassword(id uint, input *models.UpdatePasswordInput) error
	UpdateAvatar(id uint, buffer *bytes.Buffer) (string, error)
	Delete(id uint) error
	ComparePassword(user *models.User, password string) error
}



