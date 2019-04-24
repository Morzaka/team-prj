package database

import (
	"team-project/services/data"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
)

func TestGetAllTrains(t *testing.T) {
	db, mock, err := sqlmock.New()
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
	rowOK := sqlmock.NewRows([]string{"id", "departure_city", "arrival_city", "departure_time", "departure_date", "arrival_time", "arrival_date"}).AddRow(id, "Lviv", "Kiev", "13:40", "22.04.2019", "18:40", "22.04.2019")
	mock.ExpectQuery("select").WillReturnRows(rowOK)
	if _, err = Trains.GetAllTrains(); err != nil {
		t.Error("error was occured while getting train in tests ", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error("Values from GetAllTrains are different from expected", err)
	}
}

func TestAddTrain(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	Db = db
	defer db.Close()
	train := data.Train{
		DepartureCity: "Lviv",
		ArrivalCity:   "Kiev",
		DepartureTime: "14:35",
		DepartureDate: "24.04.2019",
		ArrivalTime:   "20:05",
		ArrivalDate:   "24.04.2019",
	}
	mock.ExpectExec("insert into trains").WithArgs(train.DepartureCity, train.ArrivalCity, train.DepartureTime, train.DepartureDate, train.ArrivalTime, train.ArrivalDate).WillReturnResult(sqlmock.NewResult(1, 1))
	if err := Trains.AddTrain(train); err != nil {
		t.Error("An error occured while adding train in tests ", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error("Expectations failed ", err)
	}
}

func TestGetTrain(t *testing.T) {
	db, mock, err := sqlmock.New()
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
	rowOK := sqlmock.NewRows([]string{"id", "departure_city", "arrival_city", "departure_time", "departure_date", "arrival_time", "arrival_date"}).AddRow(id, "Lviv", "Kiev", "13:40", "22.04.2019", "18:40", "22.04.2019")
	mock.ExpectQuery("select").WillReturnRows(rowOK)
	if _, err = Trains.GetTrain(s); err != nil {
		t.Error("error was occured while getting train in tests ", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error("Value from GetTrain is different from expected", err)
	}
}

func TestUpdateTrain(t *testing.T) {
	db, mock, err := sqlmock.New()
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
	train := data.Train{
		ID:            id,
		DepartureCity: "Lviv",
		ArrivalCity:   "Kiev",
		DepartureTime: "14:35",
		DepartureDate: "24.04.2019",
		ArrivalTime:   "20:05",
		ArrivalDate:   "24.04.2019",
	}
	mock.ExpectExec("update public.trains").WithArgs(train.DepartureCity, train.ArrivalCity, train.DepartureTime, train.DepartureDate, train.ArrivalTime, train.ArrivalDate, id).WillReturnResult(sqlmock.NewResult(0, 1))
	if err := Trains.UpdateTrain(train); err != nil {
		t.Error("error was occured while updating train in tests ", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error("Value from Updating is different from expected", err)
	}
}

func TestDeleteTrain(t *testing.T) {
	db, mock, err := sqlmock.New()
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
	mock.ExpectExec("delete").WithArgs(id).WillReturnResult(sqlmock.NewResult(0, 1))
	if err := Trains.DeleteTrain(id); err != nil {
		t.Error("error was occured while deleting train in tests ", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error("result from DeleteTrains is different from expected ", err)
	}
}
