package models

import "regexp"

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
	Email    string `json:"email"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type SignInInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UpdateInput struct {
	Email string `json:"email,omitempty"`
	Login string `json:"login,omitempty"`
	About string `json:"about,omitempty"`
}

type UpdatePasswordInput struct {
	NewPassword     string `json:"password,omitempty"`
	ConfirmPassword string `json:"confirm,omitempty"`
}

type ResponseSettingsUser struct {
	Email 	          string `json:"email,omitempty"`
	Login             string `json:"login"`
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

