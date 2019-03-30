package database

import (
	"github.com/google/uuid"

	"team-project/logger"
	"team-project/services/data"
)

//AddUser adds info about new user to the database
func AddUser(user data.User) {
	db := OpenDatabase()
	defer db.Close()
	//insert values to the database
	sqlStatement := `
        INSERT INTO public.user(id,name,surname,login, password,role)
        VALUES ($1, $2, $3, $4, $5, $6)
	returning id`
	id := 0
	err := db.QueryRow(sqlStatement, user.ID, user.Name, user.Surname, user.Signin.Login, user.Signin.Password, user.Role).Scan(&id)
	if err != nil {
		logger.Logger.Errorf("Error, %s", err)
	}
}

//GetUserPassword gets user's password and returns password
func GetUserPassword(login string) string {
	var password string
	db := OpenDatabase()
	defer db.Close()
	//get user's password for given login
	sqlStatement := `SELECT password FROM public.user WHERE login=$1;`
	err := db.QueryRow(sqlStatement, login).Scan(&password)
	//if there's no matches for login return empty value
	if err != nil {
		return ""
	}
	//else return password
	return password
}

//UpdateUser updates user's personal information
func UpdateUser(user data.User, id uuid.UUID) {
	db := OpenDatabase()
	defer db.Close()
	sqlStatement := `UPDATE public.user
	SET name = $2, surname = $3, login=$4, password=$5, role=$6
	WHERE id = $1;`
	_, err := db.Exec(sqlStatement, id, user.Name, user.Surname, user.Signin.Login, user.Signin.Password, user.Role)
	if err != nil {
		logger.Logger.Errorf("Error, %s", err)
	}
}

//DeleteUser deletes user's page from db
func DeleteUser(id uuid.UUID) {
	db := OpenDatabase()
	defer db.Close()
	sqlStatement := `DELETE FROM public.user
	WHERE id = $1;`
	_, err := db.Exec(sqlStatement, id)
	if err != nil {
		logger.Logger.Errorf("Error, %s", err)
	}
}
