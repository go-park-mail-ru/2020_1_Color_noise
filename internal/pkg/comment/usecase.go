package comment

import (
	"2020_1_Color_noise/internal/models"
)

type IUsecase interface {
	Create(input *models.InputComment, userId uint) (uint, error)
	GetById(id uint) (*models.Comment, error)
	GetByPinId(id uint, start int, limit int) ([]*models.Comment, error)
	GetByText(text string, start int, limit int) ([]*models.Comment, error)
}



