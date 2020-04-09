package usecase

import (
	"2020_1_Color_noise/internal/models"
	. "2020_1_Color_noise/internal/pkg/error"
	"2020_1_Color_noise/internal/pkg/notifications"
)

type Usecase struct {
	repo notifications.IRepository
}

func NewUsecase(repo notifications.IRepository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (nu *Usecase) GetNotifications(id uint, start int, limit int) ([]*models.Notification, error) {
	notifications, err := nu.repo.GetNotifications(id, start, limit)
	if err != nil {
		return nil, Wrapf(err, "GetNotifications by id error, pinId: %d", id)
	}

	return notifications, nil
}
