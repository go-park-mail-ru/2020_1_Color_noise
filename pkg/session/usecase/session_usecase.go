package usecase

import (
	"math/rand"
	"pinterest/pkg/models"
	repo "pinterest/pkg/session/repository"
	"time"
)

type SessionUsecase struct {
	sessionRepo  *repo.SessionRepository
}

func NewSessionUsecase(repo *repo.SessionRepository) *SessionUsecase {
	return &SessionUsecase{
		sessionRepo: repo,
	}
}

func (su *SessionUsecase) CreateSession(id uint) (*models.Session, error) {
	cookie := randStringRunes(30)
	token := su.CreateToken()

	session := &models.Session{
		Id: id,
		Cookie: cookie,
		Token: token,
	}
	err := su.sessionRepo.Add(session)
	if err != nil {
		return nil, err
	}
	return session, err
}

func (su *SessionUsecase) CreateToken() string {
	return randStringRunes(30)
}

func (su *SessionUsecase) GetByCookie(cookie string) (*models.Session, error) {
	session, err := su.sessionRepo.GetByCookie(cookie)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (su *SessionUsecase) UpdateToken(session *models.Session, token string) (error) {
	err := su.Delete(session.Cookie)
	if err != nil {
		return err
	}
	session.Token = token
	err = su.sessionRepo.Add(session)
	return err
}

func (su *SessionUsecase) Delete(cookie string) error {
	_, err := su.sessionRepo.Delete(cookie)
	return err
}

func randStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}