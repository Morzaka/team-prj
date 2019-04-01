package database

import (
	"github.com/google/uuid"

	"team-project/services/data"
)

var (
	insertUser = `INSERT INTO public.user (id,name,surname,login, password,role)
	VALUES ($1, $2, $3, $4, $5, $6) returning id`
	selectUser = `SELECT password FROM public.user WHERE login=$1;`
	updateUser = `UPDATE public.user SET name = $2, surname = $3, login=$4, password=$5, role=$6 WHERE id = $1;`
	deleteUser = `DELETE FROM public.user WHERE id = $1;`
)

//AddUser adds info about new user to the database
func AddUser(user data.User) (data.User, error) {
	//insert values to the database
	_, err := Db.Exec(insertUser, user.ID, user.Name, user.Surname, user.Signin.Login, user.Signin.Password, user.Role)
	if err != nil {
		return data.User{}, err
	}
	return user, nil
}

//GetUserPassword gets user's password and returns password
func GetUserPassword(login string) (string, error) {
	var password string
	//get user's password for given login
	err := Db.QueryRow(selectUser, login).Scan(&password)
	//if there's no matches for login return empty value
	if err != nil {
		return "", err
	}
	//else return password
	return password, nil
}

//UpdateUser updates user's personal information
func UpdateUser(user data.User, id uuid.UUID) (data.User, error) {
	_, err := Db.Exec(updateUser, id, user.Name, user.Surname, user.Signin.Login, user.Signin.Password, user.Role)
	if err != nil {
		return data.User{}, err
	}
	return user, nil
}

//DeleteUser deletes user's page from db
func DeleteUser(id uuid.UUID) error {
	_, err := Db.Exec(deleteUser, id)
	if err != nil {
		return err
	}
	return nil
}
