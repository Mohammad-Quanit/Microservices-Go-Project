package utils

import "golang.org/x/crypto/bcrypt"

// Takes the user input and the hashed the input and compares them.
// If the hashes match, it returns true.
func CompareHashPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
