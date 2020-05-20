package repository

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
