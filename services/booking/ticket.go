package booking

import (
	"encoding/json"
	"errors"
	"net/http"

	"team-project/database"
	"team-project/services/common"
	"team-project/services/data"
	"team-project/services/model"
)

var emptyResponse interface{}

func validateForm(tk data.Ticket) error {
	if tk.Place <= 0 || tk.TicketType == "" || tk.
		Discount == "" || tk.Price <= 0 || tk.TotalPrice <= 0 || tk.
		Name == "" || tk.Surname == "" {
		return errors.New("all fields must be complete")
	}
	return nil
}

//GetAllTickets for GETing information about all tickets
func GetAllTickets(w http.ResponseWriter, r *http.Request) {
	tkts, err := database.TicketRepo.AllTickets()
	if err != nil {
		common.RenderJSON(w, r, http.StatusInternalServerError, tkts)
		return
	}
	common.RenderJSON(w, r, http.StatusOK, tkts)
}

//GetOneTicket for GETing information about one tickets
func GetOneTicket(w http.ResponseWriter, r *http.Request) {
	id, err := model.GetID(r)
	if err != nil {
		common.RenderJSON(w, r, http.StatusBadRequest, emptyResponse)
		return
	}

	tk, err := database.TicketRepo.GetTicket(id)
	if err != nil {
		common.RenderJSON(w, r, http.StatusInternalServerError, tk)
		return
	}
	common.RenderJSON(w, r, http.StatusOK, tk)
}

//CreateTicket (POST) for creating one tickets & add to DB
func CreateTicket(w http.ResponseWriter, r *http.Request) {
	// get values from client (json format)
	tk := data.Ticket{}
	tk.ID = model.GenerateID()
	err := json.NewDecoder(r.Body).Decode(&tk)
	if err != nil {
		common.RenderJSON(w, r, http.StatusBadRequest, tk)
		return
	}

	// validate form values
	err = validateForm(tk)
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
	common.RenderJSON(w, r, http.StatusOK, tk)
}

//UpdateTicket (PATCH) for updating one tickets in DB
func UpdateTicket(w http.ResponseWriter, r *http.Request) {
	// get values from client (json format)
	id, err := model.GetID(r)
	if err != nil {
		common.RenderJSON(w, r, http.StatusBadRequest, emptyResponse)
		return
	}
	tk := data.Ticket{}
	err = json.NewDecoder(r.Body).Decode(&tk)
	if err != nil {
		common.RenderJSON(w, r, http.StatusBadRequest, tk)
		return
	}

	tk.ID = id

	// validate form values
	err = validateForm(tk)
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
	id, err := model.GetID(r)
	if err != nil {
		common.RenderJSON(w, r, http.StatusBadRequest, emptyResponse)
		return
	}
	err = database.TicketRepo.DeleteTicket(id)
	if err != nil {
		common.RenderJSON(w, r, http.StatusInternalServerError, emptyResponse)
		return
	}
	common.RenderJSON(w, r, http.StatusNotFound,
		"Ticket was deleted successfully.")
}
