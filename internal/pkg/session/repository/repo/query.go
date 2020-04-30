package repo

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
