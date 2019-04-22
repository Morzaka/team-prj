package plane

import (
	"encoding/json"
	"net/http"
	"team-project/database"
	"team-project/services/common"
	"team-project/services/data"
	"team-project/services/model"
)

var emptyResponse interface{}

// GetPlanes get all planes from database
func GetPlanes(w http.ResponseWriter, r *http.Request) {
	planes, err := database.GetPlanes()
	if err != nil {
		common.RenderJSON(w, r, http.StatusInternalServerError, planes)
		return
	}

	common.RenderJSON(w, r, http.StatusOK, planes)
}

// GetPlane get plane from database by id
func GetPlane(w http.ResponseWriter, r *http.Request) {
	id, err := model.GetID(r)
	if err != nil {
		common.RenderJSON(w, r, 400, emptyResponse)
		return
	}
	plane, err := database.GetPlane(id)
	if err != nil {
		common.RenderJSON(w, r, 404, emptyResponse)
		return
	}
	common.RenderJSON(w, r, 202, plane)
}

// CreatePlane create new plane to database
func CreatePlane(w http.ResponseWriter, r *http.Request) {
	p := data.Plane{}
	p.ID = model.GenerateID()
	json.NewDecoder(r.Body).Decode(&p)
	_, err := database.AddPlane(p)
	if err != nil {
		common.RenderJSON(w, r, 404, emptyResponse)
		return
	}
	common.RenderJSON(w, r, 202, p)
}

// UpdatePlane update plane in database by id
func UpdatePlane(w http.ResponseWriter, r *http.Request) {
	id, err := model.GetID(r)
	if err != nil {
		common.RenderJSON(w, r, 404, emptyResponse)
		return
	}
	p := data.Plane{}
	json.NewDecoder(r.Body).Decode(&p)
	p.ID = id
	_, err = database.UpdatePlane(p, p.ID)
	if err != nil {
		common.RenderJSON(w, r, 404, emptyResponse)
		return
	}
	plane, err := database.GetPlane(p.ID)
	if err != nil {
		common.RenderJSON(w, r, 404, emptyResponse)
		return
	}
	common.RenderJSON(w, r, 202, plane)
}

// DeletePlane delete plane from database by id
func DeletePlane(w http.ResponseWriter, r *http.Request) {
	id, err := model.GetID(r)
	if err != nil {
		common.RenderJSON(w, r, 404, emptyResponse)
		return
	}
	err = database.DeletePlane(id)
	if err != nil {
		common.RenderJSON(w, r, 404, emptyResponse)
		return
	}
	common.RenderJSON(w, r, 202, emptyResponse)
}
