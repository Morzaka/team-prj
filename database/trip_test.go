package database

import (
	"database/sql"
	"fmt"
	"testing"

	"team-project/services/data"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
)

var nameOfColumns = []string{
	"TripID",
	"TripName",
	"TripTicketID",
	"TripReturnTicketID",
	"TotalTripPrice",
}

var testDataTrip = data.Trip{
	TripID:             uuid.Must(uuid.Parse("fcb33af4-40a3-4c82-afb1-218731052309")),
	TripName:           "CoolWeek",
	TripTicketID:       uuid.Must(uuid.Parse("b0ffec41-eb5f-41a4-adab-4d6944a748ad")),
	TripReturnTicketID: uuid.Must(uuid.Parse("0e3763c6-a7ed-4532-afd7-420c5a48cea9")),
	TotalTripPrice:     125,
}

func openMockTrip(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %s", err)
	}
	Db = db

	return db, mock
}

//TestGet
func TestGetTrips(t *testing.T) {
	db, mock := openMockTrip(t)
	defer db.Close()

	rows := sqlmock.NewRows(nameOfColumns).AddRow(testDataTrip.TripID, testDataTrip.TripName,
		testDataTrip.TripTicketID, testDataTrip.TripReturnTicketID, testDataTrip.TotalTripPrice)
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	_, err := TripRepo.GetTrips()
	if err != nil {
		t.Errorf("error was not expected while getting tickets: %s", err)
	}
}

func TestGetTrip(t *testing.T) {
	db, mock := openMockTrip(t)
	defer db.Close()

	id := testDataTrip.TripID
	errID := uuid.New()

	rows := sqlmock.NewRows(nameOfColumns).AddRow(id, testDataTrip.TripName,
		testDataTrip.TripTicketID, testDataTrip.TripReturnTicketID, testDataTrip.TotalTripPrice)
	mock.ExpectQuery("SELECT").WithArgs(id).WillReturnRows(rows)
	mock.ExpectQuery("SELECT").WithArgs(errID).WillReturnError(fmt.Errorf(
		"no rows found"))
	if _, err := TripRepo.GetTrip(id); err != nil {
		t.Errorf("error was not expected while getting trip: %s", err)
	}
	if _, err := TripRepo.GetTrip(errID); err == nil {
		t.Errorf("error was not expected while getting trip: %s", err)
	}
	// Checks whether all queued expectations were met in order.
	// If any of them was not met - an error is returned.
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestAddTrip(t *testing.T) {
	db, mock := openMockTrip(t)
	defer db.Close()

	newTestData := testDataTrip
	newTestData.TripID = uuid.Must(uuid.Parse("11f3cb14-6750-11e9-a923-1681be663d3e"))
	mock.ExpectExec("INSERT INTO").WithArgs(newTestData.TripID, testDataTrip.TripName,
		testDataTrip.TripTicketID, testDataTrip.TripReturnTicketID, testDataTrip.TotalTripPrice).
		WillReturnResult(
			sqlmock.NewResult(1, 1))
	_, err := TripRepo.AddTrip(newTestData)
	if err != nil {
		t.Errorf("error was not expected while adding trip: %s", err)
	}
}

func TestUpdateTrip(t *testing.T) {
	db, mock := openMockTrip(t)
	defer db.Close()

	testID := uuid.Must(uuid.Parse("0e3763c6-a7ed-4532-afd7-420c5a48cea9"))
	trip := data.Trip{
		TripID:             testID,
		TripName:           "CoolWeek",
		TripTicketID:       uuid.Must(uuid.Parse("0e3763c6-a7ed-4532-afd7-420c5a48cea9")),
		TripReturnTicketID: uuid.Must(uuid.Parse("0e3763c6-a7ed-4532-afd7-420c5a48cea9")),
		TotalTripPrice:     225,
	}

	mock.ExpectExec("UPDATE").WithArgs(testID, trip.TripName,
		trip.TripTicketID, trip.TripReturnTicketID, trip.TotalTripPrice).WillReturnResult(sqlmock.NewResult(0, 1))

	if _, err := TripRepo.UpdateTrip(trip); err != nil {
		t.Errorf("error was not expected while updating user: %s", err)
	}
}

func TestDeleteTrip(t *testing.T) {
	db, mock := openMockTrip(t)
	defer db.Close()

	testID := uuid.Must(uuid.Parse("0e3763c6-a7ed-4532-afd7-420c5a48cea9"))
	mock.ExpectExec("DELETE").WithArgs(testID).WillReturnResult(sqlmock.
		NewResult(0, 1))

	if err := TripRepo.DeleteTrip(testID); err != nil {
		t.Errorf("error was not expected while deleting trip: %s", err)
	}
}
