package repository

import (
	"2020_1_Color_noise/internal/models"
	"2020_1_Color_noise/internal/pkg/database"
	. "2020_1_Color_noise/internal/pkg/error"
	"sync"
)

type Repository struct {
	db database.DBInterface
	mu *sync.Mutex
}

func NewRepo(d database.DBInterface) *Repository {
	return &Repository{
		db: d,
		mu: &sync.Mutex{},
	}
}


func (lr *Repository) GetMainList(start int, limit int) ([]*models.Pin, error) {
	p := models.DataBaseUser{}
	result, err := lr.db.GetMainFeed(p, limit, start)
	if err != nil {
		return result, PinNotFound.Newf("Pins not found, err:", err)
	}

	return result, nil
}

func (lr *Repository) GetSubList(id uint, start int, limit int) ([]*models.Pin, error) {
	p := models.DataBaseUser{Id: id}
	result, err := lr.db.GetSubFeed(p, limit, start)
	if err != nil {
		return nil, nil
	}

	return result, nil
}

func (lr *Repository) GetRecommendationList(id uint, start int, limit int) ([]*models.Pin, error) {
	p := models.DataBaseUser{}
	result, err := lr.db.GetRecFeed(p, limit, start)
	if err != nil {
		return result, PinNotFound.Newf("Pins not found, user id = %d", id)
	}

	return result, nil
}
