package database

import (
	"2020_1_Color_noise/internal/models"
	"time"
)

func (db *PgxDB) AddMessage(sid, rid int, msg string) (*models.Message, error) {
	var id uint

	sender, err := db.GetUserById(models.DataBaseUser{Id: uint(sid)})
	if err != nil {
		return &models.Message{}, err
	}

	receiver, err := db.GetUserById(models.DataBaseUser{Id: uint(rid)})
	if err != nil {
		return &models.Message{}, err
	}

	ms := models.Message{
		SendUser:  &sender,
		RecUser:   &receiver,
		Message:   msg,
		CreatedAt: time.Now(),
	}

	res := db.dbPool.QueryRow(AddMsg, ms.SendUser.Id, ms.RecUser.Id, ms.Message, ms.CreatedAt)
	err = res.Scan(&id)
	if err != nil {
		return &models.Message{}, err
	}

	return &ms, nil
}

func (db *PgxDB) GetMessages(userId uint, start int, limit int) ([]*models.Message, error) {

	var res []*models.Message

	sender, err := db.GetUserById(models.DataBaseUser{Id: uint(userId)})
	if err != nil {
		return nil, err
	}

	row, err := db.dbPool.Query(GetMsg, sender.Id, limit, start)

	if err != nil {
		return nil, err
	}

	for row.Next() {
		var tmp models.Message
		var s int
		var r int

		ok := row.Scan(&s, &r, &tmp.Message, &tmp.CreatedAt)
		if ok != nil {
			return res, nil
		}
		sender, _ := db.GetUserById(models.DataBaseUser{Id: uint(s)})
		receiver, _ := db.GetUserById(models.DataBaseUser{Id: uint(r)})

		tmp.SendUser = &sender
		tmp.RecUser = &receiver
		res = append(res, &tmp)
	}

	return res, nil
}

func (db *PgxDB) GetUsers(userId uint, start int, limit int) ([]*models.User, error) {

	var res []*models.User

	sender, err := db.GetUserById(models.DataBaseUser{Id: uint(userId)})
	if err != nil {
		return nil, err
	}

	row, err := db.dbPool.Query(GetChats, sender.Id, limit, start)

	if err != nil {
		return nil, err
	}

	for row.Next() {
		var tmp models.User
		ok := row.Scan(&tmp.Id)
		if ok != nil {
			return res, nil
		}
		receiver, _ := db.GetUserById(models.GetBUser(tmp))
		res = append(res, &receiver)
	}

	return res, nil
}
