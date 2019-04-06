package database

import (
	"github.com/google/uuid"

	"team-project/services/data"
)

var (
	addTrip = `INSERT INTO public.trip (TripID,TripName,TripTicketId,TripReturnTicketId,TotalTripPrice)
	VALUES ($1, $2, $3, $4, $5);`
	selectAllTrips = `SELECT * FROM public.trip;`
	selectTrip     = `SELECT * FROM public.trip WHERE TripID=$1`
	updateTrip     = `UPDATE public.trip SET TripName = $2, TotalTripPrice = $3 WHERE TripID = $1;`
	deleteTrip     = `DELETE FROM public.trip WHERE TripID = $1;`
)

//AddTrip function add new trip into database table
func AddTrip(trip data.Trip) (data.Trip, error) {
	_, err := Db.Exec(addTrip, trip.TripID, trip.TripName, trip.TripTicketID, trip.TripReturnTicketID, trip.TotalTripPrice)
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
		err := rows.Scan(&p.TripID, &p.TripName, &p.TripTicketID, &p.TripReturnTicketID, &p.TotalTripPrice)
		if err != nil {
			return []data.Trip{}, err
		}
		trips = append(trips, p)
	}
	return trips, nil
}

//UpdateTrip update trip name and total trip price
func UpdateTrip(trip data.Trip, id uuid.UUID) (data.Trip, error) {
	_, err := Db.Exec(updateTrip, id, trip.TripName, trip.TotalTripPrice)
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
	err := Db.QueryRow(selectTrip, id).Scan(&p.TripID, &p.TripName, &p.TripTicketID, &p.TripReturnTicketID, &p.TotalTripPrice)
	if err != nil {
		return data.Trip{}, err
	}
	return p, nil
}
