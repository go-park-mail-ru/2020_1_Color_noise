package models

import (
	"database/sql"
	"time"
)

type Board struct {
	Id          uint
	UserId      uint
	Pins	    []*Pin
	Name        string
	Description string
	CreatedAt   time.Time
}

type DataBaseBoard struct {
	Id          uint
	UserId      uint
	Name        string
	Description sql.NullString
	CreatedAt   time.Time
}

type InputBoard struct {
	Name        string `json:"name" valid:"length(0|60), required"`
	Description string `json:"description" valid:"length(0|1000), required"`
}

type ResponseBoard struct {
	Id          uint   `json:"id"`
	UserId      uint   `json:"user_id,omitempty"`
	BoardId     uint   `json:"board_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Pins        []*Pin `json:"pins,omitempty"`
}