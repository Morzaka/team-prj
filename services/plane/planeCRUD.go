package plane

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	connStr string = "host=localhost port=5432 user=postgres password=1488 dbname=team-project sslmode=disable"
)

type Plane struct {
	id    int
	route string
	seats int
	price int
}

// GetAll is a function for getting all row and column from table
func GetAll() {

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from Planes")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	planes := []Plane{}

	for rows.Next() {
		p := Plane{}
		err := rows.Scan(&p.id, &p.route, &p.seats, &p.price)
		if err != nil {
			fmt.Println(err)
			continue
		}
		planes = append(planes, p)
	}
	for _, p := range planes {
		fmt.Println(p.id, p.route, p.seats, p.price)
	}
}

// GetID is a function for getting row using id
func GetID(id int) {

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	row := db.QueryRow("select * from Planes where id = $1", id)
	p := Plane{}
	err = row.Scan(&p.id, &p.route, &p.seats, &p.price)
	if err != nil {
		panic(err)
	}
	fmt.Println(p.id, p.route, p.seats, p.price)

}

//Add is a function for adding new row to table
func Add(route string, seats int, price float64) {

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("insert into planes (route, seats,price) values ($1, $2, $3)", route, seats, price)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.RowsAffected())

}

// Update is a function for updating number of seats in train using id
func Update(seats int, id int) {

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("update Planes set seats = $1 where id = $2", seats, id)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected())

}

// Delete is a function for deleting row using id
func Delete(id int) {

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("delete from Planes where id = $1", id)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected())
}
