package database

import (
	"2020_1_Color_noise/internal/models"
	"fmt"
	"github.com/jackc/pgx"
	"time"
)

const (
	IntegrityConstraintViolation = "23000"
	RestrictViolation            = "23001"
	NotNullViolation             = "23502"
	ForeignKeyViolation          = "23503"
	UniqueViolation              = "23505"
	CheckViolation               = "23514"
)

type PgxDB struct {
	dbPool *pgx.ConnPool
}

func NewPgxDB() *PgxDB {
	return &PgxDB{}
}

func (db *PgxDB) Open() (err error) {
	//TODO: брать из конфига
	port := 5432

	connConfig := pgx.ConnConfig{
		Host:     "127.0.0.1",
		Port:     uint16(port),
		Database: "pinterest",
		User:     "postgres",
		Password: "password",
	}

	poolConfig := pgx.ConnPoolConfig{
		ConnConfig:     connConfig,
		MaxConnections: 50,
		AcquireTimeout: 10 * time.Second,
		AfterConnect:   nil,
	}

	if db.dbPool != nil {
		return fmt.Errorf("pool was created already")
	}

	db.dbPool, err = pgx.NewConnPool(poolConfig)
	if err != nil {
		return fmt.Errorf("connection is not established")
	}

	return nil
}

func (db *PgxDB) Close() error {
	db.dbPool.Close()
	return nil
}

func (db *PgxDB) Ping() error {
	if db.dbPool != nil {
		return nil
	}
	return fmt.Errorf("pool is not created")
}

func (db *PgxDB) Exec(query string, args ...interface{}) error {
	_, err := db.dbPool.Exec(query, args)
	return err
}

func (db *PgxDB) Query(query string, args ...interface{}) error {
	res, err := db.dbPool.Query(query, args)
	defer res.Close()
	return err
}

func (db *PgxDB) CreatePin(pin models.DataBasePin) error {
	_, err := db.dbPool.Exec(InsertPin, pin.UserId, pin.Name, pin.Description, pin.Image, pin.BoardId, time.Now())
	if err != nil {
		if pqError, ok := err.(pgx.PgError); ok {
			switch pqError.Code {
			case UniqueViolation:
				return fmt.Errorf("user is not unique")
			case CheckViolation:
				return fmt.Errorf("unsucseccfull check")
			}
		}
		return err
	}
	return err
}

func (db *PgxDB) UpdatePin(pin models.DataBasePin) error {
	_, err := db.dbPool.Exec(UpdatePin, pin.Name, pin.Description, pin.BoardId, pin.Id)
	if err != nil {
		if pqError, ok := err.(pgx.PgError); ok {
			switch pqError.Code {
			case ForeignKeyViolation:
				return fmt.Errorf("board not found")
			}
			return err
		}
	}
	return err
}

func (db *PgxDB) DeletePin(pin models.DataBasePin) error {
	_, err := db.dbPool.Exec(DeletePin, pin.Id)
	if err != nil {
		return err
	}
	return err
}

func (db *PgxDB) GetPinById(pin models.DataBasePin) (models.Pin, error) {
	var res models.Pin

	row := db.dbPool.QueryRow(PinById, pin.Id)
	err := row.Scan(&res.Id, &res.Name, &res.Description, &res.Image, &res.BoardId, &res.CreatedAt)
	if err != nil {
		return models.Pin{}, err
	}

	return res, nil
}

func (db *PgxDB) GetPinsByUserId(pin models.DataBasePin) (models.Pin, error) {
	var res models.Pin

	row := db.dbPool.QueryRow(PinByUser, pin.UserId)
	err := row.Scan(&res.Id, &res.Name, &res.Description, &res.Image, &res.BoardId, &res.CreatedAt)
	if err != nil {
		return models.Pin{}, err
	}

	return res, nil
}

func (db *PgxDB) GetPinByName(pin models.DataBasePin) (models.Pin, error) {
	var res models.Pin

	row := db.dbPool.QueryRow(PinByName, pin.Name)
	err := row.Scan(&res.Id, &res.Name, &res.Description, &res.Image, &res.BoardId, &res.CreatedAt)
	if err != nil {
		return models.Pin{}, err
	}
	return res, nil
}

