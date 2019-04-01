package database

import (
	"github.com/google/uuid"

	"team-project/logger"
	"team-project/services/data"
)

var (
	insertStatement = `INSERT INTO public.user (id,name,surname,login, password,role)
	VALUES ($1, $2, $3, $4, $5, $6) returning id`
	selectStatement = `SELECT password FROM public.user WHERE login=$1;`
	updateStatement = `UPDATE public.user SET name = $2, surname = $3, login=$4, password=$5, role=$6 WHERE id = $1;`
	deleteStatement = `DELETE FROM public.user WHERE id = $1;`
)

//AddUser adds info about new user to the database
func AddUser(user data.User) {
	//insert values to the database
	_, err := Db.Exec(insertStatement, user.ID, user.Name, user.Surname, user.Signin.Login, user.Signin.Password, user.Role)
	if err != nil {
		logger.Logger.Errorf("Error, %s", err)
	}
}

//GetUserPassword gets user's password and returns password
func GetUserPassword(login string) string {
	var password string
	//get user's password for given login
	err := Db.QueryRow(selectStatement, login).Scan(&password)
	//if there's no matches for login return empty value
	if err != nil {
		return ""
	}
	//else return password
	return password
}

//UpdateUser updates user's personal information
func UpdateUser(user data.User, id uuid.UUID) {
	_, err := Db.Exec(updateStatement, id, user.Name, user.Surname, user.Signin.Login, user.Signin.Password, user.Role)
	if err != nil {
		logger.Logger.Errorf("Error, %s", err)
	}
}

//DeleteUser deletes user's page from db
func DeleteUser(id uuid.UUID) {
	defer Db.Close()
	_, err := Db.Exec(deleteStatement, id)
	if err != nil {
		logger.Logger.Errorf("Error, %s", err)
	}
}
