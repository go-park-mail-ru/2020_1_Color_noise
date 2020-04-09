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

func (cr *Repository) Create(comment *models.Comment) (uint, error) {
	cr.mu.Lock()
	defer cr.mu.Unlock()
	id, err := cr.db.CreateComment(models.GetBComment(*comment))
	comment.Id = id
	_, _ = cr.db.PutNotifications(models.GetBComment(*comment))
	if err != nil {
		return 0, CommentNotFound.Newf("Comment can not be created, err %v:", err)
	}
	return id, nil
}

func (cr *Repository) GetByID(id uint) (*models.Comment, error) {
	cr.mu.Lock()
	defer cr.mu.Unlock()

	c := models.DataBaseComment{Id: id}
	com, err := cr.db.GetCommentById(c)
	if err != nil {
		return nil, CommentNotFound.Newf("Repo: Getting by id comment error, id: %d", id)
	}
	return &com, err
}

func (cr *Repository) GetByPinID(pinId uint, start int, limit int) ([]*models.Comment, error) {
	/*
		if start >= len(cr.data) {
			start = 0
		}

		if limit >= (len(cr.data) - start) {
			limit = len(cr.data)
		}
	*/

	c := models.DataBaseComment{PinId: pinId}
	comments, err := cr.db.GetCommentsByPinId(c, start, limit)
	if err != nil {
		return comments, CommentNotFound.Newf("Comments not found, pinID: %d", pinId)
	}

	return comments, nil
}

func (cr *Repository) GetByText(text string, start int, limit int) ([]*models.Comment, error) {

	/*
		if start >= len(cr.data) {
			start = 0
		}

		if limit >= (len(cr.data) - start) {
			limit = len(cr.data)
		}
	*/
	/*if len(result) == 0 {
		return result, PinNotFound.Newf("Pins not found, name: %d", name)
	}*/

	c := models.DataBaseComment{Text: text}
	comments, err := cr.db.GetCommentsByText(c, start, limit)
	if err != nil {
		return comments, CommentNotFound.Newf("Comments not found, text-like: %s", text)
	}

	return comments, nil
}

func (cr *Repository) Update(comment *models.Comment) error {
	cr.mu.Lock()
	defer cr.mu.Unlock()

	err := cr.db.UpdateComment(models.GetBComment(*comment))
	if err != nil {
		return CommentNotFound.Newf("Comment can not be updared, err %v:", err)
	}
	return nil
}

func (cr *Repository) Delete(id uint) error {
	cr.mu.Lock()
	defer cr.mu.Unlock()

	c := models.DataBaseComment{Id: id}
	err := cr.db.UpdateComment(c)
	if err != nil {
		return CommentNotFound.Newf("Comment can not be updared, err %v:", err)
	}
	return nil
}
