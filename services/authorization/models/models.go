package models

import "github.com/satori/go.uuid"

type User struct {
	Id       uuid.UUID
	Password string
	Name     string
	Surname  string
	Login    string
	Role     string
}

func NewUser(id uuid.UUID, password, name, surname, login, role string) *User {
	return &User{id, password, name, surname, login, role}
}
