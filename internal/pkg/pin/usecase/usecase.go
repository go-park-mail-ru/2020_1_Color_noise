package usecase

import (
	"encoding/base64"
	"pinterest/internal/models"
	. "pinterest/internal/pkg/error"
	"pinterest/internal/pkg/image"
	"pinterest/internal/pkg/pin"

	//"golang.org/x/crypto/bcrypt"
)

type Usecase struct {
	repo  pin.IRepository
}

func NewUsecase(repo pin.IRepository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (pu *Usecase) Create(input *models.InputPin, userId uint) (uint, error) {
	length := base64.StdEncoding.EncodedLen(len(input.Image))
	if length > 10000000 {
		return 0, Wrapf(TooMuchSize.New("Too much size image"), "Creating pin error, userId: %s", userId)
	}

	buffer := make([]byte, length)

	l, err := base64.StdEncoding.Decode(buffer, input.Image)
	if err != nil {
		err = Wrap(err, "Decoding base64")
		return 0, Wrapf(err, "Creating pin error, userId: %s", userId)
	}
	buffer = buffer[:l]

	name, err := image.SaveImage(&buffer)
	if err != nil {
		return 0, Wrapf(err, "Creating pin error, userId: %s", userId)
	}

	pin := &models.Pin{
		UserId:      userId,
		Name:        input.Name,
		Description: input.Description,
		Image:       name,
	}

	id, err := pu.repo.Add(pin)
	if err != nil {
		return 0, Wrapf(err, "Creating pin error, userId: %s", userId)
	}

	return id, nil
}

func (pu *Usecase) GetById(id uint) (*models.Pin, error) {
	pin, err := pu.repo.GetByID(id)
	if err != nil {
		return nil, Wrapf(err, "Getting pin by id error, pinId: %s", id)
	}

	return pin, nil
}

func (pu *Usecase) GetByUserId(id uint) ([]*models.Pin, error) {
	pins, err := pu.repo.GetByUserID(id)
	if err != nil {
		return nil, Wrapf(err, "Getting pin by id error, pinId: %s", id)
	}

	return pins, nil
}

func (pu *Usecase) GetByName(name string) ([]*models.Pin, error) {
	pins, err := pu.repo.GetByName(name)
	if err != nil {
		return nil, Wrapf(err, "Getting pin by id error, name: %s", name)
	}

	return pins, nil
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