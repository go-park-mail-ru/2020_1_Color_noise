package board

import (
	"2020_1_Color_noise/internal/models"
)

type IRepository interface {
	Create(pin *models.Board) (uint, error)
	GetByID(id uint) (*models.Board, error)
	GetByUserID(userId uint, start int, limit int) ([]*models.Board, error)
	GetByName(name string, start int, limit int) ([]*models.Board, error)
	//Update(board *models.Board) error
	Delete(id uint) error
}
