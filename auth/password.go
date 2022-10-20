package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func SetPassword(password string) (string, error) {
	if len(password) == 0 {
		return "", errors.New("password is empty")
	}

	bytePassword := []byte(password)

	// Make sure the second param `bcrypt generator cost` between [4, 32)
	passwordHash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(passwordHash), nil
}

func CheckPassword(password string, hashPassword string) bool {
	bytePassword := []byte(password)

	byteHashedPassword := []byte(hashPassword)

	if err := bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword); err != nil {
		return false
	}

	return true
}
