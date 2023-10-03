package test

import (
	"github.com/mirzaahmedov/golang-url-shortner/internal/models"
	"github.com/mirzaahmedov/golang-url-shortner/internal/storage"
)

type URLRepository struct {
	storage *TestStorage
}

func (s *TestStorage) URL() storage.URLRepository {
	if s.repository.url == nil {
		s.repository.url = &URLRepository{storage: s}
	}
	return s.repository.url
}

func (r *URLRepository) GetAll() ([]models.URL, error) {
	return r.storage.urls, nil
}
func (r *URLRepository) GetByID(id string) (*models.URL, error) {
	for _, u := range r.storage.urls {
		if u.ID == id {
			return &u, nil
		}
	}

	return nil, nil
}
func (r *URLRepository) Create(url *models.URL) (*models.URL, error) {
	r.storage.urls = append(r.storage.urls, *url)

	return url, nil
}
func (r *URLRepository) Update(id string, url *models.URL) (*models.URL, error) {
	for i, u := range r.storage.urls {
		if u.ID == id {
			r.storage.urls[i].Full = url.Full
			r.storage.urls[i].Short = url.Short
			return &r.storage.urls[i], nil
		}
	}

	return nil, nil
}
func (r *URLRepository) Delete(id string) (int64, error) {
	for i, u := range r.storage.urls {
		if u.ID == id {
			r.storage.urls = append(r.storage.urls[:i], r.storage.urls[i+1:]...)
			return 1, nil
		}
	}

	return 0, nil
}
