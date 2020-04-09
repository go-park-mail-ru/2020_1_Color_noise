package database

import (
	"2020_1_Color_noise/internal/models"
	"errors"
	"fmt"
	"time"
)

func (db *PgxDB) GetNotifications(user models.DataBaseUser) ([]*models.Notification, error) {
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
			return res, nil
		}
		res = append(res, &tmp)
	}

	return res, nil
}

func (db *PgxDB) PutNotifications(com models.DataBaseComment) (uint, error) {

	text := "new comment on your pin, pin №" + fmt.Sprint(com.PinId)
	pin, _ := db.GetPinById(models.DataBasePin{Id: com.PinId})
	res := db.dbPool.QueryRow(PutNoti, pin.UserId, text, com.UserId, time.Now())
	var id uint
	err := res.Scan(&id)
	if err != nil {
		return 0, errors.New("no notifications")
	}
	return id, nil
}
