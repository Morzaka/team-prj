package routing

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"team-project/logger"
	"team-project/services/common"
	"team-project/services/data"

	"github.com/thoas/go-funk"

)
//Struct of Routes
type RouteStruct  struct {
	Index  string   `json:"index"`
	Routes []string `json:"routes"`
}

//Slice of Struct of Routes
type RouteStructs [] RouteStruct

var (RouteStorage map [string] [] string;
	initialised=false)


var myClient = &http.Client{Timeout: 10 * time.Second}



//gets railroad Data
func getData () map[string][] string {
	routeS:=make(map[string][]string,100)
	r, err := myClient.Get("https://api.myjson.com/bins/x7rfk")
	if err != nil {
		log.Printf("Error while get responce ")

	}
	respBytes, err := ioutil.ReadAll(r.Body)
	if err !=nil{
		fmt.Println("Error while reading http request")
	}
	routes := new(RouteStructs)
	json.Unmarshal(respBytes,&routes)

	for _,v:=range *routes{
		routeS[v.Index]=v.Routes
	}
	initialised=true
	return routeS
}



//FindPath generates routes
func FindPath(w http.ResponseWriter, r *http.Request) {

	if (initialised==false) { RouteStorage=getData()}


	var stations data.Stations
	var pathresult []data.Routes
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		logger.Logger.Errorf("Error, %s", err)
	}
	err = json.Unmarshal(b, &stations)
	if err != nil {
		logger.Logger.Errorf("Error, %s", err)
	}
	for key, value := range RouteStorage {
		if indexStart := funk.IndexOf(value, stations.StartRoute); indexStart != -1 {
			if indexEnd := funk.IndexOf(value, stations.EndRoute); indexEnd != -1 {
				result := data.Routes{key, stations}
				pathresult = append(pathresult, result)
			}

		}
	}
	common.RenderJSON(w, r, http.StatusOK, pathresult)
}
