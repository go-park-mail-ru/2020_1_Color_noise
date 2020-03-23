package models

import (
	"time"
)

type User struct {
	Id                uint
	Email 	          string
	Login             string
	EncryptedPassword string
	About             string
	Avatar            string
	Subscriptions     int
	Subscribers       int
	CreatedAt         time.Time
}

type SignUpInput struct {
	Email    string `json:"email" valid:"email,length(0|50)"`
	Login    string `json:"login" valid:"alphanum,length(0|20)"`
	Password string `json:"password" valid:"length(6|100)"`
}

type SignInInput struct {
	Login    string `json:"login" valid:"alphanum,length(0|20)"`
	Password string `json:"password" valid:"length(6|100)"`
}

type UpdateProfileInput struct {
	Email string `json:"email" valid:"email,length(0|50),optional"`
	Login string `json:"login" valid:"alphanum,length(0|20),optional"`
}

type UpdateDescriptionInput struct {
	Description string `json:"description" valid:"length(0|1000)"`
}

type UpdatePasswordInput struct {
	Password string `json:"password" valid:"length(6|70)"`
}

type ResponseUser struct {
	Id                uint   `json:"id"`
	Email 	          string `json:"email,omitempty"`
	Login             string `json:"login"`
	About             string `json:"about,omitempty"`
	Avatar            string `json:"avatar,omitempty"`
	Subscriptions     int    `json:"subscriptions"`
	Subscribers       int    `json:"subscribers"`
}

/*
func ValidateEmail(email string) bool {
	matched, _ := regexp.MatchString(`^(([^<>()\[\]\.,;:\s@\"]+(\.[^<>()\[\]\.,;:\s@\"]+)*)|(\".+\"))@(([^<>()[\]\.,;:\s@\"]+\.)+[^<>()[\]\.,;:\s@\"]{2,})$`, email)
	return matched
}

func ValidateLogin(login string) bool {
	matched, _ := regexp.MatchString(`^[\w.-_@$]+$`, login)
	return matched
}

func ValidatePassword(password string) bool {
	return len(password) > 6
}
*/

