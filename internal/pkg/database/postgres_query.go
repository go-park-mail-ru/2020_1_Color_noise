package database

const (
	InsertPin = "INSERT INTO pins(user_id, name, description, image, board_id, created_at) " +
		"VALUES($1, $2, $3, $4, $5, $6) RETURNING id"
	UpdatePin = "UPDATE pins SET " +
		"name = $1, description = $2, board_id = $3 " +
		"WHERE id = $4"
	DeletePin  = "DELETE from pins WHERE id = $1"
	PinById    = "SELECT * FROM pins WHERE id = $1"
	PinByUser  = "SELECT * FROM pins WHERE user_id = $1"
	PinByName  = "SELECT * FROM pins WHERE name = $1"
	PinByBoard = "SELECT * FROM pins WHERE board_id = $1"
)

const (
	InsertUser = "INSERT INTO users(email, login, encrypted_password, about, avatar, " +
		"subscriptions, subscribers, created_at) " +
		"VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id"
	UpdateUser = "UPDATE users SET " +
		"email = $1, login = $2 " +
		"WHERE id = $3 RETURNING id"
	UpdateUserDesc = "UPDATE users SET " +
		"about = $1 " +
		"WHERE id = $2 RETURNING id"
	UpdateUserPs = "UPDATE users SET " +
		"encrypted_password = $1 " +
		"WHERE id = $2  RETURNING id"
	UpdateUserAv = "UPDATE users SET " +
		"avatar = $1 " +
		"WHERE id = $2 RETURNING id;"
	DeleteUser        = "DELETE FROM users WHERE id = $1;"
	UserById          = "SELECT * FROM users WHERE id = $1"
	UserByLogin       = "SELECT * FROM users WHERE login = $1 LIMIT $2 OFFSET $3"
	UserByLoginSearch = "SELECT * FROM users WHERE login = $1"
	UserByEmail       = "SELECT * FROM users WHERE email = $1"
	//кто подписан на пользователя
	UserSubscribedUsers = "SELECT users.id, email, login, encrypted_password, about, avatar, subscriptions, subscribers, created_at  " +
		" FROM users JOIN subscriptions ON users.ID = subscriptions.user_id" +
		" WHERE subscribed_at = $1"
	//на кого подписан сам пользователь
	UserSubscriptionsUsers = "SELECT users.id, email, login, encrypted_password, about, avatar, subscriptions, subscribers, created_at  " +
		" FROM users JOIN subscriptions ON users.ID = subscriptions.subscribed_at" +
		" WHERE user_id = $1"
	UserSubscribed    = "SELECT COUNT(subscribed_at) FROM subscriptions WHERE user_id = $1"
	UserSubscriptions = "SELECT COUNT(user_id) FROM subscriptions WHERE subscribed_at = $1"
	Follow            = "INSERT INTO subscriptions( user_id, subscribed_at) VALUES ($1, $2) RETURNING id;"
	Unfollow          = "DELETE FROM public.subscriptions WHERE user_id = $1 AND subscribed_at = $2 RETURNING 0;"
)

const (
	InsertComment = "INSERT INTO commentaries(user_id, pin_id, comment, created_at) " +
		"VALUES($1, $2, $3, $4) RETURNING id"
	UpdateComment = "UPDATE commentaries SET " +
		"comment = $1, created_at = $2 " +
		"WHERE id = $3"
	DeleteComment = "DELETE FROM commentaries WHERE id = $1"
	CommentById   = "SELECT * FROM commentaries WHERE id = $1"
	CommentByText = "SELECT * FROM commentaries WHERE comment LIKE $1 LIMIT $2 OFFSET $3"
	CommentByPin  = "SELECT * FROM commentaries WHERE pin_id = $1 LIMIT $2 OFFSET $3"
)

const (
	InsertBoard = "INSERT INTO boards(user_id, name, description, created_at) " +
		"VALUES($1, $2, $3, $4) RETURNING id"
	UpdateBoard = "UPDATE boards SET " +
		"name = $1, description = $2 " +
		"WHERE id = $3"
	DeleteBoard        = "DELETE FROM boards WHERE id = $1"
	BoardById          = "SELECT * FROM boards WHERE id = $1"
	BoardsByUserId     = "SELECT * FROM boards WHERE user_id = $1 LIMIT $2 OFFSET $3"
	BoardsByNameSearch = "SELECT * FROM boards WHERE name = $1 LIMIT $2 OFFSET $3"
	LastPin            = "SELECT id, user_id, name, description, image, board_id, created_at " +
		"FROM pins WHERE board_id = $1 ORDER BY created_at DESC LIMIT 1;"
)

const (
	InsertSession = "INSERT INTO sessions(" +
		"id, cookie, token, created_at, deleting_at)" +
		" VALUES ($1, $2, $3, $4, $5) RETURNING 0;"

	UpdateSession = "UPDATE sessions" +
		" SET token= $1" +
		" WHERE cookie = $2 RETURNING 0;"

	DeleteSession = "DELETE FROM sessions WHERE cookie = $1 RETURNING 0;"

	SessionByCookie = "SELECT * FROM sessions * WHERE cookie = $1; "
)

const (
	Feed = "SELECT pins.id, pins.user_id, name, description, image, board_id, created_at " +
		" FROM subscriptions JOIN pins ON subscriptions.subscribed_at = pins.user_id" +
		" WHERE subscriptions.user_id = $1  ORDER BY created_at DESC LIMIT $2 OFFSET $3;"
	Main           = "SELECT * FROM pins ORDER BY created_at DESC LIMIT $1;"
	Recommendation = "SELECT * FROM pins ORDER BY created_at DESC LIMIT $1 OFFSET $2;"
)

const (
	GetNoti = "SELECT message, from_user_id FROM notifies WHERE user_id = $1;"
	PutNoti = "INSERT INTO notifies(" +
		"user_id, message, from_user_id, created_at) VALUES($1, $2, $3, $4) RETURNING id;"
)

const (
	AddChat  = "INSERT INTO chats(sender_id, receiver_id) VALUES ($1, $2) RETURNING id;"
	GetChats = "SELECT receiver_id FROM chats WHERE sender_id = $1 LIMIT $2 OFFSET $3;"

	AddMsg = "INSERT INTO public.chat_messages(sender_id, receiver_id, message, created_at) " +
		"VALUES ($1, $2, $3, $4);"
	GetMsg = "SELECT sender_id, receiver_id, message FROM chat_messages " +
		" WHERE sender_id = $1 AND receiver_id = $2 ORDER BY created_at DESC LIMIT $3 OFFSET $4;"
)
