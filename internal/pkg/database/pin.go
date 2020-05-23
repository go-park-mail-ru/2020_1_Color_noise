package database

import (
	"2020_1_Color_noise/internal/models"
	"errors"
	"fmt"
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

func (db *PgxDB) Save(pinId, boardId uint) (bool, error) {
	var check uint

	res := db.dbPool.QueryRow(InsertBoardsPin, pinId, boardId, false)
	err := res.Scan(&check)

	if err != nil {
		if pqError, ok := err.(pgx.PgError); ok {
			switch pqError.Code {
			case ForeignKeyViolation:
				return false, errors.New("pin or board not found")
			default:
				//нарушение уникальности
				return false, errors.New("pin already added")
			}
		}
	}
	return true, err
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
	var us models.DataBaseUser

	row := db.dbPool.QueryRow(PinById, pin.Id)
	err := row.Scan(&res.Id, &res.Name, &res.Description,
		&res.Image, &res.BoardId, &res.CreatedAt, &res.Tags,
		&us.Id, &us.Login, &us.Avatar)
	if err != nil {
		return models.Pin{}, errors.New("pin not found")
	}

	update := db.UpdateViews(pin.Id)
	if update != nil {
		return models.Pin{}, errors.New("view update error")
	}
	rp := models.GetPin(res)
	ru := models.GetUser(us)
	rp.User = &ru

	return rp, nil
}

func (db *PgxDB) GetPinsByUserId(pin models.DataBasePin) ([]*models.Pin, error) {
	var res []*models.Pin

	row, err := db.dbPool.Query(PinByUser, pin.UserId)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var tmp models.DataBasePin
		var us models.DataBaseUser
		ok := row.Scan(&tmp.Id, &tmp.Name, &tmp.Description,
			&tmp.Image, &tmp.BoardId, &tmp.CreatedAt, &tmp.Tags,   &us.Id, &us.Login, &us.Avatar)
		if ok != nil {
			return nil, ok
		}
		ru := models.GetUser(us)
		p := models.GetPin(tmp)
		p.User = &ru
		res = append(res, &p)
	}
	return res, nil
}

func (db *PgxDB) GetPinsByName(pin models.DataBasePin, since time.Time, to time.Time, desc bool, st string, start int, limit int) ([]*models.Pin, error) {
	var res []*models.Pin
	var Query string

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
		var us models.DataBaseUser
		ok := row.Scan(&tmp.Id, &tmp.Name, &tmp.Description,
			&tmp.Image, &tmp.CreatedAt, &tmp.Tags, &tmp.Views, &tmp.Comment,
			&us.Id, &us.Login, &us.Avatar)
		if ok != nil {
			return nil, ok
		}
		ru := models.GetUser(us)
		p := models.GetPin(tmp)
		p.User = &ru
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
		var us models.DataBaseUser
		ok := row.Scan(&tmp.Id, &tmp.Name, &tmp.Description,
			&tmp.Image, &tmp.BoardId, &tmp.CreatedAt, &tmp.Tags,
			&us.Id, &us.Login, &us.Avatar)
		if ok != nil {
			return nil, ok
		}
		ru := models.GetUser(us)
		p := models.GetPin(tmp)
		p.User = &ru
		res = append(res, &p)
	}
	return res, nil
}

func (db *PgxDB) AddTags(pinID uint, tags []string) error {
	var check int
	fmt.Println(tags, pinID)
	res := db.dbPool.QueryRow(AddTags, tags, pinID)
	err := res.Scan(&check)

	if err != nil {
		return err
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
