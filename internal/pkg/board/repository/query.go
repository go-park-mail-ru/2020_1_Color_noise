package repository

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