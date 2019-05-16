package routing

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"team-project/configurations"
	"team-project/logger"
	"team-project/services/common"
	"team-project/services/data"
	"time"
)

var (
	//RouteStorage is a map of stations id as key and stations as value
	RouteStorage map[string][]string
	initialised  = false
	myClient     = &http.Client{Timeout: 10 * time.Second}
)

//indexOf function finds first occurence of element in the slice
func indexOf(n int, f func(int) bool) int {
	for i := 0; i < n; i++ {
		if f(i) {
			return i
		}
	}
	return -1
}

//IndexOfString gets the index at which the first occurrence of a string value is found in array or return -1
// if the value cannot be found
func IndexOfString(a []string, x string) int {
	return indexOf(len(a), func(i int) bool { return a[i] == x })
}

//getData gets railroad Data
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
	routes := new(data.RouteStructs)
	err = json.Unmarshal(respBytes, &routes)
	if err != nil {
		logger.Logger.Error("Error while unmarshaling data")
	}

	for _, v := range *routes {
		routeS[v.Index] = v.Routes
	}
	initialised = true
	return routeS
}

//FindPath generates routes
func FindPath(w http.ResponseWriter, r *http.Request) {

	if !initialised {
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
		if indexStart := IndexOfString(value, stations.StartRoute); indexStart != -1 {
			if indexEnd := IndexOfString(value, stations.EndRoute); indexEnd != -1 {
				result := data.Routes{RouteID: key, Stations: stations}
				pathResult = append(pathResult, result)
			}

		}
	}
	common.RenderJSON(w, r, http.StatusOK, pathResult)
}
