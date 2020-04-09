package database

import (
	"2020_1_Color_noise/internal/models"
	"errors"
	"github.com/jackc/pgx"
	"time"
)

func (db *PgxDB) CreatePin(pin models.DataBasePin) (uint, error) {
	var id uint

	res := db.dbPool.QueryRow(InsertPin, pin.UserId, pin.Name, pin.Description, pin.Image, pin.BoardId, time.Now())
	err := res.Scan(&id)

	if err != nil {
		return 0, errors.New("pin creation error")
	}
	return id, err
}

func (db *PgxDB) UpdatePin(pin models.DataBasePin) error {
	_, err := db.dbPool.Exec(UpdatePin, pin.Name, pin.Description, pin.BoardId, pin.Id)
	if err != nil {
		if pqError, ok := err.(pgx.PgError); ok {
			switch pqError.Code {
			case ForeignKeyViolation:
				return errors.New("board not found")
			default:
				return errors.New("pin not found")
			}
		}
	}
	return err
}

func (db *PgxDB) DeletePin(pin models.DataBasePin) error {
	_, err := db.dbPool.Exec(DeletePin, pin.Id)
	if err != nil {
		return err
	}
	return err
}

func (db *PgxDB) GetPinById(pin models.DataBasePin) (models.Pin, error) {
	var res models.DataBasePin

	row := db.dbPool.QueryRow(PinById, pin.Id)
	err := row.Scan(&res.Id, &res.UserId, &res.Name, &res.Description, &res.Image, &res.BoardId, &res.CreatedAt)
	if err != nil {
		return models.Pin{}, errors.New("pin not found")
	}

	return models.GetPin(res), nil
}

func (db *PgxDB) GetPinsByUserId(pin models.DataBasePin) ([]*models.Pin, error) {
	var res []*models.Pin

	row, err := db.dbPool.Query(PinByUser, pin.UserId)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var tmp models.DataBasePin
		ok := row.Scan(&tmp.Id, &tmp.UserId, &tmp.Name, &tmp.Description,
			&tmp.Image, &tmp.BoardId, &tmp.CreatedAt)
		if ok != nil {
			return nil, ok
		}
		p := models.GetPin(tmp)
		res = append(res, &p)
	}
	return res, nil
}

func (db *PgxDB) GetPinsByName(pin models.DataBasePin) ([]*models.Pin, error) {
	var res []*models.Pin

	row, err := db.dbPool.Query(PinByName, pin.Name)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var tmp models.DataBasePin
		ok := row.Scan(&tmp.Id, &tmp.UserId, &tmp.Name, &tmp.Description,
			&tmp.Image, &tmp.BoardId, &tmp.CreatedAt)
		if ok != nil {
			return nil, ok
		}
		p := models.GetPin(tmp)
		res = append(res, &p)
	}
	return res, nil
}

func (db *PgxDB) GetPinsByBoardID(board models.DataBaseBoard) ([]*models.Pin, error) {
	var res []*models.Pin

	row, err := db.dbPool.Query(PinByBoard, board.Id)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var tmp models.DataBasePin
		ok := row.Scan(&tmp.Id, &tmp.UserId, &tmp.Name, &tmp.Description,
			&tmp.Image, &tmp.BoardId, &tmp.CreatedAt)
		if ok != nil {
			return nil, ok
		}
		p := models.GetPin(tmp)
		res = append(res, &p)
	}
	return res, nil
}
