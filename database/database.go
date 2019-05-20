package database

import (
	"database/sql"
	"net/url"
	"os"
	"team-project-testing/logger"

	"github.com/go-redis/redis"

	//pq lib for using postgres
	_ "github.com/lib/pq"
)

var (
	//Client  for redis instance
	Client *redis.Client
)

//DBManager is a structure for singletom pattern
type DBManager struct {
	Db            *sql.DB
	isInitialized bool
}

var postgresInstance = PostgresInit()

//GetDBManager is a function that returns instance of DBManager (in our case only for postgres)
func GetDBManager() DBManager {
	return postgresInstance
}

//PostgresInit connects to postgres database
func PostgresInit() DBManager {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL")) // heroku requires to get connection from env variable
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		logger.Logger.Error("Error occured while checking connection")
	}
	return DBManager{Db: db, isInitialized: true}
}

//RedisInit initializes a new redis client
func RedisInit() error {
	env := os.Getenv("REDIS_URL")
	u, err := url.Parse(env)
	password, _ := u.User.Password()
	Client = redis.NewClient(&redis.Options{
		Addr:     u.Host,
		Password: password,
		DB:       0, // use default DB
	})
	_, err = Client.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
