package repository


import (
	"fmt"
	"pinterest/pkg/models"
)

type SessionRepository struct {
	data map[string]*models.Session
}

func NewSessionRepo() *SessionRepository {
	return &SessionRepository{
		data: make(map[string]*models.Session),
	}
}

func (sr *SessionRepository) Add(session *models.Session) (error) {
	sr.data[session.Cookie] = session
	return nil
}

func (sr *SessionRepository) GetByCookie(cookie string) (*models.Session, error) {
	session, isExist := sr.data[cookie]
	if isExist {
		return session, nil
	}
	return nil, fmt.Errorf("Not found")
}

func (sr *SessionRepository) Delete(cookie string) (bool, error) {
	delete(sr.data, cookie)
	return true, nil
}