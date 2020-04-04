package repository

import (
	"2020_1_Color_noise/internal/models"
	"2020_1_Color_noise/internal/pkg/database"
	. "2020_1_Color_noise/internal/pkg/error"
	"database/sql"
	"log"
	"sync"
)

type Repository struct {
	bd            database.DBInterface
	mu            *sync.Mutex
	subscriptions map[uint][]uint
	subscribers   map[uint][]uint
	muSub         *sync.Mutex
}

func NewRepo(bd database.DBInterface) *Repository {
	return &Repository{
		bd:            bd,
		mu:            &sync.Mutex{},
		subscriptions: make(map[uint][]uint),
		subscribers:   make(map[uint][]uint),
		muSub:         &sync.Mutex{},
	}
}

func (ur *Repository) Create(user *models.User) (uint, error) {
	_, err := ur.checkLogin(user.Login)
	if err == nil {
		return 0, LoginIsExist.New("Repo: Error in during creating")
	}

	_, err = ur.checkEmail(user.Email)
	if err == nil {
		return 0, EmailIsExist.New("Repo: Error in during creating")
	}

	ur.mu.Lock()
	defer ur.mu.Unlock()
	id, err := ur.bd.CreateUser(models.GetBUser(*user))
	if err != nil {
		return 0, UserNotFound.Newf("User can not be created")
	}

	return id, nil
}

func (ur *Repository) GetByID(id uint) (*models.User, error) {

	ur.mu.Lock()
	defer ur.mu.Unlock()

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
	ur.mu.Lock()
	defer ur.mu.Unlock()

	var us = models.DataBaseUser{
		Login: login,
	}
	log.Print("login: ", us.Login)
	user, err := ur.bd.GetUserByName(us)
	if err != nil {
		return nil, UserNotFound.New("User is not found")
	}
	return &user, err
}

func (ur *Repository) Search(login string, start int, limit int) ([]*models.User, error) {
	ur.mu.Lock()
	defer ur.mu.Unlock()

	var us = models.DataBaseUser{Login: login}
	users, err := ur.bd.GetUserByLogin(us, limit, start) //START == OFFSET
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *Repository) checkLogin(login string) (uint, error) {
	ur.mu.Lock()
	defer ur.mu.Unlock()

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
	ur.mu.Lock()
	defer ur.mu.Unlock()
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

	ur.mu.Lock()
	defer ur.mu.Unlock()

	var us = models.DataBaseUser{Id: id, Email: email, Login: login}
	err = ur.bd.UpdateUser(us)
	if err != nil {
		return UserNotFound.Newf("User to update not found, id: %d", id)
	}
	return nil
}

func (ur *Repository) UpdateDescription(id uint, description *string) error {
	ur.mu.Lock()
	defer ur.mu.Unlock()

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
	ur.mu.Lock()
	defer ur.mu.Unlock()

	var us = models.DataBaseUser{Id: id, EncryptedPassword: encryptredPassword}

	err := ur.bd.UpdateUserPassword(us)

	if err != nil {
		return UserNotFound.Newf("User to update not found, id: %d", id)
	}
	return nil
}

func (ur *Repository) UpdateAvatar(id uint, path string) error {
	ur.mu.Lock()
	defer ur.mu.Unlock()

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
	ur.mu.Lock()
	defer ur.mu.Unlock()

	var us = models.DataBaseUser{Id: id}

	err := ur.bd.DeleteUser(us)

	if err != nil {
		return UserNotFound.Newf("User to delete not found, id: %d", id)
	}
	return nil
}

//TODO: update полсе подписок
func (ur *Repository) Follow(id uint, subId uint) error {
	ur.muSub.Lock()
	defer ur.muSub.Unlock()

	_, err := ur.GetByID(id)
	if err != nil {
		return Wrapf(err, "Repo: incorrect user id, id", id)
	}

	_, err = ur.GetByID(subId)
	if err != nil {
		return Wrapf(err, "Repo: incorrect sub id, id", id)
	}

	err = ur.bd.Follow(id, subId)
	if err != nil {
		return Wrapf(err, "Repo: Error in during following, id", id)
	}

	//TODO: обновить
	return nil
}

func (ur *Repository) Unfollow(id uint, subId uint) error {
	ur.muSub.Lock()
	defer ur.muSub.Unlock()

	_, err := ur.GetByID(id)
	if err != nil {
		return Wrapf(err, "Repo: incorrect user id, id", id)
	}

	_, err = ur.GetByID(subId)
	if err != nil {
		return Wrapf(err, "Repo: incorrect sub id, id", id)
	}

	err = ur.bd.Unfollow(id, subId)
	if err != nil {
		return FollowingIsNotYetDone.Newf("Repo: Following is not yet done, id: %d", id)

	}
	//TODO: обновить
	return nil

}

func (ur *Repository) GetSubscribers(id uint, start int, limit int) ([]*models.User, error) {
	var us = models.DataBaseUser{Id: id}
	users, err := ur.bd.GetUserSubUsers(us)
	if err != nil {
		return nil, UserNotFound.Newf("User was not found, id", id)
	}
	return users, nil
}

func (ur *Repository) GetSubscriptions(id uint, start int, limit int) ([]*models.User, error) {
	var us = models.DataBaseUser{Id: id}
	users, err := ur.bd.GetUserSupUsers(us)
	if err != nil {
		return nil, UserNotFound.Newf("User was not found, id", id)
	}
	return users, nil
}
