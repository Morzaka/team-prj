package database

import (
	"database/sql"
	"fmt"
	
	//pq lib for using postgres
	_ "github.com/lib/pq"
	"github.com/go-redis/redis"

	"team-project/configurations"
)

//Db is a pointer to opened database
var (
	Db *sql.DB
//Client  for redis instance
	Client *redis.Client
)

//PostgresInit connects to postgres database
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

//RedisInit initializes a new redis client
func RedisInit() error {
	Client = redis.NewClient(&redis.Options{
		Addr:     configurations.Config.RedisAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := Client.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
