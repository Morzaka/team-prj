package database

import (
	"database/sql"
	"os"

	"github.com/go-redis/redis"
	//pq lib for using postgres
	_ "github.com/lib/pq"

	"team-project/configurations"
)

var (
	//Db is a pointer to opened database
	Db *sql.DB
	//Client  for redis instance
	Client *redis.Client
)

//PostgresInit connects to postgres database
func PostgresInit() error {
	//database connection string
	//psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", configurations.Config.PgHost, configurations.Config.PgPort, configurations.Config.PgUser, configurations.Config.PgPassword, configurations.Config.PgName)
	//connect to database
	//db, err := sql.Open("postgres", psqlInfo)
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL")) // heroku requires to get connection from env variable
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
