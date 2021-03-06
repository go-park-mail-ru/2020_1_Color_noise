package repository

import (
	"2020_1_Color_noise/internal/models"
	"2020_1_Color_noise/internal/pkg/database"
	. "2020_1_Color_noise/internal/pkg/error"
)

type Repository struct {
	db database.DBInterface
}

func NewRepository(db database.DBInterface) *Repository {
	return &Repository{db: db}
}

func (rp *Repository) AddMessage(userSentId uint, userReceivedId uint, message string, sticker string) (*models.Message, error) {
	//создать сообщение

	if message == "" && sticker == "" {
		return &models.Message{}, StupidUser.Newf("User is doing it wrong, id: %d", userSentId)
	}

	msg, err := rp.db.AddMessage(int(userSentId), int(userReceivedId), message, sticker)
	if err != nil {
		return &models.Message{}, UserNotFound.Newf("User not found, id: %d", userSentId)
	}

	return msg, nil
}

func (rp *Repository) GetUsers(userId uint, start int, limit int) ([]*models.User, error) {
	//получить чаты
	users, err := rp.db.GetUsers(userId, start, limit)
	if err != nil {
		return nil, UserNotFound.Newf("User not found, id: %v", userId)
	}
	return users, nil
}

func (rp *Repository) GetMessages(userId uint, otherId uint, start int, limit int) ([]*models.Message, error) {
	//получить сообщения
	msg, err := rp.db.GetMessages(userId, otherId, start, limit)
	if err != nil {
		return nil, UserNotFound.Newf("User not found, id: %v", userId)
	}
	return msg, nil
}
