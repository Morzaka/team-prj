package plane

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"team-project/services/model"
)

const (
	connStr string = "host=localhost port=5432 user=postgres password=1488 dbname=team-project sslmode=disable"
)

type Plane struct {
	id             uuid.UUID
	departure_city string
	arrival_city   string
}

// GetAll is a function for getting all row and column from table
func GetAll() {

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	rows, err := db.Query("select * from Planes")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	planes := []Plane{}

	for rows.Next() {
		p := Plane{}
		err := rows.Scan(&p.id, &p.departure_city, &p.arrival_city)
		if err != nil {
			fmt.Println(err)
			continue
		}
		planes = append(planes, p)
	}
	for _, p := range planes {
		fmt.Println(p.id, p.departure_city, p.arrival_city)
	}
}

// GetID is a function for getting row using id
func GetID(id uuid.UUID) {

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	row := db.QueryRow("select * from Planes where id = $1", id)
	p := Plane{}
	err = row.Scan(&p.id, &p.departure_city, &p.arrival_city)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(p.id, p.departure_city, p.arrival_city)

}

//Add is a function for adding new row to table
func Add(departure_city string, arrival_city string) {

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	result, err := db.Exec("insert into planes (id,departure_city,arrival_city) values ($1, $2, $3)", model.GenerateID(), departure_city, arrival_city)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result.RowsAffected())

}

// Update is a function for updating number of seats in train using id
func Update(arrival_city string, id uuid.UUID) {

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	result, err := db.Exec("update Planes set arrival_city = $1 where id = $2", arrival_city, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result.RowsAffected())

}

// Delete is a function for deleting row using id
func Delete(id uuid.UUID) {

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	result, err := db.Exec("delete from Planes where id = $1", id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result.RowsAffected())
}
