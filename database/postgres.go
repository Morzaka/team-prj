package database

import (
	"database/sql"
	"fmt"

	//pq lib for using postgres
	_ "github.com/lib/pq"

	"team-project/configurations"
	"team-project/logger"
)

//Database travelling
//table user (id serial, name text, surname text, login text, password text, role text)

//OpenDatabase connects to postgres database
func OpenDatabase() *sql.DB {
	fmt.Println("Connect")
	//database connection string
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		configurations.Config.PgHost, configurations.Config.PgPort, configurations.Config.PgUser, configurations.Config.PgPassword, configurations.Config.PgName)
	//connect to database
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		logger.Logger.Errorf("Error, %s", err)
	}
	//open database for operations on it
	err = db.Ping()
	if err != nil {
		logger.Logger.Errorf("Error, %s", err)
	}
	return db
}
