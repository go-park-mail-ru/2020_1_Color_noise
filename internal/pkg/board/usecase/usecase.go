package usecase

import (
	"2020_1_Color_noise/internal/models"
	"2020_1_Color_noise/internal/pkg/board"
	. "2020_1_Color_noise/internal/pkg/error"
)

type Usecase struct {
	repo board.IRepository
}

func NewUsecase(repo board.IRepository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (bu *Usecase) Create(input *models.InputBoard, userId uint) (uint, error) {
	board := &models.Board{
		UserId:      userId,
		Name:        input.Name,
		Description: input.Description,
	}

	id, err := bu.repo.Create(board)
	if err != nil {
		return 0, Wrapf(err, "Creating board error, userId: %s", userId)
	}

	return id, nil
}

func (bu *Usecase) GetById(id uint) (*models.Board, error) {
	board, err := bu.repo.GetByID(id)
	if err != nil {
		return nil, Wrapf(err, "Getting board by id error, pinId: %s", id)
	}

	return board, nil
}

func (bu *Usecase) GetByNameId(id uint) (*models.Board, error){
	board, err := bu.repo.GetByNameID(id)
	if err != nil {
		return nil, Wrapf(err, "Getting board by id error, pinId: %s", id)
	}

	return board, nil
}

func (bu *Usecase) GetByUserId(id uint, start int, limit int) ([]*models.Board, error) {
	boards, err := bu.repo.GetByUserID(id, start, limit)
	if err != nil {
		return nil, Wrapf(err, "Getting board by id error, pinId: %s", id)
	}

	return boards, nil
}

func (bu *Usecase) GetByName(name string, start int, limit int) ([]*models.Board, error) {
	boards, err := bu.repo.GetByName(name, start, limit)
	if err != nil {
		return nil, Wrapf(err, "Getting board by id error, name: %s", name)
	}

	return boards, nil
}

func (pu *Usecase) Update(input *models.InputBoard, boardId uint, userId uint) error {
	board := &models.Board{
		Id:          boardId,
		UserId:      userId,
		Name:        input.Name,
		Description: input.Description,
	}

	err := pu.repo.Update(board)
	if err != nil {
		return Wrap(err, "Updating board error")
	}

	return nil
}

func (pu *Usecase) Delete(id uint, userId uint) error {
	err := pu.repo.Delete(id, userId)
	if err != nil {
		return Wrap(err, "Deleting board error")
	}

	return nil
}
