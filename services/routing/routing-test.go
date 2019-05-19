package routing

import (
	"encoding/json"
	"strconv"
	"team-project/services/data"
	"testing"
)

var (
	stations   data.Stations
	pathResult []data.Routes
	sum        int
)

//TestRouting if path from Lviv to Kyiv contains TrainsId
func TestRouting(t *testing.T) {

	cityJSON := `{
	"start":"Lviv"",
	"end":"Kyiv"
}`

	if !initialised {
		RouteStorage = getData()
	}

	err := json.Unmarshal([]byte(cityJSON), &stations)
	if err != nil {
		t.Error("Error while unmarshaling")
	}

	for key, value := range RouteStorage {
		if indexStart := IndexOfString(value, stations.StartRoute); indexStart != -1 {
			if indexEnd := IndexOfString(value, stations.EndRoute); indexEnd != -1 {
				result := data.Routes{RouteID: key, Stations: stations}
				pathResult = append(pathResult, result)
			}
		}
	}

	for _, v := range pathResult {
		conv, _ := strconv.Atoi(v.RouteID)
		sum += conv
	}

	if sum != 1848 {
		t.Error("Router not working")
	}

}
