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

func (db *PostgresSQL) Open(config config.DataBaseConfig) error {
	var (
		err        error
		database   *sql.DB
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

func (db *PostgresSQL) Ping() error {
	return db.DB.Ping()
}

func (db *PostgresSQL) Close() error {
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
	return db.QueryRow(InsertPin,
		pin.UserId, pin.Name, pin.Description, pin.Image, pin.BoardId, time.Now())
}

//TODO:обновление пина
func (db *PostgresSQL) UpdatePin(pin models.DataBasePin) (sql.Result, error) {
	return db.Exec(UpdatePin, pin.Name, pin.Description, pin.BoardId, pin.Id)
}

func (db *PostgresSQL) DeletePin(pin models.DataBasePin) (sql.Result, error) {
	return db.Exec(DeletePin, pin.Id)
}

func (db *PostgresSQL) GetPinById(pin models.DataBasePin) *sql.Row {
	return db.QueryRow(PinById, pin.Id)
}

func (db *PostgresSQL) GetPinsByUserId(pin models.DataBasePin) (*sql.Rows, error) {
	return db.Query(PinByUser, pin.UserId)
}

//TODO: заменить на полнотекстовый поиск
func (db *PostgresSQL) GetPinByName(pin models.DataBasePin) (*sql.Rows, error) {
	return db.Query(PinByName, pin.Name)
}

//ОПЕРАЦИИ С ПОЛЬЗОВАТЕЛЕМ
func (db *PostgresSQL) CreateUser(user models.DataBaseUser) *sql.Row {
	return db.QueryRow(InsertUser,
		user.Email, user.Login, user.EncryptedPassword, user.About,
		user.Avatar, user.Subscribers, user.Subscriptions, time.Now())
}

func (db *PostgresSQL) UpdateUser(user models.DataBaseUser) (sql.Result, error) {
	return db.Exec(UpdateUser, user.Email, user.Login, user.Id)
}

func (db *PostgresSQL) UpdateUserDescription(user models.DataBaseUser) (sql.Result, error) {
	return db.Exec(UpdateUserDesc, user.About, user.Id)
}

func (db *PostgresSQL) UpdateUserPassword(user models.DataBaseUser) (sql.Result, error) {
	return db.Exec("UPDATE users SET "+
		"encrypted_password = $1 "+
		"WHERE id = $2", user.EncryptedPassword, user.Id)
}

func (db *PostgresSQL) UpdateUserAvatar(user models.DataBaseUser) (sql.Result, error) {
	return db.Exec(UpdateUserAv, user.Avatar, user.Id)
}

//WARNING: удаление пользователя вызывает нарушение внешнего ключа
func (db *PostgresSQL) DeleteUser(user models.DataBaseUser) (sql.Result, error) {
	return db.Exec(DeleteUser, user.Id)
}

func (db *PostgresSQL) GetUserById(user models.DataBaseUser) *sql.Row {
	return db.QueryRow(UserById, user.Id)
}

func (db *PostgresSQL) GetUserByLogin(user models.DataBaseUser) *sql.Row {
	return db.QueryRow("SELECT * FROM users WHERE login = $1", user.Login)
}

//TODO: выборка по имени
func (db *PostgresSQL) GetUserByName(user models.DataBaseUser) *sql.Row {
	return db.QueryRow("SELECT * FROM users WHERE login = $1", user.Login)
}

func (db *PostgresSQL) GetUserSubscriptions(user models.DataBaseUser) (*sql.Rows, error) {
	return db.Query(UserSubscriptions, user.Id)
}

func (db *PostgresSQL) GetUserSubscribers(user models.DataBaseUser) (*sql.Rows, error) {
	return db.Query(UserSubscribed, user.Id)
}

//ОПЕРАЦИИ С ДОСКАМИ

func (db *PostgresSQL) CreateBoard(board models.DataBaseBoard) *sql.Row {
	return db.QueryRow(InsertBoard,
		board.UserId, board.Name, board.Description, time.Now())
}

func (db *PostgresSQL) UpdateBoard(board models.DataBaseBoard) (sql.Result, error) {
	return db.Exec(UpdateBoard,
		board.Name, board.Description, board.Id)
}

func (db *PostgresSQL) DeleteBoard(board models.DataBaseBoard) (sql.Result, error) {
	return db.Exec(DeleteBoard, board.Id)
}

func (db *PostgresSQL) GetBoardById(board models.DataBaseBoard) *sql.Row {
	return db.QueryRow(BoardById, board.Id)
}

func (db *PostgresSQL) GetBoardsByUserId(board models.DataBaseBoard) (*sql.Rows, error) {
	return db.Query(BoardsByUserId, board.UserId)
}

func (db *PostgresSQL) GetBoardsByName(board models.DataBaseBoard) (*sql.Rows, error) {
	return db.Query(BoardsByNameSearch, board.Name)
}

//РАБОТА С КОММЕНТАРИЯМИ
func (db *PostgresSQL) CreateComment(cm models.DataBaseComment) *sql.Row {
	return db.QueryRow(InsertComment,
		cm.UserId, cm.PinId, cm.Text, time.Now())
}

func (db *PostgresSQL) UpdateComment(cm models.DataBaseComment) (sql.Result, error) {
	return db.Exec(UpdateComment, cm.Text, time.Now(), cm.Id)
}

func (db *PostgresSQL) DeleteComment(cm models.DataBaseComment) (sql.Result, error) {
	return db.Exec(DeleteComment, cm.Id)
}

func (db *PostgresSQL) GetCommentsByPinId(cm models.DataBaseComment) (*sql.Rows, error) {
	return db.Query(CommentByPin, cm.PinId)
}
