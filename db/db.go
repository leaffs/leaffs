package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Config struct {
	DatabasePath string
}

type SQLStore struct {
	db *sql.DB
}

func NewSQLStore(cfg Config) (*SQLStore, error) {
	database, err := sql.Open("sqlite3", cfg.DatabasePath)
	if err != nil {
		return nil, fmt.Errorf("could not open database: %v", err)
	}
	if err = database.Ping(); err != nil {
		return nil, fmt.Errorf("could not ping database: %v", err)
	}
	store := &SQLStore{
		db: database,
	}

	return store, nil
}

func (s *SQLStore) Close() error {
	return s.db.Close()
}
