package database

import (
	"team-project/services/data"

	"github.com/google/uuid"
)

//TicketRepository for mocking
type TicketRepository interface {
	AllTickets() ([]data.Ticket, error)
	GetTicket(id uuid.UUID) (data.Ticket, error)
	CreateTicket(tk data.Ticket) error
	UpdateTicket(tk data.Ticket) error
	DeleteTicket(id uuid.UUID) error
}

//IUser structure contains interface TicketRepository
type ticketRepository struct {
	ticketRepo TicketRepository
}

// TicketRepo is a variable for accessing to Ticket mocked Interface
var TicketRepo TicketRepository = &ticketRepository{}

var (
	getAllItems = "SELECT id, train_id, plane_id, user_id, place, " +
		"ticket_type," +
		" discount, price, total_price, name, surname FROM tickets;"
	getOneItem = "SELECT id, train_id, plane_id, user_id, place, " +
		"ticket_type," +
		" discount, price, total_price, name, " +
		"surname FROM tickets WHERE id = $1;"
	addOneItem = "INSERT INTO tickets (id, place, ticket_type, discount, " +
		"price, total_price, name, surname) " +
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
	updateItem = "UPDATE tickets SET place=$2, ticket_type=$3, " +
		"discount=$4, " +
		"price=$5, total_price=$6, name=$7, surname=$8 WHERE id=$1;"
	deleteItem = "DELETE FROM tickets WHERE id=$1;"
)

//GetAllTickets sends a query for all tickets
func (*ticketRepository) AllTickets() ([]data.Ticket, error) {
	//rows, err := configurations.DB.Query("SELECT * FROM tickets")
	rows, err := Db.Query(getAllItems)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tkts := make([]data.Ticket, 0)
	for rows.Next() {
		tk := data.Ticket{}
		err := rows.Scan(&tk.ID, &tk.TrainID, &tk.PlaneID, &tk.UserID,
			&tk.Place, &tk.TicketType, &tk.Discount, &tk.Price, &tk.TotalPrice,
			&tk.Name, &tk.Surname) /*, &tk.From_place, &tk.Departure_date,
		&tk.Departure_time, &tk.To_place, &tk.Arrival_date,
		&tk.Arrival_time*/ // order matters
		if err != nil {
			return nil, err
		}
		tkts = append(tkts, tk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return tkts, nil
}

//GetTicket sends a query for one ticket
func (*ticketRepository) GetTicket(id uuid.UUID) (data.Ticket, error) {
	tk := data.Ticket{}
	row := Db.QueryRow(getOneItem, id)

	err := row.Scan(&tk.ID, &tk.TrainID, &tk.PlaneID, &tk.UserID,
		&tk.Place, &tk.TicketType, &tk.Discount, &tk.Price, &tk.TotalPrice,
		&tk.Name, &tk.Surname) /*, &tk.From_place, &tk.Departure_date,
	&tk.Departure_time, &tk.To_place, &tk.Arrival_date,
	&tk.Arrival_time*/ // order matters
	if err != nil {
		return tk, err
	}
	return tk, nil
}

//CreateTicket sends a query for creating new one ticket
func (*ticketRepository) CreateTicket(tk data.Ticket) error {
	_, err := Db.Exec(addOneItem, tk.ID, tk.Place, tk.TicketType, tk.Discount,
		tk.Price, tk.TotalPrice, tk.Name, tk.Surname)
	if err != nil {
		return err
	}
	return nil
}

//UpdateTicket sends a query for updating one ticket by ID
func (*ticketRepository) UpdateTicket(tk data.Ticket) error {
	_, err := Db.Exec(updateItem, tk.ID, tk.Place, tk.TicketType, tk.Discount,
		tk.Price, tk.TotalPrice, tk.Name, tk.Surname)
	if err != nil {
		return err
	}
	return nil
}

//DeleteTicket sends a query for deleting one ticket by ID
func (*ticketRepository) DeleteTicket(id uuid.UUID) error {
	_, err := Db.Exec(deleteItem, id)
	if err != nil {
		return err
	}
	return nil
}
