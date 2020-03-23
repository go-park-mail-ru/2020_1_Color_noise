package repository


import (
	"2020_1_Color_noise/internal/models"
	. "2020_1_Color_noise/internal/pkg/error"
	"sync"
	"time"
)

type Repository struct {
	data []*models.Comment
	mu   *sync.Mutex
}

func NewRepo() *Repository {
	return &Repository{
		data: make([]*models.Comment, 0),
		mu:   &sync.Mutex{},
	}
}

func (cr *Repository) Create(comment *models.Comment) (uint, error) {
	cr.mu.Lock()

	comment.Id = uint(len(cr.data) + 1)
	comment.CreatedAt = time.Now()
	cr.data = append(cr.data, comment)

	cr.mu.Unlock()

	return comment.Id, nil
}

func (cr *Repository) GetByID(id uint) (*models.Comment, error) {
	cr.mu.Lock()
	defer cr.mu.Unlock()
	for _, comment := range cr.data {
		if comment.Id == id {
			return comment, nil
		}
	}

	return nil, CommentNotFound.Newf("Repo: Getting by id comment error, id: %d", id)
}

func (cr *Repository) GetByPinID(pinId uint, start int, limit int) ([]*models.Comment, error) {
	var result []*models.Comment

	if start >= len(cr.data) {
		start = 0
	}

	if limit >= (len(cr.data) - start) {
		limit = len(cr.data)
	}

	for i, comment := range cr.data {
		if comment.PinId == pinId && start >= i {
			result = append(result, comment)

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

func (cr *Repository) GetByText(text string,  start int, limit int) ([]*models.Comment, error) {
	var result []*models.Comment

	if start >= len(cr.data) {
		start = 0
	}

	if limit >= (len(cr.data) - start) {
		limit = len(cr.data)
	}

	for i, comment := range cr.data {
		if comment.Text == text && start >= i {
			result = append(result, comment)

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

/*
func (br *Repository) Update(pin *models.Board) error {
	br.mu.Lock()
	for i, oldPin := range br.data {
		if oldPin.Id == pin.Id {
			br.data[i] = pin
			br.mu.Unlock()
			return nil
		}
	}
	br.mu.Unlock()

	return BoardNotFound.Newf("Pin not found, id: %d", pin.Id)
}
*/


func (cr *Repository) Delete(id uint) error {
	cr.mu.Lock()
	for i, comment := range cr.data {
		if comment.Id == id {
			newData := cr.data[:i]
			for j := i + 1; j < len(cr.data); j++ {
				newData = append(newData, cr.data[j])
			}
			cr.data = newData
			cr.mu.Unlock()
			return nil
		}
	}
	cr.mu.Unlock()

	return PinNotFound.Newf("Pin not found, id: %d", id)
}