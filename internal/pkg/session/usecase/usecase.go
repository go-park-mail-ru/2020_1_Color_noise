package session

import (
	"pinterest/internal/models"
	"pinterest/internal/pkg/session"
	"pinterest/internal/pkg/utils"
)

type SessionUsecase struct {
	repo  session.IRepository
}

func NewUsecase(repo session.IRepository) *SessionUsecase {
	return &SessionUsecase{
		repo: repo,
	}
}

func (su *SessionUsecase) CreateSession(id uint) (*models.Session, error) {
	cookie := utils.RandStringRunes(30)
	token := utils.RandStringRunes(30)

	session := &models.Session{
		Id: id,
		Cookie: cookie,
		Token: token,
	}

	err := su.repo.Add(session)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (su *SessionUsecase) GetByCookie(cookie string) (*models.Session, error) {
	return su.repo.GetByCookie(cookie)
}

func (su *SessionUsecase) UpdateToken(session *models.Session, token string) (error) {
	session.Token = token
	return su.repo.Update(session)
}

func (su *SessionUsecase) Delete(cookie string) error {
	return su.repo.Delete(cookie)
}