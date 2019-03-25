package database

import (
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"team-project/configurations"
	"time"
)

var client *redis.Client

//Init initializes a new redis client
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

//SetRedisValue adds key/value to redis db
func SetRedisValue(key string, value string) (error) {
	err := client.Set(key, value, 15*time.Minute).Err()
	return err
}

//GetRedisValue returns value that corresponds the key from the db
func GetRedisValue(key string) (string, error) {
	res, err := client.Get(key).Result()
	return res,err
}

//DelRedisValue deletes value from db
func DelRedisValue(key string) (error) {
	return client.Del(key).Err()
}
