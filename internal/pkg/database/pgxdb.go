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

func (db *PgxDB) CreatePin(pin models.DataBasePin) (uint, error) {
	res := db.dbPool.QueryRow(InsertPin, pin.UserId, pin.Name, pin.Description, pin.Image, pin.BoardId, time.Now())
	var id uint
	err := res.Scan(&id)
	if err != nil {
		if pqError, ok := err.(pgx.PgError); ok {
			switch pqError.Code {
			case UniqueViolation:
				return 0, fmt.Errorf("user is not unique")
			case CheckViolation:
				return 0, fmt.Errorf("unsucseccfull check")
			}
		}
		return 0, err
	}
	return id, err
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
	var res models.DataBasePin

	row := db.dbPool.QueryRow(PinById, pin.Id)
	err := row.Scan(&res.Id, &res.UserId, &res.Name, &res.Description, &res.Image, &res.BoardId, &res.CreatedAt)
	if err != nil {
		return models.Pin{}, err
	}

	return models.GetPin(res), nil
}

func (db *PgxDB) GetPinsByUserId(pin models.DataBasePin) ([]*models.Pin, error) {
	var res []*models.Pin

	row, err := db.dbPool.Query(PinByUser, pin.UserId)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var tmp models.DataBasePin
		ok := row.Scan(&tmp.Id, &tmp.UserId, &tmp.Name, &tmp.Description,
			&tmp.Image, &tmp.BoardId, &tmp.CreatedAt)
		if ok != nil {
			return nil, ok
		}
		p := models.GetPin(tmp)
		res = append(res, &p)
	}
	return res, nil
}

func (db *PgxDB) GetPinsByName(pin models.DataBasePin) ([]*models.Pin, error) {
	var res []*models.Pin

	row, err := db.dbPool.Query(PinByName, pin.Name)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var tmp models.DataBasePin
		ok := row.Scan(&tmp.Id, &tmp.UserId, &tmp.Name, &tmp.Description,
			&tmp.Image, &tmp.BoardId, &tmp.CreatedAt)
		if ok != nil {
			return nil, ok
		}
		p := models.GetPin(tmp)
		res = append(res, &p)
	}
	return res, nil
}

func (db *PgxDB) CreateUser(user models.DataBaseUser) (uint, error) {
	res := db.dbPool.QueryRow(InsertUser, user.Email, user.Login, user.EncryptedPassword, user.About,
		user.Avatar, user.Subscribers, user.Subscriptions, time.Now())
	var id uint
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
	id := 0
	err := db.dbPool.QueryRow(UpdateUser, user.Email, user.Login, user.Id).Scan(&id)
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
	id := 0
	err := db.dbPool.QueryRow(UpdateUserDesc, user.About, user.Id).Scan(&id)
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
	id := 0
	err := db.dbPool.QueryRow(UpdateUserAv, user.Avatar, user.Id).Scan(&id)
	if err != nil {
		if pqError, ok := err.(pgx.PgError); ok {
			switch pqError.Code {
			case CheckViolation:
				return fmt.Errorf("unsucsessfull avatar update")
			}
			return err
		}
	}
	return err
}

