package models

import (
	"database/sql"
	"time"
)

type Pin struct {
	Id          uint
	UserId      uint
	BoardId     uint
	Name        string
	Description string
	Image       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
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
}

type InputPin struct {
	BoardId     int   `json:"board_id" valid:"int"`
	Name        string `json:"name" valid:"length(1|60), required"`
	Description string `json:"description" valid:"length(0|1000), required"`
	Image       string `json:"image" valid:"datauri,required"`
}

type UpdatePin struct {
	BoardId     int   `json:"board_id" valid:"int"`
	Name        string `json:"name" valid:"length(1|60), required"`
	Description string `json:"description" valid:"length(0|1000), required"`
}

type ResponsePin struct {
	Id          uint   `json:"id,omitempty"`
	UserId      uint   `json:"user_id,omitempty"`
	BoardId     uint   `json:"board_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Image       string `json:"image,omitempty"`

}

func GetPin(pin DataBasePin) Pin {
	tmp := Pin{
		Id:        pin.Id,
		UserId:    pin.UserId,
		BoardId:   pin.BoardId,
		Name:      pin.Name,
		Image:     pin.Image,
		CreatedAt: pin.CreatedAt,
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
		UserId:    pin.UserId,
		BoardId:   pin.BoardId,
		Name:      pin.Name,
		Image:     pin.Image,
		CreatedAt: pin.CreatedAt,
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
