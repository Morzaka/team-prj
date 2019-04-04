package train

import (
	"fmt"
	"strconv"
	"team-project/database"
	"team-project/services/data"

	_ "github.com/lib/pq" // pq lib for using postgres
)

//GetAllTrains is a method
func GetAllTrains() []data.Train {
	rows, err := database.Db.Query("select * from public.train;")
	if err != nil {
		panic(err)
	}

	trains := []data.Train{}
	for rows.Next() {
		t := data.Train{}
		err := rows.Scan(&t.ID, &t.DepartureCity, &t.ArrivalCity, &t.DepartureTime, &t.DepartureDate, &t.ArrivalTime, &t.ArrivalDate)
		if err != nil {
			trains = append(trains, t)
			continue
		}
		trains = append(trains, t)
	}
	return trains
}

//GetTrain is a method
func GetTrain(id string) (data.Train, error) {
	idint, err := strconv.Atoi(id)
	row := database.Db.QueryRow("select * from trains where id = $1", idint)
	t := data.Train{}
	err = row.Scan(&t.ID, &t.DepartureCity, &t.ArrivalCity, &t.DepartureTime, &t.DepartureDate, &t.ArrivalTime, &t.ArrivalDate)
	if err != nil {
		panic(err)
	}
	return t, nil
}

//AddTrain is a method
func AddTrain(t data.Train) {
	_, err := database.Db.Exec("insert into trains (departure_city,arrival_city,departure_time,departure_date,arrival_time,arrival_date) values ($1,$2,$3,$4,$5,$6)", t.DepartureCity, t.ArrivalCity, t.DepartureTime, t.DepartureDate, t.ArrivalTime, t.ArrivalDate)

	if err != nil {
		panic(err)
	}
}

//UpdateTrain is a method
func UpdateTrain(id string, departureCity string, arrivalCity string, departureTime string, departureDate string, arrivalTime string, arrivalDate string) {
	idint, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	_, err = database.Db.Exec("update public.trains set departure_city = $1 , arrival_city = $2, departure_time = $3, departure_date = $4, arrival_time = $5, arrival_date = $6 where id = $7", departureCity, arrivalCity, departureTime, departureDate, arrivalTime, arrivalDate, idint)

	if err != nil {
		panic(err)
	}

	fmt.Println("Updated Train with id", id)
}

//DeleteTrain is a method
func DeleteTrain(id string) {
	idint, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	_, err = database.Db.Exec("delete from trains where id = $1", idint)
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted train with id", id)
}

//GetLastTrain is a method
func GetLastTrain() {
	row := database.Db.QueryRow("select * from trains where id = $1")
	t := data.Train{}
	err := row.Scan(&t.ID, &t.DepartureCity, &t.ArrivalCity, &t.DepartureTime, &t.DepartureDate, &t.ArrivalTime, &t.ArrivalDate)
	if err != nil {
		panic(err)
	}
}
