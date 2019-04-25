package train

import "C"
import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"

	"team-project/database"
	"team-project/services/authorization"
	"team-project/services/common"
	"team-project/services/data"

	"github.com/go-zoo/bone"
	"github.com/google/uuid"
)

var (
	emptyResponse interface{}
)

//ValidateIfEmpty is a validation if train is empty
func ValidateIfEmpty(t data.Train) error {
	if t.DepartureCity == "" || t.ArrivalCity == "" || t.DepartureTime == "" || t.DepartureDate == "" || t.ArrivalTime == "" || t.ArrivalDate == "" {
		return errors.New("Some incoming data is empty =(")
	}
	return nil
}

//NameIsValid is a validation if trains name data is valid
func NameIsValid(str string) bool {
	var validName = regexp.MustCompile(`^[a-zA-Z]+(?:[\s-][a-zA-Z]+)*$`)
	return validName.MatchString(str)
}

//TimeIsValid is a validation if trains time data is valid
func TimeIsValid(str string) bool {
	var validName = regexp.MustCompile(`^(2[0-3]|[01]?[0-9]):([0-5]?[0-9])$`)
	return validName.MatchString(str)
}

//DateIsValid is a validation if trains date data is valid
func DateIsValid(str string) bool {
	var validName = regexp.MustCompile(`^\s*(3[01]|[12][0-9]|0?[1-9])\.(1[012]|0?[1-9])\.((?:19|20)\d{2})\s*$`)
	return validName.MatchString(str)
}

//Validate is a function that validates trains name, date, time data
func Validate(t data.Train) error {
	if NameIsValid(t.ArrivalCity) == false || NameIsValid(t.DepartureCity) == false || DateIsValid(t.DepartureDate) == false || DateIsValid(t.ArrivalDate) == false || TimeIsValid(t.ArrivalTime) == false || TimeIsValid(t.DepartureTime) == false {
		return errors.New("Some data is invalid")
	}
	return nil
}

//GetTrains is a handler that returns trains from db
func GetTrains(w http.ResponseWriter, r *http.Request) {
	if !authorization.AdminRole(w, r) {
		common.RenderJSON(w, r, http.StatusForbidden, emptyResponse)
		return
	}
	trains, err := database.Trains.GetAllTrains()
	if err != nil {
		common.RenderJSON(w, r, http.StatusInternalServerError, emptyResponse)
		return
	}

	for _, train := range trains {
		if err = ValidateIfEmpty(train); err != nil {
			common.RenderJSON(w, r, http.StatusNoContent, "Empty data found")
			return
		}
		if err = Validate(train); err != nil {
			common.RenderJSON(w, r, http.StatusNoContent, "Validation failed")
			return
		}
	}

	common.RenderJSON(w, r, http.StatusOK, trains)
}

//GetSingleTrain is a handler that returns single train from db
func GetSingleTrain(w http.ResponseWriter, r *http.Request) {
	if !authorization.AdminRole(w, r) {
		common.RenderJSON(w, r, http.StatusForbidden, emptyResponse)
		return
	}
	id := uuid.Must(uuid.Parse(bone.GetValue(r, "id")))
	newid := id.String()
	train, err := database.Trains.GetTrain(newid)
	if err != nil {
		common.RenderJSON(w, r, http.StatusNoContent, emptyResponse)
		return
	}
	if err = ValidateIfEmpty(train); err != nil {
		common.RenderJSON(w, r, http.StatusNoContent, "Empty data found")
		return
	}
	if err = Validate(train); err != nil {
		common.RenderJSON(w, r, http.StatusNoContent, "Validation failed")
		return
	}
	common.RenderJSON(w, r, http.StatusOK, train)
}

//CreateTrain is a handler that creates train
func CreateTrain(w http.ResponseWriter, r *http.Request) {
	if !authorization.AdminRole(w, r) {
		common.RenderJSON(w, r, http.StatusForbidden, emptyResponse)
		return
	}
	t := data.Train{}
	json.NewDecoder(r.Body).Decode(&t)
	if err := ValidateIfEmpty(t); err != nil {
		common.RenderJSON(w, r, http.StatusNoContent, "Empty data found")
		return
	}
	if err := Validate(t); err != nil {
		common.RenderJSON(w, r, http.StatusNoContent, "Validation failed")
		return
	}
	err := database.Trains.AddTrain(t)
	if err != nil {
		common.RenderJSON(w, r, http.StatusInternalServerError, "Error occured while adding train to database")
		return
	}
	common.RenderJSON(w, r, http.StatusOK, t)
}

//UpdateTrain is a handler that updates train in db
func UpdateTrain(w http.ResponseWriter, r *http.Request) {
	if !authorization.AdminRole(w, r) {
		common.RenderJSON(w, r, http.StatusForbidden, emptyResponse)
		return
	}
	id := uuid.Must(uuid.Parse(bone.GetValue(r, "id")))
	t := data.Train{}
	json.NewDecoder(r.Body).Decode(&t)
	if err := ValidateIfEmpty(t); err != nil {
		common.RenderJSON(w, r, http.StatusNoContent, "Empty data found")
		return
	}
	if err := Validate(t); err != nil {
		common.RenderJSON(w, r, http.StatusNoContent, "Validation failed")
		return
	}
	t.ID = id
	err := database.Trains.UpdateTrain(t)
	if err != nil {
		common.RenderJSON(w, r, http.StatusInternalServerError, "Couldn't update data")
		return
	}
	train, err := database.Trains.GetTrain(t.ID.String())
	if err != nil {
		common.RenderJSON(w, r, http.StatusNoContent, "Couldn't return updated data")
		return
	}
	common.RenderJSON(w, r, http.StatusOK, train)
}

//DeleteTrain is a handler that deletes train from db
func DeleteTrain(w http.ResponseWriter, r *http.Request) {
	if !authorization.AdminRole(w, r) {
		common.RenderJSON(w, r, http.StatusForbidden, emptyResponse)
		return
	}
	id := uuid.Must(uuid.Parse(bone.GetValue(r, "id")))
	err := database.Trains.DeleteTrain(id)
	if err != nil {
		common.RenderJSON(w, r, http.StatusInternalServerError, "Error occured while deleting train from database")
		return
	}
	common.RenderJSON(w, r, http.StatusOK, "Train was successfully updated")
}
