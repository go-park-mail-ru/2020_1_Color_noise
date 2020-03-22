package usecase

import (
	"2020_1_Color_noise/internal/models"
	"2020_1_Color_noise/internal/pkg/board"
	. "2020_1_Color_noise/internal/pkg/error"
)

type Usecase struct {
	repo  board.IRepository
}

func NewUsecase(repo board.IRepository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (bu *Usecase) Create(input *models.InputBoard, userId uint) (uint, error) {
	board := &models.Board{
		UserId:      userId,
		Name:        input.Name,
		Description: input.Description,
	}

	id, err := bu.repo.Create(board)
	if err != nil {
		return 0, Wrapf(err, "Creating board error, userId: %s", userId)
	}

	return id, nil
}

func (bu *Usecase) GetById(id uint) (*models.Board, error) {
	board, err := bu.repo.GetByID(id)
	if err != nil {
		return nil, Wrapf(err, "Getting board by id error, pinId: %s", id)
	}

	return board, nil
}

func (bu *Usecase) GetByUserId(id uint, start int, limit int) ([]*models.Board, error) {
	boards, err := bu.repo.GetByUserID(id, start, limit)
	if err != nil {
		return nil, Wrapf(err, "Getting board by id error, pinId: %s", id)
	}

	return boards, nil
}

func (bu *Usecase) GetByName(name string, start int, limit int) ([]*models.Board, error) {
	boards, err := bu.repo.GetByName(name, start, limit)
	if err != nil {
		return nil, Wrapf(err, "Getting board by id error, name: %s", name)
	}

	return boards, nil
}

/*
func (pu *PinUsecase) Update(id uint, pin *models.Pin) error {
	pins, err := pu.pinRepo.GetByID(id)
	if err != nil {
		return err
	}

	if len(pins) != 1 {
		return fmt.Errorf("Pin not found")
	}

	if pin.Name != "" {
		pins[0].Name = pin.Name
	}

	if pin.Description != "" {
		pins[0].Description = pin.Description
	}

	_, err = pu.pinRepo.Update(id, pins[0])
	if err != nil {
		return err
	}
	return nil
}



func (pu *PinUsecase) Delete(id uint) error {
	status, err := pu.pinRepo.Delete(id)
	if err != nil {
		return err
	}
	if !status {
		return fmt.Errorf("Pin not found")
	}
	return nil
}

*/