package usecase

import (
	"2020_1_Color_noise/internal/models"
	. "2020_1_Color_noise/internal/pkg/error"
	"2020_1_Color_noise/internal/pkg/list"
)

type Usecase struct {
	repo list.IRepository
}

func NewUsecase(repo list.IRepository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (lu *Usecase) GetMainList(start int, limit int) ([]*models.Pin, error) {
	pins, err := lu.repo.GetMainList(start, limit)
	if err != nil {
		return nil, Wrap(err, "GetMainList error")
	}

	return pins, nil
}

func (lu *Usecase) GetSubList(id uint, start int, limit int) ([]*models.Pin, error) {
	pins, err := lu.repo.GetSubList(id, start, limit)
	if err != nil {
		return nil, Wrapf(err, "GetSubList by id error, pinId: %d", id)
	}

	return pins, nil
}

func (lu *Usecase) GetRecommendationList(id uint, start int, limit int) ([]*models.Pin, error) {
	pins, err := lu.repo.GetRecommendationList(id, start, limit)
	if err != nil {
		return nil, Wrapf(err, "GetRecommendationList by id error, pinId: %d", id)
	}

	return pins, nil
}
