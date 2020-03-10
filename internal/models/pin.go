package models

type Pin struct {
	Id          uint   `json:"id,omitempty"`
	UserId      uint   `json:"user_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Image       string
}

type InputPin struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       []byte `json:"image"`
}

type ResponsePin struct {
	Id          uint   `json:"id,omitempty"`
	UserId      uint   `json:"user_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Image       string `json:"image,omitempty"`
}
