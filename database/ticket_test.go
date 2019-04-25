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

var columns = []string{
	"id",
	"train_id",
	"plane_id",
	"user_id",
	"place",
	"ticket_type",
	"discount",
	"price",
	"total_price",
	"name",
	"surname",
}

var testData = data.Ticket{
	ID:         uuid.Must(uuid.Parse("fcb33af4-40a3-4c82-afb1-218731052309")),
	TrainID:    uuid.Must(uuid.Parse("a521d12f-148a-4689-a0ff-e05ec1a40699")),
	PlaneID:    uuid.Must(uuid.Parse("b0ffec41-eb5f-41a4-adab-4d6944a748ad")),
	UserID:     uuid.Must(uuid.Parse("0e3763c6-a7ed-4532-afd7-420c5a48cea9")),
	Place:      61,
	TicketType: "Plane",
	Discount:   "-2%",
	Price:      52.60,
	TotalPrice: 51.15,
	Name:       "Oleh",
	Surname:    "Vynnyk",
}

func openMockDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %s", err)
	}
	Db = db

	return db, mock
}

func addRowItems() []driver.Value {
	var rowItems []driver.Value
	rowItems = append(rowItems, testData.Place, testData.TicketType,
		testData.Discount, testData.Price, testData.TotalPrice,
		testData.Name, testData.Surname)
	return rowItems
}

//TestGetTickets tests function that execute all tickets from DB
func TestGetAllTickets(t *testing.T) {
	db, mock := openMockDB(t)
	defer db.Close()

	rowItems := addRowItems()
	rowItems = append([]driver.Value{testData.ID, testData.TrainID,
		testData.PlaneID, testData.UserID}, rowItems...)

	rows := sqlmock.NewRows(columns).AddRow(rowItems...)
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	if _, err := TicketRepo.AllTickets(); err != nil {
		t.Errorf("error was not expected while getting tickets: %s", err)
	}
}

func TestGetTicket(t *testing.T) {
	db, mock := openMockDB(t)
	defer db.Close()

	id := testData.ID
	errID := uuid.New()

	rowItems := addRowItems()
	rowItems = append([]driver.Value{id, testData.TrainID,
		testData.PlaneID, testData.UserID}, rowItems...)

	rows := sqlmock.NewRows(columns).AddRow(rowItems...)
	mock.ExpectQuery("SELECT").WithArgs(id).WillReturnRows(rows)
	mock.ExpectQuery("SELECT").WithArgs(errID).WillReturnError(fmt.Errorf(
		"no rows found"))
	if _, err := TicketRepo.GetTicket(id); err != nil {
		t.Errorf("error was not expected while getting ticket: %s", err)
	}
	if _, err := TicketRepo.GetTicket(errID); err == nil {
		t.Errorf("error was not expected while getting ticket: %s", err)
	}
	// Checks whether all queued expectations were met in order.
	// If any of them was not met - an error is returned.
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCreateTicket(t *testing.T) {
	db, mock := openMockDB(t)
	defer db.Close()

	newTestData := testData
	newTestData.ID = uuid.Must(uuid.Parse("0e3763c6-a7ed-4532-afd7-420c5a48cea9"))
	rowItems := addRowItems()
	rowItems = append([]driver.Value{newTestData.ID}, rowItems...)
	mock.ExpectExec("INSERT INTO tickets").WithArgs(rowItems...).
		WillReturnResult(
			sqlmock.NewResult(1, 1))

	if err := TicketRepo.CreateTicket(newTestData); err != nil {
		t.Errorf("error was not expected while adding ticket: %s", err)
	}
}

func TestUpdateTicket(t *testing.T) {
	db, mock := openMockDB(t)
	defer db.Close()

	testID := uuid.Must(uuid.Parse("0e3763c6-a7ed-4532-afd7-420c5a48cea9"))
	ticket := data.Ticket{
		ID:         testID,
		Place:      92,
		TicketType: "Boat",
		Discount:   "-25%",
		Price:      100.80,
		TotalPrice: 75.60,
		Name:       "Younyj",
		Surname:    "Orel",
	}

	mock.ExpectExec("UPDATE ticket").WithArgs(testID, ticket.Place,
		ticket.TicketType, ticket.Discount, ticket.Price, ticket.TotalPrice,
		ticket.Name, ticket.Surname).WillReturnResult(sqlmock.NewResult(0, 1))

	if err := TicketRepo.UpdateTicket(ticket); err != nil {
		t.Errorf("error was not expected while updating tickets: %s", err)
	}
}

func TestDeleteTicket(t *testing.T) {
	db, mock := openMockDB(t)
	defer db.Close()

	testID := uuid.Must(uuid.Parse("0e3763c6-a7ed-4532-afd7-420c5a48cea9"))
	mock.ExpectExec("DELETE").WithArgs(testID).WillReturnResult(sqlmock.
		NewResult(0, 1))

	if err := TicketRepo.DeleteTicket(testID); err != nil {
		t.Errorf("error was not expected while deleting ticket: %s", err)
	}
}
