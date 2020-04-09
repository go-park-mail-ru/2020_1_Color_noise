package database

import (
	"2020_1_Color_noise/internal/models"
)

const (
	IntegrityConstraintViolation = "23000"
	RestrictViolation            = "23001"
	NotNullViolation             = "23502"
	ForeignKeyViolation          = "23503"
	UniqueViolation              = "23505"
	CheckViolation               = "23514"
)

func (db *PgxDB) CreateSession(s models.DataBaseSession) error {
	res := db.dbPool.QueryRow(InsertSession, s.Id, s.Cookie, s.Token, s.CreatedAt, s.DeletingAt)
	var i int
	return res.Scan(&i)
}

func (db *PgxDB) UpdateSession(s models.DataBaseSession) error {
	var tmp int
	return db.dbPool.QueryRow(UpdateSession, s.Token, s.Cookie).Scan(&tmp)
}

func (db *PgxDB) DeleteSession(s models.DataBaseSession) error {
	var tmp int
	return db.dbPool.QueryRow(DeleteSession, s.Cookie).Scan(&tmp)
}

func (db *PgxDB) GetSessionByCookie(s models.DataBaseSession) (models.Session, error) {
	session := models.DataBaseSession{}
	row := db.dbPool.QueryRow(SessionByCookie, s.Cookie)

	err := row.Scan(&session.Id, &session.Cookie, &session.Token,
		&session.CreatedAt, &session.DeletingAt)

	if err != nil {
		return models.Session{}, err
	}
	return models.GetSession(session), nil
}
