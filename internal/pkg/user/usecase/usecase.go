package usecase

import (
	"bytes"
	"golang.org/x/crypto/bcrypt"
	"pinterest/internal/models"
	. "pinterest/internal/pkg/error"
	"pinterest/internal/pkg/image"
	"pinterest/internal/pkg/user"
)

type UserUsecase struct {
	repo  user.IRepository
}

func NewUsecase(repo user.IRepository) *UserUsecase {
	return &UserUsecase{
		repo: repo,
	}
}

func (uu *UserUsecase) Create(input *models.SignUpInput) (uint, error) {
	if !models.ValidateEmail(input.Email) {
		return 0, BadEmail.New("Email incorrect")
	}

	if !models.ValidateLogin(input.Login) {
		return 0, BadLogin.New("Login incorrect")
	}

	if !models.ValidatePassword(input.Password) {
		return 0, BadPassword.New("Password should be longer than 6 characters")
	}

	if err := uu.emailIsExist(input.Email); err != nil {
		return 0, err
	}

	if err := uu.loginIsExist(input.Login); err != nil {
		return 0, err
	}

	encryptedPassword, err := encryptPassword(input.Password)
	if err != nil {
		return 0, err
	}

	user := &models.User{
		Email:             input.Email,
		Login:             input.Login,
		EncryptedPassword: encryptedPassword,
		Avatar:            "avatar.jpg",
	}

	return uu.repo.Add(user)
}

func (uu *UserUsecase) GetById(id uint) (*models.User, error) {
	return uu.repo.GetByID(id)
}

func (uu *UserUsecase) GetByLogin(login string) (*models.User, error) {
	return uu.repo.GetByLogin(login)
}

func (uu *UserUsecase) Update(id uint, input *models.UpdateInput) error {
	if !models.ValidateEmail(input.Email) {
		return BadEmail.New("Email incorrect")
	}

	if !models.ValidateLogin(input.Login) {
		return BadLogin.New("Login incorrect")
	}

	if err := uu.emailIsExist(input.Email); err != nil {
		return err
	}

	if err := uu.loginIsExist(input.Login); err != nil {
		return err
	}

	user, err := uu.GetById(id)
	if err != nil {
		return err
	}

	user.Login = input.Login
	user.Email = input.Email
	user.About = input.About

	return uu.repo.Update(user)
}

func (uu *UserUsecase) UpdatePassword(id uint, input *models.UpdatePasswordInput) error {
	if input.NewPassword != input.ConfirmPassword {
		return BadPassword.New("Passwords don't match")
	}

	if !models.ValidatePassword(input.NewPassword) {
		return BadPassword.New("Password should be longer than 6 characters")
	}

	encryptedPassword, err := encryptPassword(input.NewPassword)
	if err != nil {
		return err
	}

	user, err:= uu.repo.GetByID(id)
	if err != nil {
		return err
	}

	user.EncryptedPassword = encryptedPassword

	return uu.repo.Update(user)
}

func (uu *UserUsecase) UpdateAvatar(id uint, buffer *bytes.Buffer) (string, error) {
	bytes := buffer.Bytes()

	path, err := image.SaveImage(&bytes)
	if err != nil {
		return "", Wrapf(err, "Updating avatar error, id:%d", id)
	}

	user, err:= uu.repo.GetByID(id)
	if err != nil {
		return "", err
	}

	user.Avatar = path

	return "", uu.repo.Update(user)
}

func (uu *UserUsecase) Delete(id uint) error {
	return uu.repo.Delete(id)
}

func (uu *UserUsecase) ComparePassword(user *models.User, password string) error {
	if bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(password)) == nil {
		return nil
	}

	return Unauthorized.Newf("Incorrect password, id: %d", user.Id)
}

func (uu *UserUsecase) loginIsExist(login string) error {
	_, err := uu.repo.GetByLogin(login)
	if err == nil {
		return LoginIsExist.Newf("User with login %s already exists", login)
	}

	if GetType(err) == NotFound {
		return nil
	}

	return err
}

func (uu *UserUsecase) emailIsExist(email string) error {
	_, err := uu.repo.GetByEmail(email)
	if err == nil {
		return EmailIsExist.Newf("User with email %s already exists", email)
	}

	if GetType(err) == NotFound {
		return nil
	}

	return err
}

func encryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", Wrapf(err, "Internal error")
	}

	return string(hash), nil
}