package session

import (
	"2020_1_Color_noise/internal/models"
)

type IRepository interface {
	Add(session *models.Session) (error)
	GetByCookie(cookie string) (*models.Session, error)
	Update(session *models.Session) error
	Delete(cookie string) error
}
