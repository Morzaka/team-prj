package database

import (
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"team-project/configurations"
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
		logrus.Errorf("Redis Error - %s", err.Error())
		return
	}
	logrus.Infof("launch Redis successful.")
}
