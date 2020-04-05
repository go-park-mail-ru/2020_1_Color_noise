package repository

import (
	"2020_1_Color_noise/internal/models"
	"2020_1_Color_noise/internal/pkg/database"
	. "2020_1_Color_noise/internal/pkg/error"
	"sync"
)

type Repository struct {
	db database.DBInterface
	mu *sync.Mutex
}

func NewRepo(d database.DBInterface) *Repository {
	return &Repository{
		db: d,
		mu: &sync.Mutex{},
	}
}

func (br *Repository) Create(board *models.Board) (uint, error) {
	br.mu.Lock()
	defer br.mu.Unlock()

	id, err := br.db.CreateBoard(models.GetDBoard(*board))
	if err != nil {
		return 0, BoardNotFound.Newf("board can not be created, err: ", err)
	}

	return id, nil
}

func (br *Repository) GetByID(id uint) (*models.Board, error) {

	b := models.DataBaseBoard{
		Id: id,
	}
	board, err := br.db.GetBoardById(b)
	if err != nil {
		return nil, BoardNotFound.Newf("Board not found, board id: %d", id)
	}
	return &board, err
}

func (br *Repository) GetByUserID(userId uint, start int, limit int) ([]*models.Board, error) {
	b := models.DataBaseBoard{
		UserId: userId,
	}
	boards, err := br.db.GetBoardsByUserId(b, start, limit)
	if err != nil {
		return nil, BoardNotFound.Newf("Boards not found, user_id: %d", userId)
	}
	return boards, err
}

func (br *Repository) GetByName(name string, start int, limit int) ([]*models.Board, error) {
	//TODO: придумать проверку
	/*
		if start >= len(br.data) {
			start = 0
		}

		if limit >= (len(br.data) - start) {
			limit = len(br.data)
		}
	*/

	b := models.DataBaseBoard{
		Name: name,
	}
	boards, err := br.db.GetBoardsByName(b, start, limit)
	if err != nil {
		return nil, BoardNotFound.Newf("Boards not found, name: ", name)
	}
	return boards, err
}



func (br *Repository) Update(board *models.Board) error {
	err := br.db.UpdateBoard(models.GetDBoard(*board))
	if err != nil {
		return BoardNotFound.Newf("Board not found, id: %d", board.Id)
	}
	return nil
}




func (br *Repository) Delete(id uint) error {
	d := models.DataBaseBoard{
		Id: id,
	}
	err := br.db.DeleteBoard(d)
	if err != nil {
		return BoardNotFound.Newf("Board not found, id: %d", id)
	}

	return nil
}
