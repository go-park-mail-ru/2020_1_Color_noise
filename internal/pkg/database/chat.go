package database

import (
	"2020_1_Color_noise/internal/models"
	"sort"
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
		var tmp models.Message
		var author uint

		ok := row.Scan(&author, &tmp.Message, &tmp.CreatedAt)
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


		res = append(res, &tmp)
	}

	return res, nil
}

func (db *PgxDB) GetUsers(userId uint, start int, limit int) ([]*models.User, error) {

	var res []*models.User
	var ids []int

	sender, err := db.GetUserById(models.DataBaseUser{Id: uint(userId)})
	if err != nil {
		return nil, err
	}
	row, err := db.dbPool.Query(GetChats, sender.Id, limit, start)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var send models.User
		var receive models.User
		ok := row.Scan(&send.Id, &receive.Id)

		if ok != nil {
			return res, nil
		}
		//если мы написали
		if send.Id == sender.Id && sort.SearchInts(ids, int(send.Id)) == len(ids) {
			receiver, _ := db.GetUserById(models.GetBUser(receive))
			res = append(res, &receiver)
			ids = append(ids, int(send.Id))
		} else if sort.SearchInts(ids, int(send.Id)) == len(ids){
			//если нам написали
			receiver, _ := db.GetUserById(models.GetBUser(send))
			res = append(res, &receiver)
			ids = append(ids, int(receive.Id))
		}
		sort.Ints(ids)
	}

	return res, nil
}
