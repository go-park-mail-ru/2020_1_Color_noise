package database

import (
	"2020_1_Color_noise/internal/models"
	"2020_1_Color_noise/internal/pkg/config"
	"database/sql"
)

type DBInterface interface {
	//соединиться
	Open(config config.DataBaseConfig) error
	//проверить
	Ping() error
	//закрыть соединение
	Close() error
	//выполнить что-то
	Exec(query string, args ...interface{}) (sql.Result, error)
	//получить выборку значений
	Query(query string, args ...interface{}) (*sql.Rows, error)
	//получить одно значение
	QueryRow(query string, args ...interface{}) *sql.Row
	//методы для пинов
	CreatePin(pin models.DataBasePin) *sql.Row
	UpdatePin(pin models.DataBasePin) (sql.Result, error)
	DeletePin(pin models.DataBasePin) (sql.Result, error)

	GetPinById(pin models.DataBasePin) *sql.Row
	GetPinsByUserId(pin models.DataBasePin) (*sql.Rows, error)
	GetPinByName(pin models.DataBasePin) (*sql.Rows, error)

	//методы для пользователей
	CreateUser(user models.DataBaseUser)  *sql.Row
	UpdateUser(user models.DataBaseUser)  (sql.Result, error)
	UpdateUserDescription(user models.DataBaseUser)  (sql.Result, error)
	UpdateUserPassword(user models.DataBaseUser)  (sql.Result, error)
	UpdateUserAvatar(user models.DataBaseUser)  (sql.Result, error)
	DeleteUser (user models.DataBaseUser) (sql.Result, error)
	GetUserById(user models.DataBaseUser) *sql.Row
	GetUserByLogin(user models.DataBaseUser) *sql.Row
	GetUserByName(user models.DataBaseUser) *sql.Row
	GetUserSubscriptions(user models.DataBaseUser) (*sql.Rows, error)
	GetUserSubscribers(user models.DataBaseUser) (*sql.Rows, error)

	//методы для досок
	CreateBoard(board models.DataBaseBoard)  *sql.Row
	UpdateBoard(board models.DataBaseBoard)  (sql.Result, error)
	DeleteBoard(board models.DataBaseBoard) (sql.Result, error)
	GetBoardById(board models.DataBaseBoard) *sql.Row
	GetBoardsByUserId(board models.DataBaseBoard) (*sql.Rows, error)
	GetBoardsByName(board models.DataBaseBoard) (*sql.Rows, error)

	//методы для комментариев
	CreateComment(cm models.DataBaseComment)  *sql.Row
	UpdateComment(cm models.DataBaseComment) (sql.Result, error)
	DeleteComment(cm models.DataBaseComment) (sql.Result, error)
	GetCommentsByPinId(cm models.DataBaseComment) (*sql.Rows, error)
}

