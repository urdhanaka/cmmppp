package utils

import "golang.org/x/crypto/bcrypt"

// hash password with a work factor of workFactor
func HashPassword(password string, workFactor int) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), workFactor)
	if err != nil {
		return "", err
	}

	return string(hashedBytes), nil
}

// compare pass with actualPassword. return true if pass matches the hash
func ComparePassword(pass, actualPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(pass), []byte(actualPassword))

	return err == nil
}
