package test

import (
	"github.com/mirzaahmedov/golang-url-shortner/internal/models"
)

type TestStorage struct {
	users []models.User
	urls  []models.URL

	repository struct {
		user *UserRepository
		url  *URLRepository
	}
}

func NewStorage() *TestStorage {
	return &TestStorage{
		urls:  []models.URL{},
		users: []models.User{},
	}
}
