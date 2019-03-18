package train

import (
	"database/sql"
	"fmt"
	train "team-project/services/train/model"

	_ "github.com/lib/pq"
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
func GetAllTrains() []train.Train {
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

	trains := []train.Train{}
	for rows.Next() {
		t := train.Train{}
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
func GetTrain(id int) train.Train {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	row := db.QueryRow("select * from Train where trainid = $1", id)
	t := train.Train{}
	err = row.Scan(&t.ID, &t.Route)
	if err != nil {
		panic(err)
	}
	return t
}

//AddTrain is a method
func AddTrain(t train.Train) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("insert into Train (route) values ($1)", t.Route)

	if err != nil {
		panic(err)
	}
	fmt.Println("Added Train: ")
	t.Print()
}

//UpdateTrain is a method
func UpdateTrain(id int, route int) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("update public.train set route = $1 where trainid = $2", route, id)

	if err != nil {
		panic(err)
	}

	fmt.Println("Updated Train with id", id)
}

//DeleteTrain is a method
func DeleteTrain(id int) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("delete from Train where trainid = $1", id)
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted train with id", id)
}
