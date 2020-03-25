package database

import (
	"2020_1_Color_noise/internal/models"
	"2020_1_Color_noise/internal/pkg/config"
	"database/sql"
	_ "github.com/lib/pq"
	"time"
)

//THIS IS FOR POSTGRES ONLY

type PostgresSQL struct {
	DB *sql.DB
}

func (db *PostgresSQL) Open(config config.DataBaseConfig)  error{
	var (
		err error
		database *sql.DB
		connection string
	)

	connection = config.ConnString

	if database, err = sql.Open("postgres", connection); err != nil {
		return err
	}

	db.DB = database
	db.DB.SetMaxOpenConns(config.MaxConns)
	//проверяем подключение
	err = db.Ping()
	return err
}

func (db *PostgresSQL) Ping()  error{
	return db.DB.Ping()
}

func (db *PostgresSQL) Close()  error{
	return db.DB.Close()
}

func (db *PostgresSQL) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.DB.Exec(query, args...)
}

func (db *PostgresSQL) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return db.DB.Query(query, args...)
}

func (db *PostgresSQL) QueryRow(query string, args ...interface{}) *sql.Row {
	return db.DB.QueryRow(query, args...)
}

//TODO:создание пина
func (db *PostgresSQL) CreatePin(pin models.DataBasePin) *sql.Row {
	return db.QueryRow("INSERT INTO pins(user_id, name, description, image, board_id, created_at) " +
		"VALUES($1, $2, $3, $4, $5, $6) RETURNING id",
		pin.UserId, pin.Name, pin.Description, pin.Image, pin.BoardId, time.Now())
}

//TODO:обновление пина
func (db *PostgresSQL) UpdatePin(pin models.DataBasePin) (sql.Result, error) {
	return db.Exec("UPDATE pins SET " +
	"name = $1, description = $2, board_id = $3 " +
	"WHERE id = $4", pin.Name, pin.Description, pin.BoardId, pin.Id)
}

func (db *PostgresSQL) DeletePin(pin models.DataBasePin) (sql.Result, error) {
	return db.Exec("DELETE from pins WHERE id = $1", pin.Id)
}

func (db *PostgresSQL) GetPinById(pin models.DataBasePin) *sql.Row {
	return db.QueryRow("SELECT * FROM pins WHERE id = $1", pin.Id)
}

func (db *PostgresSQL) GetPinsByUserId(pin models.DataBasePin) (*sql.Rows, error) {
	return db.Query("SELECT * FROM pins WHERE user_id = $1", pin.UserId)
}

//TODO: заменить на полнотекстовый поиск
func (db *PostgresSQL) GetPinByName(pin models.DataBasePin) (*sql.Rows, error) {
	return db.Query("SELECT * FROM pins WHERE name = $1", pin.Name)
}

//ОПЕРАЦИИ С ПОЛЬЗОВАТЕЛЕМ
func (db *PostgresSQL) CreateUser(user models.DataBaseUser)  *sql.Row {
	return db.QueryRow("INSERT INTO users(email, login, encrypted_password, about, avatar, " +
		"subscriptions, subscribers, created_at) " +
		"VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id",
		user.Email, user.Login, user.EncryptedPassword, user.About,
		user.Avatar, user.Subscribers, user.Subscriptions, time.Now())
}

func (db *PostgresSQL) UpdateUser(user models.DataBaseUser)  (sql.Result, error) {
	return db.Exec("UPDATE users SET " +
		"email = $1, login = $2 " +
		"WHERE id = $3", user.Email, user.Login, user.Id)
}

func (db *PostgresSQL) UpdateUserDescription(user models.DataBaseUser)  (sql.Result, error) {
	return db.Exec("UPDATE users SET " +
		"about = $1 " +
		"WHERE id = $2", user.About, user.Id)
}

func (db *PostgresSQL) UpdateUserPassword(user models.DataBaseUser)  (sql.Result, error) {
	return db.Exec("UPDATE users SET " +
		"encrypted_password = $1 " +
		"WHERE id = $2", user.EncryptedPassword, user.Id)
}

