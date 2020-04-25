package models

import "time"

type Message struct {
	SendUser  *User
	RecUser   *User
	Message   string
	CreatedAt time.Time
}

type InputMessage struct {
	UserRecviredId uint    `json:"user_id" valid:"int"`
	Message        string `json:"message" valid:"string"`
}

type ResponseMessage struct {
	SendUser  *ResponseUser `json:"user_send"`
	Message   string        `json:"message"`
	CreatedAt time.Time     `json:"created_at"`
}

type ResponseMessages struct {
	SendUser  *ResponseUser `json:"user_send"`
	RecUser   *ResponseUser `json:"user_rec"`
	Message   string        `json:"message"`
	CreatedAt time.Time     `json:"created_at"`
}
