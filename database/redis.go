package database

import (
	"github.com/go-redis/redis"

	"team-project/configurations"
	"team-project/logger"
)

//Client  for redis instance
var Client *redis.Client

//init initializes a new redis client
func init() {
	Client = redis.NewClient(&redis.Options{
		Addr:     configurations.Config.RedisAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := Client.Ping().Result()
	if err != nil {
		logger.Logger.Errorf("Error, %s", err)
	}
	//	logger.Logger.Info("launch Redis successful.")
}
