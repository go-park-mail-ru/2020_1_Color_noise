package pin

import (
	"2020_1_Color_noise/internal/models"
	"bytes"
)

type IUsecase interface {
	//Create(input *models.InputPin, userId uint) (uint, error)
	CreatePin(input *models.InputPin, id uint, userId uint) (uint, error)
	SaveImage(userId uint, buffer *bytes.Buffer) (uint, *[]string, error)
	Save(pinId uint, boardId uint) (bool, error)
	GetById(id uint, userId uint) (*models.Pin, error)
	GetByUserId(id uint, start int, limit int) ([]*models.Pin, error)
	GetByName(name string, start int, limit int, date string, desc bool, most string) ([]*models.Pin, error)
	Update(input *models.UpdatePin, pinId uint, userId uint) error
	Delete(pinId uint, userId uint) error
}
