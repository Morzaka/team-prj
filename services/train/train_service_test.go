package train_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"team-project/database"
	"team-project/services"
	"team-project/services/data"
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

func TestGetTrains(t *testing.T) {
	tests := []TrainsTestCases{
		{
			tcase:        "GetTrains200",
			url:          "/api/v1/trains",
			expected:     http.StatusOK,
			mockedTrains: []data.Train{},
			mockedErr:    nil,
		},
		{
			tcase:        "GetTrains204",
			url:          "/api/v1/trains",
			expected:     http.StatusNoContent,
			mockedTrains: []data.Train{},
			mockedErr:    errors.New("db error"),
		},
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
}

func TestGetSingleTrain(t *testing.T) {
	s := "08307904-f18e-4fb8-9d18-29cfad38ffaf"
	id, err := uuid.Parse(s)
	if err != nil {
		t.Fatal(err)
	}
	test := []TrainsTestCases{
		/*{
			tcase:       "GetTrainOK",
			url:         "/api/v1/train/08307904-f18e-4fb8-9d18-29cfad38ffaf",
			expected:    http.StatusOK,
			mockedTrain: data.Train{},
			mockedErr:   nil,
		},*/
		{
			tcase:       "GetTrain204",
			url:         "/api/v1/train/08307904-f18e-4fb8-9d18-29cfad38ffaf",
			expected:    http.StatusNoContent,
			mockedTrain: data.Train{},
			mockedErr:   errors.New("db error , no data found"),
		},
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	for _, tc := range test {
		t.Run(tc.tcase, func(t *testing.T) {
			trainMock := database.NewMockTrainCrud(mockCtrl)
			trainMock.EXPECT().GetTrain(s).Return(data.Train{
				ID: id,
			}, tc.mockedErr)
			database.Trains = trainMock
			rec := httptest.NewRecorder()
			req, err := http.NewRequest("GET", tc.url, nil)
			if err != nil {
				t.Fatal(err)
			}
			router.ServeHTTP(rec, req)
			if rec.Code != tc.expected {
				t.Error("expected: ", tc.expected, " got: ", rec.Code)
			}
		})
	}
}

func TestCreateTrain(t *testing.T) {
	test := []TrainsTestCases{
		{
			tcase:       "GetTrainOK",
			url:         "/api/v1/train",
			expected:    http.StatusOK,
			mockedTrain: data.Train{},
			mockedErr:   nil,
		},
		{
			tcase:       "GetTrainOK",
			url:         "/api/v1/train",
			expected:    http.StatusNoContent,
			mockedTrain: data.Train{},
			mockedErr:   errors.New("failed to create"),
		},
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	for _, tc := range test {
		t.Run(tc.tcase, func(t *testing.T) {
			trainMock := database.NewMockTrainCrud(mockCtrl)
			train := data.Train{}
			req, err := http.NewRequest("POST", tc.url, nil)
			if err != nil {
				t.Fatal(err)
			}
			json.NewDecoder(req.Body).Decode(&train)
			trainMock.EXPECT().AddTrain(train).Return(tc.mockedErr)
			database.Trains = trainMock
			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)
			if rec.Code != tc.expected {
				t.Error("expected: ", tc.expected, " got: ", rec.Code)
			}
		})
	}
}
