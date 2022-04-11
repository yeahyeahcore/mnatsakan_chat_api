package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const (
	cost = 14
)

// GeneratePasswordHash ...
func GeneratePasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)

	return fmt.Sprintf("%x", bytes), err
}

// CheckPasswordHash ...
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
