package board

import (
	"2020_1_Color_noise/internal/models"
)

type IUsecase interface {
	Create(input *models.InputBoard, userId uint) (uint, error)
	GetById(id uint) (*models.Board, error)
	GetByNameId(id uint) (*models.Board, error)
	GetByUserId(id uint, start int, limit int) ([]*models.Board, error)
	GetByName(name string, start int, limit int) ([]*models.Board, error)
	Update(input *models.InputBoard, id uint, userId uint) error
	Delete(id uint, userId uint) error
}
