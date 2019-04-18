package train_test

import (
	"net/http"
	"net/http/httptest"
	"team-project/services"
	"testing"
)

var router = services.NewRouter()

type TrainsTestCases struct {
	tcase       string
	url         string
	expected    int
	testTrainID string
}

func TestGetTrains(t *testing.T) {
	tests := []TrainsTestCases{
		{
			tcase:    "GetTrainsOk",
			url:      "api/v1/trains",
			expected: 404,
		},
	}
	for _, tc := range tests {
		t.Run(tc.tcase, func(t *testing.T) {
			req, err := http.NewRequest("GET", tc.url, nil)
			if err != nil {
				t.Fatal(err)
			}
			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)
			rec.WriteHeader(tc.expected)
			if rec.Code != tc.expected {
				t.Error("expected: ", tc.expected, " want: ", rec.Code)
			}
		})
	}
}
