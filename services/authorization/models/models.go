package models


type User struct {
	Password string
	Name     string
	Surname  string
	Login    string
	Role     string
}

//NewUser creates new user with data, received from registration
func NewUser(password, name, surname, login, role string) User {
	return User{password, name, surname, login, role}
}
