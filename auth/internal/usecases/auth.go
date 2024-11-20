package usecases

import (
	"auth/internal/entities"
	"fmt"
)

type AuthUseCase struct {
	// Внедрите зависимости, такие как репозитории
}

func NewAuthUseCase() *AuthUseCase {
	return &AuthUseCase{}
}

func (uc *AuthUseCase) VerifyToken(token string) (bool, string, error) {
	if token == "valid-token" {
		return true, "12345", nil
	}
	return false, "", fmt.Errorf("invalid token")
}

func (uc *AuthUseCase) RegisterUser(username, password string) error {
	user := &entities.User{
		ID:       username,
		Username: username,
		Password: password,
	}
	return uc.userRepo.CreateUser(user)
}

func (uc *AuthUseCase) Login(username, password string) (string, error) {
	user, err := uc.userRepo.FindUser(username)
	if err != nil {
		return "", err
	}

	if user.Password != password {
		return "", fmt.Errorf("invalid credentials")
	}

	return "valid-token", nil
}
