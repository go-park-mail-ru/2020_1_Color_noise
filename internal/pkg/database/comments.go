package database

import (
	"2020_1_Color_noise/internal/models"
	"errors"
	"time"
)

func (db *PgxDB) CreateComment(cm models.DataBaseComment) (uint, error) {
	var id uint
	row := db.dbPool.QueryRow(InsertComment,
		cm.UserId, cm.PinId, cm.Text, time.Now())
	err := row.Scan(&id)
	if err != nil {
		return 0, errors.New("comment error")
	}

	update := db.UpdateComments(cm.PinId)
	if update != nil {
		return 0, errors.New("comment update error")
	}

	return id, err
}

func (db *PgxDB) UpdateComment(cm models.DataBaseComment) error {
	_, err := db.dbPool.Exec(UpdateComment, cm.Text, time.Now(), cm.Id)
	if err != nil {
		return errors.New("comment not found")
	}
	return err
}

func (db *PgxDB) DeleteComment(cm models.DataBaseComment) error {
	_, err := db.dbPool.Exec(DeleteComment, cm.Id)
	if err != nil {
		return errors.New("board not found")
	}
	return err
}

func (db *PgxDB) GetCommentById(cm models.DataBaseComment) (models.Comment, error) {
	var r models.Comment
	row := db.dbPool.QueryRow(CommentById, cm.Id)

	ok := row.Scan(&r.Id, &r.UserId, &r.PinId, &r.Text, &r.CreatedAt)
	if ok != nil {
		return models.Comment{}, errors.New("comment not found")
	}
	return r, nil
}

func (db *PgxDB) GetCommentsByPinId(cm models.DataBaseComment, start, limit int) ([]*models.Comment, error) {
	var res []*models.Comment
	r, err := db.dbPool.Query(CommentByPin, cm.PinId, limit, start)
	if err != nil {
		return nil, errors.New("fatal error")
	}

	for r.Next() {
		var tmp models.Comment
		ok := r.Scan(&tmp.Id, &tmp.UserId, &tmp.PinId, &tmp.Text, &tmp.CreatedAt)
		if ok != nil {
			return nil, errors.New("comment not found")
		}
		res = append(res, &tmp)
	}
	return res, nil
}

func (db *PgxDB) GetCommentsByText(cm models.DataBaseComment, start, limit int) ([]*models.Comment, error) {
	var res []*models.Comment
	r, err := db.dbPool.Query(CommentByText, cm.Text, limit, start)
	if err != nil {
		return nil, errors.New("db error")
	}

	for r.Next() {
		var tmp models.Comment
		ok := r.Scan(&tmp.Id, &tmp.UserId, &tmp.PinId, &tmp.Text, &tmp.CreatedAt)
		if ok != nil {
			return nil, errors.New("comment not found")
		}
		res = append(res, &tmp)
	}
	return res, nil
}
