package database

import (
	 "team-project/services/authorization/models"
	 "team-project/services/database"
)

//AddUser adds info about new user to the database
func AddUser(user models.User) int {
        db:=OpenDatabase()
	defer db.Close()
        //insert values to the database
        sqlStatement := `
        INSERT INTO users(name,surname,login, password,role)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id`
        id := 0
        err = db.QueryRow(sqlStatement, user.Name, user.Surname, user.Login, user.Password, user.Role).Scan(&id)
        if err != nil {
                panic(err)
        }
        return id
}

//GetUserPassword gets user's password and returns password
func GetUserPassword(login string) string {
        var password string
        db:=OpenDatabase()
	defer db.Close()
        //get user's password for given login
        sqlStatement := `SELECT password FROM users WHERE login=$1;`
        err = db.QueryRow(sqlStatement, login).Scan(&password)
        //if there's no matches for login return empty value
        if err != nil {
                 return ""
        }
        //else return password
        return password
}
