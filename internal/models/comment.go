package models

import "time"

type Comment struct {
	Id        uint
	User      *User
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
	PinId uint   `json:"pin_id" valid:"int"`
	Text  string `json:"comment" valid:"length(1|2000), required"`
}

type ResponseComment struct {
	Id        uint       `json:"id,omitempty"`
	User      *User      `json:"user,omitempty"`
	PindId    uint       `json:"pin_id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Text      string     `json:"comment,omitempty"`
}

func GetBComment(c Comment) DataBaseComment {
	tmp := DataBaseComment{
	}

	tmp.Id = c.Id
	tmp.PinId = c.PinId
	tmp.Text = c.Text
	tmp.UserId = c.User.Id
	return tmp
}

func GetComment(c DataBaseComment) Comment {
	tmp := Comment{
	}

	tmp.Id = c.Id
	tmp.PinId = c.PinId
	tmp.Text = c.Text
	return tmp
}
