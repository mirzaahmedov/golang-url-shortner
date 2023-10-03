package test

import (
	"github.com/mirzaahmedov/golang-url-shortner/internal/models"
	"github.com/mirzaahmedov/golang-url-shortner/internal/storage"
)

type UserRepository struct {
	storage *TestStorage
}

func (s *TestStorage) User() storage.UserRepository {
	if s.repository.user == nil {
		s.repository.user = &UserRepository{
			storage: s,
		}
	}

	return s.repository.user
}

func (r *UserRepository) Create(user *models.User) (*models.User, error) {
	r.storage.users = append(r.storage.users, *user)

	return user, nil
}
func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	for _, u := range r.storage.users {
		if u.Email == email {
			return &u, nil
		}
	}

	return nil, nil
}
