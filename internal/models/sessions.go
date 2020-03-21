package models

import "time"

type Session struct {
	Id         uint
	Cookie     string
	Token      string
	CreatedAt  time.Time
	DeletingAt time.Time
}
