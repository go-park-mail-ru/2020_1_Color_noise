package models

import (
	"database/sql"
	"time"
)

type Board struct {
	Id          uint
	UserId      uint
	Pins        []*Pin
	Name        string
	Description string
	CreatedAt   time.Time
	LastPin     Pin
}

type DataBaseBoard struct {
	Id          uint
	UserId      uint
	Name        string
	Description sql.NullString
	CreatedAt   time.Time
}

type InputBoard struct {
	Name        string `json:"name" valid:"length(1|60), required"`
	Description string `json:"description" valid:"length(0|1000), required"`
}

type ResponseBoard struct {
	Id          uint           `json:"id"`
	UserId      uint           `json:"user_id,omitempty"`
	Name        string         `json:"name,omitempty"`
	Description string         `json:"description"`
	Pins        []*ResponsePin `json:"pins,omitempty"`
	LastPin     *ResponsePin   `json:"last_pin,omitempty"`
}

func GetDBoard(board Board) DataBaseBoard {
	tmp := DataBaseBoard{
		Id:        board.Id,
		UserId:    board.UserId,
		Name:      board.Name,
		CreatedAt: board.CreatedAt,
	}

	if board.Description != "" {
		tmp.Description.Valid = true
		tmp.Description.String = board.Description
	}
	return tmp
}

func GetBoard(board DataBaseBoard) Board {
	tmp := Board{
		Id:        board.Id,
		UserId:    board.UserId,
		Name:      board.Name,
		CreatedAt: board.CreatedAt,
	}

	if board.Description.Valid {
		tmp.Description = board.Description.String
	}
	return tmp
}
