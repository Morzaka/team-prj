package train_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"team-project/database"
	"team-project/services"
	"team-project/services/authorization"
	"team-project/services/data"
	"team-project/services/train"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

var router = services.NewRouter()

type TrainsTestCases struct {
	tcase        string
	url          string
	expected     int
	testTrainID  string
	mockedTrains []data.Train
	mockedTrain  data.Train
	mockedErr    error
}

var testTrain = data.Train{
	ID:            uuid.Must(uuid.Parse("08307904-f18e-4fb8-9d18-29cfad38ffaf")),
	DepartureCity: "Lviv",
	ArrivalCity:   "Kiev",
	DepartureTime: "8:40",
	DepartureDate: "29.04.2019",
	ArrivalTime:   "14:05",
	ArrivalDate:   "29.04.2019",
}

func TestValidateIfEmpty(t *testing.T) {
	v := train.ValidateIfEmpty(testTrain)
	if v != nil {
		t.Error("Expcected nil while validatingIfEmptty , got ", v)
	}
}

func TestNameIsValid(t *testing.T) {
	for i := 1; i < 3; i++ {
		switch i {
		case 1:
			v := train.NameIsValid(testTrain.DepartureCity)
			if v != true {
				t.Error("Expected that name validation was true, got ", v)
			}
		case 2:
			v := train.NameIsValid(testTrain.ArrivalCity)
			if v != true {
				t.Error("Expected that name validation was true, got ", v)
			}
		}
	}
}

func TestTimeIsValid(t *testing.T) {
	for i := 1; i < 3; i++ {
		switch i {
		case 1:
			v := train.TimeIsValid(testTrain.DepartureTime)
			if v != true {
				t.Error("Expected that name validation was true, got ", v)
			}
		case 2:
			v := train.TimeIsValid(testTrain.ArrivalTime)
			if v != true {
				t.Error("Expected that name validation was true, got ", v)
			}
		}
	}
}

func TestDateIsValid(t *testing.T) {
	for i := 1; i < 3; i++ {
		switch i {
		case 1:
			v := train.DateIsValid(testTrain.DepartureDate)
			if v != true {
				t.Error("Expected that name validation was true, got ", v)
			}
		case 2:
			v := train.DateIsValid(testTrain.ArrivalDate)
			if v != true {
				t.Error("Expected that name validation was true, got ", v)
			}
		}
	}
}
func TestGetTrains(t *testing.T) {
	tests := []TrainsTestCases{
		{
			tcase:        "OkGetTrains",
			url:          "/api/v1/trains",
			expected:     http.StatusOK,
			mockedTrains: []data.Train{},
			mockedErr:    nil,
		},
		{
			tcase:        "InternalServerError",
			url:          "/api/v1/trains",
			expected:     http.StatusBadRequest,
			mockedTrains: []data.Train{},
			mockedErr:    errors.New("db error"),
		},
	}
	authorization.AdminRole = func(http.ResponseWriter, *http.Request) bool {
		return true
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	for _, tc := range tests {
		t.Run(tc.tcase, func(t *testing.T) {
			trainMock := database.NewMockTrainCrud(mockCtrl)
			trainMock.EXPECT().GetAllTrains().Return(tc.mockedTrains, tc.mockedErr)
			database.Trains = trainMock
			rec := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodGet, tc.url, nil)
			if err != nil {
				t.Fatal(err)
			}
			router.ServeHTTP(rec, req)
			if rec.Code != tc.expected {
				t.Error("expected: ", tc.expected, " got: ", rec.Code)
			}
		})
	}
	authorization.AdminRole = authorization.CheckAdmin
}

func TestGetSingleTrain(t *testing.T) {
	test := []TrainsTestCases{
		{
			tcase:       "GetTrainOK",
			url:         "/api/v1/train/08307904-f18e-4fb8-9d18-29cfad38ffaf",
			expected:    http.StatusOK,
			testTrainID: "08307904-f18e-4fb8-9d18-29cfad38ffaf",
			mockedTrain: testTrain,
			mockedErr:   nil,
		},
		{
			tcase:       "NoContentGetTrain",
			url:         "/api/v1/train/08307904-f18e-4fb8-9d18-29cfad38aaaf",
			expected:    http.StatusBadRequest,
			testTrainID: "08307904-f18e-4fb8-9d18-29cfad38aaaf",
			mockedTrain: data.Train{},
			mockedErr:   errors.New("db error , no data found"),
		},
	}
	authorization.AdminRole = func(http.ResponseWriter, *http.Request) bool {
		return true
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	for _, tc := range test {
		t.Run(tc.tcase, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, err := http.NewRequest("GET", tc.url, nil)
			if err != nil {
				t.Fatal(err)
			}
			trainMock := database.NewMockTrainCrud(mockCtrl)
			trainMock.EXPECT().GetTrain(tc.testTrainID).Return(tc.mockedTrain, tc.mockedErr)
			database.Trains = trainMock
			router.ServeHTTP(rec, req)
			if rec.Code != tc.expected {
				t.Error("expected: ", tc.expected, " got: ", rec.Code)
			}
		})
	}
	authorization.AdminRole = authorization.CheckAdmin
}

