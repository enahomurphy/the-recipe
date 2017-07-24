package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword takes is a password and returns a hash
func HashPassword(password string) string {
	const DefaultCost = 10
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), DefaultCost)

	return string(hash)
}

// ValidatePassword takes in a hash and a password
// and validates the hash against the password. if there is match
// true is returned else false
func ValidatePassword(password string, hash string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return false
	}
	return true
}
