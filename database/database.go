package database

import (
	"database/sql"
	"net/url"
	"os"

	"github.com/go-redis/redis"

	//pq lib for using postgres
	_ "github.com/lib/pq"
)

var (
	//Db is a pointer to opened database
	Db *sql.DB
	//Client  for redis instance
	Client *redis.Client
)

//PostgresInit connects to postgres database
func PostgresInit() error {
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
	env := os.Getenv("REDIS_URL")
	u, _ := url.Parse(env)
	password, _ := u.User.Password()
	Client = redis.NewClient(&redis.Options{
		Addr:     u.Host,
		Password: password,
		DB:       0, // use default DB
	})
	_, err := Client.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
