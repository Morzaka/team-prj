package database

import (
	"database/sql"
	"fmt"
	"team-project/configurations"
	//pq lib for using postgres
	_ "github.com/lib/pq"
)

//Db is a pointer to opened database
var Db *sql.DB

//PostgresInit connects to postgres database
func PostgresInit() error {
	//database connection string
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", configurations.Config.PgHost, configurations.Config.PgPort, configurations.Config.PgUser, configurations.Config.PgPassword, configurations.Config.PgName)
	//connect to database
	db, err := sql.Open("postgres", psqlInfo)
	//db, err := sql.Open("postgres", os.Getenv("DATABASE_URL")) // heroku requires to get connection from env variable
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	Db = db

	return nil
}
