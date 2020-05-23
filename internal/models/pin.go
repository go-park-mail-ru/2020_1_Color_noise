package models

import (
	"database/sql"
	"time"
)

type Pin struct {
	Id          uint
	User        *User
	BoardId     uint
	Name        string
	Description string
	Image       string
	Tags        []string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Views       uint
	Comment     uint
}

type DataBasePin struct {
	Id          uint
	UserId      uint
	BoardId     uint
	Name        string
	Description sql.NullString //не гарантируется, что есть описание
	Image       string
	CreatedAt   time.Time
	UpdatedAt   sql.NullTime //не гарантируется, что пин был обновлен
	Tags        []string
	Views       uint
	Comment     uint
}

type InputPin struct {
	BoardId     int    `json:"board_id" valid:"int"`
	Name        string `json:"name" valid:"length(1|60), required"`
	Description string `json:"description" valid:"length(0|1000), required"`
	Image       string `json:"image" valid:"datauri,required"`
}

type UpdatePin struct {
	BoardId     int    `json:"board_id" valid:"int"`
	Name        string `json:"name" valid:"length(1|60), required"`
	Description string `json:"description" valid:"length(1|1000), required"`
}

type ResponsePin struct {
	Id          uint     `json:"id,omitempty"`
	User        *User    `json:"user,omitempty"`
	BoardId     uint     `json:"board_id,omitempty"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description"`
	Image       string   `json:"image,omitempty"`
	Tags        []string `json:"tags,omitempty"`
}

func GetPin(pin DataBasePin) Pin {
	tmp := Pin{
		Id:        pin.Id,
		BoardId:   pin.BoardId,
		Name:      pin.Name,
		Image:     pin.Image,
		CreatedAt: pin.CreatedAt,
		Tags:      pin.Tags,
		Views:     pin.Views,
		Comment:   pin.Comment,
	}

	if pin.Description.Valid {
		tmp.Description = pin.Description.String
	}

	if pin.UpdatedAt.Valid {
		tmp.UpdatedAt = pin.UpdatedAt.Time
	}
	return tmp
}

func GetBPin(pin Pin) DataBasePin {
	tmp := DataBasePin{
		Id:        pin.Id,
		BoardId:   pin.BoardId,
		Name:      pin.Name,
		Image:     pin.Image,
		CreatedAt: pin.CreatedAt,
		Tags:      pin.Tags,
		Views:     pin.Views,
		Comment:   pin.Comment,
	}
	if !pin.UpdatedAt.IsZero() {
		tmp.UpdatedAt.Valid = true
		tmp.UpdatedAt.Time = pin.UpdatedAt
	}

	if pin.Description != "" {
		tmp.Description.Valid = true
		tmp.Description.String = pin.Description
	}
	return tmp
}