func (db *PgxDB) UpdateUserPassword(user models.DataBaseUser) error {
	id := 0
	err := db.dbPool.QueryRow(UpdateUserPs, user.EncryptedPassword, user.Id).Scan(&id)
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

func (db *PgxDB) DeleteUser(user models.DataBaseUser) error {
	return db.Exec(DeleteUser, user.Id)
}

func (db *PgxDB) GetUserById(user models.DataBaseUser) (models.User, error) {
	var res models.DataBaseUser
	row := db.dbPool.QueryRow(UserById, user.Id)
	err := row.Scan(&res.Id, &res.Email, &res.Login, &res.EncryptedPassword,
		&res.About, &res.Avatar, &res.Subscribers, &res.Subscriptions, &res.CreatedAt)

	if err != nil {
		return models.User{}, err
	}
	return models.GetUser(res), nil
}

func (db *PgxDB) GetUserByLogin(user models.DataBaseUser, start int, offset int) ([]*models.User, error) {

	var users []*models.User
	row, err := db.dbPool.Query(UserByLogin, user.Login, start, offset)

	if err != nil {
		return nil, err
	}

	for row.Next() {
		var res models.DataBaseUser
		ok := row.Scan(&res.Id, &res.Email, &res.Login, &res.EncryptedPassword,
			&res.About, &res.Avatar, &res.Subscribers, &res.Subscriptions, &res.CreatedAt)
		if ok != nil {
			return nil, ok
		}
		r := models.GetUser(res)
		users = append(users, &r)
	}

	if err != nil {
		return nil, err
	}
	return users, nil
}

func (db *PgxDB) GetUserByName(user models.DataBaseUser) (models.User, error) {
	var res models.DataBaseUser

	row := db.dbPool.QueryRow(UserByLoginSearch, user.Login)
	err := row.Scan(&res.Id, &res.Email, &res.Login, &res.EncryptedPassword,
		&res.About, &res.Avatar, &res.Subscribers, &res.Subscriptions, &res.CreatedAt)
	if err != nil {
		return models.User{}, err
	}
	return models.GetUser(res), nil
}

func (db *PgxDB) GetUserByEmail(user models.DataBaseUser) (models.User, error) {
	var res models.DataBaseUser

	row := db.dbPool.QueryRow(UserByEmail, user.Email)
	err := row.Scan(&res.Id, &res.Email, &res.Login, &res.EncryptedPassword,
		&res.About, &res.Avatar, &res.Subscribers, &res.Subscriptions, &res.CreatedAt)

	if err != nil {
		return models.User{}, err
	}
	return models.GetUser(res), nil
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

func (db *PgxDB) GetUserSubUsers(user models.DataBaseUser) ([]*models.User, error) {
	var users []*models.User
	row, err := db.dbPool.Query(UserSubscribedUsers, user.Id)

	if err != nil {
		return nil, err
	}

	for row.Next() {
		var res models.DataBaseUser
		ok := row.Scan(&res.Id, &res.Email, &res.Login, &res.EncryptedPassword,
			&res.About, &res.Avatar, &res.Subscribers, &res.Subscriptions, &res.CreatedAt)
		if ok != nil {
			return nil, ok
		}
		r := models.GetUser(res)
		users = append(users, &r)
	}

	if err != nil {
		return nil, err
	}
	return users, nil
}

func (db *PgxDB) GetUserSupUsers(user models.DataBaseUser) ([]*models.User, error) {
	var users []*models.User
	row, err := db.dbPool.Query(UserSubscriptionsUsers, user.Id)

	if err != nil {
		return nil, err
	}

	for row.Next() {
		var res models.DataBaseUser
		ok := row.Scan(&res.Id, &res.Email, &res.Login, &res.EncryptedPassword,
			&res.About, &res.Avatar, &res.Subscribers, &res.Subscriptions, &res.CreatedAt)
		if ok != nil {
			return nil, ok
		}
		r := models.GetUser(res)
		users = append(users, &r)
	}

	if err != nil {
		return nil, err
	}
	return users, nil
}

//комментарии не обернуты, так как модели полностью совместимы
func (db *PgxDB) CreateComment(cm models.DataBaseComment) (uint, error) {
	var id uint
	row := db.dbPool.QueryRow(InsertComment,
		cm.UserId, cm.PinId, cm.Text, time.Now())
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func (db *PgxDB) UpdateComment(cm models.DataBaseComment) error {
	return db.Exec(UpdateComment, cm.Text, time.Now(), cm.Id)
}

func (db *PgxDB) DeleteComment(cm models.DataBaseComment) error {
	return db.Exec(DeleteComment, cm.Id)
}

func (db *PgxDB) GetCommentById(cm models.DataBaseComment) (models.Comment, error) {
	var r models.Comment
	row := db.dbPool.QueryRow(CommentById, cm.Id)

	ok := row.Scan(&r.Id, &r.UserId, &r.PinId, &r.Text, &r.CreatedAt)
	if ok != nil {
		return r, ok
	}
	return r, nil
}

func (db *PgxDB) GetCommentsByPinId(cm models.DataBaseComment, start, limit int) ([]*models.Comment, error) {
	var res []*models.Comment
	r, err := db.dbPool.Query(CommentByPin, cm.PinId, limit, start)
	if err != nil {
		return nil, err
	}

	for r.Next() {
		var tmp models.Comment
		ok := r.Scan(&tmp.Id, &tmp.UserId, &tmp.PinId, &tmp.Text, &tmp.CreatedAt)
		if ok != nil {
			return nil, ok
		}
		res = append(res, &tmp)
	}
	return res, nil
}

func (db *PgxDB) GetCommentsByText(cm models.DataBaseComment, start, limit int) ([]*models.Comment, error) {
	var res []*models.Comment
	r, err := db.dbPool.Query(CommentByText, cm.Text, limit, start)
	if err != nil {
		return nil, err
	}

	for r.Next() {
		var tmp models.Comment
		ok := r.Scan(&tmp.Id, &tmp.UserId, &tmp.PinId, &tmp.Text, &tmp.CreatedAt)
		if ok != nil {
			return nil, ok
		}
		res = append(res, &tmp)
	}
	return res, nil
}

func (db *PgxDB) CreateBoard(board models.DataBaseBoard) (uint, error) {
	var id uint
	row := db.dbPool.QueryRow(InsertBoard,
		board.UserId, board.Name, board.Description, time.Now())
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func (db *PgxDB) UpdateBoard(board models.DataBaseBoard) error {
	return db.Exec(UpdateBoard,
		board.Name, board.Description, board.Id)
}

func (db *PgxDB) DeleteBoard(board models.DataBaseBoard) error {
	return db.Exec(DeleteBoard, board.Id)
}

func (db *PgxDB) GetBoardById(board models.DataBaseBoard) (models.Board, error) {
	var res models.DataBaseBoard
	row := db.dbPool.QueryRow(BoardById, board.Id)
	err := row.Scan(&res.Id, &res.UserId, &res.Name, &res.Description, &res.CreatedAt)

	if err != nil {
		return models.Board{}, err
	}
	return models.GetBoard(res), nil
}

func (db *PgxDB) GetBoardsByUserId(board models.DataBaseBoard, start, limit int) ([]*models.Board, error) {
	var res []*models.Board
	row, err := db.dbPool.Query(BoardsByUserId, board.UserId, limit, start)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		var tmp models.DataBaseBoard
		ok := row.Scan(&tmp.Id, &tmp.UserId, &tmp.Name, &tmp.Description, &tmp.CreatedAt)
		if ok != nil {
			return nil, ok
		}

		r := models.GetBoard(tmp)
		res = append(res, &r)
	}
	return res, nil
}

func (db *PgxDB) GetBoardsByName(board models.DataBaseBoard, start, limit int) ([]*models.Board, error) {
	var res []*models.Board
	row, err := db.dbPool.Query(BoardsByNameSearch, board.Name, limit, start)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		var tmp models.DataBaseBoard
		ok := row.Scan(&tmp.Id, &tmp.UserId, &tmp.Name, &tmp.Description, &tmp.CreatedAt)
		if ok != nil {
			return nil, ok
		}
		r := models.GetBoard(tmp)
		res = append(res, &r)
	}
	return res, nil
}

func (db *PgxDB) GetBoardLastPin(board models.DataBaseBoard) (models.Pin, error) {
	var res models.DataBasePin

	row := db.dbPool.QueryRow(LastPin, board.Id)
	err := row.Scan(&res.Id, &res.UserId, &res.Name, &res.Description, &res.Image, &res.BoardId, &res.CreatedAt)
	if err != nil {
		return models.Pin{}, err
	}

	return models.GetPin(res), nil
}

func (db *PgxDB) CreateSession(s models.DataBaseSession) error {
	res := db.dbPool.QueryRow(InsertSession, s.Id, s.Cookie, s.Token, s.CreatedAt, s.DeletingAt)
	var i int
	return res.Scan(&i)
}

func (db *PgxDB) UpdateSession(s models.DataBaseSession) error {
	var tmp int
	return db.dbPool.QueryRow(UpdateSession, s.Token, s.Cookie).Scan(&tmp)
}

func (db *PgxDB) DeleteSession(s models.DataBaseSession) error {
	var tmp int
	return db.dbPool.QueryRow(DeleteSession, s.Cookie).Scan(&tmp)
}

func (db *PgxDB) GetSessionByCookie(s models.DataBaseSession) (models.Session, error) {
	session := models.DataBaseSession{}
	row := db.dbPool.QueryRow(SessionByCookie, s.Cookie)

	err := row.Scan(&session.Id, &session.Cookie, &session.Token,
		&session.CreatedAt, &session.DeletingAt)

	if err != nil {
		return models.Session{}, err
	}
	return models.GetSession(session), nil
}

func (db *PgxDB) GetSubFeed( user models.DataBaseUser, start, limit int) ([]*models.Pin, error){
	var res []*models.Pin

	row, err := db.dbPool.Query(Feed, user.Id, limit, start)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var tmp models.DataBasePin
		ok := row.Scan(&tmp.Id, &tmp.UserId, &tmp.Name, &tmp.Description,
			&tmp.Image, &tmp.BoardId, &tmp.CreatedAt)
		if ok != nil {
			return nil, ok
		}
		p := models.GetPin(tmp)
		res = append(res, &p)
	}
	return res, nil
}

func (db *PgxDB) GetMainFeed( user models.DataBaseUser, start, limit int) ([]*models.Pin, error){
	var res []*models.Pin

	row, err := db.dbPool.Query(Main, limit)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var tmp models.DataBasePin
		ok := row.Scan(&tmp.Id, &tmp.UserId, &tmp.Name, &tmp.Description,
			&tmp.Image, &tmp.BoardId, &tmp.CreatedAt)
		if ok != nil {
			return nil, ok
		}
		p := models.GetPin(tmp)
		res = append(res, &p)
	}
	return res, nil
}

