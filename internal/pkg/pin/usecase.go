package pin

import (
	"pinterest/internal/models"
)

type IUsecase interface {
	Add(input *models.InputPin) (uint, error)
	GetById(id uint) (*models.Pin, error)
	GetByUserId(id uint) ([]*models.Pin, error)
	GetByName(name string) ([]*models.Pin, error)
}



