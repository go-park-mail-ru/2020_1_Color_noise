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
	pinRepo  pin.IRepository
}

func NewUsecase(repo pin.IRepository) *Usecase {
	return &Usecase{
		pinRepo: repo,
	}
}

func (pu *Usecase) Add(input *models.InputPin) (uint, error) {
	length := base64.StdEncoding.EncodedLen(len(input.Image))
	if length > 10000000 {
		return 0, TooMuchSize.New("Too much size image")
	}

	if input.Name == "" || input.Description == "" || length == 0 {
		return 0, BadPin.New("Fill in all the fields")
	}

	buffer := make([]byte, length)

	l, err := base64.StdEncoding.Decode(buffer, input.Image)
	if err != nil {
		return 0, Wrap(err, "Decoding error")
	}
	buffer = buffer[:l]

	name, err := image.SaveImage(&buffer)
	if err != nil {
		return 0, nil
	}

	pin := &models.Pin{
		Name:        input.Name,
		Description: input.Description,
		Image:       "static/" + name,
	}

	return pu.pinRepo.Add(pin)
}

func (pu *Usecase) GetById(id uint) (*models.Pin, error) {
	return pu.pinRepo.GetByID(id)
}

func (pu *Usecase) GetByUserId(id uint) ([]*models.Pin, error) {
	return pu.pinRepo.GetByUserID(id)
}

func (pu *Usecase) GetByName(name string) ([]*models.Pin, error) {
	return pu.pinRepo.GetByName(name)
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