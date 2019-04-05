package database

import (
	"os"

	"github.com/go-redis/redis"
)

//Client  for redis instance
var Client *redis.Client

//RedisInit initializes a new redis client
func RedisInit() error {
	Client = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := Client.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
