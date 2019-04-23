package train

import (
	"encoding/json"
	"github.com/go-zoo/bone"
	"github.com/google/uuid"
	"net/http"
	"team-project/database"
	"team-project/services/common"
	"team-project/services/data"
)

var emptyResponse interface{}

//GetTrains is a method
func GetTrains(w http.ResponseWriter, r *http.Request) {
	trains, err := database.GetAllTrains()
	if err != nil {
		common.RenderJSON(w, r, 404, emptyResponse)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trains)
}

//GetSingleTrain is a method
func GetSingleTrain(w http.ResponseWriter, r *http.Request) {
	params := bone.GetAllValues(r)
	train, err := database.GetTrain(params["id"])
	if err != nil {
		common.RenderJSON(w, r, 404, emptyResponse)
		return
	}
	common.RenderJSON(w, r, 202, train)
}

//CreateTrain is a method
func CreateTrain(w http.ResponseWriter, r *http.Request) {
	t := data.Train{}
	json.NewDecoder(r.Body).Decode(&t)
	err := database.AddTrain(t)
	if err != nil {
		common.RenderJSON(w, r, 404, emptyResponse)
		return
	}
	common.RenderJSON(w, r, 202, t)
}

//UpdateTrain is a method
func UpdateTrain(w http.ResponseWriter, r *http.Request) {
	id := uuid.Must(uuid.Parse(bone.GetValue(r, "id")))
	t := data.Train{}
	json.NewDecoder(r.Body).Decode(&t)
	t.ID = id
	err := database.UpdateTrain(t.ID, t.DepartureCity, t.ArrivalCity,
		t.DepartureDate, t.DepartureTime, t.ArrivalTime, t.ArrivalDate)
	if err != nil {
		common.RenderJSON(w, r, 404, emptyResponse)
		return
	}
	train, err := database.GetTrain(t.ID.String())
	if err != nil {
		common.RenderJSON(w, r, 404, emptyResponse)
		return
	}
	common.RenderJSON(w, r, 202, train)
}

//DeleteTrain is a method
func DeleteTrain(w http.ResponseWriter, r *http.Request) {
	id := uuid.Must(uuid.Parse(bone.GetValue(r, "id")))
	err := database.DeleteTrain(id)
	if err != nil {
		common.RenderJSON(w, r, 404, emptyResponse)
		return
	}
	common.RenderJSON(w, r, 202, emptyResponse)
}
