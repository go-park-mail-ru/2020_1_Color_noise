package database

import (
	"2020_1_Color_noise/internal/models"
)

type DBInterface interface {
	Open() (err error)

	Close() error
	Ping() error
	Exec(query string, args ...interface{}) error
	Query(query string, args ...interface{}) error

	CreatePin(pin models.DataBasePin) error
	UpdatePin(pin models.DataBasePin) error
	DeletePin(pin models.DataBasePin) error
	GetPinById(pin models.DataBasePin) (models.Pin, error)
	GetPinsByUserId(pin models.DataBasePin) (models.Pin, error)
	GetPinByName(pin models.DataBasePin) (models.Pin, error)

	CreateUser(user models.DataBaseUser) (int, error)
	UpdateUser(user models.DataBaseUser) error
	UpdateUserDescription(user models.DataBaseUser) error
	UpdateUserPassword(user models.DataBaseUser) error
	UpdateUserAvatar(user models.DataBaseUser) error
	DeleteUser(user models.DataBaseUser) error
	GetUserById(user models.DataBaseUser) (models.User, error)
	GetUserByLogin(user models.DataBaseUser, start int, limit int) ([]*models.User, error)
	GetUserByName(user models.DataBaseUser) (models.User, error)
	GetUserByEmail(user models.DataBaseUser) (models.User, error)
	GetUserSubscriptions(user models.DataBaseUser) (models.User, error)
	GetUserSubscribers(user models.DataBaseUser) (models.User, error)
	Follow(who, whom uint) error
	Unfollow(who, whom uint) error

	CreateComment(cm models.DataBaseComment) error
	UpdateComment(cm models.DataBaseComment) error
	DeleteComment(cm models.DataBaseComment) error
	GetCommentsByPinId(cm models.DataBaseComment) ([]models.Comment, error)

	CreateBoard(board models.DataBaseBoard) error
	UpdateBoard(board models.DataBaseBoard) error
	DeleteBoard(board models.DataBaseBoard) error
	GetBoardById(board models.DataBaseBoard) (models.Board, error)
	GetBoardsByUserId(board models.DataBaseBoard) ([]models.Board, error)
	GetBoardsByName(board models.DataBaseBoard) ([]models.Board, error)
}
