package usecase

import (
	"2020_1_Color_noise/internal/models"
	. "2020_1_Color_noise/internal/pkg/error"
	"2020_1_Color_noise/internal/pkg/image"
	"2020_1_Color_noise/internal/pkg/pin"
	"encoding/base64"
	"strings"
)

type Usecase struct {
	repo pin.IRepository
}

func NewUsecase(repo pin.IRepository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (pu *Usecase) Create(input *models.InputPin, userId uint) (uint, error) {
	b64data := input.Image[strings.IndexByte(input.Image, ',')+1:]

	length := base64.StdEncoding.EncodedLen(len(b64data))
	if length > 10000000 {
		return 0, Wrapf(TooMuchSize.New("Too much size image"), "Creating pin error, userId: %s", userId)
	}

	buffer, err := base64.StdEncoding.DecodeString(b64data)
	if err != nil {
		err = Wrap(err, "Decoding base64")
		return 0, Wrapf(err, "Creating pin error, userId: %s", userId)
	}

	name, err := image.SaveImage(&buffer)
	if err != nil {
		return 0, Wrapf(err, "Creating pin error, userId: %s", userId)
	}

	pin := &models.Pin{
		UserId:      userId,
		BoardId:     uint(input.BoardId),
		Name:        input.Name,
		Description: input.Description,
		Image:       name,
	}

	id, err := pu.repo.Create(pin)
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

func (pu *Usecase) GetByUserId(id uint, start int, limit int) ([]*models.Pin, error) {
	pins, err := pu.repo.GetByUserID(id, start, limit)
	if err != nil {
		return nil, Wrapf(err, "Getting pin by id error, pinId: %s", id)
	}

	return pins, nil
}

func (pu *Usecase) GetByName(name string, start int, limit int) ([]*models.Pin, error) {
	pins, err := pu.repo.GetByName(name, start, limit)
	if err != nil {
		return nil, Wrapf(err, "Getting pin by id error, name: %s", name)
	}

	return pins, nil
}

func (pu *Usecase) Update(input *models.UpdatePin, pinId uint, userId uint) error {
	pin := &models.Pin{
		Id:          pinId,
		UserId:      userId,
		BoardId:     uint(input.BoardId),
		Name:        input.Name,
		Description: input.Description,
	}

	err := pu.repo.Update(pin)
	if err != nil {
		return Wrap(err, "Updating pin error")
	}

	return nil
}

func (pu *Usecase) Delete(pinId uint, userId uint) error {
	err := pu.repo.Delete(pinId, userId)
	if err != nil {
		return Wrap(err, "Deleting pin error")
	}

	return nil
}
