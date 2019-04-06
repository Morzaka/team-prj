package database

import (
	"github.com/google/uuid"

	"team-project/services/data"
)

var (
	addTrip = `INSERT INTO public.trip (trip_id,trip_name,trip_ticket_id,trip_return_ticket_id,total_trip_price)
	VALUES ($1, $2, $3, $4, $5);`
	selectAllTrips = `SELECT * FROM public.trip;`
	selectTrip     = `SELECT * FROM public.trip WHERE Trip_id=$1`
	updateTrip     = `UPDATE public.trip SET trip_name = $2, total_trip_price = $3 WHERE id = $1;`
	deleteTrip     = `DELETE FROM public.trip WHERE id = $1;`
)

//AddTrip function add new trip into database table
func AddTrip(trip data.Trip) (data.Trip, error) {
	_, err := Db.Exec(addTrip, trip.Trip_id, trip.Trip_name, trip.Trip_ticket_id, trip.Trip_return_ticket_id, trip.Total_trip_price)
	if err != nil {
		return data.Trip{}, err
	}
	return trip, err
}

//GetTrips return all trips which exist in table
func GetTrips() ([]data.Trip, error) {
	rows, err := Db.Query(selectAllTrips)
	if err != nil {
		return []data.Trip{}, err
	}
	defer rows.Close()
	trips := []data.Trip{}
	for rows.Next() {
		p := data.Trip{}
		err := rows.Scan(&p.Trip_id, &p.Trip_name, &p.Trip_ticket_id, &p.Trip_return_ticket_id, &p.Total_trip_price)
		if err != nil {
			return []data.Trip{}, err
		}
		trips = append(trips, p)
	}
	return trips, nil
}

//UpdateTrip update trip name and total trip price
func UpdateTrip(trip data.Trip, id uuid.UUID) (data.Trip, error) {
	_, err := Db.Exec(updateTrip, id, trip.Trip_name, trip.Total_trip_price)
	if err != nil {
		return data.Trip{}, err
	}
	return trip, err
}

//DeleteTrip delete trip from table
func DeleteTrip(id uuid.UUID) error {
	_, err := Db.Exec(deleteTrip, id)
	if err != nil {
		return err
	}
	return nil
}

//GetTrip return element which Trip_if equal to id
func GetTrip(id uuid.UUID) (data.Trip, error) {
	p := data.Trip{}
	err := Db.QueryRow(selectTrip, id).Scan(&p.Trip_ticket_id, &p.Trip_name, &p.Trip_ticket_id, &p.Trip_return_ticket_id, &p.Total_trip_price)
	if err != nil {
		return data.Trip{}, err
	}
	return p, nil
}
