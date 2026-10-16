package database

import (
	"database/sql"
	"os"
)

func ConnDB() (*sql.DB, error) {
	dsn := os.Getenv("DB_DSN")

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
