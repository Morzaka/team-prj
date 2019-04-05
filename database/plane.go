package database

import (
	"github.com/google/uuid"
	"team-project/database"

	"team-project/services/data"
)

var (
	insertPlane = `INSERT INTO public.plane (id,departureCity,arrivalCity)
	VALUES ($1, $2, $3) returning id`
	selectPlane = `SELECT password FROM public.user WHERE login=$1;`
	updatePlane = `UPDATE public.plane SET name = $2, surname = $3, login=$4, password=$5, role=$6 WHERE id = $1;`
	deletePlane = `DELETE FROM public.plane WHERE id = $1;`
)

func AddPlane(plane data.Plane) (data.Plane, error) {
	//insert values to the database
	_, err := Db.Exec(insertPlane, plane.)
	if err != nil {
		return data.Plane{}, err
	}
	return plane, nil
}

// DeletePlane is a function for deleting row using id
func DeletePlane(id uuid.UUID) error {

	_, err := database.Db.Exec(deletePlane, id)
	if err != nil {
		return err
	}
	return nil
}
