package database

import (
	"database/sql"
	"fmt"
	"log"
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

var (
	DB                  *sql.DB
	IsPostgresConnected bool

)

type Type string


type Info struct {
	// Database type
	Type []Type
	// Postgres info if used
	PostgreSQL PostgreSQLInfo

}

type PostgreSQLInfo struct {
	Hostname     string
	Port         int
	DatabaseName string
	Username     string
	Password     string
}

func SetPostgresConnected() {
	IsPostgresConnected = true
}

func DSN(ci PostgreSQLInfo) string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		ci.Hostname, ci.Port, ci.Username, ci.Password, ci.DatabaseName)
}
func SetupPostgres(d Info) (*sql.DB, error) {
	if IsPostgresConnected {
		return DB, nil
	}
	db, err := sql.Open("postgres", DSN(d.PostgreSQL))
	if err != nil {
		log.Println(err)
	}

	err = db.Ping()
	if err != nil {
		log.Println(err)
	}
	SetPostgresConnected()
	return db, err
}