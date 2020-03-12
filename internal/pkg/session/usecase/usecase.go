package session

import (
	"pinterest/internal/models"
	. "pinterest/internal/pkg/error"
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

	if err := su.repo.Add(session); err != nil {
		return nil, Wrapf(err, "Creating session error, id: %d", id)
	}

	return session, nil
}

func (su *SessionUsecase) GetByCookie(cookie string) (*models.Session, error) {
	session, err := su.repo.GetByCookie(cookie)
	if err != nil {
		return nil, Wrapf(err, "Getting session by cookie error, cookie: %d", cookie)
	}
	return session, nil
}

func (su *SessionUsecase) UpdateToken(session *models.Session, token string) (error) {
	session.Token = token

	if err := su.repo.Update(session); err != nil {
		return Wrapf(err, "Updating token error, sess_id: %d", session.Id)
	}

	return nil
}

func (su *SessionUsecase) Delete(cookie string) error {
	if err := su.repo.Delete(cookie); err != nil {
		return Wrapf(err, "Deleting session error, cookie: %d", cookie)
	}

	return nil
}