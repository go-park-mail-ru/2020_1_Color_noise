package models

import (
	"database/sql"
	"time"
)

type Session struct {
	Id         uint
	Cookie     string
	Token      string
	CreatedAt  time.Time
	DeletingAt time.Time
}

type DataBaseSession struct {
	Id         uint
	Cookie     string
	Token      string
	CreatedAt  time.Time
	DeletingAt sql.NullTime
}

func GetSession(s DataBaseSession) Session {
	tmp := Session{
		Id:        s.Id,
		Cookie:    s.Cookie,
		Token:     s.Token,
		CreatedAt: s.CreatedAt,
	}

	if s.DeletingAt.Valid {
		tmp.DeletingAt = s.DeletingAt.Time
	}
	return tmp
}

func GetBSession(s Session) DataBaseSession {
	tmp := DataBaseSession{
		Id:        s.Id,
		Cookie:    s.Cookie,
		Token:     s.Token,
		CreatedAt: s.CreatedAt,
	}

	if !s.DeletingAt.IsZero() {
		tmp.DeletingAt.Valid = true
		tmp.DeletingAt.Time = s.DeletingAt
	}
	return tmp
}
