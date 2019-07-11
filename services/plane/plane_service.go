package plane

import (
	"encoding/json"
	"net/http"
	"team-project/database"
	"team-project/services/common"
	"team-project/services/data"

	"github.com/go-zoo/bone"
	"github.com/google/uuid"
)

var emptyResponse interface{}

// GetPlanes get all planes from database
func GetPlanes(w http.ResponseWriter, r *http.Request) {
	planes, err := database.PlaneRepo.GetPlanes()
	if err != nil {
		common.RenderJSON(w, r, http.StatusInternalServerError, planes)
		return
	}

	common.RenderJSON(w, r, http.StatusOK, planes)
}

// GetPlane get plane from database by id
func GetPlane(w http.ResponseWriter, r *http.Request) {
	id := uuid.Must(uuid.Parse(bone.GetValue(r, "id")))
	plane, err := database.PlaneRepo.GetPlane(id)
	if err != nil {
		common.RenderJSON(w, r, 500, emptyResponse)
		return
	}
	common.RenderJSON(w, r, 200, plane)
}

// CreatePlane create new plane to database
func CreatePlane(w http.ResponseWriter, r *http.Request) {
	p := data.Plane{}
	p.ID = uuid.New()
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		common.RenderJSON(w, r, http.StatusBadRequest, emptyResponse)
		return
	}
	_, err = database.PlaneRepo.AddPlane(p)
	if err != nil {
		common.RenderJSON(w, r, http.StatusBadRequest, emptyResponse)
		return
	}
	common.RenderJSON(w, r, http.StatusCreated, p)
}

// UpdatePlane update plane in database by id
func UpdatePlane(w http.ResponseWriter, r *http.Request) {
	id := uuid.Must(uuid.Parse(bone.GetValue(r, "id")))
	p := data.Plane{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		common.RenderJSON(w, r, 500, emptyResponse)
		return
	}
	p.ID = id
	_, err = database.PlaneRepo.UpdatePlane(p, p.ID)
	if err != nil {
		common.RenderJSON(w, r, 500, emptyResponse)
		return
	}

	common.RenderJSON(w, r, 200, p)
}

// DeletePlane delete plane from database by id
func DeletePlane(w http.ResponseWriter, r *http.Request) {
	id := uuid.Must(uuid.Parse(bone.GetValue(r, "id")))
	err := database.PlaneRepo.DeletePlane(id)
	if err != nil {
		common.RenderJSON(w, r, http.StatusInternalServerError, emptyResponse)
		return
	}
	common.RenderJSON(w, r, http.StatusOK, emptyResponse)
}
