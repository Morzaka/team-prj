package plane_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"team-project/database"
	"team-project/services"
	"team-project/services/data"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

var testDataPlane = data.Plane{
	ID:            uuid.Must(uuid.Parse("fcb33af4-40a3-4c82-afb1-218731052309")),
	DepartureCity: "Lviv",
	ArrivalCity:   "Kyiv",
}

var router = services.NewRouter()

type ListPlaneTestCase struct {
	name         string
	id           uuid.UUID
	url          string
	want         int
	mockedPlane  data.Plane
	mockedPlanes []data.Plane
	mockedError  error
}

func TestGetPlanes(t *testing.T) {
	tests := []ListPlaneTestCase{
		{
			name:         "Get_Planes_200",
			url:          "/api/v1/planes",
			want:         http.StatusOK,
			mockedPlanes: []data.Plane{},
			mockedError:  nil,
		},
		{
			name:         "Get_Planes_500",
			url:          "/api/v1/planes",
			want:         http.StatusInternalServerError,
			mockedPlanes: []data.Plane{},
			mockedError:  errors.New("db error"),
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, p := range tests {
		t.Run(p.name, func(t *testing.T) {
			planeMock := database.NewMockPlaneRepository(mockCtrl)
			planeMock.EXPECT().GetPlanes().Return(p.mockedPlanes, p.mockedError)
			database.PlaneRepo = planeMock

			rw := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, p.url, nil)
			//http.HandlerFunc(.GetPlanes).ServeHTTP(rw, req)
			router.ServeHTTP(rw, req)

			if rw.Code != p.want {
				t.Errorf("Expected: %d , got %d", p.want, rw.Code)
			}
		})
	}
}

func TestGetPlane(t *testing.T) {
	tests := []ListPlaneTestCase{
		{
			name:        "Get_Plane_200",
			id:          uuid.Must(uuid.Parse("fcb33af4-40a3-4c82-afb1-218731052309")),
			url:         "/api/v1/plane/fcb33af4-40a3-4c82-afb1-218731052309",
			want:        http.StatusOK,
			mockedPlane: testDataPlane,
			mockedError: nil,
		},
		{
			name:        "Get_Plane_500",
			id:          uuid.Must(uuid.Parse("0e3763c6-a7ed-4532-afd7-420c5a480000")),
			url:         "/api/v1/plane/0e3763c6-a7ed-4532-afd7-420c5a480000",
			want:        http.StatusInternalServerError,
			mockedPlane: data.Plane{},
			mockedError: errors.New("db error"),
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, p := range tests {
		t.Run(p.name, func(t *testing.T) {
			if p.name != "Get_Plane_400" {
				planeMock := database.NewMockPlaneRepository(mockCtrl)
				planeMock.EXPECT().GetPlane(p.id).Return(p.mockedPlane,
					p.mockedError)
				database.PlaneRepo = planeMock
			}

			rw := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, p.url, nil)
			//http.HandlerFunc(GetOnePlane).ServeHTTP(rw, req)
			router.ServeHTTP(rw, req)
			if rw.Code != p.want {
				t.Errorf("Expected: %d , got %d", p.want, rw.Code)
			}
		})
	}
}

func TestCreatePlane(t *testing.T) {
	plane := data.Plane{
		ID:            uuid.Must(uuid.Parse("fcb33af4-40a3-4c82-afb1-218731052309")),
		DepartureCity: "Kyiv",
		ArrivalCity:   "Lviv",
	}
	tests := []ListPlaneTestCase{
		{
			name:        "Post_Plane_201",
			url:         "/api/v1/plane",
			want:        http.StatusCreated,
			mockedPlane: plane,
			mockedError: nil,
		},
		{
			name:        "Post_Plane_500",
			url:         "/api/v1/plane",
			want:        http.StatusBadRequest,
			mockedPlane: plane,
			mockedError: errors.New("db error"),
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, p := range tests {
		t.Run(p.name, func(t *testing.T) {
			if p.name != "Post_Plane_406" {
				planeMock := database.NewMockPlaneRepository(mockCtrl)
				planeMock.EXPECT().AddPlane(p.mockedPlane).Return(p.mockedPlane, p.mockedError)
				database.PlaneRepo = planeMock
			}
			b, err := json.Marshal(p.mockedPlane)
			if err != nil {
				t.Error(err)
			}

			rw := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, p.url, bytes.NewBuffer(b))
			router.ServeHTTP(rw, req)
			if rw.Code != p.want {
				t.Errorf("Expected: %d , got %d", p.want, rw.Code)
			}
		})
	}
}

func TestUpdatePlane(t *testing.T) {
	plane := data.Plane{
		ID:            uuid.Must(uuid.Parse("fcb33af4-40a3-4c82-afb1-218731052309")),
		DepartureCity: "Odessa",
		ArrivalCity:   "Lviv",
	}
	tests := []ListPlaneTestCase{
		{
			name:        "Patch_Plane_200",
			url:         "/api/v1/plane/fcb33af4-40a3-4c82-afb1-218731052309",
			want:        http.StatusOK,
			mockedPlane: plane,
			mockedError: nil,
		},
		{
			name:        "Patch_Plane_500",
			url:         "/api/v1/plane/fcb33af4-40a3-4c82-afb1-218731052309",
			want:        http.StatusInternalServerError,
			mockedPlane: plane,
			mockedError: errors.New("db error"),
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, p := range tests {
		t.Run(p.name, func(t *testing.T) {
			if p.name != "Patch_Plane_406" {
				planeMock := database.NewMockPlaneRepository(mockCtrl)
				planeMock.EXPECT().UpdatePlane(p.mockedPlane, p.mockedPlane.ID).Return(p.mockedPlane, p.mockedError)
				database.PlaneRepo = planeMock
			}
			b, err := json.Marshal(p.mockedPlane)
			if err != nil {
				t.Error(err)
			}

			rw := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPatch, p.url, bytes.NewBuffer(b))
			router.ServeHTTP(rw, req)
			if rw.Code != p.want {
				t.Errorf("Expected: %d , got %d", p.want, rw.Code)
			}
		})
	}
}

func TestDeleteTicket(t *testing.T) {
	tests := []ListPlaneTestCase{
		{
			name:        "Delete_Plane_200",
			id:          uuid.Must(uuid.Parse("fcb33af4-40a3-4c82-afb1-218731052309")),
			url:         "/api/v1/plane/fcb33af4-40a3-4c82-afb1-218731052309",
			want:        http.StatusOK,
			mockedError: nil,
		},
		{
			name:        "Delete_Plane_500",
			id:          uuid.Must(uuid.Parse("fcb33af4-40a3-4c82-afb1-218731052309")),
			url:         "/api/v1/plane/fcb33af4-40a3-4c82-afb1-218731052309",
			want:        http.StatusInternalServerError,
			mockedError: errors.New("db error"),
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, p := range tests {
		t.Run(p.name, func(t *testing.T) {

			planeMock := database.NewMockPlaneRepository(mockCtrl)
			planeMock.EXPECT().DeletePlane(p.id).Return(p.mockedError)
			database.PlaneRepo = planeMock

			rw := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodDelete, p.url, nil)
			router.ServeHTTP(rw, req)
			if rw.Code != p.want {
				t.Errorf("Expected: %d , got %d", p.want, rw.Code)
			}
		})
	}
}
