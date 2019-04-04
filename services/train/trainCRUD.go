package train

import (
	"database/sql"
	"fmt"
	"strconv"
	"team-project/services/data"

	_ "github.com/lib/pq" // pq lib for using postgres
)

const (
	connStr string = "host=localhost port=5432 user=postgres password=1488 dbname=team-project sslmode=disable"
)

//ConnectToDb is a method
func ConnectToDb() *sql.DB {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return db
}

//GetAllTrains is a method
func GetAllTrains() []data.Train {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	rows, err := db.Query("select * from public.train;")
	if err != nil {
		panic(err)
	}

	trains := []data.Train{}
	for rows.Next() {
		t := data.Train{}
		err := rows.Scan(&t.ID, &t.Route)
		if err != nil {
			t.Route = 0
			trains = append(trains, t)
			continue
		}
		trains = append(trains, t)
	}
	return trains
}

//GetTrain is a method
func GetTrain(id string) (data.Train, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	idint, err := strconv.Atoi(id)
	row := db.QueryRow("select * from Train where trainid = $1", idint)
	t := data.Train{}
	err = row.Scan(&t.ID, &t.Route)
	if err != nil {
		panic(err)
	}
	return t, nil
}

//AddTrain is a method
func AddTrain(t data.Train) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("insert into Train (route) values ($1)", t.Route)

	if err != nil {
		panic(err)
	}
}

//UpdateTrain is a method
func UpdateTrain(id string, route string) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	idint, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	routeint, err := strconv.Atoi(route)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("update public.train set route = $1 where trainid = $2", routeint, idint)

	if err != nil {
		panic(err)
	}

	fmt.Println("Updated Train with id", id)
}

//DeleteTrain is a method
func DeleteTrain(id string) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	idint, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("delete from trains where id = $1", idint)
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted train with id", id)
}

//GetLastTrain is a method
func GetLastTrain() {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	row := db.QueryRow("select id from trains where trainid = $1")
	t := data.Train{}
	err = row.Scan(&t.ID, &t.Route)
	if err != nil {
		panic(err)
	}
}
