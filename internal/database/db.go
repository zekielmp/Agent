package database

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)
func ConnDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(30 * time.Minute)
	db.SetMaxOpenConns(15)
	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(5 * time.Minute)

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