func (db *PgxDB) CreateUser(user models.DataBaseUser) (int, error) {
	res := db.dbPool.QueryRow(InsertUser, user.Email, user.Login, user.EncryptedPassword, user.About,
		user.Avatar, user.Subscribers, user.Subscriptions, time.Now())
	var id int
	err := res.Scan(&id)

	if pqError, ok := err.(pgx.PgError); ok {
		switch pqError.Code {
		case UniqueViolation:
			return 0, fmt.Errorf("user is not unique")
		case CheckViolation:
			return 0, fmt.Errorf("unsucseccfull check")
		}
	}
	return id, err
}

func (db *PgxDB) UpdateUser(user models.DataBaseUser) error {
	err := db.Exec(UpdateUser, user.Email, user.Login, user.Id)
	if err != nil {
		if pqError, ok := err.(pgx.PgError); ok {
			switch pqError.Code {
			case CheckViolation:
				return fmt.Errorf("unsucsessfull user update")
			}
			return err
		}
	}
	return err
}

func (db *PgxDB) UpdateUserDescription(user models.DataBaseUser) error {
	err := db.Exec(UpdateUserDesc, user.About, user.Id)
	if err != nil {
		if pqError, ok := err.(pgx.PgError); ok {
			switch pqError.Code {
			case CheckViolation:
				return fmt.Errorf("unsucsessfull description update")
			}
			return err
		}
	}
	return err
}

func (db *PgxDB) UpdateUserAvatar(user models.DataBaseUser) error {
	return db.Exec(UpdateUserAv, user.Avatar, user.Id)
}

func (db *PgxDB) UpdateUserPassword(user models.DataBaseUser) error {
	return db.Exec(UpdateUserPs, user.EncryptedPassword, user.Id)
}

func (db *PgxDB) DeleteUser(user models.DataBaseUser) error {
	return db.Exec(DeleteUser, user.Id)
}

func (db *PgxDB) GetUserById(user models.DataBaseUser) (models.User, error) {
	var res models.User
	row := db.dbPool.QueryRow(UserById, user.Id)
	err := row.Scan(&res.Id, &res.Email, &res.Login, &res.EncryptedPassword,
		&res.About, &res.Avatar, &res.Subscribers, &res.Subscriptions, &res.CreatedAt)

	if err != nil {
		return models.User{}, err
	}
	return res, nil
}

func (db *PgxDB) GetUserByLogin(user models.DataBaseUser, limit, offset int) ([]*models.User, error) {

	var users []*models.User
	row, err := db.dbPool.Query(UserByLogin, user.Login, limit, offset)

	if err != nil {
		return nil, err
	}

	for row.Next() {
		var res models.User
		ok := row.Scan(&res.Id, &res.Email, &res.Login, &res.EncryptedPassword,
			&res.About, &res.Avatar, &res.Subscribers, &res.Subscriptions, &res.CreatedAt)
		if ok != nil {
			return nil, ok
		}
		users = append(users, &res)
	}

	if err != nil {
		return nil, err
	}
	return users, nil
}

func (db *PgxDB) GetUserByName(user models.DataBaseUser) (models.User, error) {
	var res models.User

	row := db.dbPool.QueryRow(UserByLoginSearch, user.Id)
	err := row.Scan(&res.Id, &res.Email, &res.Login, &res.EncryptedPassword,
		&res.About, &res.Avatar, &res.Subscribers, &res.Subscriptions, &res.CreatedAt)

	if err != nil {
		return models.User{}, err
	}
	return res, nil
}

func (db *PgxDB) GetUserByEmail(user models.DataBaseUser) (models.User, error) {
	var res models.User

	row := db.dbPool.QueryRow(UserByEmail, user.Id)
	err := row.Scan(&res.Id, &res.Email, &res.Login, &res.EncryptedPassword,
		&res.About, &res.Avatar, &res.Subscribers, &res.Subscriptions, &res.CreatedAt)

	if err != nil {
		return models.User{}, err
	}
	return res, nil
}

