package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(password string) (string, error) {
	decodedPass := []byte(password)
	hashed, err := bcrypt.GenerateFromPassword(decodedPass, bcrypt.DefaultCost)

	return string(hashed), err
}

func CompareHashPassword(password string, claimedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(claimedPassword))

	return err == nil
}