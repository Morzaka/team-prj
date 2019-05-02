package database

import (
	"database/sql"
	"fmt"

	"github.com/go-redis/redis"
	"net/url"
	"os"
	//pq lib for using postgres
	_ "github.com/lib/pq"
)

var (
	//Db is a pointer to opened database
	Db *sql.DB
	//Client  for redis instance
	Client *redis.Client
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "admin"
	dbname   = "postgres"
)

//PostgresInit connects to postgres database
func PostgresInit() error {
	connection := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connection)
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
