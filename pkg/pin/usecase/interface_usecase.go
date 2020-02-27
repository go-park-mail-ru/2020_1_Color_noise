package usecase

import (
	"pinterest/pkg/models"
)

type IPinUsecase interface {
	Add(pin *models.Pin) (uint, error)
	Get(id uint) (*models.Pin, error)
	SaveImage(pin *models.Pin) (error)
}



