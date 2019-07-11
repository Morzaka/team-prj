package trip_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"team-project/services/trip"
	"testing"

	"team-project/database"
	"team-project/services"
	"team-project/services/data"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

var testDataTripHandlers = data.Trip{
	TripID:             uuid.Must(uuid.Parse("fcb33af4-40a3-4c82-afb1-218731052309")),
	TripName:           "CoolWeek",
	TripTicketID:       uuid.Must(uuid.Parse("b0ffec41-eb5f-41a4-adab-4d6944a748ad")),
	TripReturnTicketID: uuid.Must(uuid.Parse("0e3763c6-a7ed-4532-afd7-420c5a48cea9")),
	TotalTripPrice:     125,
}

var router = services.NewRouter()

type ListTripTestCase struct {
	name        string
	id          uuid.UUID
	url         string
	want        int
	mockedTrip  data.Trip
	mockedTrips []data.Trip
	mockedError error
}

//TestValidateFormTrip function for test validation
func TestValidateFormTrip(t *testing.T) {
	for i := 0; i < 4; i++ {
		switch i {
		case 0:
			v := trip.ValidateFormTrip(testDataTripHandlers)
			if v != nil {
				t.Error("Expected nil, got ", v)
			}
		case 1:
			testDataTripHandlers.TotalTripPrice = 0
			v := trip.ValidateFormTrip(testDataTripHandlers)
			if v == nil {
				t.Error("Expected nil, got ", v)
			}
		case 2:
			testDataTripHandlers.TripName = ""
			v := trip.ValidateFormTrip(testDataTripHandlers)
			if v == nil {
				t.Error("Expected nil, got ", v)
			}

		}
	}
}