func (db *PgxDB) Follow(who, whom uint) error {

	_, err := db.dbPool.Exec(Follow, who, whom)
	if err != nil {
		return err
	}
	return err
}

func (db *PgxDB) Unfollow(who, whom uint) error {

	_, err := db.dbPool.Exec(Unfollow, who, whom)
	if err != nil {
		return err
	}
	return err
}

func (db *PgxDB) GetUserSubscriptions(user models.DataBaseUser) (models.User, error) {
	var res models.User

	row := db.dbPool.QueryRow(UserSubscriptions, user.Id)
	err := row.Scan(&res.Subscriptions)

	if err != nil {
		return models.User{}, err
	}
	return res, nil
}

func (db *PgxDB) GetUserSubscribers(user models.DataBaseUser) (models.User, error) {
	var res models.User

	row := db.dbPool.QueryRow(UserSubscriptions, user.Id)
	err := row.Scan(&res.Subscriptions)

	if err != nil {
		return models.User{}, err
	}

	return res, nil
}

func (db *PgxDB) CreateComment(cm models.DataBaseComment) error {
	return db.Exec(InsertComment,
		cm.UserId, cm.PinId, cm.Text, time.Now())
}

func (db *PgxDB) UpdateComment(cm models.DataBaseComment) error {
	return db.Exec(UpdateComment, cm.Text, time.Now(), cm.Id)
}

func (db *PgxDB) DeleteComment(cm models.DataBaseComment) error {
	return db.Exec(DeleteComment, cm.Id)
}

func (db *PgxDB) GetCommentsByPinId(cm models.DataBaseComment) ([]models.Comment, error) {
	var res []models.Comment
	r, err := db.dbPool.Query(CommentByPin, cm.PinId)
	if err != nil {
		return nil, err
	}

	for r.Next() {
		var tmp models.Comment
		ok := r.Scan(&tmp.Id, &tmp.UserId, &tmp.PinId, &tmp.Text, &tmp.CreatedAt)
		if ok != nil {
			return nil, ok
		}
		res = append(res, tmp)
	}
	return res, nil
}

func (db *PgxDB) CreateBoard(board models.DataBaseBoard) error {
	return db.Exec(InsertBoard,
		board.UserId, board.Name, board.Description, time.Now())
}

func (db *PgxDB) UpdateBoard(board models.DataBaseBoard) error {
	return db.Exec(UpdateBoard,
		board.Name, board.Description, board.Id)
}

func (db *PgxDB) DeleteBoard(board models.DataBaseBoard) error {
	return db.Exec(DeleteBoard, board.Id)
}

func (db *PgxDB) GetBoardById(board models.DataBaseBoard) (models.Board, error) {
	var res models.Board
	row := db.dbPool.QueryRow(BoardById, board.Id)
	err := row.Scan(&res.Id, &res.UserId, &res.Name, &res.Description, &res.CreatedAt)

	if err != nil {
		return models.Board{}, err
	}
	return res, nil
}

func (db *PgxDB) GetBoardsByUserId(board models.DataBaseBoard) ([]models.Board, error) {
	var res []models.Board
	row, err := db.dbPool.Query(BoardsByUserId, board.UserId)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		var tmp models.Board
		ok := row.Scan(&tmp.Id, &tmp.UserId, &tmp.Name, &tmp.Description, &tmp.CreatedAt)
		if ok != nil {
			return nil, ok
		}
		res = append(res, tmp)
	}
	return res, nil
}

func (db *PgxDB) GetBoardsByName(board models.DataBaseBoard) ([]models.Board, error) {
	var res []models.Board
	row, err := db.dbPool.Query(BoardsByNameSearch, board.UserId)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		var tmp models.Board
		ok := row.Scan(&tmp.Id, &tmp.UserId, &tmp.Name, &tmp.Description, &tmp.CreatedAt)
		if ok != nil {
			return nil, ok
		}
		res = append(res, tmp)
	}
	return res, nil
}
