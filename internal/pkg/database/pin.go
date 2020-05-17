package database

import (
	"2020_1_Color_noise/internal/models"
	"errors"
	"github.com/jackc/pgx"
	"time"
)

func (db *PgxDB) CreatePin(pin models.DataBasePin) (uint, error) {
	var id, check uint

	res := db.dbPool.QueryRow(InsertPin, pin.UserId, pin.Name, pin.Description, pin.Image, time.Now(), []string{}, 0, 0)
	err := res.Scan(&id)

	if err != nil {
		return 0, errors.New("pin creation error")
	}
	res = db.dbPool.QueryRow(InsertBoardsPin, id, pin.BoardId, true)
	err = res.Scan(&check)
	return id, err
}

func (db *PgxDB) Save(pinId, boardId uint) ( error) {
	var check uint

	res := db.dbPool.QueryRow(InsertBoardsPin, pinId, boardId, false)
	err := res.Scan(&check)
	return err
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
	err := row.Scan(&res.Id, &res.UserId, &res.Name, &res.Description, &res.Image, &res.BoardId, &res.CreatedAt, &res.Tags)
	if err != nil {
		return models.Pin{}, errors.New("pin not found")
	}

	db.UpdateViews(pin.Id)
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
			&tmp.Image, &tmp.BoardId, &tmp.CreatedAt, &tmp.Tags)
		if ok != nil {
			return nil, ok
		}
		p := models.GetPin(tmp)
		res = append(res, &p)
	}
	return res, nil
}

func (db *PgxDB) GetPinsByName(pin models.DataBasePin, since time.Time, to time.Time, desc bool, st string, start int, limit int) ([]*models.Pin, error) {
	var res []*models.Pin
	Query := PinByName

	if desc {
		switch st {
		case "popular":
			Query = PopularDesc
		case "comment":
			Query = CommentsDesc
		default:
			Query = IdDesc
		}
	} else {
		switch st {
		case "popular":
			Query = PopularAsc
		case "comment":
			Query = CommentsAsc
		default:
			Query = IdAsc
		}
	}

	row, err := db.dbPool.Query(Query, pin.Name, since, to, start, limit)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var tmp models.DataBasePin
		ok := row.Scan(&tmp.Id, &tmp.UserId, &tmp.Name, &tmp.Description,
			&tmp.Image, &tmp.BoardId, &tmp.CreatedAt, &tmp.Tags)
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
			&tmp.Image, &tmp.BoardId, &tmp.CreatedAt, &tmp.Tags)
		if ok != nil {
			return nil, ok
		}
		p := models.GetPin(tmp)
		res = append(res, &p)
	}
	return res, nil
}

func (db *PgxDB) AddTags(pinID uint, tags []string) error {
	var check int

	res := db.dbPool.QueryRow(AddTags, tags, pinID)
	err := res.Scan(&check)

	if err != nil {
		return errors.New("tags adding error")
	}
	return nil
}

func (db *PgxDB) UpdateViews(pinID uint) error {

	_, err := db.dbPool.Exec(UpdateViews, pinID)

	if err != nil {
		return errors.New("views error")
	}
	return nil
}

func (db *PgxDB) UpdateComments(pinID uint) error {

	_, err := db.dbPool.Exec(UpdateComments, pinID)
	if err != nil {
		return errors.New("comments error")
	}
	return nil
}