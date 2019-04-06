package model

import (
	"encoding/json"
	"fmt"
	"net/http"

	"team-project/logger"

	"github.com/go-zoo/bone"
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

//JSONEncoding encode data to JSON
func JSONEncoding(w http.ResponseWriter, data interface{}) {
	dataJSON, err := json.Marshal(data)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
	}
	_, err = w.Write(dataJSON)
	if err != nil {
		logger.Logger.Errorf("Error, %s", err)
	}
}

//GetID parse id from request
func GetID(r *http.Request) (uuid.UUID, error) {
	id, err := uuid.Parse(bone.GetValue(r, "id"))
	if err != nil {
		return id, err
	}
	return id, nil
}
