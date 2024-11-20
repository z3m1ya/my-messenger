package repo

import (
	"sync"

	"chats/internal/entities"
)

type InMemoryUserRepository struct {
	users map[string]*entities.User
	mu    sync.RWMutex
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*entities.User),
	}
}

func (repo *InMemoryUserRepository) AddUser(user *entities.User) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.users[user.ID] = user
}

func (repo *InMemoryUserRepository) GetUserByID(id string) (*entities.User, bool) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()
	user, exists := repo.users[id]
	return user, exists
}

func (repo *InMemoryUserRepository) GetUserByUsername(username string) (*entities.User, bool) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()
	for _, user := range repo.users {
		if user.Username == username {
			return user, true
		}
	}
	return nil, false
}
