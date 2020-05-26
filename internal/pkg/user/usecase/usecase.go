package usecase

import (
	"2020_1_Color_noise/internal/models"
	. "2020_1_Color_noise/internal/pkg/error"
	"2020_1_Color_noise/internal/pkg/user"
	"2020_1_Color_noise/internal/pkg/utils"
	"github.com/asaskevich/govalidator"
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
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		err = WithMessage(BadRequest.Wrap(err, "Creating new user error"),
			"Password should be longer than 6 characters and shorter 100. "+
				"Login should be letters and numbers, and shorter than 20 characters "+
				"Email should be like hello@example.com and shorter than 50 characters.")
		return nil, err
	}

	encryptedPassword, err := utils.EncryptPassword(input.Password)
	if err != nil {
		return nil, Wrap(err, "Creating new user error")
	}

	us := &models.User{
		Email:             input.Email,
		Login:             input.Login,
		EncryptedPassword: encryptedPassword,
		Avatar:            "avatar.jpg",
	}

	us, err = uu.repo.Create(us)
	if err != nil {
		return nil, Wrap(err, "Creating by id user error")
	}

	return us, nil
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
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		err = WithMessage(BadRequest.Wrap(err, "Updating user profile error"),
			"Login should be letters and numbers, shorter than 20 characters "+
				"Email should be like hello@example.com and shorter than 50 characters")
		return err
	}

	err = uu.repo.UpdateProfile(id, input.Email, input.Login)
	if err != nil {
		return Wrap(err, "Updating user profile error")
	}

	return nil
}

func (uu *UserUsecase) UpdatePassword(id uint, input *models.UpdatePasswordInput) error {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		err = WithMessage(BadRequest.Wrap(err, "request"),
			"Password should be longer than 6 characters and shorter 100.")
		return err
	}

	encryptedPassword, err := utils.EncryptPassword(input.Password)
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
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		err = WithMessage(BadRequest.Wrap(err, "Updating user description error"),
			"Description should be shorter than 1000 characters.")
		return err
	}

	err = uu.repo.UpdateDescription(id, &input.Description)
	if err != nil {
		return Wrap(err, "Updating user description error")
	}

	return nil
}

func (uu *UserUsecase) UpdateAvatar(id uint, buffer []byte) (string, error) {
	path, err := utils.SaveImage(&buffer)
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
		err = WithMessage(FollowingIsAlreadyDone.Wrap(err, "You are following this user"),
			"You are following this user")
		return err
	}

	return nil
}

func (uu *UserUsecase) IsFollowed(id uint, subId uint) (bool, error) {
	s, err := uu.repo.IsFollowed(id, subId)
	if err != nil {
		return false, Wrap(err, "Error in during check follow")
	}

	return s, nil
}

func (uu *UserUsecase) Unfollow(id uint, subId uint) error {
	if err := uu.repo.Unfollow(id, subId); err != nil {
		err = WithMessage(FollowingIsNotYetDone.Wrap(err, "You are not following this user"),
			"You are not following this user")
		return err
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

func (uu *UserUsecase) UpdatePreferences(userId uint, preferences []string) error {
	u, err := uu.repo.GetByID(userId)
	if err != nil {
		return Wrap(err, "UpdatePreferences error")
	}

	for _, tag := range u.Tags {
		if len(preferences) == 10 {
			break
		}
		preferences = append(preferences, tag)
	}

	if err := uu.repo.UpdatePreferences(userId, preferences); err != nil {
		return Wrap(err, "UpdatePreferences error")
	}

	return nil
}