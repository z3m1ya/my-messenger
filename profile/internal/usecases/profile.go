package usecases

import (
	"profile/internal/entities"
	"profile/internal/repo"
)

type ProfileUsecase interface {
	GetProfile(id string) (*entities.Profile, error)
	UpdateProfile(id, name, email string) error
	GetFriends(id string) ([]string, error)
	SendFriendRequest(userId, targetId string) error
}

type profileUsecase struct {
	repo repo.ProfileRepository
}

func NewProfileUsecase(repo repo.ProfileRepository) ProfileUsecase {
	return &profileUsecase{repo: repo}
}

func (u *profileUsecase) GetProfile(id string) (*entities.Profile, error) {
	return u.repo.GetProfile(id)
}

func (u *profileUsecase) UpdateProfile(id, name, email string) error {
	return u.repo.UpdateProfile(id, name, email)
}

func (u *profileUsecase) GetFriends(id string) ([]string, error) {
	return u.repo.GetFriends(id)
}

func (u *profileUsecase) SendFriendRequest(userId, targetId string) error {
	return u.repo.SendFriendRequest(userId, targetId)
}
