package pin

import (
	"2020_1_Color_noise/internal/models"
)

type IUsecase interface {
	Create(input *models.InputPin, userId uint) (uint, error)
	GetById(id uint) (*models.Pin, error)
	GetByUserId(id uint, start int, limit int) ([]*models.Pin, error)
	GetByName(name string, start int, limit int) ([]*models.Pin, error)
	Update(input *models.UpdatePin, pinId uint, userId uint) error
	Delete(pinId uint, userId uint) error
}



