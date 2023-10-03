package postgres

import (
	"database/sql"
)

type PostgresStorage struct {
	db          *sql.DB
	databaseURL string

	repository struct {
		user *UserRepository
		url  *URLRepository
	}
}

func NewStorage(databaseURL string) *PostgresStorage {
	return &PostgresStorage{
		databaseURL: databaseURL,
	}
}

func (s *PostgresStorage) Open() error {
	db, err := sql.Open("postgres", s.databaseURL)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	s.db = db

	return nil
}

func (s *PostgresStorage) Close() error {
	return s.db.Close()
}
