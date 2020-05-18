package database

import (
	"2020_1_Color_noise/internal/models"
	"time"
)

func (db *PgxDB) AddMessage(sid, rid int, msg string, sticker string) (*models.Message, error) {
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
		Stickers: sticker,
		CreatedAt: time.Now(),
	}

	dms := models.GetDMessage(ms)

	res := db.dbPool.QueryRow(AddMsg, dms.SendUser.Id, dms.RecUser.Id, dms.Message, dms.Stickers, dms.CreatedAt)
	err = res.Scan(&id)
	if err != nil {
		return &models.Message{}, err
	}

	return &ms, nil
}

func (db *PgxDB) GetMessages(userId, otherId uint, start int, limit int) ([]*models.Message, error) {

	var res []*models.Message

	sender, err := db.GetUserById(models.DataBaseUser{Id: uint(userId)})
	if err != nil {
		return nil, err
	}

	receiver, err := db.GetUserById(models.DataBaseUser{Id: uint(otherId)})
	if err != nil {
		return nil, err
	}

	row, err := db.dbPool.Query(GetMsg, sender.Id, receiver.Id, limit, start)

	if err != nil {
		return nil, err
	}

	for row.Next() {
		var tmp models.DMessage
		var author uint

		ok := row.Scan(&author, &tmp.Message, &tmp.Stickers, &tmp.CreatedAt)
		if ok != nil {
			return res, nil
		}

		if author == sender.Id {
			tmp.SendUser = &sender
			tmp.RecUser = &receiver
		} else {
			tmp.SendUser = &receiver
			tmp.RecUser = &sender
		}

		ms := models.GetMessage(tmp)
		res = append(res, &ms)
	}

	return res, nil
}

func (db *PgxDB) GetUsers(userId uint, start int, limit int) ([]*models.User, error) {
	sender, err := db.GetUserById(models.DataBaseUser{Id: uint(userId)})
	if err != nil {
		return nil, err
	}
	row, err := db.dbPool.Query(GetChats, sender.Id, limit, start)
	if err != nil {
		return nil, err
	}

	var res []*models.User
	m := make(map[uint]*models.User)

	for row.Next() {
		var send models.User
		var receive models.User
		ok := row.Scan(&send.Id, &receive.Id)

		if ok != nil {
			return res, nil
		}

		if send.Id == userId {
			m[receive.Id] = &receive
		} else {
			m[send.Id] = &send
		}
	}
	for _, val := range m {
		res = append(res, val)
	}

	return res, nil
}
