package database

const (
	InsertPin = "INSERT INTO pins(user_id, name, description, image, created_at, tags, views, comments) " +
		"VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id"
	InsertBoardsPin = "INSERT INTO boards_pins(image_id, board_id, original) VALUES ($1, $2, $3) RETURNING 0;"
	UpdatePin       = "UPDATE pins SET " +
		" name = $1, description = $2, board_id = $3 " +
		" WHERE id = $4"
	UpdateViews      = "UPDATE pins SET " +
		" views = views + 1" +
		" WHERE id = $1"
	UpdateComments     = "UPDATE pins SET " +
		" comments = comments + 1" +
		" WHERE id = $1"
	DeletePin = "DELETE from pins WHERE id = $1 CASCADE;"
	PinById   = "SELECT id, user_id, name, description, image, board_id, created_at, tags " +
		" FROM pins JOIN boards_pins ON pins.id = boards_pins.image_id " +
		" WHERE original = true AND id = $1"
	PinByUser = "SELECT id, user_id, name, description, image, board_id, created_at, tags " +
		" FROM pins JOIN boards_pins ON pins.id = boards_pins.image_id " +
		" WHERE original = true AND user_id = $1  ORDER BY id DESC"
	PinByName = "SELECT id, user_id, name, description, image, board_id, created_at, tags " +
		" FROM pins JOIN boards_pins ON pins.id = boards_pins.image_id " +
		" WHERE original = true AND  LOWER(name) = LOWER($1);"
	PinByBoard = "SELECT id, user_id, name, description, image, board_id, created_at, tags " +
		" FROM pins JOIN boards_pins ON pins.id = boards_pins.image_id " +
		" WHERE original = true AND  board_id = $1"
	AddTags = "UPDATE pins SET tags= $1 WHERE id = $2 RETURNING 0;"

	PopularDesc = "SELECT id, user_id, name, description, image, created_at, tags, views, comments" +
		" FROM pins  WHERE name LIKE $1  AND created_at BETWEEN $2 AND $3" +
		"ORDER BY views DESC OFFSET $4 LIMIT $5;"
	PopularAsc = "SELECT id, user_id, name, description, image, created_at, tags, views, comments" +
		" FROM pins  WHERE name LIKE $1  AND created_at BETWEEN $2 AND $3" +
		"ORDER BY views ASC OFFSET $4 LIMIT $5;"
	CommentsDesc = "SELECT id, user_id, name, description, image, created_at, tags, views, comments" +
		" FROM pins  WHERE name LIKE $1  AND created_at BETWEEN $2 AND $3" +
		"ORDER BY comments DESC OFFSET $4 LIMIT $5;"
	CommentsAsc = "SELECT id, user_id, name, description, image, created_at, tags, views, comments" +
		" FROM pins  WHERE name LIKE $1  AND created_at BETWEEN $2 AND $3" +
		"ORDER BY comments ASC OFFSET $4 LIMIT $5;"
	IdDesc = "SELECT id, user_id, name, description, image, created_at, tags, views, comments" +
		" FROM pins  WHERE name LIKE $1  AND created_at BETWEEN $2 AND $3" +
		"ORDER BY id DESC OFFSET $4 LIMIT $5;"
	IdAsc = "SELECT id, user_id, name, description, image, created_at, tags, views, comments" +
		" FROM pins  WHERE name LIKE $1  AND created_at BETWEEN $2 AND $3" +
		"ORDER BY id ASC OFFSET $4 LIMIT $5;"
)

const (
	InsertUser = "INSERT INTO users(email, login, encrypted_password, about, avatar, " +
		"subscriptions, subscribers, created_at, tags) " +
		"VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id"
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
	DeleteUser = "DELETE FROM users WHERE id = $1;"
	UserById   = "SELECT id, email, login, encrypted_password, about, avatar, subscriptions, subscribers, created_at, tags FROM users WHERE id = $1"
	//это поиск
	UserByLogin = "SELECT * FROM users WHERE LOWER(login) = LOWER($1) LIMIT $2 OFFSET $3"
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
	UpdateUnfollowA   = "UPDATE users SET subscriptions = subscriptions - 1 WHERE id = $1 RETURNING 0;"
	UpdateFollowA     = "UPDATE users SET subscriptions = subscriptions + 1 WHERE id = $1 RETURNING 0;"
	UpdateUnfollowB   = "UPDATE users SET subscribers = subscribers - 1 WHERE id = $1 RETURNING 0;"
	UpdateFollowB     = "UPDATE users SET subscribers = subscribers + 1 WHERE id =  $1 RETURNING 0;"
	AddUserTags = "UPDATE users SET tags= $1 WHERE id = $2 RETURNING 0;"
	IsFollowing =  "SELECT id FROM subscriptions WHERE user_id = $1 AND subscribed_at = $2;"
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
		" VALUES($1, $2, $3, $4) RETURNING id"
	UpdateBoard = "UPDATE boards SET " +
		" name = $1, description = $2 " +
		" WHERE id = $3"
	DeleteBoard        = "DELETE FROM boards WHERE id = $1"
	BoardById          = "SELECT * FROM boards WHERE id = $1"
	BoardsByUserId     = "SELECT * FROM boards WHERE user_id = $1 ORDER BY id ASC LIMIT $2 OFFSET $3"
	BoardsByNameSearch = "SELECT * FROM boards WHERE name = $1 LIMIT $2 OFFSET $3"
	LastPin            = "SELECT id, user_id, name, description, image, board_id, created_at " +
		" FROM pins JOIN boards_pins ON pins.id = boards_pins.image_id WHERE board_id = $1 ORDER BY created_at DESC LIMIT 1;"
)

const (
	InsertSession = "INSERT INTO sessions(" +
		" id, cookie, token, created_at, deleting_at)" +
		" VALUES ($1, $2, $3, $4, $5) RETURNING 0;"

	UpdateSession = "UPDATE sessions" +
		" SET token= $1" +
		" WHERE cookie = $2 RETURNING 0;"

	DeleteSession = "DELETE FROM sessions WHERE cookie = $1 RETURNING 0;"

	SessionByCookie = "SELECT * FROM sessions * WHERE cookie = $1; "
)

const (
	Feed = "SELECT pins.id, pins.user_id, name, description, image, board_id, created_at " +
		" FROM pins JOIN boards_pins ON pins.id = boards_pins.image_id " +
		" JOIN subscriptions ON subscriptions.subscribed_at = pins.user_id" +
		" WHERE subscriptions.user_id = $1  AND original = true ORDER BY created_at DESC LIMIT $2 OFFSET $3;"
	Main  = "SELECT id, user_id, name, description, image, board_id, created_at " +
		" FROM pins JOIN boards_pins ON pins.id = boards_pins.image_id " +
		" WHERE original = true ORDER BY created_at DESC LIMIT $1  OFFSET $2;"
	Recommendation = "SELECT id, user_id, name, description, image, board_id, created_at " +
		" FROM pins JOIN boards_pins ON pins.id = boards_pins.image_id " +
		" WHERE original = true ORDER BY created_at DESC LIMIT $1  OFFSET $2;"
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
