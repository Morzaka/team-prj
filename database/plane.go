package database

import (
	"github.com/google/uuid"

	"team-project/services/data"
)

var (
	selectPlanes = `SELECT * FROM public.plane;`
	selectPlane  = `SELECT * FROM public.plane WHERE login=$1;`
	insertPlane  = `INSERT INTO public.plane (id,departureCity,arrivalCity) VALUES ($1, $2, $3)`
	updatePlane  = `UPDATE public.plane SET departureCity = $2, arrivalCity = $3 WHERE id = $1;`
	deletePlane  = `DELETE FROM public.plane WHERE id = $1;`
)

// GetPlanes is a function for getting all Planes from table
func GetPlanes() ([]data.Plane, error) {
	rows, err := Db.Query(selectPlanes)
	if err != nil {
		return []data.Plane{}, err
	}
	defer rows.Close()
	planes := []data.Plane{}
	for rows.Next() {
		p := data.Plane{}
		err := rows.Scan(&p.Id, &p.DepartureCity, &p.ArrivalCity)
		if err != nil {
			return []data.Plane{}, err
		}
		planes = append(planes, p)
	}
	return planes, nil
}

// GetPlane is a function for getting Plane using id
func GetPlane(id uuid.UUID) (data.Plane, error) {
	p := data.Plane{}
	err := Db.QueryRow(selectPlane, id).Scan(&p.Id, &p.DepartureCity, &p.ArrivalCity)
	if err != nil {
		return data.Plane{}, err
	}
	return p, nil
}

//Update is a function for updating Plane using id
func UpdatePlane(plane data.Plane, id uuid.UUID) (data.Plane, error) {
	_, err := Db.Exec(updatePlane, id, plane.DepartureCity, plane.ArrivalCity)
	if err != nil {
		return data.Plane{}, err
	}
	return plane, nil
}

// AddPlane is a function for adding new Plane to table
func AddPlane(plane data.Plane) (data.Plane, error) {
	_, err := Db.Exec(insertPlane, plane.Id, plane.DepartureCity, plane.ArrivalCity)
	if err != nil {
		return data.Plane{}, err
	}
	return plane, nil
}

// DeletePlane is a function for deleting Plane using id
func DeletePlane(id uuid.UUID) error {
	_, err := Db.Exec(deletePlane, id)
	if err != nil {
		return err
	}
	return nil
}
