package models

type Pin struct {
	Id          uint   `json:"id,omitempty"`
	UserId      uint   `json:"user_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Image       []byte `json:"-"`
	ImageAdress string `json:"image"`
}
