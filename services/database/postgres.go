package database

import (
	"database/sql"
	"fmt"
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

//AddUser adds info about new user to the database
func OpenDatabase() *sql.DB {
	//database connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, dbuser, dbpassword, dbname)
	//connect to database
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	//open database for operations on it
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}


