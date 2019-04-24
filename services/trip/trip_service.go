package trip

import (
	"encoding/json"
	"errors"
	"net/http"
	"team-project/database"
	"team-project/services/common"
	"team-project/services/data"
	"team-project/services/model"
)

//emptyResponse variable will be sent in response
var emptyResponse interface{}

//ValidateForm function for data validation
func ValidateForm(trip data.Trip) error {
	if trip.TripName == "" || trip.TotalTripPrice == 0 || trip.TripName != string(trip.TripName) || trip.
		TotalTripPrice != float32(trip.TotalTripPrice) {
		return errors.New("Fields should be not empty or invalid data type of field.")
	}
	return nil
}

//GetTrips function return all trips from trip table
func GetTrips(w http.ResponseWriter, r *http.Request) {

	trips, err := database.GetTrips()
	if err != nil {
		common.RenderJSON(w, r, 404, emptyResponse)
		return
	}
	common.RenderJSON(w, r, 202, trips)

}

//GetTrip function return trip from trip table request specified id
func GetTrip(w http.ResponseWriter, r *http.Request) {
	id, err := model.GetID(r)
	trip, err := database.GetTrip(id)
	if err != nil {
		common.RenderJSON(w, r, 404, emptyResponse)
		return
	}
	common.RenderJSON(w, r, 202, trip)
}

//CreateTrip function add trip to trip table and generate uuid for new object
func CreateTrip(w http.ResponseWriter, r *http.Request) {
	trip := data.Trip{}
	err := json.NewDecoder(r.Body).Decode(&trip)
	if err != nil{
		common.RenderJSON(w, r, 404, emptyResponse)
		return
	}
	trip.TripID = model.GenerateID()

	err = ValidateForm(trip)
	if err != nil {
		common.RenderJSON(w, r, 404, emptyResponse)
		return
	}

	_, err = database.AddTrip(trip)
	if err != nil {
		common.RenderJSON(w, r, 404, emptyResponse)
		return
	}
	common.RenderJSON(w, r, 202, trip)
}

//UpdateTrip function update trip data using id from request
func UpdateTrip(w http.ResponseWriter, r *http.Request) {
	id, err := model.GetID(r)
	if err != nil {
		common.RenderJSON(w, r, 404, emptyResponse)
		return
	}
	trip := data.Trip{}
	_ = json.NewDecoder(r.Body).Decode(&trip)
	trip.TripID = id

	err = ValidateForm(trip)
	if err != nil {
		common.RenderJSON(w, r, 404, emptyResponse)
		return
	}
	_, err = database.UpdateTrip(trip)
	if err != nil {
		common.RenderJSON(w, r, 404, emptyResponse)
		return
	}
	common.RenderJSON(w, r, 202, trip)
}

//DeleteTrip function delete data from trip table using id from request
func DeleteTrip(w http.ResponseWriter, r *http.Request) {
	id, err := model.GetID(r)
	if err != nil {
		common.RenderJSON(w, r, 404, emptyResponse)
		return
	}

	err = database.DeleteTrip(id)
	if err != nil {
		common.RenderJSON(w, r, 404, emptyResponse)
		return
	}
	common.RenderJSON(w, r, 202, emptyResponse)
}
