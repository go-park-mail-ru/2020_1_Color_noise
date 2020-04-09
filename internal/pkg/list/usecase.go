package list

import (
	"2020_1_Color_noise/internal/models"
)

type IUsecase interface {
	GetMainList(start int, limit int) ([]*models.Pin, error)
	GetSubList(id uint, start int, limit int) ([]*models.Pin, error)
	GetRecommendationList(id uint, start int, limit int) ([]*models.Pin, error)
}
