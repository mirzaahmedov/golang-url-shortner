package postgres

import (
	"github.com/mirzaahmedov/golang-url-shortner/internal/models"
	"github.com/mirzaahmedov/golang-url-shortner/internal/storage"
)

type URLRepository struct {
	storage *PostgresStorage
}

func (s *PostgresStorage) URL() storage.URLRepository {
	if s.repository.url == nil {
		s.repository.url = &URLRepository{storage: s}
	}
	return s.repository.url
}

func (r *URLRepository) GetAll() ([]models.URL, error) {
	urls := []models.URL{}

	rows, err := r.storage.db.Query(
		`
                  SELECT id, full, short
                  FROM urls
                `,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		url := models.URL{}
		rows.Scan(
			&url.ID,
			&url.Full,
			&url.Short,
		)
	}

	return urls, nil
}
func (r *URLRepository) GetByID(id string) (*models.URL, error) {
	url := &models.URL{}

	err := r.storage.db.QueryRow(
		`
                  SELECT id, full, short FROM urls WHERE id = $1
                `,
		id,
	).Scan(
		&url.ID,
		&url.Full,
		&url.Short,
	)
	if err != nil {
		return nil, err
	}

	return url, nil
}
func (r *URLRepository) Create(url *models.URL) (*models.URL, error) {
	err := r.storage.db.QueryRow(
		`
                  INSERT INTO users (full, short)
                  VALUES ($1, $2)
                  RETURNING id, full, short
                `,
		url.Full,
		url.Short,
	).Scan(
		url.ID,
		url.Full,
		url.Short,
	)
	if err != nil {
		return nil, err
	}

	return url, nil
}
func (r *URLRepository) Update(id string, url *models.URL) (*models.URL, error) {
	err := r.storage.db.QueryRow(
		`
                  UPDATE users
                  SET full = $1, short = $2
                  WHERE id = $3
                  RETURNING id, full, short
                `,
		url.Full,
		url.Short,
		id,
	).Scan(
		url.ID,
		url.Full,
		url.Short,
	)
	if err != nil {
		return nil, err
	}

	return url, nil
}
func (r *URLRepository) Delete(id string) (int64, error) {
	result, err := r.storage.db.Exec(
		`
                  DELETE users
                  WHERE id = $1
                `,
		id,
	)
	n, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return n, err
}
