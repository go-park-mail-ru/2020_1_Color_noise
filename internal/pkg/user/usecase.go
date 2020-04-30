package user

import (
	"2020_1_Color_noise/internal/models"
)

type IUsecase interface {
	Create(input *models.SignUpInput) (*models.User, error)
	GetById(id uint) (*models.User, error)
	UpdateProfile(id uint, input *models.UpdateProfileInput) error
	UpdateDescription(id uint, input *models.UpdateDescriptionInput) error
	UpdatePassword(id uint, input *models.UpdatePasswordInput) error
	UpdateAvatar(id uint, buffer []byte) (string, error)
	Delete(id uint) error
	Follow(id uint, subId uint) error
	Unfollow(id uint, subId uint) error
	GetByLogin(login string) (*models.User, error)
	Search(login string, start int, limit int) ([]*models.User, error)
	GetSubscribers(id uint, start int, limit int) ([]*models.User, error)
	GetSubscriptions(id uint, start int, limit int) ([]*models.User, error)
}
