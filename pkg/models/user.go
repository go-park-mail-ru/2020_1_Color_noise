package models

type User struct {
	Id                uint   `json:"id"`
	Email 	          string `json:"email,omitempty"`
	Login             string `json:"login"`
	Password          string `json:"password,omitempty"`
	EncryptedPassword string `json:"-"`
	About             string `json:"about,omitempty"`
	Avatar            string `json:"avatar,omitempty"`
	Image			  []byte  `json:"-"`
}


