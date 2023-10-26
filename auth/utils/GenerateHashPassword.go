package utils

import "golang.org/x/crypto/bcrypt"

// Returns a hash value generated from it using a one-way hashing algorithm
func GenerateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
