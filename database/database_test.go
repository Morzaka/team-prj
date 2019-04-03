package database

import (
	"errors"
	"team-project/configurations"
	"testing"
)

func TestPostgresInit(t *testing.T) {
	testData := []struct {
		config configurations.Configuration
		errexp error
	}{
		{configurations.Configuration{
			PgHost:     "localhost",
			PgPort:     "5432",
			PgUser:     "postgres",
			PgPassword: "postgres",
			PgName:     "travel_test",
		}, nil},
		{configurations.Configuration{
			PgHost:     "localhost",
			PgPort:     "5434", //wrong port
			PgUser:     "postgres",
			PgPassword: "postgres",
			PgName:     "travel_test",
		}, errors.New("dial tcp 127.0.0.1:5434: connect: connection refused")},
		{configurations.Configuration{}, errors.New("dial tcp: lookup port=: Temporary failure in name resolution")},
	}
	for _, testCase := range testData {
		configurations.Config = testCase.config
		err := PostgresInit()
		if testCase.errexp == nil {
			if err != testCase.errexp {
				t.Errorf("Failed postgres connection %s", err)
			}
		} else if err.Error() != testCase.errexp.Error() {
			t.Errorf("Failed postgres connection %s", err)
		}
	}
}

func TestRedisInit(t *testing.T) {
	testData := []struct {
		config configurations.Configuration
		errexp error
	}{
		{configurations.Configuration{
			RedisAddr: "localhost:6379",
		}, nil},
		{configurations.Configuration{
			RedisAddr: "localhost:6374", //wrong port,
		}, errors.New("dial tcp 127.0.0.1:6374: connect: connection refused")},
	}
	for _, testCase := range testData {
		configurations.Config = testCase.config
		err := RedisInit()
		if testCase.errexp == nil {
			if err != testCase.errexp {
				t.Errorf("Failed redis connection %s", err)
			}
		} else if err.Error() != testCase.errexp.Error() {
			t.Errorf("Failed redis connection %s", err)
		}
	}
}
