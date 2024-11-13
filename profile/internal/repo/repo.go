package repo

import (
	"errors"
	"profile/internal/entities"
)

type ProfileRepository struct {
	profiles map[string]*entities.Profile
}

func (r *ProfileRepository) GetProfile(id string) (*entities.Profile, error) {
	profile, exists := r.profiles[id]
	if !exists {
		return nil, errors.New("profile not found")
	}
	return profile, nil
}

func (r *ProfileRepository) UpdateProfile(id, name, email string) error {
	profile, exists := r.profiles[id]
	if !exists {
		return errors.New("profile not found")
	}
	profile.Name = name
	profile.Email = email
	return nil
}

func (r *ProfileRepository) GetFriends(id string) ([]string, error) {
	profile, exists := r.profiles[id]
	if !exists {
		return nil, errors.New("profile not found")
	}
	return profile.Friends, nil
}

func (r *ProfileRepository) SendFriendRequest(userId, targetId string) error {
	profile, exists := r.profiles[userId]
	if !exists {
		return errors.New("profile not found")
	}

	for _, friendId := range profile.Friends {
		if friendId == targetId {
			return errors.New("already friends")
		}
	}

	profile.Friends = append(profile.Friends, targetId)
	return nil
}
