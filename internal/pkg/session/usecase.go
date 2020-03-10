package session

import "pinterest/internal/models"

type IUsecase interface {
	CreateSession(id uint) (*models.Session, error)
	GetByCookie(cookie string) (*models.Session, error)
	UpdateToken(session *models.Session, token string) error
	Delete(cookie string) error
}
