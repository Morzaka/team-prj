package database

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
)

func testGetAllTrains(t *testing.T) {
	db, mock, err = sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	Db = db
	defer db.Close()
	s := "08307904-f18e-4fb8-9d18-29cfad38ffaf"
	id, err := uuid.Parse(s)
	if err != nil {
		t.Fatal(err)
	}
}
