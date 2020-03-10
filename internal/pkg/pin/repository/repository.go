package repository


import (
	"pinterest/internal/models"
	. "pinterest/internal/pkg/error"
)

type Repository struct {
	data []*models.Pin
}

func NewRepo() *Repository {
	return &Repository{
		data: make([]*models.Pin, 0),
	}
}

func (pr *Repository) Add(pin *models.Pin) (uint, error) {
	pin.Id = uint(len(pr.data) + 1)
	pr.data = append(pr.data, pin)

	return pin.Id, nil
}

func (pr *Repository) GetByID(id uint) (*models.Pin, error) {
	for _, pin := range pr.data {
		if pin.Id == id {
			return pin, nil
		}
	}

	return nil, NotFound.Newf("Pin not found, id: %d", id)
}

func (pr *Repository) GetByUserID(userId uint) ([]*models.Pin, error) {
	result := []*models.Pin{}
	for _, pin := range pr.data {
		if pin.UserId == userId {
			result = append(result, pin)
		}
	}

	if len(result) == 0 {
		return result, NotFound.Newf("Pins not found, userId: %d", userId)
	}

	return result, nil
}

func (pr *Repository) GetByName(name string) ([]*models.Pin, error) {
	result := []*models.Pin{}
	for _, pin := range pr.data {
		if pin.Name == name {
			result = append(result, pin)
		}
	}

	if len(result) == 0 {
		return result, NotFound.Newf("Pins not found, name: %d", name)
	}

	return result, nil
}

func (pr *Repository) Update(pin *models.Pin) error {
	for i, oldPin := range pr.data {
		if oldPin.Id == pin.Id {
			pr.data[i] = pin
			return nil
		}
	}

	return NotFound.Newf("Pin not found, id: %d", pin.Id)
}

func (pr *Repository) Delete(id uint) error {
	for i, pin := range pr.data {
		if pin.Id == id {
			newData := pr.data[:i]
			for j := i + 1; j < len(pr.data); j++ {
				newData = append(newData, pr.data[j])
			}
			pr.data = newData
			return nil
		}
	}

	return NotFound.Newf("Pin not found, id: %d", id)
}