package model

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

//GenerateID generates unique id
func GenerateID() uuid.UUID {
	id := uuid.New()
	return id
}

//HashPassword function hashes user's password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPasswordHash function valides user's password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
