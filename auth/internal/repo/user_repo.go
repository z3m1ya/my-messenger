package repository

import (
	"auth/internal/entities"
	"fmt"
)

type InMemoryUserRepository struct {
	users map[string]*entities.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*entities.User),
	}
}

func (repo *InMemoryUserRepository) CreateUser(user *entities.User) error {
	if _, exists := repo.users[user.Username]; exists {
		return fmt.Errorf("user already exists")
	}
	repo.users[user.Username] = user
	return nil
}

func (repo *InMemoryUserRepository) FindUser(username string) (*entities.User, error) {
	user, exists := repo.users[username]
	if !exists {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}
