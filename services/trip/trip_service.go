package trip

import (
	"encoding/json"
	"errors"
	"github.com/go-zoo/bone"
	"github.com/google/uuid"
	"net/http"
	"team-project/database"
	"team-project/services/common"
	"team-project/services/data"
)

//emptyResponse variable will be sent in response
var emptyResponse interface{}

//ValidateFormTrip function for data validation
func ValidateFormTrip(trip data.Trip) error {
	if trip.TripName == "" || trip.TotalTripPrice <= 0 || trip.TripName != string(trip.TripName) || trip.
		TotalTripPrice != float32(trip.TotalTripPrice) {
		return errors.New("fields should be not empty or invalid data type of field")
	}
	return nil
}

//GetTrips function return all trips from trip table
func GetTrips(w http.ResponseWriter, r *http.Request) {

	trips, err := database.TripRepo.GetTrips()
	if err != nil {
		common.RenderJSON(w, r, http.StatusInternalServerError, emptyResponse)
		return
	}
	common.RenderJSON(w, r, http.StatusOK, trips)

}

//GetTrip function return trip from trip table request specified id
func GetTrip(w http.ResponseWriter, r *http.Request) {
	id := uuid.Must(uuid.Parse(bone.GetValue(r, "id")))
	trip, err := database.TripRepo.GetTrip(id)
	if err != nil {
		common.RenderJSON(w, r, http.StatusInternalServerError, emptyResponse)
		return
	}
	common.RenderJSON(w, r, http.StatusOK, trip)
}

//CreateTrip function add trip to trip table and generate uuid for new object
func CreateTrip(w http.ResponseWriter, r *http.Request) {
	trip := data.Trip{}
	trip.TripID = uuid.New()
	err := json.NewDecoder(r.Body).Decode(&trip)
	if err != nil {
		common.RenderJSON(w, r, http.StatusInternalServerError, emptyResponse)
		return
	}

	err = ValidateFormTrip(trip)
	if err != nil {
		common.RenderJSON(w, r, http.StatusInternalServerError, emptyResponse)
		return
	}

	_, err = database.TripRepo.AddTrip(trip)
	if err != nil {
		common.RenderJSON(w, r, http.StatusInternalServerError, emptyResponse)
		return
	}
	common.RenderJSON(w, r, http.StatusOK, trip)
}

//UpdateTrip function update trip data using id from request
func UpdateTrip(w http.ResponseWriter, r *http.Request) {
	id := uuid.Must(uuid.Parse(bone.GetValue(r, "id")))
	trip := data.Trip{}
	trip.TripID = id
	_ = json.NewDecoder(r.Body).Decode(&trip)

	err := ValidateFormTrip(trip)
	if err != nil {
		common.RenderJSON(w, r, http.StatusInternalServerError, emptyResponse)
		return
	}
	_, err = database.TripRepo.UpdateTrip(trip)
	if err != nil {
		common.RenderJSON(w, r, http.StatusInternalServerError, emptyResponse)
		return
	}
	common.RenderJSON(w, r, http.StatusOK, trip)
}

//DeleteTrip function delete data from trip table using id from request
func DeleteTrip(w http.ResponseWriter, r *http.Request) {
	id := uuid.Must(uuid.Parse(bone.GetValue(r, "id")))
	err := database.TripRepo.DeleteTrip(id)
	if err != nil {
		common.RenderJSON(w, r, http.StatusInternalServerError, emptyResponse)
		return
	}
	common.RenderJSON(w, r, http.StatusOK, emptyResponse)
}
