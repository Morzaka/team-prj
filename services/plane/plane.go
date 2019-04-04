package plane

import (
	"fmt"
	"github.com/google/uuid"
	"team-project/database"
	"team-project/services/model"
)

// Plane struct contains plane data
type Plane struct {
	id             uuid.UUID
	departureCity string
	arrivalCity   string
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
		err := rows.Scan(&p.id, &p.departureCity, &p.arrivalCity)
		if err != nil {
			fmt.Println(err)
			continue
		}
		planes = append(planes, p)
	}
	for _, p := range planes {
		fmt.Println(p.id, p.departureCity, p.arrivalCity)
	}
}

// GetID is a function for getting row using id
func GetID(id uuid.UUID) {

	row := database.Db.QueryRow("select * from Planes where id = $1", id)
	p := Plane{}
	err := row.Scan(&p.id, &p.departureCity, &p.arrivalCity)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(p.id, p.departureCity, p.arrivalCity)

}

//Add is a function for adding new row to table
func Add(departureCity string, arrivalCity string) {

	result, err := database.Db.Exec("insert into planes (id,departure_city,arrival_city) values ($1, $2, $3)", model.GenerateID(), departureCity, arrivalCity)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result.RowsAffected())

}

// Update is a function for updating number of seats in train using id
func Update(arrivalCity string, id uuid.UUID) {

	result, err := database.Db.Exec("update Planes set arrival_city = $1 where id = $2", arrivalCity, id)
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
