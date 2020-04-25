package usecase

import (
	"2020_1_Color_noise/internal/models"
	"2020_1_Color_noise/internal/pkg/chat"
	. "2020_1_Color_noise/internal/pkg/error"
)

type Usecase struct {
	repo chat.IRepository
}

func NewUsecase(repo chat.IRepository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) AddMessage(userSentId uint, input *models.InputMessage) (*models.Message, error) {
	message, err := u.repo.AddMessage(userSentId, input.UserRecviredId, input.Message)
	if err != nil {
		return nil, NoType.Wrap(err, "AddMessge error")
	}
	message.RecUser = &models.User{
		Id: input.UserRecviredId,
	}

	return message, nil
}

func (u *Usecase) GetUsers(userId uint, start int, limit int) ([]*models.User, error) {
	users, err := u.repo.GetUsers(userId, start, limit)
	if err != nil {
		return nil, NoType.Wrap(err, "GetUsers error")
	}

	return users, nil
}

func (u *Usecase) GetMessages(userId uint, start int, limit int) ([]*models.Message, error) {
	messages, err := u.repo.GetMessages(userId, start, limit)
	if err != nil {
		return nil, NoType.Wrap(err, "GetMessages error")
	}

	return messages, nil
}

