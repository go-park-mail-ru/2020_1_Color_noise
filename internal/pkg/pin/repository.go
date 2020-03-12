package pin

import (
	"2020_1_Color_noise/internal/models"
)

type IRepository interface {
	Add(pin *models.Pin) (uint, error)
	GetByID(id uint) (*models.Pin, error)
	GetByUserID(userId uint) ([]*models.Pin, error)
	GetByName(name string) ([]*models.Pin, error)
	Update(pin *models.Pin) error
	Delete(id uint) error
}
