package models

import (
	"database/sql"
	"time"
)

type Message struct {
	SendUser  *User
	RecUser   *User
	Message   string
	Stickers  string
	CreatedAt time.Time
}

type DMessage struct {
	SendUser  *User
	RecUser   *User
	Message   sql.NullString
	Stickers  sql.NullString
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

func GetDMessage(message Message)  DMessage{
	tmp := DMessage{
		SendUser:  message.SendUser,
		RecUser:   message.RecUser,
		CreatedAt: message.CreatedAt,
	}

	if message.Message != "" {
		tmp.Message.Valid = true
		tmp.Message.String = message.Message
	}

	if message.Stickers != "" {
		tmp.Stickers.Valid = true
		tmp.Stickers.String = message.Stickers
	}

	return tmp
}

func GetMessage(message DMessage)  Message{
	tmp := Message{
		SendUser:  message.SendUser,
		RecUser:   message.RecUser,
		CreatedAt: message.CreatedAt,
	}

	if message.Message.Valid {
		tmp.Message = message.Message.String
	}

	if message.Stickers.Valid {
		tmp.Stickers = message.Stickers.String
	}

	return tmp
}