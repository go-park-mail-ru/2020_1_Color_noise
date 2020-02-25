package usecase

import (
	"math/rand"
	"pinterest/pkg/models"
	repo "pinterest/pkg/session/repository"
)

type SessionUsecase struct {
	sessionRepo  *repo.SessionRepository
}

func NewSessionUsecase(repo *repo.SessionRepository) *SessionUsecase {
	return &SessionUsecase{
		sessionRepo: repo,
	}
}
/*
func (sh *SessionUsecase) IsActive (cookie string) (bool, error) {
	session, err := sh.GetByCookie(cookie)
	if err != nil {
		return false, err
	}
	return true, nil
}*/

func (su *SessionUsecase) CreateSession(id uint) (*models.Session, error) {
	cookie := randStringRunes(30)

	session := &models.Session{
		Id: id,
		Cookie: cookie,
	}

	err := su.sessionRepo.Add(session)
	if err != nil {
		return nil, err
	}
	return session, err
}

func (su *SessionUsecase) GetByCookie(cookie string) (*models.Session, error) {
	session, err := su.sessionRepo.GetByCookie(cookie)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (su *SessionUsecase) Delete(cookie string) error {
	_, err := su.sessionRepo.Delete(cookie)
	return err
}

var (
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}