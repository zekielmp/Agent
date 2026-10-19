package models

import (
	"time"
)

type UserRole string

const (
	UserRoleLandlord UserRole = "landlord"
	UserRoleRenter   UserRole = "renter"
	userRoleAdmiin   UserRole = "admin"
)

// User represents a user in the system.
type User struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"-"  validate:"required"`
	FirstName string    `json:"first_name" validate:"required"`
	LastName  string    `json:"last_name" validate:"required"`
	Phone     string    `json:"phone" validate:"required"`
	IsActive  bool      `json:"is_active" default:"true"`
	Role      UserRole  `json:"role" default:"renter"`
	CreatedAt time.Time `json:"created_at" `
	UpdatedAt time.Time `json:"updated_at" `
}
