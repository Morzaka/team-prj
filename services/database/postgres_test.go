package database


import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

type databaseTestCase struct {
	name        string
	isConnected bool
	want        bool
}

func TestSetPostgresConnected(t *testing.T) {
	IsPostgresConnected = false
	SetPostgresConnected()
	if IsPostgresConnected != true {
		t.Error("Method SetPostgresConnected is not switching value to true")
	}
}

func TestSetupPostgres(t *testing.T) {
	tests := []databaseTestCase{
		{
			name:        "Connect postgres first time",
			isConnected: false,
			want:        false,
		},
		{
			name:        "Connect postgres second time",
			isConnected: true,
			want:        true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			IsPostgresConnected = tc.isConnected
			switch tc.want {
			case false:
				//Error skipped because we are checking reusing of connection, not if db is up
				db, _, _ := sqlmock.NewWithDSN(DSN(PostgreSQLInfo{})) //.Config.Database.PostgreSQL))
				DB, _ = SetupPostgres(Info{})
				if tc.want != (DB == db) {
					t.Error("Error in test: ", tc.name)
				}
			case true:
				DB, _ = SetupPostgres(Info{})
				db, err := SetupPostgres(Info{})
				if err != nil {
					t.Log("Postgres is not runing")
				}
				if DB != db {
					t.Error("SetupPostgres is not reusing connection, test failed: ", tc.name)
				}
			}

		})
	}
}