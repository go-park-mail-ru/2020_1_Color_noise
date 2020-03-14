package models

type Pin struct {
	Id          uint
	UserId      uint
	Name        string
	Description string
	Image       string
}

type InputPin struct {
	Name        string `json:"name" valid:"required"`
	Description string `json:"description" valid:"required"`
	Image       string `json:"image" valid:"datauri,required"`
}

type ResponsePin struct {
	Id          uint   `json:"id,omitempty"`
	UserId      uint   `json:"user_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Image       string `json:"image,omitempty"`
}
