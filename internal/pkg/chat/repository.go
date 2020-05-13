package chat

import "2020_1_Color_noise/internal/models"

type IRepository interface {
	AddMessage(userSentId uint, userReceivedId uint, message string, sticker string) (*models.Message, error)
	GetUsers(userId uint, start int, limit int) ([]*models.User, error)
	GetMessages(userId uint, otherId uint, start int, limit int) ([]*models.Message, error)
}
