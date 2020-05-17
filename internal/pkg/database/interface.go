package database

import (
	"2020_1_Color_noise/internal/models"
	"2020_1_Color_noise/internal/pkg/config"
)

type DBInterface interface {
	Open(c config.DataBaseConfig) (err error)

	Close() error
	Ping() error

	CreateSession(s models.DataBaseSession) error
	DeleteSession(s models.DataBaseSession) error
	UpdateSession(s models.DataBaseSession) error
	GetSessionByCookie(s models.DataBaseSession) (models.Session, error)

	CreatePin(pin models.DataBasePin) (uint, error)
	Save(pinId, boardId uint) error
	UpdatePin(pin models.DataBasePin) error
	DeletePin(pin models.DataBasePin) error
	GetPinById(pin models.DataBasePin) (models.Pin, error)
	GetPinsByUserId(pin models.DataBasePin) ([]*models.Pin, error)
	GetPinsByName(pin models.DataBasePin) ([]*models.Pin, error)
	GetPinsByBoardID(board models.DataBaseBoard) ([]*models.Pin, error)
	AddTags(pinID uint, tags []string) error

	CreateUser(user models.DataBaseUser) (uint, error)
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
	GetUserSubUsers(user models.DataBaseUser) ([]*models.User, error)
	GetUserSupUsers(user models.DataBaseUser) ([]*models.User, error)
	Follow(who, whom uint) error
	Unfollow(who, whom uint) error
	AddUserTags(userID uint, tags []string) error

	CreateComment(cm models.DataBaseComment) (uint, error)
	UpdateComment(cm models.DataBaseComment) error
	DeleteComment(cm models.DataBaseComment) error
	GetCommentById(cm models.DataBaseComment) (models.Comment, error)
	GetCommentsByPinId(cm models.DataBaseComment, start int, limit int) ([]*models.Comment, error)
	GetCommentsByText(cm models.DataBaseComment, start int, limit int) ([]*models.Comment, error)

	CreateBoard(board models.DataBaseBoard) (uint, error)
	UpdateBoard(board models.DataBaseBoard) error
	DeleteBoard(board models.DataBaseBoard) error
	GetBoardById(board models.DataBaseBoard) (models.Board, error)
	GetBoardsByUserId(board models.DataBaseBoard, start, offset int) ([]*models.Board, error)
	GetBoardsByName(board models.DataBaseBoard, start, offset int) ([]*models.Board, error)
	GetBoardLastPin(board models.DataBaseBoard) (models.Pin, error)

	GetSubFeed(user models.DataBaseUser, start, limit int) ([]*models.Pin, error)
	GetMainFeed(user models.DataBaseUser, start, limit int) ([]*models.Pin, error)
	GetRecFeed(user models.DataBaseUser, start, limit int) ([]*models.Pin, error)

	GetNotifications(user models.DataBaseUser, start, limit int) ([]*models.Notification, error)
	PutNotifications(cm models.DataBaseComment) (uint, error)

	AddMessage(send, rec int, mess string, sticker string) (*models.Message, error)
	GetMessages(userId, otherId uint, start int, limit int) ([]*models.Message, error)
	GetUsers(userId uint, start int, limit int) ([]*models.User, error)
}
