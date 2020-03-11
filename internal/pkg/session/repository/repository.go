package repository

import (
	"pinterest/internal/models"
	. "pinterest/internal/pkg/error"
	"sync"
)

type Repository struct {
	data map[string]*models.Session
	mu   *sync.Mutex
}

func NewRepo() *Repository {
	return &Repository{
		data: make(map[string]*models.Session),
		mu:   &sync.Mutex{},
	}
}

func (sr *Repository) Add(session *models.Session) error {
	sr.mu.Lock()
	sr.data[session.Cookie] = session
	sr.mu.Unlock()
	return nil
}

func (sr *Repository) GetByCookie(cookie string) (*models.Session, error) {
	session, isExist := sr.data[cookie]
	if isExist {
		return session, nil
	}

	return nil, Unauthorized.Newf("Session is not found, cookie: %s", cookie)
}

func (sr *Repository) Update(session *models.Session) error {
	sr.mu.Lock()
	sr.data[session.Cookie] = session
	sr.mu.Unlock()
	return nil
}

func (sr *Repository) Delete(cookie string) error {
	sr.mu.Lock()
	delete(sr.data, cookie)
	sr.mu.Unlock()
	return nil
}