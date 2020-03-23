package repository


import (
	"2020_1_Color_noise/internal/models"
	. "2020_1_Color_noise/internal/pkg/error"
	"sync"
	"time"
)

type Repository struct {
	data []*models.Pin
	mu   *sync.Mutex
}

func NewRepo() *Repository {
	return &Repository{
		data: make([]*models.Pin, 0),
		mu:   &sync.Mutex{},
	}
}

func (pr *Repository) Create(pin *models.Pin) (uint, error) {
	pr.mu.Lock()

	pin.Id = uint(len(pr.data) + 1)
	pr.data = append(pr.data, pin)

	pr.mu.Unlock()

	return pin.Id, nil
}

func (pr *Repository) GetByID(id uint) (*models.Pin, error) {
	pr.mu.Lock()
	defer pr.mu.Unlock()

	for _, pin := range pr.data {
		if pin.Id == id {
			return pin, nil
		}
	}

	return nil, PinNotFound.Newf("Pin not found, id: %d", id)
}

func (pr *Repository) GetByUserID(userId uint, start int, limit int) ([]*models.Pin, error) {
	var result []*models.Pin

	if start >= len(pr.data) {
		start = 0
	}

	if limit >= (len(pr.data) - start) {
		limit = len(pr.data)
	}

	pr.mu.Lock()
	defer pr.mu.Unlock()

	for i, pin := range pr.data {
		if pin.UserId == userId && start >= i {
			result = append(result, pin)

			if limit == len(result){
				break
			}
		}
	}

	/*if len(result) == 0 {
		return result, PinNotFound.Newf("Pins not found, userId: %d", userId)
	}*/

	return result, nil
}

func (pr *Repository) GetByName(name string,  start int, limit int) ([]*models.Pin, error) {
	var result []*models.Pin

	if start >= len(pr.data) {
		start = 0
	}

	if limit >= (len(pr.data) - start) {
		limit = len(pr.data)
	}

	pr.mu.Lock()
	defer pr.mu.Unlock()

	for i, pin := range pr.data {
		if pin.Name == name && start >= i {
			result = append(result, pin)

			if limit == len(result){
				break
			}
		}
	}

	/*if len(result) == 0 {
		return result, PinNotFound.Newf("Pins not found, name: %d", name)
	}*/

	return result, nil
}

func (pr *Repository) Update(pin *models.Pin) error {
	pr.mu.Lock()
	defer pr.mu.Unlock()

	for i, oldPin := range pr.data {
		if oldPin.Id == pin.Id && oldPin.UserId == pin.UserId{
			pin.CreatedAt = oldPin.CreatedAt
			pin.UpdatedAt = time.Now()
			pin.Image = oldPin.Image

			pr.data[i] = pin
			return nil
		}
	}

	return PinNotFound.Newf("Pin not found, id: %d", pin.Id)
}

func (pr *Repository) Delete(pinId uint, userId uint) error {
	pr.mu.Lock()
	defer pr.mu.Unlock()

	for i, pin := range pr.data {
		if pin.Id == pinId && pin.UserId == userId {
			newData := pr.data[:i]
			for j := i + 1; j < len(pr.data); j++ {
				newData = append(newData, pr.data[j])
			}
			pr.data = newData
			return nil
		}
	}

	return PinNotFound.Newf("Pin not found, id: %d", userId)
}