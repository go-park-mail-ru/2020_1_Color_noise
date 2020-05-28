package repository

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
