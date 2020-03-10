package repository

import (
	"pinterest/internal/models"
	. "pinterest/internal/pkg/error"
)

type Repository struct {
	data map[string]*models.Session
}

func NewRepo() *Repository {
	return &Repository{
		data: make(map[string]*models.Session),
	}
}

func (sr *Repository) Add(session *models.Session) error {
	sr.data[session.Cookie] = session
	return nil
}

func (sr *Repository) GetByCookie(cookie string) (*models.Session, error) {
	session, isExist := sr.data[cookie]
	if isExist {
		return session, nil
	}

	return nil, BadCookie.Newf("Session not found, cookie: %s", cookie)
}

func (sr *Repository) Update(session *models.Session) error {
	sr.data[session.Cookie] = session
	return nil
}

func (sr *Repository) Delete(cookie string) error {
	delete(sr.data, cookie)
	return nil
}