package repository


import (
	"pinterest/pkg/models"
)

type PinRepository struct {
	data []*models.Pin
}

func NewPinRepo() *PinRepository {
	return &PinRepository{
		data: make([]*models.Pin, 0),
	}
}

func (pr *PinRepository) Add(pin *models.Pin) (uint, error) {
	pin.Id = uint(len(pr.data) + 1)
	pr.data = append(pr.data, pin)
	return pin.Id, nil
}

func (pr *PinRepository) GetByID(id uint) ([]*models.Pin, error) {
	for _, pin := range pr.data {
		if pin.Id == id {
			return []*models.Pin{pin}, nil
		}
	}
	return []*models.Pin{}, nil
}

func (pr *PinRepository) GetByUserID(userId uint) ([]*models.Pin, error) {
	result := []*models.Pin{}
	for _, pin := range pr.data {
		if pin.UserId == userId {
			result = append(result, pin)
		}
	}
	return result, nil
}

func (pr *PinRepository) GetByName(name string) ([]*models.Pin, error) {
	result := []*models.Pin{}
	for _, pin := range pr.data {
		if pin.Name == name {
			result = append(result, pin)
		}
	}
	return result, nil
}

func (pr *PinRepository) Update(id uint, pin *models.Pin) (bool, error) {
	for i, pin := range pr.data {
		if pin.Id == id {
			pr.data[i] = pin
			return true, nil
		}
	}
	return false, nil
}

func (pr *PinRepository) Delete(id uint) (bool, error) {
	for i, pin := range pr.data {
		if pin.Id == id {
			newData := pr.data[:i]
			for j := i + 1; j < len(pr.data); j++ {
				newData = append(newData, pr.data[j])
			}
			pr.data = newData
			return true, nil
		}
	}
	return false, nil
}