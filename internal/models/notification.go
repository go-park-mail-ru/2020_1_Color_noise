package models

type Notification struct {
	User    User
	Message string
}

type ResponseNotification struct {
	User    User   `json:"user,omitempty"`
	Message string `json:"message,omitempty"`
}
