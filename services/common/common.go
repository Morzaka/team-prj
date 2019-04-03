package common

import (
	"encoding/json"
	"net/http"

	"team-project/logger"
)

//RenderJSON render json data to user
func RenderJSON(w http.ResponseWriter, r *http.Request, status int, response interface{}) {
	data, err := json.Marshal(response)
	if err != nil {
		logger.Logger.Errorf("Error, %s", err)
	}
	w.WriteHeader(status)
	if data == nil {
		return
	}
	w.Header().Set("content-type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		logger.Logger.Errorf("Error, %s", err)
	}
}
