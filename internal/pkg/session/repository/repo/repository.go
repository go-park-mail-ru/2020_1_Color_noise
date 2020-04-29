package repo

import (
	"2020_1_Color_noise/internal/models"
	"2020_1_Color_noise/internal/pkg/database"
	. "2020_1_Color_noise/internal/pkg/error"
)

type Repository struct {
	db  database.DBInterface
}

func NewRepo(d database.DBInterface) *Repository {
	return &Repository{
		db: d,
	}
}

func (sr *Repository) Add(session *models.Session) error {
	sr.db.CreateSession(models.GetBSession(*session))
	return nil
}

func (sr *Repository) GetByCookie(cookie string) (*models.Session, error) {
	s := models.DataBaseSession{
		Cookie: cookie,
	}
	session, isExist := sr.db.GetSessionByCookie(s)
	if isExist != nil {
		return nil, Unauthorized.Newf("Session is not found, cookie: %s", cookie)
	}

	return &session, nil
}

func (sr *Repository) Update(session *models.Session) error {
	_ = sr.db.UpdateSession(models.GetBSession(*session))
	return nil
}

func (sr *Repository) Delete(cookie string) error {
	dbs := models.DataBaseSession{
		Cookie: cookie,
	}
	sr.db.DeleteSession(dbs)
	return nil
}
