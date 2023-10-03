package postgres

import (
	"github.com/mirzaahmedov/golang-url-shortner/internal/models"
	"github.com/mirzaahmedov/golang-url-shortner/internal/storage"
)

type UserRepository struct {
	storage *PostgresStorage
}

func (s *PostgresStorage) User() storage.UserRepository {
	if s.repository.user == nil {
		s.repository.user = &UserRepository{storage: s}
	}
	return s.repository.user
}

func (r *UserRepository) Create(user *models.User) (*models.User, error) {
	err := r.storage.db.QueryRow(
		`
                  INSERT INTO users (fullname, email, PASSWORD)
                  VALUES ($1, $2, $3)
                  RETURNING id, fullname, email, password
                `,
		user.Fullname,
		user.Email,
		user.Password,
	).Scan(
		&user.ID,
		&user.Fullname,
		&user.Email,
		&user.Password,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}
func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	user := &models.User{}

	err := r.storage.db.QueryRow(
		`
                  SELECT id, fullname, email, password FROM users WHERE email = $1;
                `,
		email,
	).Scan(
		&user.ID,
		&user.Fullname,
		&user.Email,
		&user.Password,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}
