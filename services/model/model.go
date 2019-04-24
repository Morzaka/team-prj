package model

import (
	"golang.org/x/crypto/bcrypt"
)

//Model for mocking
type Model interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

//IModel struct contains interface Model
type IModel struct {
	ModelMethods Model
}

//HelperModel is an instance presented IModel
var HelperModel Model = &IModel{}

//HashPassword function hashes user's password
func (*IModel) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPasswordHash function valid user's password
func (*IModel) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
