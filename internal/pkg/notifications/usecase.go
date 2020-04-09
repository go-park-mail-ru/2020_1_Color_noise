package notifications

import (
	"2020_1_Color_noise/internal/models"
)

type IUsecase interface {
	GetNotifications(id uint, start int, limit int) ([]*models.Notification, error)
}
