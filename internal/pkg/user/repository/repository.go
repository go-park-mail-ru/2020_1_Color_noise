package repository

import (
	"2020_1_Color_noise/internal/models"
	. "2020_1_Color_noise/internal/pkg/error"
	"sync"
)

type Repository struct {
	data          []*models.User
	mu            *sync.Mutex
	subscriptions map[uint][]uint
	subscribers   map[uint][]uint
	muSub         *sync.Mutex
}

func NewRepo() *Repository {
	return &Repository{
		data:          make([]*models.User, 0),
		mu:            &sync.Mutex{},
		subscriptions: make(map[uint][]uint),
		subscribers:   make(map[uint][]uint),
		muSub:         &sync.Mutex{},
	}
}

func (ur *Repository) Create(user *models.User) (uint, error) {
	_, err := ur.GetByLogin(user.Login)
	if err == nil {
		return 0, LoginIsExist.New("Repo: Error in during creating")
	}

	_, err = ur.GetByEmail(user.Email)
	if err == nil {
		return 0, EmailIsExist.New("Repo: Error in during creating")
	}

	ur.mu.Lock()
	user.Id = uint(len(ur.data) + 1)
	ur.data = append(ur.data, user)
	ur.mu.Unlock()

	return user.Id, nil
}

func (ur *Repository) GetByID(id uint) (*models.User, error) {
	ur.mu.Lock()
	defer ur.mu.Unlock()

	for _, user := range ur.data {
		if user.Id == id {
			user.Subscriptions = len(ur.subscriptions[id])
			user.Subscribers = len(ur.subscribers[id])
			return user, nil
		}
	}

	return nil, UserNotFound.Newf("User to get not found, id: %d", id)
}

func (ur *Repository) GetByLogin(login string) (*models.User, error) {
	ur.mu.Lock()
	defer ur.mu.Unlock()

	for _, user := range ur.data {
		if user.Login == login {
			return user, nil
		}
	}

	return nil, BadLogin.Newf("User to get not found, login: %s", login)
}

