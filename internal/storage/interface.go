package storage

import (
	"github.com/mirzaahmedov/golang-url-shortner/internal/models"
)

type Storage interface {
	User() UserRepository
	URL() URLRepository
}

type UserRepository interface {
	Create(*models.User) (*models.User, error)
	GetByEmail(string) (*models.User, error)
}
type URLRepository interface {
	GetAll() ([]models.URL, error)
	GetByID(string) (*models.URL, error)
	Create(*models.URL) (*models.URL, error)
	Update(string, *models.URL) (*models.URL, error)
	Delete(string) (int64, error)
}
