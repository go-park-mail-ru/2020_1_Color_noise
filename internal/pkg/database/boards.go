package database

import (
	"2020_1_Color_noise/internal/models"
	"errors"
	"time"
)

func (db *PgxDB) CreateBoard(board models.DataBaseBoard) (uint, error) {
	var id uint
	row := db.dbPool.QueryRow(InsertBoard,
		board.UserId, board.Name, board.Description, time.Now())
	err := row.Scan(&id)
	if err != nil {
		return 0, errors.New("user not found")
	}
	return id, err
}

func (db *PgxDB) UpdateBoard(board models.DataBaseBoard) error {
	_, err := db.dbPool.Exec(UpdateBoard,
		board.Name, board.Description, board.Id)
	if err != nil {
		return errors.New("board not found")
	}
	return err
}

func (db *PgxDB) DeleteBoard(board models.DataBaseBoard) error {
	_, err := db.dbPool.Exec(DeleteBoard, board.Id)
	if err != nil {
		return errors.New("board not found")
	}
	return err
}

func (db *PgxDB) GetBoardById(board models.DataBaseBoard) (models.Board, error) {
	var res models.DataBaseBoard
	row := db.dbPool.QueryRow(BoardById, board.Id)
	err := row.Scan(&res.Id, &res.UserId, &res.Name, &res.Description, &res.CreatedAt)

	if err != nil {
		return models.Board{}, errors.New("board not found")
	}
	return models.GetBoard(res), nil
}

func (db *PgxDB) GetBoardsByUserId(board models.DataBaseBoard, start, limit int) ([]*models.Board, error) {
	var res []*models.Board
	row, err := db.dbPool.Query(BoardsByUserId, board.UserId, limit, start)
	if err != nil {
		return nil, errors.New("db problem")
	}

	for row.Next() {
		var tmp models.DataBaseBoard
		ok := row.Scan(&tmp.Id, &tmp.UserId, &tmp.Name, &tmp.Description, &tmp.CreatedAt)
		if ok != nil {
			return res, nil
		}
		r := models.GetBoard(tmp)
		res = append(res, &r)
	}
	return res, nil
}

func (db *PgxDB) GetBoardsByName(board models.DataBaseBoard, start, limit int) ([]*models.Board, error) {
	var res []*models.Board
	row, err := db.dbPool.Query(BoardsByNameSearch, board.Name, limit, start)
	if err != nil {
		return nil, errors.New("db problem")
	}

	for row.Next() {
		var tmp models.DataBaseBoard
		ok := row.Scan(&tmp.Id, &tmp.UserId, &tmp.Name, &tmp.Description, &tmp.CreatedAt)
		if ok != nil {
			return res, nil
		}
		r := models.GetBoard(tmp)
		res = append(res, &r)
	}
	return res, nil
}

func (db *PgxDB) GetBoardLastPin(board models.DataBaseBoard) (models.Pin, error) {
	var res models.DataBasePin
	row := db.dbPool.QueryRow(LastPin, board.Id)
	err := row.Scan(&res.Id, &res.UserId, &res.Name, &res.Description, &res.Image, &res.BoardId, &res.CreatedAt)

	if err != nil {
		return models.Pin{}, errors.New("board not found")
	}
	p := models.GetPin(res)
	us, _ := db.GetUserById(models.DataBaseUser{Id:res.UserId})
	p.User = &us

	return p, nil
}
