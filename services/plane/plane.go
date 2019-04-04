package plane

import (
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"team-project/database"
	"team-project/services/model"
)

type Plane struct {
	id             uuid.UUID
	departure_city string
	arrival_city   string
}

// GetAll is a function for getting all row and column from table
func GetAll() {

	rows, err := database.Db.Query("select * from Planes")
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

	row := database.Db.QueryRow("select * from Planes where id = $1", id)
	p := Plane{}
	err := row.Scan(&p.id, &p.departure_city, &p.arrival_city)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(p.id, p.departure_city, p.arrival_city)

}

//Add is a function for adding new row to table
func Add(departure_city string, arrival_city string) {

	result, err := database.Db.Exec("insert into planes (id,departure_city,arrival_city) values ($1, $2, $3)", model.GenerateID(), departure_city, arrival_city)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result.RowsAffected())

}

// Update is a function for updating number of seats in train using id
func Update(arrival_city string, id uuid.UUID) {

	result, err := database.Db.Exec("update Planes set arrival_city = $1 where id = $2", arrival_city, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result.RowsAffected())

}

// Delete is a function for deleting row using id
func Delete(id uuid.UUID) {

	result, err := database.Db.Exec("delete from Planes where id = $1", id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result.RowsAffected())
}
