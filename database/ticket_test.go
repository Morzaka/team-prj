package database

import (
	"github.com/google/uuid"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

//TestGetTickets tests function GetAllUsers
func TestGetAllTickets(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %s", err)
	}
	Db = db
	defer db.Close()
	strID := "98a57b8e-c081-4234-9c29-f29e74f82221"
	id, err := uuid.Parse(strID)
	if err != nil {
		t.Errorf("Error while parsing string to uuid, %s \n", err)
		return
	}

	//TrainID    uuid.UUID `json:"train_id"`
	//	PlaneID    uuid.UUID `json:"plane_id"`
	//	UserID     uuid.UUID `json:"user_id"`
	rows := sqlmock.NewRows([]string{"id", "train_id", "plane_id", "user_id", "place",
		"ticket_type", "discount",
		"price", "total_price", "name", "surname"}).
		AddRow(id, id, id, id, 23, "Bus", "-10%", 2.60, 3.45, "Ivan",
			"Ivanyshyn")

	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	if _, err = GetAllTickets(); err != nil {
		t.Errorf("error was not expected while getting tickets: %s", err)
	}
}