package common

import (
	"encoding/json"
	"net/http"

	"team-project/logger"
)

//go:generate mockgen -destination=../../mocks/mock_common.go -package=mocks team-project/services/common Common

//Common inerface contains methods, provides mocking
type Common interface {
	RenderJSON(w http.ResponseWriter, r *http.Request, status int, response interface{})
}

//ICommon struct contains interface Common
type ICommon struct {
	CommonMethod Common
}

//RenderJSON render json data to user
func (*ICommon) RenderJSON(w http.ResponseWriter, r *http.Request, status int, response interface{}) {
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
