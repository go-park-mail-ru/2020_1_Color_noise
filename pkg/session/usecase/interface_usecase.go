package usecase

import "pinterest/pkg/models"

type ISessionUsecase interface {
	CreateSession(id uint) (*models.Session, error)
	CreateToken() string
	GetByCookie(cookie string) (*models.Session, error)
	UpdateToken(session *models.Session, token string) error
	Delete(cookie string) error
}
