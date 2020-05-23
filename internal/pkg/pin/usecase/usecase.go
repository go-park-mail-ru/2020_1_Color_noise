package usecase

import (
	"2020_1_Color_noise/internal/models"
	"2020_1_Color_noise/internal/pkg/board"
	. "2020_1_Color_noise/internal/pkg/error"
	"2020_1_Color_noise/internal/pkg/pin"
	imageService "2020_1_Color_noise/internal/pkg/proto/image"
	userService "2020_1_Color_noise/internal/pkg/proto/user"
	"2020_1_Color_noise/internal/pkg/utils"
	"context"
	"encoding/base64"
	"log"
	"strings"
)

type Usecase struct {
	repoPin      pin.IRepository
	repoBoard    board.IRepository
	imageService imageService.ImageServiceClient
	userServ     userService.UserServiceClient
}

func NewUsecase(repoPin pin.IRepository, repoBoard board.IRepository, imageServ imageService.ImageServiceClient, userServ userService.UserServiceClient) *Usecase {
	return &Usecase{
		repoPin,
		repoBoard,
		imageServ,
		userServ,
	}
}

func (pu *Usecase) Create(input *models.InputPin, userId uint) (uint, error) {
	b64data := input.Image[strings.IndexByte(input.Image, ',')+1:]

	length := base64.StdEncoding.EncodedLen(len(b64data))
	if length > 10000000 {
		return 0, Wrapf(TooMuchSize.New("Too much size image"), "Creating pin error, userId: %d", userId)
	}

	buffer, err := base64.StdEncoding.DecodeString(b64data)
	if err != nil {
		err = Wrap(err, "Decoding base64")
		return 0, Wrapf(err, "Creating pin error, userId: %d", userId)
	}

	name, err := utils.SaveImage(&buffer)
	if err != nil {
		return 0, Wrapf(err, "Creating pin error, userId: %d", userId)
	}

	pin := &models.Pin{
		User:        &models.User{
			Id: userId,
		},
		BoardId:     uint(input.BoardId),
		Name:        input.Name,
		Description: input.Description,
		Image:       name,
	}

	id, err := pu.repoPin.Create(pin)
	if err != nil {
		return 0, Wrapf(err, "Creating pin error, userId: %d", userId)
	}

	go func() {
		err := pu.Analyze(id, name)
		if err != nil {
			log.Println(err)
		}
	}()

	return id, nil
}

func (pu *Usecase) Save(pinId uint, boardId uint) (bool, error) {
	status, err := pu.repoPin.Save(pinId, boardId)
	if err != nil {
		return false, Wrapf(err, "Saving pin error, boardId: %d", boardId)
	}

	return status, nil
}

func (pu *Usecase) GetById(id uint) (*models.Pin, error) {
	pin, err := pu.repoPin.GetByID(id)
	if err != nil {
		return nil, Wrapf(err, "Getting pin by id error, pinId: %d", id)
	}

	return pin, nil
}

func (pu *Usecase) GetByUserId(id uint, start int, limit int) ([]*models.Pin, error) {
	pins, err := pu.repoPin.GetByUserID(id, start, limit)
	if err != nil {
		return nil, Wrapf(err, "Getting pin by id error, pinId: %d", id)
	}

	return pins, nil
}

func (pu *Usecase) GetByName(name string, start int, limit int, date string, desc bool, most string) ([]*models.Pin, error) {
	pins, err := pu.repoPin.GetByName(name, start, limit, date, desc, most)
	if err != nil {
		return nil, Wrapf(err, "Getting pin by id error, name: %s", name)
	}

	return pins, nil
}

func (pu *Usecase) Update(input *models.UpdatePin, pinId uint, userId uint) error {
	pin := &models.Pin{
		Id:          pinId,
		User:        &models.User{
			Id: userId,
		},
		BoardId:     uint(input.BoardId),
		Name:        input.Name,
		Description: input.Description,
	}

	err := pu.repoPin.Update(pin)
	if err != nil {
		return Wrap(err, "Updating pin error")
	}

	return nil
}

func (pu *Usecase) Delete(pinId uint, userId uint) error {
	err := pu.repoPin.Delete(pinId, userId)
	if err != nil {
		return Wrap(err, "Deleting pin error")
	}

	return nil
}

func (pu *Usecase) Analyze(pinId uint, name string) error {
	tags, err := pu.imageService.Analyze(context.Background(), &imageService.Address{Image: name})
	if err != nil {
		return Wrap(err, "analyzing pin error")
	}

	err = pu.repoPin.AddTags(pinId, tags.Tags)
	if err != nil {
		return Wrap(err, "adding tags pin error")
	}

	return nil
}
