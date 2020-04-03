package database

const (
	InsertPin = "INSERT INTO pins(user_id, name, description, image, board_id, created_at) " +
		"VALUES($1, $2, $3, $4, $5, $6) RETURNING id"
	UpdatePin = "UPDATE pins SET " +
		"name = $1, description = $2, board_id = $3 " +
		"WHERE id = $4"
	DeletePin = "DELETE from pins WHERE id = $1"
	PinById   = "SELECT * FROM pins WHERE id = $1"
	PinByUser = "SELECT * FROM pins WHERE user_id = $1"
	PinByName = "SELECT * FROM pins WHERE name = $1"
)

const (
	InsertUser = "INSERT INTO users(email, login, encrypted_password, about, avatar, " +
		"subscriptions, subscribers, created_at) " +
		"VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id"
	UpdateUser = "UPDATE users SET " +
		"email = $1, login = $2 " +
		"WHERE id = $3"
	UpdateUserDesc = "UPDATE users SET " +
		"about = $1 " +
		"WHERE id = $2"
	UpdateUserPs = "UPDATE users SET " +
		"encrypted_password = $1 " +
		"WHERE id = $2"
	UpdateUserAv = "UPDATE users SET " +
		"avatar = $1 " +
		"WHERE id = $2"
	DeleteUser        = "DELETE FROM users WHERE id = $1 CASCADE"
	UserById          = "SELECT * FROM users WHERE id = $1"
	UserByLogin       = "SELECT * FROM users WHERE login = $1 LIMIT $2 OFFSET $3"
	UserByLoginSearch = "SELECT * FROM users WHERE login = $1"
	UserByEmail       = "SELECT * FROM users WHERE email = $1"
	UserSubscribed    = "SELECT COUNT(subscribed_at) FROM subscriptions WHERE user_id = $1"
	UserSubscriptions = "SELECT COUNT(user_id) FROM subscriptions WHERE subscribed_at = $1"
	Follow            = "INSERT INTO subscriptions( user_id, subscribed_at) VALUES ($1, $2, $3);"
	Unfollow          = "DELETE FROM public.subscriptions WHERE user_id $1 = AND subscribed_at = $2;"
)

const (
	InsertComment = "INSERT INTO commentaries(user_id, pin_id, comment, created_at) " +
		"VALUES($1, $2, $3, $4) RETURNING id"
	UpdateComment = "UPDATE commentaries SET " +
		"comment = $1, created_at = $2 " +
		"WHERE id = $3"
	DeleteComment = "DELETE FROM commentaries WHERE id = $1"
	CommentByPin  = "SELECT * FROM commentaries WHERE pin_id = $1"
)

const (
	InsertBoard = "INSERT INTO boards(user_id, name, description, created_at) " +
		"VALUES($1, $2, $3, $4) RETURNING id"
	UpdateBoard = "UPDATE boards SET " +
		"name = $1, description = $2 " +
		"WHERE id = $3"
	DeleteBoard        = "DELETE FROM boards WHERE id = $1"
	BoardById          = "SELECT * FROM boards WHERE id = $1"
	BoardsByUserId     = "SELECT * FROM boards WHERE user_id = $1"
	BoardsByNameSearch = "SELECT * FROM boards WHERE name = $1"
)
