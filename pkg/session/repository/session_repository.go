package repository


import (
	"fmt"
	"pinterest/pkg/models"
)

type SessionRepository struct {
	data map[string]uint
}

func NewSessionRepo() *SessionRepository {
	return &SessionRepository{
		data: make(map[string]uint),
	}
}

func (sr *SessionRepository) Add(session *models.Session) (error) {
	sr.data[session.Cookie] = session.Id
	fmt.Println(sr.data)
	return nil
}

func (sr *SessionRepository) GetByCookie(cookie string) (*models.Session, error) {
	id, isExist := sr.data[cookie]
	if isExist {
		return &models.Session{Id: id, Cookie: cookie}, nil
	}
	return nil, fmt.Errorf("Not found")
}

func (sr *SessionRepository) Delete(cookie string) (bool, error) {
	delete(sr.data, cookie)
	return true, nil
}