package database

import (
	"database/sql"
	"fmt"
	"team-project/services/authorization/models"
	_ "github.com/lib/pq"
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

//AddUser ads info about new user to the database
func AddUser(user models.User) int {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, dbuser, dbpassword, dbname)
	fmt.Println(psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
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
//GetUser get's user's password and returns password
func GetUser(login string) string {
	var password string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, dbuser, dbpassword, dbname)
	fmt.Println(psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	sqlStatement := `SELECT password FROM users WHERE login=$1;`
	err = db.QueryRow(sqlStatement, login).Scan(&password)
	if err != nil {
		panic(err)
	}
	return password
}

