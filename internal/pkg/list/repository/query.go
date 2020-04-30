package repository

const (
	Feed = "SELECT pins.id, pins.user_id, name, description, image, board_id, created_at " +
		" FROM subscriptions JOIN pins ON subscriptions.subscribed_at = pins.user_id" +
		" WHERE subscriptions.user_id = $1  ORDER BY created_at DESC LIMIT $2 OFFSET $3;"
	Main           = "SELECT * FROM pins ORDER BY created_at DESC LIMIT $1  OFFSET $2;"
	Recommendation = "SELECT * FROM pins ORDER BY created_at DESC LIMIT $1 OFFSET $2;"
)
