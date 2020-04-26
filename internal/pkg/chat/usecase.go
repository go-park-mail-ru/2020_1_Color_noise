package chat

import (
	"2020_1_Color_noise/internal/models"
)

type IUsecase interface {
	AddMessage(userSentId uint, input *models.InputMessage) (*models.Message, error)
	GetUsers(userId uint, start int, limit int) ([]*models.User, error)
	GetMessages(userId uint, otherId uint, start int, limit int) ([]*models.Message, error)
}
