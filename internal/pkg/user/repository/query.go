package repository

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
	UserSubscribed    = "SELECT COUNT(subscribed_at) FROM subscriptions WHERE user_id = $1"
	UserSubscriptions = "SELECT COUNT(user_id) FROM subscriptions WHERE subscribed_at = $1"
	Follow            = "INSERT INTO subscriptions( user_id, subscribed_at) VALUES ($1, $2) RETURNING id;"
	Unfollow          = "DELETE FROM public.subscriptions WHERE user_id = $1 AND subscribed_at = $2 RETURNING 0;"
)
