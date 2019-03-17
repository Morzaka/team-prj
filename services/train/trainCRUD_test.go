package train

import (
	"database/sql"
	"testing"
)

// TestConnectToDb is a function to test database connection
func TestConnectToDb(t *testing.T) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
}
