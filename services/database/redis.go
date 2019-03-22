package database

import (
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"team-project/configurations"
)

var client *redis.Client

func Init() {
	client = redis.NewClient(&redis.Options{
		Addr:     configurations.Config.RedisAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		logrus.Errorf("Redis Error - %s", err.Error())
		panic(err)
	}
	logrus.Infof("launch Redis successful.")
}
