package models

import (
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"

	"team-project/logger"
)

//GenerateID generates unique id
func GenerateID() uuid.UUID {
	id, err := uuid.NewV4()
	if err != nil {
		logger.Logger.Errorf("Error, %s", err)
	}
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
