package services

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/zekielmp/Agent/internal/dto"
	"github.com/zekielmp/Agent/internal/models"
)

var db *sql.DB

func CreateUser(user *models.User) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `insert into users
	 (email,password,first_name, last_name,phone)
	 values ($1, $2, $3, $4, $5, $6)
	 returning id
	`

	err := db.QueryRowContext(ctx, query, user.Email, user.Password, user.FirstName, user.LastName, user.Phone).Scan(&user.ID)
	if err != nil {
		return "", fmt.Errorf("failed to create user:%w", err)
	}

	return user.ID, nil
}

func GetUser(payload dto.LoginRequest) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		select id, first_name, last_name, password, email, created_at, updated_at 
		from users where email = $1;`

	var user models.User
	rows := db.QueryRowContext(ctx, query, payload.Email)
	if err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &user, nil
}
