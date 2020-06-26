package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword ...
func HashPassword(password string) (string, error) {
	hb, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MaxCost)
	return string(hb), err
}

// CheckPassword ...
func CheckPassword(password string, hashed []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashed, []byte(password))
	return err == nil
}
