package database

import (
	"github.com/google/uuid"

	"team-project/logger"
	"team-project/services/data"
)

// PlaneRepository for mocking
type PlaneRepository interface {
	GetPlanes() ([]data.Plane, error)
	GetPlane(id uuid.UUID) (data.Plane, error)
	AddPlane(plane data.Plane) (data.Plane, error)
	UpdatePlane(plane data.Plane, id uuid.UUID) (data.Plane, error)
	DeletePlane(id uuid.UUID) error
}

// planeRepository structure contains interface PlaneRepository
type planeRepository struct {
	planeRepo PlaneRepository
}

// PlaneRepo is a variable
var PlaneRepo PlaneRepository = &planeRepository{}

var (
	selectPlanes = `SELECT * FROM public.plane;`
	selectPlane  = `SELECT * FROM public.plane WHERE id=$1;`
	insertPlane  = `INSERT INTO public.plane (id, departureCity, arrivalCity) VALUES ($1, $2, $3)`
	updatePlane  = `UPDATE public.plane SET departureCity = $2, arrivalCity = $3 WHERE id = $1;`
	deletePlane  = `DELETE FROM public.plane WHERE id = $1;`
)

// GetPlanes is a function for getting all Planes from table
func (*planeRepository) GetPlanes() ([]data.Plane, error) {
	rows, err := Db.Query(selectPlanes)
	if err != nil {
		return []data.Plane{}, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			logger.Logger.Errorf("Error, %s", err)
		}
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
func (*planeRepository) GetPlane(id uuid.UUID) (data.Plane, error) {
	p := data.Plane{}
	row := Db.QueryRow(selectPlane, id)
	err := row.Scan(&p.ID, &p.DepartureCity, &p.ArrivalCity)
	if err != nil {
		return p, err
	}
	return p, nil
}

// UpdatePlane is a function for updating Plane using id
func (*planeRepository) UpdatePlane(plane data.Plane, id uuid.UUID) (data.Plane, error) {
	_, err := Db.Exec(updatePlane, id, plane.DepartureCity, plane.ArrivalCity)
	if err != nil {
		return data.Plane{}, err
	}
	return plane, nil
}

// AddPlane is a function for adding new Plane to table
func (*planeRepository) AddPlane(plane data.Plane) (data.Plane, error) {
	_, err := Db.Exec(insertPlane, plane.ID, plane.DepartureCity, plane.ArrivalCity)
	if err != nil {
		return data.Plane{}, err
	}
	return plane, nil
}

// DeletePlane is a function for deleting Plane using id
func (*planeRepository) DeletePlane(id uuid.UUID) error {
	_, err := Db.Exec(deletePlane, id)
	if err != nil {
		return err
	}
	return nil
}
