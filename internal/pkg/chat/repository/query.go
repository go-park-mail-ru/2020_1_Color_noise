package repository

const (
	AddChat  = "INSERT INTO chats(sender_id, receiver_id) VALUES ($1, $2) RETURNING id;"
	GetChats = "SELECT sender_id, receiver_id FROM chat_messages WHERE sender_id = $1 OR receiver_id = $1 " +
		" GROUP BY sender_id, receiver_id LIMIT $2 OFFSET $3;"

	AddMsg = "INSERT INTO chat_messages(sender_id, receiver_id, message, created_at) " +
		"VALUES ($1, $2, $3, $4) RETURNING 0;"
	GetMsg = "SELECT sender_id, message, created_at FROM chat_messages " +
		" WHERE sender_id = $1 AND receiver_id = $2 OR sender_id = $2 AND receiver_id = $1 " +
		" ORDER BY created_at ASC LIMIT $3 OFFSET $4;"
)
