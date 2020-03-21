package models

import (
	"regexp"
)

type User struct {
	Id                uint
	Email 	          string
	Login             string
	EncryptedPassword string
	About             string
	Avatar            string
	Image			  []byte
}

type SignUpInput struct {
	Email    string `json:"email" valid:"email"`
	Login    string `json:"login" valid:"alphanum"`
	Password string `json:"password" valid:"length(6|100)"`
}

type SignInInput struct {
	Login    string `json:"login" valid:"alphanum"`
	Password string `json:"password" valid:"length(6|100)"`
}

type UpdateProfileInput struct {
	Email string `json:"email" valid:"email,optional"`
	Login string `json:"login" valid:"alphanum,optional"`
}

type ResponseSettingsUser struct {
	Email 	          string `json:"email,omitempty"`
	Login             string `json:"login,omitempty"`
	About             string `json:"about,omitempty"`
	Avatar            string `json:"avatar,omitempty"`
}

type ResponseUser struct {
	Login             string `json:"login"`
	About             string `json:"about,omitempty"`
	Avatar            string `json:"avatar,omitempty"`
}

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

