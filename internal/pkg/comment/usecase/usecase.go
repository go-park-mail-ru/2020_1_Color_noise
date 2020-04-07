package usecase

import (
	"2020_1_Color_noise/internal/models"
	"2020_1_Color_noise/internal/pkg/comment"
	. "2020_1_Color_noise/internal/pkg/error"
)

type Usecase struct {
	repo comment.IRepository
}

func NewUsecase(repo comment.IRepository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (cu *Usecase) Create(input *models.InputComment, userId uint) (uint, error) {
	comment := &models.Comment{
		UserId: userId,
		PinId:  uint(input.PinId),
		Text:   input.Text,
	}

	id, err := cu.repo.Create(comment)
	if err != nil {
		return 0, Wrapf(err, "Creating comment error, userId: %s", userId)
	}

	return id, nil
}

func (cu *Usecase) GetById(id uint) (*models.Comment, error) {
	comment, err := cu.repo.GetByID(id)
	if err != nil {
		return nil, Wrapf(err, "Getting comment by id error, pinId: %s", id)
	}

	return comment, nil
}

func (cu *Usecase) GetByPinId(id uint, start int, limit int) ([]*models.Comment, error) {
	comments, err := cu.repo.GetByPinID(id, start, limit)
	if err != nil {
		return nil, Wrapf(err, "Getting comment by id error, pinId: %s", id)
	}

	return comments, nil
}

func (cu *Usecase) GetByText(text string, start int, limit int) ([]*models.Comment, error) {
	comments, err := cu.repo.GetByText(text, start, limit)
	if err != nil {
		return nil, Wrap(err, "Getting comment by id error")
	}

	return comments, nil
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
