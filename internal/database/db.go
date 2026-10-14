package database

import (
	"database/sql"
	"os"
)

func connDB() (*sql.DB, error) {
	dsn := os.Getenv("DB_DSN")

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
