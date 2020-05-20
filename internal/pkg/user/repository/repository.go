package repository

import (
	"2020_1_Color_noise/internal/models"
	"2020_1_Color_noise/internal/pkg/database"
	. "2020_1_Color_noise/internal/pkg/error"
	"database/sql"
	"fmt"
)

type Repository struct {
	bd database.DBInterface
}

func NewRepo(bd database.DBInterface) *Repository {
	return &Repository{
		bd: bd,
	}
}

func (ur *Repository) UpdatePreferences(userId uint, preferences []string) error {

	return ur.bd.AddTags(userId, preferences)
}

func (ur *Repository) Create(user *models.User) (*models.User, error) {
	fmt.Println(*user)

	_, err := ur.checkLogin(user.Login)
	if err == nil {
		return nil, LoginIsExist.New("Repo: Error in during creating")
	}

	_, err = ur.checkEmail(user.Email)
	if err == nil {
		return nil, EmailIsExist.New("Repo: Error in during creating")
	}

	id, err := ur.bd.CreateUser(models.GetBUser(*user))
	if err != nil {
		return nil, UserNotFound.Newf("User can not be created")
	}

	user.Id = id

	return user, nil
}

func (ur *Repository) GetByID(id uint) (*models.User, error) {

	var us = models.DataBaseUser{
		Id: id,
	}
	user, err := ur.bd.GetUserById(us)
	if err != nil {
		return nil, UserNotFound.Newf("User to get not found, id: %d", id)
	}
	return &user, err
}

func (ur *Repository) GetByLogin(login string) (*models.User, error) {
	var us = models.DataBaseUser{
		Login: login,
	}
	user, err := ur.bd.GetUserByName(us)
	if err != nil {
		return nil, UserNotFound.New("User is not found")
	}
	return &user, err
}

func (ur *Repository) Search(login string, start int, limit int) ([]*models.User, error) {

	var us = models.DataBaseUser{Login: login}
	users, err := ur.bd.GetUserByLogin(us, start, limit) //START == OFFSET
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *Repository) checkLogin(login string) (uint, error) {
	var us = models.DataBaseUser{
		Login: login,
	}
	user, err := ur.bd.GetUserByName(us)
	if err != nil {
		return 0, BadLogin.Newf("User to get not found, login: %s", login)
	}
	return user.Id, err
}

func (ur *Repository) checkEmail(email string) (uint, error) {
	var us = models.DataBaseUser{
		Email: email,
	}
	user, err := ur.bd.GetUserByEmail(us)

	if err != nil {
		return 0, BadEmail.Newf("User to get not found, email: %s", email)
	}
	return user.Id, err

}

func (ur *Repository) UpdateProfile(id uint, email string, login string) error {
	userId, err := ur.checkLogin(login)
	if err == nil && userId != id {
		return LoginIsExist.New("Repo: Error in during updating profile")
	}

	userId, err = ur.checkEmail(email)
	if err == nil && userId != id {
		return EmailIsExist.New("Repo: Error in during updating profile")
	}

	var us = models.DataBaseUser{Id: id, Email: email, Login: login}
	err = ur.bd.UpdateUser(us)
	if err != nil {
		return UserNotFound.Newf("User to update not found, id: %d", id)
	}
	return nil
}

func (ur *Repository) UpdateDescription(id uint, description *string) error {
	var us = models.DataBaseUser{Id: id, About: struct {
		String string
		Valid  bool
	}{
		String: *description,
		Valid:  true,
	}}

	err := ur.bd.UpdateUserDescription(us)

	if err != nil {
		return UserNotFound.Newf("User to update not found, id: %d", id)
	}
	return nil
}

func (ur *Repository) UpdatePassword(id uint, encryptredPassword string) error {
	var us = models.DataBaseUser{Id: id, EncryptedPassword: encryptredPassword}

	err := ur.bd.UpdateUserPassword(us)

	if err != nil {
		return UserNotFound.Newf("User to update not found, id: %d", id)
	}
	return nil
}

func (ur *Repository) UpdateAvatar(id uint, path string) error {
	var us = models.DataBaseUser{Id: id, Avatar: sql.NullString{
		String: path,
		Valid:  true,
	}}

	err := ur.bd.UpdateUserAvatar(us)

	if err != nil {
		return UserNotFound.Newf("User to update not found, id: %d", id)
	}
	return nil
}

func (ur *Repository) Delete(id uint) error {
	var us = models.DataBaseUser{Id: id}

	err := ur.bd.DeleteUser(us)

	if err != nil {
		return UserNotFound.Newf("User to delete not found, id: %d", id)
	}
	return nil
}

//TODO: update полсе подписок
func (ur *Repository) Follow(id uint, subId uint) error {
	_, err := ur.GetByID(id)
	if err != nil {
		return UserNotFound.Newf("User not found, id: %d", id)
	}

	_, err = ur.GetByID(subId)
	if err != nil {
		return UserNotFound.Newf("User not found, id: %d", id)
	}

	err = ur.bd.Follow(id, subId)
	if err != nil {
		return FollowingIsAlreadyDone.New("Following is already done")
	}

	//TODO: обновить

	return nil
}

func (ur *Repository) Unfollow(id uint, subId uint) error {
	_, err := ur.GetByID(id)
	if err != nil {
		return UserNotFound.Newf("User not found, id: %d", id)
	}

	_, err = ur.GetByID(subId)
	if err != nil {
		return UserNotFound.Newf("User not found, id: %d", id)
	}

	err = ur.bd.Unfollow(id, subId)
	if err != nil {
		return FollowingIsNotYetDone.New("Following is not yet done")

	}
	//TODO: обновить

	return nil

}

func (ur *Repository) GetSubscribers(id uint, start int, limit int) ([]*models.User, error) {
	var us = models.DataBaseUser{Id: id}
	users, err := ur.bd.GetUserSubUsers(us)
	if err != nil {
		return nil, UserNotFound.Newf("User was not found, id %d", id)
	}
	return users, nil
}

func (ur *Repository) GetSubscriptions(id uint, start int, limit int) ([]*models.User, error) {
	var us = models.DataBaseUser{Id: id}
	users, err := ur.bd.GetUserSupUsers(us)
	if err != nil {
		return nil, UserNotFound.Newf("User was not found, id %d", id)
	}
	return users, nil
}

func (ur *Repository) IsFollowed(id uint, subId uint) (bool, error) {

	return ur.bd.IsFollowing(id, subId)
}
