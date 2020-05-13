package models

import "time"

type Message struct {
	SendUser  *User
	RecUser   *User
	Message   string
	Stickers  string
	CreatedAt time.Time
}

type InputMessage struct {
	UserRecviredId uint   `json:"user_id" valid:"int"`
	Message        string `json:"message" valid:"string"`
	Stickers       string `json:"stickers"`
}


type ResponseMessage struct {
	SendUser  *ResponseUser `json:"user_send"`
	RecUser   *ResponseUser `json:"user_rec"`
	Message   string        `json:"message"`
	Stickers  string        `json:"stickers"`
	CreatedAt time.Time     `json:"created_at"`
}
