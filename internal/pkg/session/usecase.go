package session

import "2020_1_Color_noise/internal/models"

type IUsecase interface {
	Create(id uint) (*models.Session, error)
	Login(u *models.User, password string) (*models.Session, error)
	GetByCookie(cookie string) (*models.Session, error)
	Update(session *models.Session) error
	Delete(cookie string) error
}
