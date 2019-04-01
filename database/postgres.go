package database

import (
	"database/sql"
	"fmt"

	//pq lib for using postgres
	_ "github.com/lib/pq"

	"team-project/configurations"
)

var Db *sql.DB

//OpenDatabase connects to postgres database
func PostgresInit() error {
	//database connection string
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		configurations.Config.PgHost, configurations.Config.PgPort, configurations.Config.PgUser, configurations.Config.PgPassword, configurations.Config.PgName)
	//connect to database
	db, err := sql.Open("postgres", psqlInfo)
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
