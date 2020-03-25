package models

import "time"

type Comment struct {
	Id        uint
	UserId    uint
	PinId     uint
	Text      string
	CreatedAt time.Time
}

type DataBaseComment struct {
	Id        uint
	UserId    uint
	PinId     uint
	Text      string
	CreatedAt time.Time
}

type InputComment struct {
	PinId     uint   `json:"pin_id" valid:"int"`
	Text      string `json:"comment" valid:"length(1|2000), required"`
}

type ResponseComment struct {
	Id        uint      `json:"id,omitempty"`
	UserId    uint      `json:"user_id,omitempty"`
	PindId    uint      `json:"pin_id,omitempty"`
	CreatedAt time.Time `json:"time"`
	Text      string    `json:"comment"`
}

