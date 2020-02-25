package models

type Pin struct {
	Id          uint   `json:"id"`
	UserId      uint   `json:"user_id"`
	Name        string `json:"name"`
	Data        []byte `json:"-"`
	ImageAdress string `json:"image"`
	Description string `json:"description"`
}
