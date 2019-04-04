package train

import (
	"encoding/json"
	"net/http"
	"team-project/services/data"

	"github.com/go-zoo/bone"
)

//GetTrains is a method
func GetTrains(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	trains := GetAllTrains()
	json.NewEncoder(w).Encode(trains)
}

//GetSingleTrain is a method
func GetSingleTrain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := bone.GetAllValues(r)
	train, err := GetTrain(params["id"])
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(train)
}

//CreateTrain is a method
func CreateTrain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	t := data.Train{}
	json.NewDecoder(r.Body).Decode(&t)
	//previous := db.Query
}
