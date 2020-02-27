package usecase

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"os"
	"pinterest/pkg/models"
	repo "pinterest/pkg/user/repository"
	"regexp"
	"time"
)

type UserUsecase struct {
	userRepo  *repo.UserRepository
}

func NewUserUsecase(repo *repo.UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepo: repo,
	}
}

func (uu *UserUsecase) Add(user *models.User) (uint, error) {
	if !validateLogin(user.Login) {
		return 0, fmt.Errorf("Change Login")
	}
	if !validatePassword(user.Password) {
		return 0, fmt.Errorf("Change Password")
	}
	err := uu.CheckLogin(user)
	if err != nil {
		return 0, err
	}
	encryptedPassword, err := encryptPassword(user.Password)
	if err != nil {
		return 0, err
	}
	user.EncryptedPassword = encryptedPassword
	user.Password = ""
	user.Avatar = "/avatar.jpg"
	id, err := uu.userRepo.Add(user)
	return id, err
}

func (uu *UserUsecase) GetById(id uint) (*models.User, error) {
	users, err := uu.userRepo.GetByID(id)
	if err != nil {
		return &models.User{}, err
	}
	if len(users) != 1 {
		return &models.User{}, fmt.Errorf("User not found")
	}
	return users[0], nil
}

func (uu *UserUsecase) GetByLogin(login string) (*models.User, error) {
	user, err := uu.userRepo.GetByLogin(login)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, fmt.Errorf("User not found")
	}
	return user, nil
}

func (uu *UserUsecase) Update(user *models.User) error {
	users, err := uu.userRepo.GetByID(user.Id)
	if err != nil {
		return err
	}
	fmt.Println(users, "hello")

	if len(users) != 1 {
		return fmt.Errorf("User not found")
	}

	if user.Login != "" {
		err = uu.CheckLogin(user)
		if err != nil {
			return err
		}
		if !validateLogin(user.Login) {
			return fmt.Errorf("Change Login")
		}
		users[0].Login = user.Login
	}

	if user.Email != "" {
		err = uu.CheckEmail(user)
		if err != nil {
			return err
		}
		if !validateEmail(user.Email) {
			return fmt.Errorf("Change Email")
		}
		users[0].Email = user.Email
	}

	if user.Password != "" {
		if !validatePassword(user.Password) {
			return fmt.Errorf("Change Password")
		}
		encryptedPassword, err := encryptPassword(user.Password)
		if err != nil {
			return err
		}
		users[0].EncryptedPassword = encryptedPassword
	}

	if len(user.Image) > 0 {
		err = uu.SaveAvatar(user)
		user.DataAvatar = []byte{}
		if err != nil {
			return nil
		}
		users[0].Avatar = user.Avatar
	}

	if user.About != "" {
		users[0].About = user.About
	}
	_, err = uu.userRepo.Update(user.Id, users[0])
	return err
}

func (uu *UserUsecase) Delete(id uint) error {
	status, err := uu.userRepo.Delete(id)
	if err != nil {
		return err
	}
	if !status {
		return fmt.Errorf("User not found")
	}
	return nil
}

func (uu *UserUsecase) ComparePassword(user *models.User, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(password)) == nil
}

func (uu *UserUsecase) CheckLogin(user *models.User) error {
	oldUser, _ := uu.userRepo.GetByLogin(user.Login)
	if oldUser == nil {
		return nil
	}
	if oldUser.Id == user.Id {
		return nil
	}
	return fmt.Errorf("Change login")
}

func (uu *UserUsecase) CheckEmail(user *models.User) error {
	oldUser, _ := uu.userRepo.GetByEmail(user.Email)
	if oldUser == nil {
		return nil
	}
	if oldUser.Id == user.Id {
		return nil
	}
	return fmt.Errorf("Change email")
}

func encryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(hash), err
}

func (uu *UserUsecase) SaveAvatar(user *models.User) (error) {
	name := "/" + randStringRunes(30) + ".jpg"
	file, err := os.Create(name)
	if err != nil{
		return err
	}
	defer file.Close()

	_, err = file.Write(user.Image)
	if err == nil {
		user.Avatar = name
	}
	return err
}

func randStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func validateEmail(email string) bool {
	matched, _ := regexp.MatchString(`^(([^<>()\[\]\.,;:\s@\"]+(\.[^<>()\[\]\.,;:\s@\"]+)*)|(\".+\"))@(([^<>()[\]\.,;:\s@\"]+\.)+[^<>()[\]\.,;:\s@\"]{2,})$`, email)
	return matched
}

func validateLogin(login string) bool {
	matched, _ := regexp.MatchString(`^[\w.-_@$]+$`, login)
	return matched
}

func validatePassword(password string) bool {
	return len(password) > 5
}

/*
func (uu *UserUsecase) ChangePassword(id uint, newPassword string) error {
	users, err:= uu.userRepo.GetByID(id)
	if err != nil {
		return err
	}
	if len(users) != 1 {
		return fmt.Errorf("Not found")
	}
	newEncryptedPassword, err := encryptPassword(newPassword)
	if err != nil {
		return err
	}
	users[0].EncryptedPassword = newEncryptedPassword
	_, err = uu.userRepo.Update(id, users[0])
	return err
}

func (uu *UserUsecase) ChangeLogin(id uint, newLogin string) error {
	if err := uu.checkLogin(newLogin); err != nil {
		return err
	}
	users, err:= uu.userRepo.GetByID(id)
	if err != nil {
		return err
	}
	if len(users) != 1 {
		return fmt.Errorf("Not found")
	}
	users[0].Login = newLogin
	_, err = uu.userRepo.Update(id, users[0])
	return err
}

func (uu *UserUsecase) ChangeEmail(id uint, newEmail string) error {
	if err := uu.checkEmail(newEmail); err != nil {
		return err
	}
	users, err:= uu.userRepo.GetByID(id)
	if err != nil {
		return err
	}
	if len(users) != 1 {
		return fmt.Errorf("Not found")
	}
	users[0].Email = newEmail
	_, err = uu.userRepo.Update(id, users[0])
	return err
}

func (uu *UserUsecase) ChangeDescription(id uint, newDescriptoin string) error {
	users, err:= uu.userRepo.GetByID(id)
	if err != nil {
		return err
	}
	if len(users) != 1 {
		return fmt.Errorf("Not found")
	}
	users[0].About = newDescriptoin
	_, err = uu.userRepo.Update(id, users[0])
	return err
}

func (uu *UserUsecase) ChangeAvatar(id uint, newDescriptoin string) error {
	users, err:= uu.userRepo.GetByID(id)
	if err != nil {
		return err
	}
	if len(users) != 1 {
		return fmt.Errorf("Not found")
	}
	users[0].About = newDescriptoin
	_, err = uu.userRepo.Update(id, users[0])
	return err
}
*/
