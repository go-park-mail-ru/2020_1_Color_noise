package models

import (
	"database/sql"
	"time"
)

type User struct {
	Id                uint
	Email             string
	Login             string
	EncryptedPassword string
	About             string
	Avatar            string
	Subscriptions     int
	Subscribers       int
	Preferences       []string
	CreatedAt         time.Time
	Tags              []string
}

type DataBaseUser struct {
	Id                uint
	Email             string
	Login             string
	EncryptedPassword string
	About             sql.NullString
	Avatar            sql.NullString
	Subscriptions     int
	Subscribers       int
	CreatedAt         time.Time
	Tags              []string
}

type SignUpInput struct {
	Email    string `json:"email" valid:"email,length(1|50)"`
	Login    string `json:"login" valid:"alphanum,length(1|20)"`
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
	Id            uint   `json:"id"`
	Email         string `json:"email,omitempty"`
	Login         string `json:"login"`
	About         string `json:"about"`
	Avatar        string `json:"avatar,omitempty"`
	Subscriptions int    `json:"subscriptions"`
	Subscribers   int    `json:"subscribers"`
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

func GetUser(u DataBaseUser) User {
	tmp := User{
		Id:                u.Id,
		Email:             u.Email,
		Login:             u.Login,
		EncryptedPassword: u.EncryptedPassword,
		Subscriptions:     u.Subscriptions,
		Subscribers:       u.Subscribers,
		CreatedAt:         u.CreatedAt,
		Tags:              u.Tags,
	}

	if u.Avatar.Valid {
		tmp.Avatar = u.Avatar.String
	}
	if u.About.Valid {
		tmp.About = u.About.String
	}

	return tmp
}

func GetBUser(u User) DataBaseUser {
	tmp := DataBaseUser{
		Id:                u.Id,
		Email:             u.Email,
		Login:             u.Login,
		EncryptedPassword: u.EncryptedPassword,
		Subscriptions:     u.Subscriptions,
		Subscribers:       u.Subscribers,
		CreatedAt:         u.CreatedAt,
		Tags:              u.Tags,
	}

	if u.Avatar != "" {
		tmp.Avatar.String = u.Avatar
		tmp.Avatar.Valid = true
	}
	if u.About != "" {
		tmp.About.String = u.About
		tmp.About.Valid = true
	}
	return tmp
}
