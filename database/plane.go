package database

import (
	"github.com/google/uuid"

	"team-project/services/data"
)

var (
	selectPlanes = `SELECT * FROM public.plane;`
	selectPlane  = `SELECT * FROM public.plane WHERE id=$1;`
	insertPlane  = `INSERT INTO public.plane (id, departure_City, arrival_City) VALUES ($1, $2, $3)`
	updatePlane  = `UPDATE public.plane SET departure_City = $2, arrival_City = $3 WHERE id = $1;`
	deletePlane  = `DELETE FROM public.plane WHERE id = $1;`
)

// GetPlanes is a function for getting all Planes from table
func GetPlanes() ([]data.Plane, error) {
	rows, err := Db.Query(selectPlanes)
	if err != nil {
		return []data.Plane{}, err
	}
	defer func() error {
		err := rows.Close()
		if err != nil {
			return err
		}
		return nil
	}()
	planes := []data.Plane{}
	for rows.Next() {
		p := data.Plane{}
		err := rows.Scan(&p.ID, &p.DepartureCity, &p.ArrivalCity)
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
	row := Db.QueryRow(selectPlane, id)
	err := row.Scan(&p.ID, &p.DepartureCity, &p.ArrivalCity)
	if err != nil {
		return p, err
	}
	return p, nil
}

// UpdatePlane is a function for updating Plane using id
func UpdatePlane(plane data.Plane, id uuid.UUID) (data.Plane, error) {
	_, err := Db.Exec(updatePlane, id, plane.DepartureCity, plane.ArrivalCity)
	if err != nil {
		return data.Plane{}, err
	}
	return plane, nil
}

// AddPlane is a function for adding new Plane to table
func AddPlane(plane data.Plane) (data.Plane, error) {
	_, err := Db.Exec(insertPlane, plane.ID, plane.DepartureCity, plane.ArrivalCity)
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
