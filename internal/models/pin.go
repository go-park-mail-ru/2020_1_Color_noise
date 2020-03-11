package models

type Pin struct {
	Id          uint   `json:"id,omitempty"`
	UserId      uint   `json:"user_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Image       string
}

type InputPin struct {
	Name        string `json:"name" valid:"required"`
	Description string `json:"description" valid:"required"`
	Image       []byte `json:"image" valid:"base64,required"`
}

type ResponsePin struct {
	Id          uint   `json:"id,omitempty"`
	UserId      uint   `json:"user_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Image       string `json:"image,omitempty"`
}
