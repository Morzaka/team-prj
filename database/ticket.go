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
		err := rows.Scan(&tk.Id, &tk.TrainId, &tk.PlaneId, &tk.UserId,
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

func OneTicket(id uuid.UUID) (data.Ticket, error) {
	tk := data.Ticket{}
	row := Db.QueryRow(getOneItem, id)

	err := row.Scan(&tk.Id, &tk.TrainId, &tk.PlaneId, &tk.UserId,
		&tk.Place, &tk.TicketType, &tk.Discount, &tk.Price, &tk.TotalPrice,
		&tk.Name, &tk.Surname) /*, &tk.From_place, &tk.Departure_date,
	&tk.Departure_time, &tk.To_place, &tk.Arrival_date,
	&tk.Arrival_time*/// order matters
	if err != nil {
		return tk, err
	}
	return tk, nil
}

func PutTicket(tk data.Ticket) error {
	_, err := Db.Exec(addOneItem, tk.Id, tk.Place, tk.TicketType, tk.Discount,
		tk.Price, tk.TotalPrice, tk.Name, tk.Surname)
	if err != nil {
		return errors.New("500. Internal Server Error." + err.Error())
	}
	return nil
}

func UpdTicket(tk data.Ticket) error {
	_, err := Db.Exec(updateItem, tk.Id, tk.Place, tk.TicketType, tk.Discount,
		tk.Price, tk.TotalPrice, tk.Name, tk.Surname)
	if err != nil {
		return err
	}
	return nil
}

func DelTicket(id uuid.UUID) error {
	_, err := Db.Exec(deleteItem, id)
	if err != nil {
		return errors.New("500. Internal Server Error")
	}
	return nil
}
