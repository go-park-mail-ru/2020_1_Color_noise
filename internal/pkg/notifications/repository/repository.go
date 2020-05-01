package repository

import (
	"2020_1_Color_noise/internal/models"
	"2020_1_Color_noise/internal/pkg/database"
	"fmt"
)

type Repository struct {
	db database.DBInterface
}

func NewRepo(d database.DBInterface) *Repository {
	return &Repository{
		db: d,
	}
}

func (nr *Repository) GetNotifications(id uint, start int, limit int) ([]*models.Notification, error) {
	user := models.DataBaseUser{Id: id}
	nts, err := nr.db.GetNotifications(user, start, limit)
	if err != nil {
		return nil, fmt.Errorf("no notifications")
	}
	return nts, nil
}
