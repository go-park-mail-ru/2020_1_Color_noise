package repository

import (
	"2020_1_Color_noise/internal/models"
	"2020_1_Color_noise/internal/pkg/database"
	. "2020_1_Color_noise/internal/pkg/error"
	"time"
)

type Repository struct {
	db database.DBInterface
}

func NewRepo(d database.DBInterface) *Repository {
	return &Repository{
		db: d,
	}
}

func (pr *Repository) Create(pin *models.Pin) (uint, error) {
	//добавить в пины
	//добавить в таблицу
	id, err := pr.db.CreatePin(models.GetBPin(*pin))

	if err != nil {
		return 0, PinNotFound.Newf("Pin can not be created, err: %v", err)
	}

	return id, err
}

func (pr *Repository) Save(pinId uint, boardId uint) (error) {
	//добавить в пины
	//добавить в таблицу
	err := pr.db.Save(pinId, boardId)

	if err != nil {
		return PinNotFound.Newf("Pin can not be saved, err: %v", err)
	}

	return err
}

func (pr *Repository) GetByID(id uint) (*models.Pin, error) {
	p := models.DataBasePin{Id: id}
	pin, err := pr.db.GetPinById(p)

	if err != nil {
		return nil, PinNotFound.Newf("Pin not found, id: %d", id)
	}
	return &pin, err
}

func (pr *Repository) GetByUserID(userId uint, start int, limit int) ([]*models.Pin, error) {


	p := models.DataBasePin{UserId: userId}
	result, err := pr.db.GetPinsByUserId(p)
	if err != nil {
		return result, PinNotFound.Newf("Pins not found, user id = %d", userId)
	}

	return result, nil
}

func (pr *Repository) GetByName(name string, start int, limit int, date string, desc bool, most string) ([]*models.Pin, error) {
	var since, to time.Time

	switch date {
	case "":
		since = time.Time{}
		to = time.Now()
	case "day":
		to = time.Now()
		since = to.AddDate(0, 0, -1)
	case "week":
		to = time.Now()
		since = to.AddDate(0, 0, -7)
	case "month":
		to = time.Now()
		since = to.AddDate(0, -1, 0)
	}

	p := models.DataBasePin{Name: "%" + name + "%"}
	result, err := pr.db.GetPinsByName(p, since, to, desc, most, start, limit)
	if err != nil {
		return result, PinNotFound.Newf("Pins not found, err = %v", err)
	}

	return result, nil
}

func (pr *Repository) Update(pin *models.Pin) error {

	err := pr.db.UpdatePin(models.GetBPin(*pin))
	if err != nil {
		return PinNotFound.Newf("Pin not found, id: %d", pin.Id)
	}
	return nil

}

func (pr *Repository) Delete(pinId uint, userId uint) error {

	p := models.DataBasePin{
		Id: pinId,
	}

	err := pr.db.DeletePin(p)
	if err != nil {
		return PinNotFound.Newf("Pin not found, id: %d", pinId)
	}
	return nil
}

func (pr *Repository) AddTags(pinId uint, tags []string) error {
	err := pr.db.AddTags(pinId, tags)
	if err != nil {
		return NoType.Newf("Tags error: %v", err)
	}
	return nil
}