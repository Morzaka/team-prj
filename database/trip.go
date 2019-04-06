package database

import (
	"team-project/services/data"
)

var (
	addTrip = `INSERT INTO trip (trip_id,trip_name,trip_ticket_id,trip_return_ticket_id,hotel_ticket_id,total_trip_price)
	VALUES ($1, $2, $3, $4, $5, $6) returning trip_id;`
	selectTrip = `SELECT * FROM trip WHERE trip_id=$1;`
	deleteTrip = `DELETE FROM trip WHERE id = $1;`
)

func AddTrip(trip data.Trip) error {
	_, err := Db.Exec(addTrip, trip.Trip_id, trip.Trip_name, trip.Trip_ticket_id, trip.Trip_return_ticket_id, trip.Trip_hotel_id, trip.Total_trip_price)
	return err
}
