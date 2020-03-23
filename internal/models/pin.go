package models

import "time"

type Pin struct {
	Id          uint
	UserId      uint
	BoardId	    uint
	Name        string
	Description string
	Image       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type InputPin struct {
	BoardId     uint   `json:"board_id" valid:"int"`
	Name        string `json:"name" valid:"length(1|60), required"`
	Description string `json:"description" valid:"length(0|1000), required"`
	Image       string `json:"image" valid:"datauri,required"`
}

type UpdatePin struct {
	BoardId     uint   `json:"board_id" valid:"int"`
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
