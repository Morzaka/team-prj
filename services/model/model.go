package model

import (
	"net/http"
	"net/url"

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

//CheckPasswordHash function valid user's password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//GetID parse id from request
func GetID(r *http.Request) (uuid.UUID, error) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		return uuid.Must(uuid.Parse("")), err
	}
	id, err := uuid.Parse(u.Query().Get("id"))
	if err != nil {
		return id, err
	}
	return id, nil
}
