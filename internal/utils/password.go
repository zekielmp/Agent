package utils

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// Hash return a string from bcypt hash
func Hash(password string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), err
}

// Compare, compares hashed password with password
func Compare(hashedpassword, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(password)); err != nil {
		log.Println(err.Error())
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return ErrMismatchedHashAndPassword
		default:
			return ErrIncorrectPassword
		}
	}

	return nil
}
