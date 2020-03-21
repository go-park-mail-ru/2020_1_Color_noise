package usecase

import (
	"2020_1_Color_noise/internal/models"
	. "2020_1_Color_noise/internal/pkg/error"
	"2020_1_Color_noise/internal/pkg/image"
	"2020_1_Color_noise/internal/pkg/user"
	"bytes"
	"golang.org/x/crypto/bcrypt"
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
	/*if !models.ValidateEmail(input.Email) {
		return 0, BadEmail.New("Email incorrect")
	}

	if !models.ValidateLogin(input.Login) {
		return 0, BadLogin.New("Login incorrect")
	}

	if !models.ValidatePassword(input.Password) {
		fmt.Println("pass")
		return 0, BadPassword.New("Password should be longer than 6 characters")
	}*/
	/*
	if err := uu.emailIsExist(input.Email); err != nil {
		return 0, Wrap(err, "Creating new user error")
	}

	if err := uu.loginIsExist(input.Login); err != nil {
		return 0, Wrap(err, "Creating new user error")
	}
	*/
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

	id, err := uu.repo.Create(user)
	if err != nil {
		return 0, Wrap(err, "Creating new user error")
	}

	return id, nil
}

func (uu *UserUsecase) GetById(id uint) (*models.User, error) {
	user, err := uu.repo.GetByID(id)
	if err != nil {
		return nil, Wrap(err, "Getting by id user error")
	}

	return user, nil
}

func (uu *UserUsecase) GetByLogin(login string) (*models.User, error) {
	user, err := uu.repo.GetByLogin(login)
	if err != nil {
		return nil, Wrap(err, "Getting by login user error")
	}

	return user, nil

}

func (uu *UserUsecase) UpdateProfile(id uint, input *models.UpdateProfileInput) error {
	err := uu.repo.UpdateProfile(id, input.Email, input.Login)
	if err != nil {
		return Wrap(err, "Updating user profile error")
	}

	return nil
}

func (uu *UserUsecase) UpdatePassword(id uint, input *models.UpdatePasswordInput) error {
	encryptedPassword, err := encryptPassword(input.Password)
	if err != nil {
		return Wrapf(err, "Updating user password error, id: %d", id)
	}

	err = uu.repo.UpdatePassword(id, encryptedPassword)
	if err != nil {
		return Wrap(err, "Updating user password error")
	}

	return uu.repo.UpdatePassword(id, encryptedPassword)
}

func (uu *UserUsecase) UpdateDescription(id uint, input *models.UpdateDescriptionInput) error {
	err := uu.repo.UpdateDescription(id, &input.Description)
	if err != nil {
		return Wrap(err, "Updating user description error")
	}

	return nil
}

func (uu *UserUsecase) UpdateAvatar(id uint, buffer *bytes.Buffer) (string, error) {
	bytes := buffer.Bytes()

	path, err := image.SaveImage(&bytes)
	if err != nil {
		return "", Wrapf(err, "Updating avatar error, id:%d", id)
	}

	err = uu.repo.UpdateAvatar(id, path)
	if err != nil {
		return "", Wrap(err, "Updating avatar error")
	}

	return path, nil
}

func (uu *UserUsecase) Delete(id uint) error {
	if err := uu.repo.Delete(id); err != nil {
		return Wrap(err, "Deleting error")
	}

	return nil
}

func (uu *UserUsecase) ComparePassword(user *models.User, password string) error {
	if bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(password)) == nil {
		return nil
	}

	return BadPassword.Newf("Password is incorrect, id: %d", user.Id)
}

/*
func (uu *UserUsecase) loginIsExist(login string) error {
	_, err := uu.repo.GetByLogin(login)
	if err == nil {
		return LoginIsExist.Newf("User with login %s already exists", login)
	}

	if GetType(err) == BadLogin {
		return nil
	}

	return err
}

func (uu *UserUsecase) emailIsExist(email string) error {
	_, err := uu.repo.GetByEmail(email)
	if err == nil {
		return EmailIsExist.Newf("User with email %s already exists", email)
	}

	if GetType(err) == BadEmail {
		return nil
	}

	return err
}
 */

func encryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", Wrapf(err, "Error in during encrypting password")
	}

	return string(hash), nil
}