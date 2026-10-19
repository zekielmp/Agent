package utils

import "golang.org/x/crypto/bcrypt"

func Hash(password string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), err
}

func Compare(password, hashedpassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(password)); err != nil {
		return nil
	}

	return nil
}
