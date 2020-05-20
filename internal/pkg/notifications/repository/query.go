package repository

const (
	GetNoti = "SELECT message, from_user_id FROM notifies WHERE user_id = $1 ORDER BY created_at DESC;"
	PutNoti = "INSERT INTO notifies(" +
		"user_id, message, from_user_id, created_at) VALUES($1, $2, $3, $4) RETURNING id;"
)
