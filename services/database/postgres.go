package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // pq lib for using postgres
	"team-project/services/authorization/models"
)

const (
	host       = "localhost"
	port       = 5432
	dbuser     = "postgres"
	dbpassword = "postgres"
	dbname     = "travelling"
)

//Database travelling
//table user (id serial, name text, surname text, login text, password text, role text)

//AddUser adds info about new user to the database
func AddUser(user models.User) int {
	//database connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, dbuser, dbpassword, dbname)
	//connect to database
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//open database for operations on it
	err = db.Ping()
	if err != nil {
		panic(err)
	}
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

//GetUser gets user's password and returns password
func GetUser(login string) string {
	var password string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, dbuser, dbpassword, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
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
