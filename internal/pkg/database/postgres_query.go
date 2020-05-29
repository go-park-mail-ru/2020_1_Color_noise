package database

const (
	InsertPin = "INSERT INTO pins(user_id, name, description, image, created_at, tags, views, comments) " +
		"VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id"

	InsertImage = "INSERT INTO pins(user_id, image, created_at, tags, views, comments, visible) " +
		"VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id"

	PinCreation = "UPDATE pins SET " +
		" name = $2, description = $3, visible = $4" +
		" WHERE id = $1 and user_id = $5"

	InsertBoardsPin = "INSERT INTO boards_pins(image_id, board_id, original) VALUES ($1, $2, $3) RETURNING 0;"
	UpdatePin       = "UPDATE pins SET " +
		" name = $1, description = $2, board_id = $3 " +
		" WHERE id = $4"
	UpdateViews = "UPDATE pins SET " +
		" views = views + 1" +
		" WHERE id = $1"
	UpdateComments = "UPDATE pins SET " +
		" comments = comments + 1" +
		" WHERE id = $1"
	DeletePin = "DELETE from pins CASCADE WHERE id = $1 AND user_id = $2"
	PinById   = "SELECT pins.id, name, description, image, board_id, pins.created_at, pins.tags, users.id, users.login, users.avatar" +
		" FROM pins JOIN boards_pins ON pins.id = boards_pins.image_id" +
		" JOIN users ON pins.user_id = users.id" +
		" WHERE original = true AND pins.id = $1 AND visible = true"
	ImageById   = "SELECT pins.id, name, description, image, pins.created_at, pins.tags, " +
		" users.id, users.login, users.avatar" +
		" FROM pins JOIN users ON pins.user_id = users.id" +
		" WHERE pins.id = $1 "

	PinByUser =  "SELECT pins.id, name, description, image, board_id, pins.created_at, pins.tags, users.id, users.login, users.avatar" +
		" FROM pins JOIN boards_pins ON pins.id = boards_pins.image_id" +
		" JOIN users ON pins.user_id = users.id" +
		" WHERE original = true AND visible = true AND user_id = $1  ORDER BY pins.id DESC OFFSET $2 LIMIT $3;"
	PinByBoard =  "SELECT pins.id, name, description, image, board_id, pins.created_at, pins.tags, users.id, users.login, users.avatar" +
		" FROM pins JOIN boards_pins ON pins.id = boards_pins.image_id" +
		" JOIN users ON pins.user_id = users.id" +
		" WHERE board_id = $1 AND visible = true"
	AddTags = "UPDATE pins SET tags= $1 WHERE id = $2 RETURNING 0;"

	PopularDesc = "SELECT pins.id, name, description, image, pins.created_at, pins.tags, pins.views, pins.comments," +
		" users.id, users.login, users.avatar" +
		" FROM pins " +
		" JOIN users ON pins.user_id = users.id" +
		" WHERE make_tsvector(name) @@ to_tsquery($1)" +
		" AND pins.created_at BETWEEN $2 AND $3 AND visible = true" +
		" ORDER BY pins.views DESC OFFSET $4 LIMIT $5;"
	PopularAsc = "SELECT pins.id, name, description, image, pins.created_at, pins.tags, pins.views, pins.comments," +
		" users.id, users.login, users.avatar" +
		" FROM pins " +
		" JOIN users ON pins.user_id = users.id" +
		" WHERE make_tsvector(name) @@ to_tsquery($1)  AND pins.created_at BETWEEN $2 AND $3 AND visible = true" +
		" ORDER BY pins.views ASC OFFSET $4 LIMIT $5;"
	CommentsDesc = "SELECT pins.id, name, description, image, pins.created_at, pins.tags, pins.views, pins.comments," +
		" users.id, users.login, users.avatar" +
		" FROM pins " +
		" JOIN users ON pins.user_id = users.id" +
		" WHERE name LIKE $1  AND pins.created_at BETWEEN $2 AND $3 AND visible = true" +
		" ORDER BY pins.comments DESC OFFSET $4 LIMIT $5;"
	CommentsAsc = "SELECT pins.id, name, description, image, pins.created_at, pins.tags, pins.views, pins.comments," +
		" users.id, users.login, users.avatar" +
		" FROM pins " +
		" JOIN users ON pins.user_id = users.id" +
		" WHERE make_tsvector(name) @@ to_tsquery($1)  AND pins.created_at BETWEEN $2 AND $3 AND visible = true" +
		" ORDER BY pins.comments ASC OFFSET $4 LIMIT $5;"
	IdDesc = "SELECT pins.id, name, description, image, pins.created_at, pins.tags, pins.views, pins.comments," +
		" users.id, users.login, users.avatar" +
		" FROM pins " +
		" JOIN users ON pins.user_id = users.id" +
		" WHERE make_tsvector(name) @@ to_tsquery($1)  AND pins.created_at BETWEEN $2 AND $3 AND visible = true" +
		" ORDER BY pins.id DESC OFFSET $4 LIMIT $5;"
	IdAsc = "SELECT pins.id, name, description, image, pins.created_at, pins.tags, pins.views, pins.comments," +
		" users.id, users.login, users.avatar" +
		" FROM pins  " +
		" JOIN users ON pins.user_id = users.id" +
		" WHERE make_tsvector(name) @@ to_tsquery($1)  AND pins.created_at BETWEEN $2 AND $3 AND visible = true" +
		" ORDER BY pins.id ASC OFFSET $4 LIMIT $5;"

	PinByTag = "SELECT DISTINCT pins.id, name, description, image, board_id, pins.created_at, pins.tags, pins.views, users.id, users.login, users.avatar" +
	" FROM pins JOIN boards_pins ON pins.id = boards_pins.image_id" +
	" JOIN users ON pins.user_id = users.id" +
	" WHERE pins.tags[0] = ANY ($1 ) OR pins.tags[1] = ANY ($1 ) AND visible = true" +
		" ORDER BY pins.views OFFSET $2 LIMIT $3;"
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
	UserByLogin = "SELECT * FROM users " +
		" WHERE make_tsvector(login) @@ to_tsquery($1) " +
		" LIMIT $2 OFFSET $3"
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
	AddUserTags       = "UPDATE users SET tags= $1 WHERE id = $2 RETURNING 0;"
	IsFollowing       = "SELECT id FROM subscriptions WHERE user_id = $1 AND subscribed_at = $2;"
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
		" FROM pins JOIN boards_pins ON pins.id = boards_pins.image_id WHERE board_id = $1 AND visible = true ORDER BY created_at DESC LIMIT 1;"
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
		" WHERE subscriptions.user_id = $1  AND original = true AND visible = true ORDER BY created_at DESC LIMIT $2 OFFSET $3;"
	Main = "SELECT id, user_id, name, description, image, board_id, created_at " +
		" FROM pins JOIN boards_pins ON pins.id = boards_pins.image_id " +
		" WHERE original = true AND visible = true ORDER BY views DESC LIMIT $1  OFFSET $2;"
	Recommendation = "SELECT id, user_id, name, description, image, board_id, created_at " +
		" FROM pins JOIN boards_pins ON pins.id = boards_pins.image_id " +
		" WHERE original = true AND visible = true ORDER BY created_at DESC LIMIT $1  OFFSET $2;"
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
