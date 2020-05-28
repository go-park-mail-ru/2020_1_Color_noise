package user

import (
	"2020_1_Color_noise/internal/models"
)

type IRepository interface {
	Create(user *models.User) (*models.User, error)
	GetByID(id uint) (*models.User, error)
	//GetByEmail(email string) (*models.User, error)
	UpdateProfile(id uint, email string, login string) error
	UpdateDescription(id uint, description *string) error
	UpdatePassword(id uint, encryptredPassword string) error
	UpdateAvatar(id uint, path string) error
	Follow(id uint, subId uint) error
	//TODO
	IsFollowed(id uint, subId uint) (bool, error)
	Unfollow(id uint, subId uint) error
	GetByLogin(login string) (*models.User, error)
	Search(login string, start int, limit int) ([]*models.User, error)
	GetSubscribers(id uint, start int, limit int) ([]*models.User, error)
	GetSubscriptions(id uint, start int, limit int) ([]*models.User, error)
	Delete(id uint) error
	//TODO
	UpdatePreferences(userId uint, preferences []string) error
}
