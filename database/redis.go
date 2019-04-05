package database

import (
	"net/url"
	"os"

	"github.com/go-redis/redis"
)

//Client  for redis instance
var Client *redis.Client

//RedisInit initializes a new redis client
func RedisInit() error {
	env := os.Getenv("REDIS_URL")
	u, err := url.Parse(env)
	password, _ := u.User.Password()
	Client = redis.NewClient(&redis.Options{
		Addr:     u.Host,
		Password: password, // no password set
		DB:       0,        // use default DB
	})

	_, err = Client.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
