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
	repo user.IRepository
}

func NewUsecase(repo user.IRepository) *UserUsecase {
	return &UserUsecase{
		repo: repo,
	}
}

func (uu *UserUsecase) Create(input *models.SignUpInput) (*models.User, error) {
	encryptedPassword, err := encryptPassword(input.Password)
	if err != nil {
		return nil, Wrap(err, "Creating new user error")
	}

	user := &models.User{
		Email:             input.Email,
		Login:             input.Login,
		EncryptedPassword: encryptedPassword,
		Avatar:            "avatar.jpg",
	}

	user, err = uu.repo.Create(user)
	if err != nil {
		return nil, Wrap(err, "Creating new user error")
	}

	return user, nil
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

func (uu *UserUsecase) Search(login string, start int, limit int) ([]*models.User, error) {
	users, err := uu.repo.Search(login, start, limit)
	if err != nil {
		return nil, Wrap(err, "Searching by login user error")
	}

	return users, nil
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
		return Wrap(err, "Updating user password error")
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

func (uu *UserUsecase) Follow(id uint, subId uint) error {
	if err := uu.repo.Follow(id, subId); err != nil {
		return Wrap(err, "Following error")
	}

	return nil
}

func (uu *UserUsecase) Unfollow(id uint, subId uint) error {
	if err := uu.repo.Unfollow(id, subId); err != nil {
		return Wrap(err, "Unfollowing error")
	}

	return nil
}

func (uu *UserUsecase) GetSubscribers(id uint, start int, limit int) ([]*models.User, error) {
	users, err := uu.repo.GetSubscribers(id, start, limit)
	if err != nil {
		return users, Wrap(err, "Getting subscribers error")
	}

	return users, nil
}

func (uu *UserUsecase) GetSubscriptions(id uint, start int, limit int) ([]*models.User, error) {
	users, err := uu.repo.GetSubscriptions(id, start, limit)
	if err != nil {
		return users, Wrap(err, "Getting subscriptions error")
	}

	return users, nil
}

func (uu *UserUsecase) ComparePassword(user *models.User, password string) error {
	if bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(password)) == nil {
		return nil
	}

	return BadPassword.Newf("Password is incorrect")
}

func encryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", Wrapf(err, "Error in during encrypting password")
	}

	return string(hash), nil
}
