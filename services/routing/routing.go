package routing

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"team-project/configurations"
	"time"

	"team-project/logger"
	"team-project/services/common"
	"team-project/services/data"

	"github.com/thoas/go-funk"
)

//RouteStruct is a struct of Routes
type RouteStruct struct {
	Index  string   `json:"index"`
	Routes []string `json:"routes"`
}

// RouteStructs is a slice of Struct of Routes
type RouteStructs []RouteStruct

var (
	//RouteStorage is a map of stations id as key and stations as value
	RouteStorage map[string][]string
	initialised  = false
)

var myClient = &http.Client{Timeout: 10 * time.Second}

//gets railroad Data
func getData() map[string][]string {
	routeS := make(map[string][]string, 100)
	r, err := myClient.Get(configurations.Config.JSONApi)
	if err != nil {
		logger.Logger.Error("Error while geting responce from JSONAPI ")

	}
	respBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Logger.Error("Error while reading http request to JSONApi")
	}
	routes := new(RouteStructs)
	json.Unmarshal(respBytes, &routes)

	for _, v := range *routes {
		routeS[v.Index] = v.Routes
	}
	initialised = true
	return routeS
}

//FindPath generates routes
func FindPath(w http.ResponseWriter, r *http.Request) {

	if initialised == false {
		RouteStorage = getData()
	}

	var stations data.Stations
	var pathResult []data.Routes
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		logger.Logger.Errorf("Error, %s", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	err = json.Unmarshal(b, &stations)
	if err != nil {
		logger.Logger.Errorf("Error, %s", err)
	}
	for key, value := range RouteStorage {
		if indexStart := funk.IndexOf(value, stations.StartRoute); indexStart != -1 {
			if indexEnd := funk.IndexOf(value, stations.EndRoute); indexEnd != -1 {
				result := data.Routes{key, stations}
				pathResult = append(pathResult, result)
			}

		}
	}
	common.RenderJSON(w, r, http.StatusOK, pathResult)
}
