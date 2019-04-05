package booking

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"team-project/database"
	"team-project/logger"
	"team-project/services/data"
	"team-project/services/model"
)

func validateForm(tk data.Ticket) error {
	if tk.Place == 0 || tk.TicketType == "" || tk.
		Discount == "" || tk.Price == 0 || tk.TotalPrice == 0 || tk.
		Name == "" || tk.Surname == "" {
		return errors.New("all fields must be complete")
	}
	return nil
}

//GetAllTickets for GETing information about all tickets
func GetAllTickets(w http.ResponseWriter, r *http.Request) {
	tkts, err := database.GetAllTickets()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	model.JSONEncoding(w, tkts)
}

//GetOneTicket for GETing information about one tickets
func GetOneTicket(w http.ResponseWriter, r *http.Request) {
	id, err := model.GetID(r)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
	}

	tk, err := database.GetTicket(id)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	model.JSONEncoding(w, tk)
}

//CreateTicket (POST) for creating one tickets & add to DB
func CreateTicket(w http.ResponseWriter, r *http.Request) {
	// get values from client (json format)
	tk := data.Ticket{}
	tk.ID = model.GenerateID()
	err := json.NewDecoder(r.Body).Decode(&tk)
	if err != nil {
		logger.Logger.Errorf("Error, %s", err)
	}

	// validate form values
	err = validateForm(tk)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		w.Write([]byte(string(err.Error())))
		return
	}

	//validate real number values
	if tk.Price != float32(tk.Price) || tk.TotalPrice != float32(tk.TotalPrice) {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		_, err = w.Write([]byte(string("Not Acceptable. Price must be a number.")))
	}

	//insert data to DB 'ticket' table
	err = database.CreateTicket(tk)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		logger.Logger.Errorf("Error, %s", err)
		return
	}
}

//UpdateTicket (PATCH) for updating one tickets in DB
func UpdateTicket(w http.ResponseWriter, r *http.Request) {
	// get values from client (json format)
	id, err := model.GetID(r)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
	}
	tk := data.Ticket{}
	err = json.NewDecoder(r.Body).Decode(&tk)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
	}

	tk.ID = id

	// validate form values
	err = validateForm(tk)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		w.Write([]byte(string(err.Error())))
		return
	}

	//validate real number values
	if tk.Price != float32(tk.Price) || tk.TotalPrice != float32(tk.TotalPrice) {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		_, err = w.Write([]byte(string("Not Acceptable. Price must be a number.")))
	}

	err = database.UpdateTicket(tk)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		logger.Logger.Errorf("Error, %s", err)
		return
	}
}

//DeleteTicket (DELETE) for deleting one tickets in DB by id
func DeleteTicket(w http.ResponseWriter, r *http.Request) {
	id, err := model.GetID(r)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
	}
	err = database.DeleteTicket(id)
	if err != nil {
		logger.Logger.Errorf("Error, %s", err)
	}
}