func (db *PostgresSQL) UpdateUserAvatar(user models.DataBaseUser)  (sql.Result, error) {
	return db.Exec("UPDATE users SET " +
		"avatar = $1 " +
		"WHERE id = $2", user.Avatar, user.Id)
}

//WARNING: удаление пользователя вызывает нарушение внешнего ключа
func (db *PostgresSQL) DeleteUser (user models.DataBaseUser) (sql.Result, error) {
	return db.Exec("DELETE FROM users WHERE id = $1", user.Id)
}

func (db *PostgresSQL) GetUserById(user models.DataBaseUser) *sql.Row {
	return db.QueryRow("SELECT * FROM users WHERE id = $1", user.Id)
}

func (db *PostgresSQL) GetUserByLogin(user models.DataBaseUser) *sql.Row {
	return db.QueryRow("SELECT * FROM users WHERE login = $1", user.Login)
}

//TODO: выборка по имени
func (db *PostgresSQL) GetUserByName(user models.DataBaseUser) *sql.Row {
	return db.QueryRow("SELECT * FROM users WHERE login = $1", user.Login)
}

func (db *PostgresSQL) GetUserSubscriptions(user models.DataBaseUser) (*sql.Rows, error)  {
	return db.Query("SELECT subscribed_at FROM subscriptions WHERE user_id = $1", user.Id)
}

func (db *PostgresSQL) GetUserSubscribers(user models.DataBaseUser) (*sql.Rows, error)  {
	return db.Query("SELECT user_id FROM subscriptions WHERE subscribed_at = $1", user.Id)
}

//ОПЕРАЦИИ С ДОСКАМИ

func (db *PostgresSQL) CreateBoard(board models.DataBaseBoard)  *sql.Row {
	return db.QueryRow("INSERT INTO boards(user_id, name, description, created_at) " +
		"VALUES($1, $2, $3, $4) RETURNING id",
		board.UserId, board.Name, board.Description, time.Now())
}

func (db *PostgresSQL) UpdateBoard(board models.DataBaseBoard)  (sql.Result, error) {
	return db.Exec("UPDATE boards SET " +
		"name = $1, description = $2 " +
		"WHERE id = $3",
		board.Name, board.Description, board.Id)
}

func (db *PostgresSQL) DeleteBoard(board models.DataBaseBoard) (sql.Result, error) {
	return db.Exec("DELETE FROM boards WHERE id = $1", board.Id)
}

func (db *PostgresSQL) GetBoardById(board models.DataBaseBoard) *sql.Row {
	return db.QueryRow("SELECT * FROM boards WHERE id = $1", board.Id)
}

func (db *PostgresSQL) GetBoardsByUserId(board models.DataBaseBoard) (*sql.Rows, error) {
	return db.Query("SELECT * FROM boards WHERE user_id = $1", board.UserId)
}

func (db *PostgresSQL) GetBoardsByName(board models.DataBaseBoard) (*sql.Rows, error) {
	return db.Query("SELECT * FROM boards WHERE name = $1", board.Name)
}

//РАБОТА С КОММЕНТАРИЯМИ
func (db *PostgresSQL) CreateComment(cm models.DataBaseComment)  *sql.Row {
	return db.QueryRow("INSERT INTO commentaries(user_id, pin_id, comment, created_at) " +
		"VALUES($1, $2, $3, $4) RETURNING id",
		cm.UserId, cm.PinId, cm.Text, time.Now())
}

func (db *PostgresSQL) UpdateComment(cm models.DataBaseComment)  (sql.Result, error) {
	return db.Exec("UPDATE commentaries SET " +
		"comment = $1, created_at = $2 " +
		"WHERE id = $3", cm.Text, time.Now(), cm.Id)
}

func (db *PostgresSQL) DeleteComment(cm models.DataBaseComment) (sql.Result, error) {
	return db.Exec("DELETE FROM commentaries WHERE id = $1", cm.Id)
}

func (db *PostgresSQL) GetCommentsByPinId(cm models.DataBaseComment) (*sql.Rows, error) {
	return db.Query("SELECT * FROM commentaries WHERE pin_id = $1", cm.PinId)
}
