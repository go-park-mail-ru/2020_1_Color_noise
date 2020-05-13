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
	PinByName  = "SELECT * FROM pins WHERE LOWER(name) = LOWER($1);"
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
	UserById          = "SELECT id, email, login, encrypted_password, about, avatar, subscriptions, subscribers, created_at FROM users WHERE id = $1"
	//это поиск
	UserByLogin       = "SELECT * FROM users WHERE LOWER(login) = LOWER($1) LIMIT $2 OFFSET $3"
	//это точный поиск
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
	UserSubscribed    = "SELECT subscribers FROM users WHERE id = $1;"
	UserSubscriptions = "SELECT subscriptions FROM users WHERE id = $1;"
	Follow            = "INSERT INTO subscriptions( user_id, subscribed_at) VALUES ($1, $2) RETURNING id;"
	Unfollow          = "DELETE FROM subscriptions WHERE user_id = $1 AND subscribed_at = $2 RETURNING 0;"
	UpdateUnfollowA = "UPDATE users SET subscriptions = subscriptions - 1 WHERE id = $1 RETURNING 0;"
	UpdateFollowA = "UPDATE users SET subscriptions = subscriptions + 1 WHERE id = $1 RETURNING 0;"
	UpdateUnfollowB = "UPDATE users SET subscribers = subscribers - 1 WHERE id = $1 RETURNING 0;"
	UpdateFollowB = "UPDATE users SET subscribers = subscribers + 1 WHERE id =  $1 RETURNING 0;"
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
	BoardsByUserId     = "SELECT * FROM boards WHERE user_id = $1 LIMIT $2 OFFSET $3 ORDER BY id ASC"
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
	Main           = "SELECT * FROM pins ORDER BY created_at DESC LIMIT $1  OFFSET $2;"
	Recommendation = "SELECT * FROM pins ORDER BY created_at DESC LIMIT $1 OFFSET $2;"
)

const (
	GetNoti = "SELECT message, from_user_id FROM notifies WHERE user_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3;"
	PutNoti = "INSERT INTO notifies(" +
		"user_id, message, from_user_id, created_at) VALUES($1, $2, $3, $4) RETURNING id;"
)

const (
	AddChat  = "INSERT INTO chats(sender_id, receiver_id) VALUES ($1, $2) RETURNING id;"
	GetChats = "SELECT sender_id, receiver_id FROM chat_messages WHERE sender_id = $1 OR receiver_id = $1 " +
		" GROUP BY sender_id, receiver_id LIMIT $2 OFFSET $3;"

	AddMsg = "INSERT INTO chat_messages(sender_id, receiver_id, message, sticker, created_at) " +
		"VALUES ($1, $2, $3, $4, $5) RETURNING 0;"
	GetMsg = "SELECT sender_id, message, sticker, created_at FROM chat_messages " +
		" WHERE sender_id = $1 AND receiver_id = $2 OR sender_id = $2 AND receiver_id = $1 " +
		" ORDER BY created_at ASC LIMIT $3 OFFSET $4;"
)
