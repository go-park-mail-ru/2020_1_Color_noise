package pin

import (
	"2020_1_Color_noise/internal/models"
)

type IRepository interface {
	//Create(pin *models.Pin) (uint, error)
	CreatePin(pin *models.Pin) (uint, error)
	SaveImage(pin *models.Pin) (uint, error)
	GetByID(id uint) (*models.Pin, error)
	GetByUserID(userId uint, start int, limit int) ([]*models.Pin, error)
	GetByName(name string, start int, limit int, date string, desc bool, most string) ([]*models.Pin, error)
	Update(pin *models.Pin) error
	Delete(pinId uint, userId uint) error
	AddTags(pinId uint, tags []string) error
}