//TestGetTrips function for test GetTrips route
func TestGetTrips(t *testing.T) {
	trips := []ListTripTestCase{
		{
			name:        "Get_Trip_202",
			url:         "/api/v1/trips",
			want:        http.StatusOK,
			mockedTrips: []data.Trip{},
			mockedError: nil,
		},
		{
			name:        "Get_Trips_500",
			url:         "/api/v1/trips",
			want:        http.StatusInternalServerError,
			mockedTrips: []data.Trip{},
			mockedError: errors.New("db error"),
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, tr := range trips {
		t.Run(tr.name, func(t *testing.T) {
			tripMock := database.NewMockTripRepository(mockCtrl)
			tripMock.EXPECT().GetTrips().Return(tr.mockedTrips, tr.mockedError)
			database.TripRepo = tripMock

			rw := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tr.url, nil)
			//http.HandlerFunc(.GetAllTickets).ServeHTTP(rw, req)
			router.ServeHTTP(rw, req)

			if rw.Code != tr.want {
				t.Errorf("Expected: %d , got %d", tr.want, rw.Code)
			}
		})
	}
}

//TestGetTrip function for test GetTrip route
func TestGetTrip(t *testing.T) {
	tests := []ListTripTestCase{
		{
			name:        "Get_Ticket_200",
			id:          uuid.Must(uuid.Parse("fcb33af4-40a3-4c82-afb1-218731052309")),
			url:         "/api/v1/trip/fcb33af4-40a3-4c82-afb1-218731052309",
			want:        http.StatusOK,
			mockedTrip:  testDataTripHandlers,
			mockedError: nil,
		},
		{
			name:        "Get_Trips_500",
			id:          uuid.Must(uuid.Parse("9ed94c7c-6767-11e9-a923-1681be663d3e")),
			url:         "/api/v1/trip/9ed94c7c-6767-11e9-a923-1681be663d3e",
			want:        http.StatusInternalServerError,
			mockedTrip:  testDataTripHandlers,
			mockedError: errors.New("db error"),
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, tr := range tests {
		t.Run(tr.name, func(t *testing.T) {
			if tr.name != "Get_Trips_404" {
				tripMock := database.NewMockTripRepository(mockCtrl)
				tripMock.EXPECT().GetTrip(tr.id).Return(tr.mockedTrip,
					tr.mockedError)
				database.TripRepo = tripMock
			}

			rw := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tr.url, nil)
			//http.HandlerFunc(GetOneTicket).ServeHTTP(rw, req)
			router.ServeHTTP(rw, req)
			if rw.Code != tr.want {
				t.Errorf("Expected: %d , got %d", tr.want, rw.Code)
			}
		})
	}
}

//TestCreateTrip function for test CreateTrip route
func TestCreateTrip(t *testing.T) {
	trip := data.Trip{
		TripName:           "NewTrip",
		TripTicketID:       uuid.Must(uuid.Parse("82ca5cea-675a-11e9-a923-1681be663d3e")),
		TripReturnTicketID: uuid.Must(uuid.Parse("88a48bae-675a-11e9-a923-1681be663d3e")),
		TotalTripPrice:     545,
	}
	tests := []ListTripTestCase{
		{
			name:        "Post_Trip_200",
			url:         "/api/v1/trip",
			want:        http.StatusOK,
			mockedTrip:  trip,
			mockedError: nil,
		},
		{
			name:        "Post_Trip_500",
			url:         "/api/v1/trip",
			want:        http.StatusInternalServerError,
			mockedTrip:  trip,
			mockedError: errors.New("db error"),
		},
		{
			name:        "Post_Trip_404",
			url:         "/api/v1/trip",
			want:        http.StatusInternalServerError,
			mockedTrip:  data.Trip{},
			mockedError: errors.New("validation failure"),
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, tr := range tests {
		t.Run(tr.name, func(t *testing.T) {
			if tr.name != "Post_Trip_404" {
				tripMock := database.NewMockTripRepository(mockCtrl)
				tripMock.EXPECT().AddTrip(tr.mockedTrip).Return(tr.mockedTrip, tr.mockedError)
				database.TripRepo = tripMock
			}
			b, err := json.Marshal(tr.mockedTrip)
			if err != nil {
				t.Error(err)
			}

			rw := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, tr.url, bytes.NewBuffer(b))
			router.ServeHTTP(rw, req)
			if rw.Code != tr.want {
				t.Errorf("Expected: %d , got %d", tr.want, rw.Code)
			}
		})
	}
}

//TestUpdateTrip function for test UpdateTrip route
func TestUpdateTrip(t *testing.T) {
	trip := data.Trip{
		TripID:             uuid.Must(uuid.Parse("36e48772-675c-11e9-a923-1681be663d3e")),
		TripName:           "CoolWeek",
		TripTicketID:       uuid.Must(uuid.Parse("3ba020d2-675c-11e9-a923-1681be663d3e")),
		TripReturnTicketID: uuid.Must(uuid.Parse("40ead3fc-675c-11e9-a923-1681be663d3e")),
		TotalTripPrice:     2112,
	}
	tests := []ListTripTestCase{
		{
			name:        "Patch_Trip_200",
			url:         "/api/v1/trip/36e48772-675c-11e9-a923-1681be663d3e",
			want:        http.StatusOK,
			mockedTrip:  trip,
			mockedError: nil,
		},
		{
			name:        "Patch_Trip_404",
			url:         "/api/v1/trip/36e48772-675c-11e9-a923-1681be663d3e",
			want:        http.StatusInternalServerError,
			mockedTrip:  trip,
			mockedError: errors.New("db error"),
		},
		{
			name:        "Patch_Trip_500",
			url:         "/api/v1/trip/36e48772-675c-11e9-a923-1681be663d3e",
			want:        http.StatusInternalServerError,
			mockedTrip:  data.Trip{},
			mockedError: errors.New("validation failure"),
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, tr := range tests {
		t.Run(tr.name, func(t *testing.T) {
			if tr.name != "Patch_Trip_500" {
				tripMock := database.NewMockTripRepository(mockCtrl)
				tripMock.EXPECT().UpdateTrip(tr.mockedTrip).Return(tr.mockedTrip, tr.mockedError)
				database.TripRepo = tripMock
			}
			b, err := json.Marshal(tr.mockedTrip)
			if err != nil {
				t.Error(err)
			}

			rw := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPatch, tr.url, bytes.NewBuffer(b))
			router.ServeHTTP(rw, req)
			if rw.Code != tr.want {
				t.Errorf("Expected: %d , got %d", tr.want, rw.Code)
			}
		})
	}
}

//TestDeleteTrip function for test DeleteTrip route
func TestDeleteTrip(t *testing.T) {
	trips := []ListTripTestCase{
		{
			name:        "Delete_Trip_404",
			id:          uuid.Must(uuid.Parse("fcb33af4-40a3-4c82-afb1-218731052309")),
			url:         "/api/v1/trip/fcb33af4-40a3-4c82-afb1-218731052309",
			want:        http.StatusOK,
			mockedError: nil,
		},
		{
			name:        "Delete_Trip_500",
			id:          uuid.Must(uuid.Parse("fcb33af4-40a3-4c82-afb1-218731052309")),
			url:         "/api/v1/trip/fcb33af4-40a3-4c82-afb1-218731052309",
			want:        http.StatusInternalServerError,
			mockedError: errors.New("db error"),
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, tr := range trips {
		t.Run(tr.name, func(t *testing.T) {

			tripMock := database.NewMockTripRepository(mockCtrl)
			tripMock.EXPECT().DeleteTrip(tr.id).Return(tr.mockedError)
			database.TripRepo = tripMock

			rw := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodDelete, tr.url, nil)
			router.ServeHTTP(rw, req)
			if rw.Code != tr.want {
				t.Errorf("Expected: %d , got %d", tr.want, rw.Code)
			}
		})
	}
}
