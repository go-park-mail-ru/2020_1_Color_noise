package comment

import (
	"2020_1_Color_noise/internal/models"
)

type IRepository interface {
	Create(comment *models.Comment) (uint, error)
	GetByID(id uint) (*models.Comment, error)
	GetByPinID(pinId uint, start int, limit int) ([]*models.Comment, error)
	GetByText(text string, start int, limit int) ([]*models.Comment, error)
	Delete(id uint) error
}