func TestCreateTrain(t *testing.T) {
	test := []TrainsTestCases{
		{
			tcase:       "CreateTrainOK",
			url:         "/api/v1/train",
			expected:    http.StatusOK,
			mockedTrain: testTrain,
			mockedErr:   nil,
		},
		{
			tcase:       "CreateTrainNoContent",
			url:         "/api/v1/train",
			expected:    http.StatusBadRequest,
			mockedTrain: testTrain,
			mockedErr:   errors.New("failed to create"),
		},
	}
	authorization.AdminRole = func(http.ResponseWriter, *http.Request) bool {
		return true
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	for _, tc := range test {
		t.Run(tc.tcase, func(t *testing.T) {
			trainMock := database.NewMockTrainCrud(mockCtrl)
			b, err := json.Marshal(tc.mockedTrain)
			if err != nil {
				t.Fatal(err)
			}
			req, err := http.NewRequest(http.MethodPost, tc.url, bytes.NewBuffer(b))
			if err != nil {
				t.Fatal(err)
			}
			database.Trains = trainMock
			trainMock.EXPECT().AddTrain(tc.mockedTrain).Return(tc.mockedErr)
			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)
			if rec.Code != tc.expected {
				t.Error("expected: ", tc.expected, " got: ", rec.Code)
			}
		})
	}
	authorization.AdminRole = authorization.CheckAdmin
}

func TestUpdateTrain(t *testing.T) {
	tests := []TrainsTestCases{
		{
			tcase:       "UpdateTrainOk",
			url:         "/api/v1/train/08307904-f18e-4fb8-9d18-29cfad38ffaf",
			expected:    http.StatusOK,
			mockedTrain: testTrain,
			mockedErr:   nil,
		},
	}
	authorization.AdminRole = func(http.ResponseWriter, *http.Request) bool {
		return true
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	for _, tc := range tests {
		t.Run(tc.tcase, func(t *testing.T) {
			trainMock := database.NewMockTrainCrud(mockCtrl)
			b, err := json.Marshal(tc.mockedTrain)
			if err != nil {
				t.Fatal(err)
			}
			req, _ := http.NewRequest(http.MethodPatch, tc.url, bytes.NewBuffer(b))
			rec := httptest.NewRecorder()
			database.Trains = trainMock
			trainMock.EXPECT().UpdateTrain(tc.mockedTrain).Return(tc.mockedErr)
			trainMock.EXPECT().GetTrain(tc.mockedTrain.ID.String()).Return(tc.mockedTrain, tc.mockedErr)
			router.ServeHTTP(rec, req)
			if rec.Code != tc.expected {
				t.Error("expected: ", tc.expected, " got: ", rec.Code)
			}
		})
	}
	authorization.AdminRole = authorization.CheckAdmin
}

func TestDeleteTrain(t *testing.T) {
	tests := []TrainsTestCases{
		{
			tcase:     "DeleteTrainOk",
			url:       "/api/v1/train/08307904-f18e-4fb8-9d18-29cfad38ffaf",
			expected:  http.StatusOK,
			mockedErr: nil,
		},
		{
			tcase:     "DeleteTrainError",
			url:       "/api/v1/train/08307904-f18e-4fb8-9d18-29cfad38ffaf",
			expected:  http.StatusBadRequest,
			mockedErr: errors.New("db error"),
		},
	}
	authorization.AdminRole = func(http.ResponseWriter, *http.Request) bool {
		return true
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	for _, tc := range tests {
		t.Run(tc.tcase, func(t *testing.T) {
			trainMock := database.NewMockTrainCrud(mockCtrl)
			req, _ := http.NewRequest(http.MethodDelete, tc.url, nil)
			rec := httptest.NewRecorder()
			database.Trains = trainMock
			trainMock.EXPECT().DeleteTrain(uuid.Must(uuid.Parse("08307904-f18e-4fb8-9d18-29cfad38ffaf"))).Return(tc.mockedErr)
			router.ServeHTTP(rec, req)
			if rec.Code != tc.expected {
				t.Error("expected: ", tc.expected, " got: ", rec.Code)
			}
		})
	}
	authorization.AdminRole = authorization.CheckAdmin
}
