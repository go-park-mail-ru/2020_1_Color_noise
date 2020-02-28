package usecase

import (
	"fmt"
	"math/rand"
	"os"
	"pinterest/pkg/models"
	"time"
	//"golang.org/x/crypto/bcrypt"
	repo "pinterest/pkg/pin/repository"
)

type PinUsecase struct {
	pinRepo  *repo.PinRepository
}

func NewPinUsecase(repo *repo.PinRepository) *PinUsecase {
	return &PinUsecase{
		pinRepo: repo,
	}
}

func (pu *PinUsecase) Add(pin *models.Pin) (uint, error) {
	pins, err := pu.pinRepo.GetByID(pin.Id)
	if err != nil {
		return 0, err
	}
	if len(pins) != 1 {
		return 0, fmt.Errorf("Internal error")
	}
	pins[0].Name = pin.Name
	pins[0].Description = pin.Description
	ok, err := pu.pinRepo.Update(pins[0])
	if err != nil {
		return 0, err
	}
	if !ok {
		return 0, fmt.Errorf("Internal error")
	}
	return 0, nil
}

func (pu *PinUsecase) Get(id uint) (*models.Pin, error) {
	pins, err := pu.pinRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if len(pins) != 1 {
		return nil, fmt.Errorf("Pin not found")
	}
	return pins[0], nil
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

func (pu *PinUsecase) SaveImage(pin *models.Pin) (*models.Pin, error) {
	name := randStringRunes(30) + ".jpg"
	file, err := os.Create("static/" + name)
	if err != nil{
		return nil, err
	}
	defer file.Close()

	_, err = file.Write(pin.Image)
	if err != nil {
		return nil, err
	}

	pin.Image = []byte{}
	pin.ImageAdress = name
	id, err := pu.pinRepo.Add(pin)
	if err != nil {
		return nil, err
	}
	pin.Id = id
	return pin, nil
}

var (
	letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func randStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}