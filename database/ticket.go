package database

import (
	"errors"

	"team-project/services/data"

	"github.com/google/uuid"
)

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

//AllTickets sends a query for all tickets
func AllTickets() (data.Tickets, error) {
	//rows, err := configurations.DB.Query("SELECT * FROM tickets")
	rows, err := Db.Query(getAllItems)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tkts := make(data.Tickets, 0)
	for rows.Next() {
		tk := data.Ticket{}
		err := rows.Scan(&tk.ID, &tk.TrainID, &tk.PlaneID, &tk.UserID,
			&tk.Place, &tk.TicketType, &tk.Discount, &tk.Price, &tk.TotalPrice,
			&tk.Name, &tk.Surname) /*, &tk.From_place, &tk.Departure_date,
		&tk.Departure_time, &tk.To_place, &tk.Arrival_date,
		&tk.Arrival_time*/// order matters
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

//OneTicket sends a query for one ticket
func OneTicket(id uuid.UUID) (data.Ticket, error) {
	tk := data.Ticket{}
	row := Db.QueryRow(getOneItem, id)

	err := row.Scan(&tk.ID, &tk.TrainID, &tk.PlaneID, &tk.UserID,
		&tk.Place, &tk.TicketType, &tk.Discount, &tk.Price, &tk.TotalPrice,
		&tk.Name, &tk.Surname) /*, &tk.From_place, &tk.Departure_date,
	&tk.Departure_time, &tk.To_place, &tk.Arrival_date,
	&tk.Arrival_time*/// order matters
	if err != nil {
		return tk, err
	}
	return tk, nil
}

//PutTicket sends a query for creating new one ticket
func PutTicket(tk data.Ticket) error {
	_, err := Db.Exec(addOneItem, tk.ID, tk.Place, tk.TicketType, tk.Discount,
		tk.Price, tk.TotalPrice, tk.Name, tk.Surname)
	if err != nil {
		return errors.New("500. Internal Server Error." + err.Error())
	}
	return nil
}

//UpdTicket sends a query for updating one ticket by ID
func UpdTicket(tk data.Ticket) error {
	_, err := Db.Exec(updateItem, tk.ID, tk.Place, tk.TicketType, tk.Discount,
		tk.Price, tk.TotalPrice, tk.Name, tk.Surname)
	if err != nil {
		return err
	}
	return nil
}

//DelTicket sends a query for deleting one ticket by ID
func DelTicket(id uuid.UUID) error {
	_, err := Db.Exec(deleteItem, id)
	if err != nil {
		return errors.New("500. Internal Server Error")
	}
	return nil
}
