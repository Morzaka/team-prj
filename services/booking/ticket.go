package booking

import (
	"encoding/json"
	"errors"
	"net/http"
	"team-project/services/authorization"

	"team-project/database"
	"team-project/services/common"
	"team-project/services/data"

	"github.com/go-zoo/bone"
	"github.com/google/uuid"
)

var emptyResponse interface{}

//ValidateForm function validate incoming data from client
func ValidateForm(tk data.Ticket) error {
	if tk.Place <= 0 || tk.TicketType == "" || tk.
		Discount == "" || tk.Price <= 0 || tk.TotalPrice <= 0 || tk.
		Name == "" || tk.Surname == "" {
		return errors.New("all fields must be complete")
	}
	return nil
}

//GetAllTickets for GETing information about all tickets
func GetAllTickets(w http.ResponseWriter, r *http.Request) {
	if !authorization.AdminRole(w, r) {
		common.RenderJSON(w, r, http.StatusForbidden, emptyResponse)
		return
	}
	tkts, err := database.TicketRepo.AllTickets()
	if err != nil {
		common.RenderJSON(w, r, http.StatusInternalServerError, tkts)
		return
	}
	common.RenderJSON(w, r, http.StatusOK, tkts)
}

//GetOneTicket for GETing information about one tickets
func GetOneTicket(w http.ResponseWriter, r *http.Request) {
	if !authorization.LoggedIn(w, r) {
		common.RenderJSON(w, r, http.StatusForbidden, emptyResponse)
		return
	}
	id := uuid.Must(uuid.Parse(bone.GetValue(r, "id")))
	tk, err := database.TicketRepo.GetTicket(id)
	if err != nil {
		common.RenderJSON(w, r, http.StatusInternalServerError, tk)
		return
	}
	common.RenderJSON(w, r, http.StatusOK, tk)
}

//CreateTicket (POST) for creating one tickets & add to DB
func CreateTicket(w http.ResponseWriter, r *http.Request) {
	if !authorization.AdminRole(w, r) {
		common.RenderJSON(w, r, http.StatusForbidden, emptyResponse)
		return
	}
	tk := data.Ticket{}
	tk.ID = uuid.New()
	err := json.NewDecoder(r.Body).Decode(&tk)
	if err != nil {
		common.RenderJSON(w, r, http.StatusBadRequest, tk)
		return
	}

	err = ValidateForm(tk)
	if err != nil {
		common.RenderJSON(w, r, http.StatusNotAcceptable, err.Error())
		return
	}

	//validate real number values
	if tk.Price != float32(tk.Price) || tk.TotalPrice != float32(tk.TotalPrice) {
		common.RenderJSON(w, r, http.StatusNotAcceptable, "Not Acceptable. Price must be a number.")
		return
	}

	//insert data to DB 'ticket' table
	err = database.TicketRepo.CreateTicket(tk)
	if err != nil {
		common.RenderJSON(w, r, http.StatusInternalServerError, tk)
		return
	}
	common.RenderJSON(w, r, http.StatusCreated, tk)
}

//UpdateTicket (PATCH) for updating one tickets in DB
func UpdateTicket(w http.ResponseWriter, r *http.Request) {
	if !authorization.AdminRole(w, r) {
		common.RenderJSON(w, r, http.StatusForbidden, emptyResponse)
		return
	}
	// get ID value from client (json format)
	id := uuid.Must(uuid.Parse(bone.GetValue(r, "id")))
	tk := data.Ticket{}
	err := json.NewDecoder(r.Body).Decode(&tk)
	if err != nil {
		common.RenderJSON(w, r, http.StatusBadRequest, tk)
		return
	}

	tk.ID = id
	err = ValidateForm(tk)
	if err != nil {
		common.RenderJSON(w, r, http.StatusNotAcceptable, err.Error())
		return
	}

	//validate real number values
	if tk.Price != float32(tk.Price) || tk.TotalPrice != float32(tk.TotalPrice) {
		common.RenderJSON(w, r, http.StatusNotAcceptable, "Not Acceptable. Price must be a number.")
		return
	}

	//insert data to DB 'ticket' table
	err = database.TicketRepo.UpdateTicket(tk)
	if err != nil {
		common.RenderJSON(w, r, http.StatusInternalServerError, tk)
		return
	}
	common.RenderJSON(w, r, http.StatusOK, tk)
}

//DeleteTicket (DELETE) for deleting one tickets in DB by id
func DeleteTicket(w http.ResponseWriter, r *http.Request) {
	if !authorization.AdminRole(w, r) {
		common.RenderJSON(w, r, http.StatusForbidden, emptyResponse)
		return
	}
	id := uuid.Must(uuid.Parse(bone.GetValue(r, "id")))
	err := database.TicketRepo.DeleteTicket(id)
	if err != nil {
		common.RenderJSON(w, r, http.StatusInternalServerError, emptyResponse)
		return
	}
	common.RenderJSON(w, r, http.StatusOK,
		"Ticket was deleted successfully.")
}
