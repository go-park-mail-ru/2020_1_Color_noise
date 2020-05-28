package database

import (
	"2020_1_Color_noise/internal/models"
	"errors"
)

func (db *PgxDB) GetSubFeed(user models.DataBaseUser, start, limit int) ([]*models.Pin, error) {
	var res []*models.Pin

	row, err := db.dbPool.Query(Feed, user.Id, limit, start)
	defer row.Close()
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var tmp models.DataBasePin
		ok := row.Scan(&tmp.Id, &tmp.UserId, &tmp.Name, &tmp.Description,
			&tmp.Image, &tmp.BoardId, &tmp.CreatedAt)
		if ok != nil {
			return nil, errors.New("pin not found")
		}
		p := models.GetPin(tmp)
		us, _ := db.GetUserById(models.DataBaseUser{Id:tmp.UserId})
		p.User = &us
		res = append(res, &p)
	}
	return res, nil
}

func (db *PgxDB) GetMainFeed(user models.DataBaseUser, start, limit int) ([]*models.Pin, error) {
	var res []*models.Pin

	row, err := db.dbPool.Query(Main, limit, start)
	defer row.Close()
	if err != nil {
		return nil, errors.New("db problem")
	}
	for row.Next() {
		var tmp models.DataBasePin
		ok := row.Scan(&tmp.Id, &tmp.UserId, &tmp.Name, &tmp.Description,
			&tmp.Image, &tmp.BoardId, &tmp.CreatedAt)
		if ok != nil {
			return res, nil
		}
		p := models.GetPin(tmp)
		us, _ := db.GetUserById(models.DataBaseUser{Id:tmp.UserId})
		p.User = &us
		res = append(res, &p)
	}
	return res, nil
}

func (db *PgxDB) GetRecFeed(user models.DataBaseUser, start, limit int) ([]*models.Pin, error) {
	var res []*models.Pin

	row, err := db.dbPool.Query(Recommendation, limit, start)
	defer row.Close()
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var tmp models.DataBasePin
		ok := row.Scan(&tmp.Id, &tmp.UserId, &tmp.Name, &tmp.Description,
			&tmp.Image, &tmp.BoardId, &tmp.CreatedAt)
		if ok != nil {
			return res, nil
		}
		p := models.GetPin(tmp)
		us, _ := db.GetUserById(models.DataBaseUser{Id:tmp.UserId})
		p.User = &us
		res = append(res, &p)
	}
	return res, nil
}
