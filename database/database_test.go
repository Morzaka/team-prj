package database

import(
	"testing"
	"team-project/configurations"
)

func TestPostgresInit(t *testing.T){
	configurations.Config = configurations.Configuration{
		PgHost: "localhost",
		PgPort: "5432",
		PgUser: "postgres",
		PgPassword: "postgres",
		PgName: "travel_test",
	}
	err:=PostgresInit()
	if err!=nil{
		t.Errorf("Failed postgres connection")
	}
	configurations.Config = configurations.Configuration{}
	err=PostgresInit()
	if err==nil{
		t.Errorf("Failed: database not active")
	}
}

func TestRedisInit(t *testing.T){
	configurations.Config = configurations.Configuration{
		RedisAddr: "localhost:6379",
	}
	err:=RedisInit()
	if err!=nil{
		t.Errorf("Failed redis connection")
	}
	configurations.Config = configurations.Configuration{
		RedisAddr: "localhost:5432",
	}
	err=RedisInit()
	if err==nil{
		t.Errorf("Failed: no active database")
	}
}