package database

import (
	"2020_1_Color_noise/internal/models"
	"errors"
	"github.com/jackc/pgx"
	"time"
)

func (db *PgxDB) CreateUser(user models.DataBaseUser) (uint, error) {
	res := db.dbPool.QueryRow(InsertUser, user.Email, user.Login, user.EncryptedPassword, user.About,
		user.Avatar, user.Subscribers, user.Subscriptions, time.Now())
	var id uint
	err := res.Scan(&id)

	if err != nil {

		if pqError, ok := err.(pgx.PgError); ok {
			switch pqError.Code {
			case UniqueViolation:
				return 0, errors.New("user is not unique")
			}
		}
		return 0, errors.New("scan error")
	}
	return id, nil
}

func (db *PgxDB) UpdateUser(user models.DataBaseUser) error {
	id := 0
	row := db.dbPool.QueryRow(UpdateUser, user.Email, user.Login, user.Id)
	err := row.Scan(&id)
	if err != nil {
		return errors.New("user can not be updated")
	}
	return nil
}

func (db *PgxDB) UpdateUserDescription(user models.DataBaseUser) error {
	id := 0
	err := db.dbPool.QueryRow(UpdateUserDesc, user.About, user.Id).Scan(&id)
	if err != nil {
		return errors.New("user not found")
	}
	return err
}

func (db *PgxDB) UpdateUserAvatar(user models.DataBaseUser) error {
	id := 0
	err := db.dbPool.QueryRow(UpdateUserAv, user.Avatar, user.Id).Scan(&id)
	if err != nil {
		return errors.New("user not found")
	}
	return err
}

func (db *PgxDB) UpdateUserPassword(user models.DataBaseUser) error {
	id := 0
	err := db.dbPool.QueryRow(UpdateUserPs, user.EncryptedPassword, user.Id).Scan(&id)
	if err != nil {
		return errors.New("user not found")
	}
	return err
}

func (db *PgxDB) DeleteUser(user models.DataBaseUser) error {
	_, err := db.dbPool.Exec(DeleteUser, user.Id)
	return err
}

func (db *PgxDB) GetUserById(user models.DataBaseUser) (models.User, error) {
	var res models.DataBaseUser
	row := db.dbPool.QueryRow(UserById, user.Id)
	err := row.Scan(&res.Id, &res.Email, &res.Login, &res.EncryptedPassword,
		&res.About, &res.Avatar, &res.Subscribers, &res.Subscriptions, &res.CreatedAt)

	if err != nil {
		return models.User{}, errors.New("user not found")
	}
	return models.GetUser(res), nil
}

func (db *PgxDB) GetUserByLogin(user models.DataBaseUser, start int, limit int) ([]*models.User, error) {

	var users []*models.User
	row, err := db.dbPool.Query(UserByLogin, user.Login, limit, start)

	if err != nil {
		return nil, errors.New("db error")
	}

	for row.Next() {
		var res models.DataBaseUser
		ok := row.Scan(&res.Id, &res.Email, &res.Login, &res.EncryptedPassword,
			&res.About, &res.Avatar, &res.Subscribers, &res.Subscriptions, &res.CreatedAt)
		if ok != nil {
			return nil, errors.New("user not found")
		}
		r := models.GetUser(res)
		users = append(users, &r)
	}
	return users, nil
}

func (db *PgxDB) GetUserByName(user models.DataBaseUser) (models.User, error) {
	var res models.DataBaseUser

	row := db.dbPool.QueryRow(UserByLoginSearch, user.Login)
	err := row.Scan(&res.Id, &res.Email, &res.Login, &res.EncryptedPassword,
		&res.About, &res.Avatar, &res.Subscribers, &res.Subscriptions, &res.CreatedAt)
	if err != nil {
		return models.User{}, errors.New("user not found")
	}
	return models.GetUser(res), nil
}

func (db *PgxDB) GetUserByEmail(user models.DataBaseUser) (models.User, error) {
	var res models.DataBaseUser

	row := db.dbPool.QueryRow(UserByEmail, user.Email)
	err := row.Scan(&res.Id, &res.Email, &res.Login, &res.EncryptedPassword,
		&res.About, &res.Avatar, &res.Subscribers, &res.Subscriptions, &res.CreatedAt)

	if err != nil {
		return models.User{}, errors.New("user not found")
	}
	return models.GetUser(res), nil
}

func (db *PgxDB) Follow(who, whom uint) error {

	var id uint
	row := db.dbPool.QueryRow(Follow, who, whom)
	err := row.Scan(&id)
	if err != nil {
		return errors.New("follow error")
	}

	row = db.dbPool.QueryRow(UpdateFollowA, who)
	_ = row.Scan(&id)

	row = db.dbPool.QueryRow(UpdateFollowB, whom)
	_ = row.Scan(&id)

	return err
}

func (db *PgxDB) Unfollow(who, whom uint) error {
	var id uint
	row := db.dbPool.QueryRow(Unfollow, who, whom)
	err := row.Scan(&id)
	if err != nil {
		return errors.New("unfollow error")
	}

	row = db.dbPool.QueryRow(UpdateUnfollowA, who)
	_ = row.Scan(&id)

	row = db.dbPool.QueryRow(UpdateUnfollowB, whom)
	_ = row.Scan(&id)

	return err
}

func (db *PgxDB) GetUserSubscriptions(user models.DataBaseUser) (models.User, error) {
	var res models.User

	row := db.dbPool.QueryRow(UserSubscriptions, user.Id)
	err := row.Scan(&res.Subscriptions)

	if err != nil {
		return models.User{}, errors.New("user not found")
	}
	return res, nil
}

func (db *PgxDB) GetUserSubscribers(user models.DataBaseUser) (models.User, error) {
	var res models.User

	row := db.dbPool.QueryRow(UserSubscriptions, user.Id)
	err := row.Scan(&res.Subscriptions)

	if err != nil {
		return models.User{}, errors.New("user not found")
	}

	return res, nil
}

func (db *PgxDB) GetUserSubUsers(user models.DataBaseUser) ([]*models.User, error) {
	var users []*models.User
	row, err := db.dbPool.Query(UserSubscribedUsers, user.Id)

	if err != nil {
		return nil, errors.New("db error")
	}

	for row.Next() {
		var res models.DataBaseUser
		ok := row.Scan(&res.Id, &res.Email, &res.Login, &res.EncryptedPassword,
			&res.About, &res.Avatar, &res.Subscribers, &res.Subscriptions, &res.CreatedAt)
		if ok != nil {
			return users, nil
		}
		r := models.GetUser(res)
		users = append(users, &r)
	}
	return users, nil
}

func (db *PgxDB) GetUserSupUsers(user models.DataBaseUser) ([]*models.User, error) {
	var users []*models.User
	row, err := db.dbPool.Query(UserSubscriptionsUsers, user.Id)

	if err != nil {
		return nil, errors.New("db error")
	}

	for row.Next() {
		var res models.DataBaseUser
		ok := row.Scan(&res.Id, &res.Email, &res.Login, &res.EncryptedPassword,
			&res.About, &res.Avatar, &res.Subscribers, &res.Subscriptions, &res.CreatedAt)
		if ok != nil {
			return users, nil
		}
		r := models.GetUser(res)
		users = append(users, &r)
	}
	return users, nil
}