func (ur *Repository) GetByEmail(email string) (*models.User, error) {
	ur.mu.Lock()
	defer ur.mu.Unlock()

	for _, user := range ur.data {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, BadEmail.Newf("User to get not found, email: %s", email)
}

func (ur *Repository) UpdateProfile(id uint, email string, login string) error {
	user, err := ur.GetByLogin(login)
	if err == nil && user.Id != id {
		return LoginIsExist.New("Repo: Error in during updating profile")
	}

	user, err = ur.GetByEmail(email)
	if err == nil && user.Id != id {
		return EmailIsExist.New("Repo: Error in during updating profile")
	}

	ur.mu.Lock()
	defer ur.mu.Unlock()

	for i, user := range ur.data {
		if user.Id == id {
			ur.data[i].Email = email
			ur.data[i].Login = login
			return nil
		}
	}

	return UserNotFound.Newf("User to update not found, id: %d", id)
}

func (ur *Repository) UpdateDescription(id uint, description *string) error {
	ur.mu.Lock()
	defer ur.mu.Unlock()

	for i, user := range ur.data {
		if user.Id == id {
			ur.data[i].About = *description
			return nil
		}
	}

	return UserNotFound.Newf("User to update not found, id: %d", id)
}

func (ur *Repository) UpdatePassword(id uint, encryptredPassword string) error {
	ur.mu.Lock()
	defer ur.mu.Unlock()
	for i, user := range ur.data {
		if user.Id == id {
			ur.data[i].EncryptedPassword = encryptredPassword
			return nil
		}
	}

	return UserNotFound.Newf("User to update not found, id: %d", id)
}

func (ur *Repository) UpdateAvatar(id uint, path string) error {
	ur.mu.Lock()
	defer ur.mu.Unlock()

	for i, user := range ur.data {
		if user.Id == id {
			ur.data[i].Avatar = path
			return nil
		}
	}

	return UserNotFound.Newf("User to update not found, id: %d", id)
}

func (ur *Repository) Delete(id uint) error {
	ur.mu.Lock()
	defer ur.mu.Unlock()

	for i, user := range ur.data {
		if user.Id == id {
			newData := ur.data[:i]
			for j := i + 1; j < len(ur.data); j++ {
				newData = append(newData, ur.data[j])
			}
			ur.data = newData
			return nil
		}
	}

	return UserNotFound.Newf("User to delete not found, id: %d", id)
}

func (ur *Repository) Follow(id uint, subId uint) error {
	ur.muSub.Lock()
	defer ur.muSub.Unlock()

	_, err := ur.GetByID(subId)
	if err != nil {
		return Wrapf(err, "Repo: Error in during following, id", id)
	}

	subscriptions := ur.subscriptions[id]
	for _, subscriptionId := range subscriptions {
		if subId == subscriptionId {
			return FollowingIsAlreadyDone.Newf("Repo: Error in during following, id: %d", id)
		}
	}

	subscriptions = append(subscriptions, subId)
	ur.subscriptions[id] = subscriptions

	subscribers := ur.subscribers[subId]
	subscribers = append(subscribers, id)
	ur.subscribers[subId] = subscribers

	return nil
}

func (ur *Repository) Unfollow(id uint, subId uint) error {
	ur.muSub.Lock()
	defer ur.muSub.Unlock()

	_, err := ur.GetByID(subId)
	if err != nil {
		return Wrapf(err, "Repo: Error in during following, id", id)
	}

	subscriptions := ur.subscriptions[id]
	for i, subscriptionId := range subscriptions {
		if subId == subscriptionId {
			newSubscriptions := subscriptions[:i]

			for j := i + 1; j < len(subscriptions); j++ {
				newSubscriptions = append(newSubscriptions , subscriptions[j])
			}
			ur.subscriptions[id] = newSubscriptions

			subscribers := ur.subscribers[subId]
			for k, subscriberId := range subscribers {
				if id == subscriberId {
					newSubscribers := subscribers[:k]

					for m := k + 1; m < len(subscribers); m++ {
						newSubscribers = append(newSubscribers , subscribers[m])
					}
					ur.subscribers[subId] = newSubscribers
				}
			}
			return nil
		}
	}

	return FollowingIsNotYetDone.Newf("Repo: Following is not yet done, id: %d", id)
}

func (ur *Repository) GetSubscribers(id uint, start int, limit int) ([]*models.User, error) {
	_, err := ur.GetByID(id)
	if err != nil {
		return []*models.User{}, Wrapf(err, "Error in during unfollowing, id", id)
	}

	ur.muSub.Lock()
	usersId := ur.subscribers[id]
	ur.muSub.Unlock()

	if start >= len(usersId) {
		start = 0
	}
	usersId = usersId[start:]

	if limit >= len(usersId) {
		limit = len(usersId)
	}
	usersId = usersId[:limit]

	users := make([]*models.User, len(usersId))
	for i, id := range usersId {
		user, _ := ur.GetByID(id)
		users[i] = user
	}

	return users, nil
}

func (ur *Repository) GetSubscriptions(id uint, start int, limit int) ([]*models.User, error) {
	_, err := ur.GetByID(id)
	if err != nil {
		return []*models.User{}, Wrapf(err, "Error in during unfollowing, id", id)
	}

	ur.muSub.Lock()
	usersId := ur.subscriptions[id]
	ur.muSub.Unlock()

	if start >= len(usersId) {
		start = 0
	}
	usersId = usersId[start:]

	if limit >= len(usersId) {
		limit = len(usersId)
	}
	usersId = usersId[:limit]

	users := make([]*models.User, len(usersId))
	for i, id := range usersId {
		user, _ := ur.GetByID(id)
		users[i] = user
	}

	return users, nil
}