func (db *PgxDB) GetRecFeed( user models.DataBaseUser, start, limit int) ([]*models.Pin, error){
	var res []*models.Pin

	row, err := db.dbPool.Query(Recommendation, limit, start)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var tmp models.DataBasePin
		ok := row.Scan(&tmp.Id, &tmp.UserId, &tmp.Name, &tmp.Description,
			&tmp.Image, &tmp.BoardId, &tmp.CreatedAt)
		if ok != nil {
			return nil, ok
		}
		p := models.GetPin(tmp)
		res = append(res, &p)
	}
	return res, nil
}

func (db *PgxDB) GetNotifications( user models.DataBaseUser) ([]*models.Notification, error){
	var res []*models.Notification

	row, err := db.dbPool.Query(GetNoti, user.Id)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var tmp models.Notification
		var fromUser = models.User{}
		ok := row.Scan(&tmp.Message, &fromUser.Id)
		tmp.User, _ = db.GetUserById(models.GetBUser(fromUser))
		if ok != nil {
			return nil, ok
		}
		res = append(res, &tmp)
	}

	return res, nil
}

func (db *PgxDB) PutNotifications(com models.DataBaseComment) (uint, error){

	text := "new comment on your pin, pin №" + fmt.Sprint(com.PinId)
	pin, _ := db.GetPinById(models.DataBasePin{Id:com.PinId})
	res := db.dbPool.QueryRow(PutNoti, pin.UserId, text, com.UserId,  time.Now())
	var id uint
	err := res.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("No notifications")
	}
	return id, nil
}