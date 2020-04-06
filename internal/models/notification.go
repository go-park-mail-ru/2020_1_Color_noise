package models

type Notification struct {
	User    User
	Message string
}

type ResponseNotification struct {
	User    ResponseUser `json:"user,omitempty"`
	Message string       `json:"message,omitempty"`
}
