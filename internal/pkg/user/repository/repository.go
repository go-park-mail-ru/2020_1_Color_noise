package repository

import (
	"2020_1_Color_noise/internal/models"
	. "2020_1_Color_noise/internal/pkg/error"
	"sync"
)

type Repository struct {
	data []*models.User
	mu   *sync.Mutex
}

func NewRepo() *Repository {
	return &Repository{
		data: make([]*models.User, 0),
		mu:   &sync.Mutex{},
	}
}

func (ur *Repository) Create(user *models.User) (uint, error) {
	ur.mu.Lock()
	user.Id = uint(len(ur.data) + 1)
	ur.data = append(ur.data, user)
	ur.mu.Unlock()

	return user.Id, nil
}

func (ur *Repository) GetByID(id uint) (*models.User, error) {
	for _, user := range ur.data {
		if user.Id == id {
			return user, nil
		}
	}

	return nil, UserNotFound.Newf("User to get not found, id: %d", id)
}

func (ur *Repository) GetByLogin(login string) (*models.User, error) {
	for _, user := range ur.data {
		if user.Login == login {
			return user, nil
		}
	}

	return nil, BadLogin.Newf("User to get not found, login: %s", login)
}

func (ur *Repository) GetByEmail(email string) (*models.User, error) {
	for _, user := range ur.data {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, BadEmail.Newf("User to get not found, email: %s", email)
}

func (ur *Repository) UpdateProfile(id uint, email string, login string) error {
	ur.mu.Lock()
	for i, user := range ur.data {
		if user.Id == id {
			ur.data[i].Email = email
			ur.data[i].Login = login
			ur.mu.Unlock()
			return nil
		}
	}
	ur.mu.Unlock()

	return UserNotFound.Newf("User to update not found, id: %d", id)
}

func (ur *Repository) UpdateDescription(id uint, description *string) error {
	ur.mu.Lock()
	for i, user := range ur.data {
		if user.Id == id {
			ur.data[i].About = *description
			ur.mu.Unlock()
			return nil
		}
	}
	ur.mu.Unlock()

	return UserNotFound.Newf("User to update not found, id: %d", id)
}

func (ur *Repository) UpdatePassword(id uint, encryptredPassword string) error {
	ur.mu.Lock()
	for i, user := range ur.data {
		if user.Id == id {
			ur.data[i].EncryptedPassword = encryptredPassword
			ur.mu.Unlock()
			return nil
		}
	}
	ur.mu.Unlock()

	return UserNotFound.Newf("User to update not found, id: %d", id)
}

func (ur *Repository) UpdateAvatar(id uint, path string) error {
	ur.mu.Lock()
	for i, user := range ur.data {
		if user.Id == id {
			ur.data[i].Avatar = path
			ur.mu.Unlock()
			return nil
		}
	}
	ur.mu.Unlock()

	return UserNotFound.Newf("User to update not found, id: %d", id)
}

func (ur *Repository) Delete(id uint) error {
	ur.mu.Lock()
	for i, user := range ur.data {
		if user.Id == id {
			newData := ur.data[:i]
			for j := i + 1; j < len(ur.data); j++ {
				newData = append(newData, ur.data[j])
			}
			ur.data = newData
			ur.mu.Unlock()
			return nil
		}
	}
	ur.mu.Unlock()

	return UserNotFound.Newf("User to delete not found, id: %d", id)
}

func (ur *Repository) Follow(id uint, subId uint) error {
	return nil
}

func (ur *Repository) Unfollow(id uint, subId uint) error {
	return nil
}