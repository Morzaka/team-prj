package common

import (
	"encoding/json"
	"net/http"

	"team-project/logger"
)

//RenderJSON render json data to user
func RenderJSON(w http.ResponseWriter, r *http.Request, status int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if response == nil {
		return
	}
	data, err := json.Marshal(response)
	if err != nil {
		logger.Logger.Errorf("Error, %s", err)
	}
	_, err = w.Write(data)
	if err != nil {
		logger.Logger.Errorf("Error, %s", err)
	}
}

//REnderJSON render json data to user
func REnderJSON(w http.ResponseWriter, r *http.Request, status int, response interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if response == nil {
		return
	}
	data, err := json.Marshal(response)
	if err != nil {
		logger.Logger.Errorf("Error, %s", err)
	}
	_, err = w.Write(data)
	if err != nil {
		logger.Logger.Errorf("Error, %s", err)
	}
}

// Error contains the message about error
type Error struct {
	Message string `json:"message"`
}

// ErrorMessage contains the error
type ErrorMessage struct {
	Error Error `json:"error"`
}

// SendMethodNotAllowed sends Internal Server Error Completed and logs an error if it exists
func SendMethodNotAllowed(w http.ResponseWriter, r *http.Request, message string, err error) {
	SendError(w, r, http.StatusMethodNotAllowed, message, err)
}

//SendUnsupportedMediaType sends Internal Server Error Completed and logs an error if it exists
func SendUnsupportedMediaType(w http.ResponseWriter, r *http.Request, message string, err error) {
	SendError(w, r, http.StatusUnsupportedMediaType, message, err)
}

//SendBadRequest sends Internal Server Error Completed and logs an error if it exists
func SendBadRequest(w http.ResponseWriter, r *http.Request, message string, err error) {
	SendError(w, r, http.StatusBadRequest, message, err)
}

//SendNotFound sends Internal Server Error Completed and logs an error if it exists
func SendNotFound(w http.ResponseWriter, r *http.Request, message string, err error) {
	SendError(w, r, http.StatusNotFound, message, err)
}

// SendInternalServerError sends Internal Server Error Completed and logs an error if it exists
func SendInternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	SendError(w, r, http.StatusInternalServerError, "", err)
}

// SendError writes a defined string as an error message
// with appropriate headers to the HTTP response
func SendError(w http.ResponseWriter, r *http.Request, status int, message string, errMessage error) {
	var data []byte
	var err error

	if message != "" {
		data, err = json.Marshal(ErrorMessage{Error{message}})
		if err != nil {
			SendInternalServerError(w, r, err)
			return
		}
	}

	if errMessage != nil {
		logger.Logger.Printf(`"%s %s" err: %s`, r.Method, r.URL, errMessage)
	}
	REnderJSON(w, r, status, data)
}

//RENDERJSON is used for rendering JSON response body with appropriate headers
func RENDERJSON(w http.ResponseWriter, r *http.Request, response interface{}) {
	switch r.Method {
	case http.MethodPost:
		REnderJSON(w, r, http.StatusCreated, response)
	case http.MethodDelete:
		REnderJSON(w, r, http.StatusNoContent, response)
	case http.MethodGet:
		REnderJSON(w, r, http.StatusOK, response)
	}
}
