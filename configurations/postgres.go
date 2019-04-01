package configurations

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var (
	DB       *sql.DB
	host     = "host=" + Config.PgHost
	port     = " port=" + Config.PgPort
	user     = " user=" + Config.PgUser
	password = " password=" + Config.PgPassword
	dbname   = " dbname=" + Config.PgName
)

func init() {

	var err error
	connStr := host + port + user + password + dbname + " sslmode=disable"

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}
}
