package services

import (
	"context"
	"database/sql"
	"time"

	"github.com/zekielmp/Agent/internal/dto"
)

var db *sql.DB

func createUSer(dto *dto.RegisterRequest) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `insert into users
	 (email,password,first_name, last_name,phone)
	 values ($1, $2, $3, $4, $5, $6)
	 returning id
	`

}
