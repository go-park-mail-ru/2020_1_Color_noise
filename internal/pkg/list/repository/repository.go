package repository

import (
	"2020_1_Color_noise/internal/models"
	"2020_1_Color_noise/internal/pkg/database"
	. "2020_1_Color_noise/internal/pkg/error"
)

type Repository struct {
	db database.DBInterface
}

func NewRepo(d database.DBInterface) *Repository {
	return &Repository{
		db: d,
	}
}

func (lr *Repository) GetMainList(start int, limit int) ([]*models.Pin, error) {
	p := models.DataBaseUser{}
	result, err := lr.db.GetMainFeed(p, start, limit)
	if err != nil {
		return result, PinNotFound.Newf("Pins not found, err: %v", err)
	}

	return result, nil
}

func (lr *Repository) GetSubList(id uint, start int, limit int) ([]*models.Pin, error) {
	p := models.DataBaseUser{Id: id}
	result, err := lr.db.GetSubFeed(p, start, limit)
	if err != nil {
		return nil, nil
	}

	return result, nil
}

func (lr *Repository) GetRecommendationList(id uint, start int, limit int) ([]*models.Pin, error) {
	p := models.DataBaseUser{Id: id}
	user, err := lr.db.GetUserById(p)
	if err != nil {
		return nil, UserNotFound.Newf("User not found, user id = %d", user.Id)
	}


	pins, ok := lr.db.GetPinsByTag(user.Tags, start, limit)

	return pins, ok
}
