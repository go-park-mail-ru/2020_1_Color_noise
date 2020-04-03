package repository


import (
	"2020_1_Color_noise/internal/models"
	. "2020_1_Color_noise/internal/pkg/error"
	"sync"
)

type Repository struct {
	data []*models.Board
	mu   *sync.Mutex
}

func NewRepo() *Repository {
	return &Repository{
		data: make([]*models.Board, 0),
		mu:   &sync.Mutex{},
	}
}

func (br *Repository) Create(board *models.Board) (uint, error) {
	br.mu.Lock()
	board.Id = uint(len(br.data) + 1)
	br.data = append(br.data, board)
	br.mu.Unlock()

	return board.Id, nil
}

func (br *Repository) GetByID(id uint) (*models.Board, error) {
	for _, board := range br.data {
		if board.Id == id {
			return board, nil
		}
	}

	return nil, BoardNotFound.Newf("Pin not found, id: %d", id)
}

func (br *Repository) GetByUserID(userId uint, start int, limit int) ([]*models.Board, error) {
	var result []*models.Board

	if start >= len(br.data) {
		start = 0
	}

	if limit >= (len(br.data) - start) {
		limit = len(br.data)
	}

	for i, board := range br.data {
		if board.UserId == userId && start >= i {
			result = append(result, board)

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

func (br *Repository) GetByName(name string,  start int, limit int) ([]*models.Board, error) {
	var result []*models.Board

	if start >= len(br.data) {
		start = 0
	}

	if limit >= (len(br.data) - start) {
		limit = len(br.data)
	}

	for i, board := range br.data {
		if board.Name == name && start >= i {
			result = append(result, board)

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


func (br *Repository) Update(board *models.Board) error {
	br.mu.Lock()
	defer br.mu.Unlock()
	for i, oldBoard := range br.data {
		if oldBoard.Id == board.Id {
			br.data[i].Name = board.Name
			br.data[i].Description = board.Description
			return nil
		}
	}

	return BoardNotFound.Newf("Pin not found, id: %d", pin.Id)
}



func (br *Repository) Delete(id uint) error {
	br.mu.Lock()
	defer br.mu.Unlock()
	for i, board := range br.data {
		if board.Id == id {
			newData := br.data[:i]
			for j := i + 1; j < len(br.data); j++ {
				newData = append(newData, br.data[j])
			}
			br.data = newData
			return nil
		}
	}

	return PinNotFound.Newf("Pin not found, id: %d", id)
}