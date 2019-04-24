package train

import "C"
import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"

	"team-project/database"
	"team-project/services/authorization"
	"team-project/services/common"
	"team-project/services/data"

	"github.com/go-zoo/bone"
)

var (
	emptyResponse interface{}
)

//GetTrains is a method
func GetTrains(w http.ResponseWriter, r *http.Request) {
	if !authorization.CheckAdmin(w, r) {
		common.RenderJSON(w, r, http.StatusForbidden, emptyResponse)
		return
	}
	trains, err := database.GetAllTrains()
	if err != nil {
		common.RenderJSON(w, r, http.StatusNotFound, emptyResponse)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trains)
}

//GetSingleTrain is a method
func GetSingleTrain(w http.ResponseWriter, r *http.Request) {
	if !authorization.CheckAdmin(w, r) {
		common.RenderJSON(w, r, http.StatusForbidden, emptyResponse)
		return
	}
	params := bone.GetAllValues(r)
	train, err := database.GetTrain(params["id"])
	if err != nil {
		common.RenderJSON(w, r, http.StatusNotFound, emptyResponse)
		return
	}
	common.RenderJSON(w, r, http.StatusOK, train)
}

//CreateTrain is a method
func CreateTrain(w http.ResponseWriter, r *http.Request) {
	if !authorization.CheckAdmin(w, r) {
		common.RenderJSON(w, r, http.StatusForbidden, emptyResponse)
		return
	}
	t := data.Train{}
	json.NewDecoder(r.Body).Decode(&t)
	err := database.AddTrain(t)
	if err != nil {
		common.RenderJSON(w, r, http.StatusNotFound, emptyResponse)
		return
	}
	common.RenderJSON(w, r, http.StatusOK, t)
}

//UpdateTrain is a method
func UpdateTrain(w http.ResponseWriter, r *http.Request) {
	if !authorization.CheckAdmin(w, r) {
		common.RenderJSON(w, r, http.StatusForbidden, emptyResponse)
		return
	}
	id := uuid.Must(uuid.Parse(bone.GetValue(r, "id")))
	t := data.Train{}
	json.NewDecoder(r.Body).Decode(&t)
	t.ID = id
	err := database.UpdateTrain(t.ID, t.DepartureCity, t.ArrivalCity, t.DepartureDate, t.DepartureTime, t.ArrivalTime, t.ArrivalDate)
	if err != nil {
		common.RenderJSON(w, r, http.StatusNotFound, emptyResponse)
		return
	}
	train, err := database.GetTrain(t.ID.String())
	if err != nil {
		common.RenderJSON(w, r, http.StatusNotFound, emptyResponse)
		return
	}
	common.RenderJSON(w, r, http.StatusOK, train)
}

//DeleteTrain is a method
func DeleteTrain(w http.ResponseWriter, r *http.Request) {
	if !authorization.CheckAdmin(w, r) {
		common.RenderJSON(w, r, http.StatusForbidden, emptyResponse)
		return
	}
	id := uuid.Must(uuid.Parse(bone.GetValue(r, "id")))
	err := database.DeleteTrain(id)
	if err != nil {
		common.RenderJSON(w, r, http.StatusNotFound, emptyResponse)
		return
	}
	common.RenderJSON(w, r, http.StatusOK, emptyResponse)
}
