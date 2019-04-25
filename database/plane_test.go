package database

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"testing"

	"team-project/services/data"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
)

var PlaneColumns = []string{
	"id",
	"departure_City",
	"arrival_City",
}

var testDataPlane = data.Plane{
	ID:         	uuid.Must(uuid.Parse("fcb33af4-40a3-4c82-afb1-218731052309")),
	DepartureCity:	"Lviv",
	ArrivalCity:	"Kharkiv",
}

func openMockPlane(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %s", err)
	}
	Db = db

	return db, mock
}

func Items() []driver.Value {
	var rowItems []driver.Value
	rowItems = append(rowItems, testDataPlane.ID, testDataPlane.DepartureCity, testDataPlane.ArrivalCity)
	return rowItems
}

//TestGetTickets tests function that execute all tickets from DB
func TestGetPlanes(t *testing.T) {
	db, mock := openMockPlane(t)
	defer db.Close()

	rows := sqlmock.NewRows(PlaneColumns).AddRow(testDataPlane.ID,testDataPlane.DepartureCity, testDataPlane.ArrivalCity)
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	if _, err := PlaneRepo.GetPlanes(); err != nil {
		t.Errorf("error was not expected while getting tickets: %s", err)
	}
}

func TestGetPlane(t *testing.T) {
	db, mock := openMockPlane(t)
	defer db.Close()

	id := testDataPlane.ID
	errID := uuid.New()

	rows := sqlmock.NewRows(PlaneColumns).AddRow(testDataPlane.ID,testDataPlane.DepartureCity, testDataPlane.ArrivalCity)
	mock.ExpectQuery("SELECT").WithArgs(id).WillReturnRows(rows)
	mock.ExpectQuery("SELECT").WithArgs(errID).WillReturnError(fmt.Errorf(
		"no rows found"))
	if _, err := PlaneRepo.GetPlane(id); err != nil {
		t.Errorf("error was not expected while getting user: %s", err)
	}
	if _, err := PlaneRepo.GetPlane(errID); err == nil {
		t.Errorf("error was not expected while getting user: %s", err)
	}
	// Checks whether all queued expectations were met in order.
	// If any of them was not met - an error is returned.
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCreatePlane(t *testing.T) {
	db, mock := openMockPlane(t)
	defer db.Close()

	newTestData := testDataPlane
	newTestData.ID = uuid.Must(uuid.Parse("0e3763c6-a7ed-4532-afd7-420c5a48cea9"))
	mock.ExpectExec("INSERT INTO public.plane").WithArgs(newTestData.ID,testDataPlane.DepartureCity, testDataPlane.ArrivalCity).
		WillReturnResult(
			sqlmock.NewResult(1, 1))

	if _,err := PlaneRepo.AddPlane(newTestData); err != nil {
		t.Errorf("error was not expected while adding ticket: %s", err)
	}
}

func TestUpdatePlane(t *testing.T) {
	db, mock := openMockPlane(t)
	defer db.Close()

	testID := uuid.Must(uuid.Parse("0e3763c6-a7ed-4532-afd7-420c5a48cea9"))
	plane := data.Plane{
		ID:         	testID,
		DepartureCity:	"Kyiv",
		ArrivalCity:	"Lviv",
	}

	mock.ExpectExec("UPDATE public.plane").WithArgs(testID, plane.DepartureCity,
		plane.ArrivalCity).WillReturnResult(sqlmock.NewResult(0, 1))

	if _,err := PlaneRepo.UpdatePlane(plane,testID); err != nil {
		t.Errorf("error was not expected while deleting user: %s", err)
	}
}

func TestDeletePlane(t *testing.T) {
	db, mock := openMockPlane(t)
	defer db.Close()

	testID := uuid.Must(uuid.Parse("0e3763c6-a7ed-4532-afd7-420c5a48cea9"))
	mock.ExpectExec("DELETE").WithArgs(testID).WillReturnResult(sqlmock.
		NewResult(0, 1))

	if err := PlaneRepo.DeletePlane(testID); err != nil {
		t.Errorf("error was not expected while deleting user: %s", err)
	}
}

