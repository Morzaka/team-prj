package common

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"team-project/services/data"
)

var commonTest *ICommon

//TestRenderJSON tests function RenderJSON
func TestRenderJSON(t *testing.T) {
	r, err := http.NewRequest("POST", "/api/v1", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	var empty interface{}
	testData := []struct {
		status       int
		response     interface{}
		statusExpect int
	}{
		{http.StatusUnauthorized, empty, http.StatusUnauthorized},
		{http.StatusOK, data.Signin{Login: "Golang", Password: "Golang"}, http.StatusOK},
	}
	for _, testCase := range testData {
		commonTest.RenderJSON(w, r, testCase.status, testCase.response)
		if testCase.response == nil && w.Body.Len() != 0 {
			t.Error("ResponseRecorder body should be empty")
		}
		if w.Code != testCase.statusExpect && (testCase.response != nil && w.Body.Len() == 0) {
			t.Error("Function broke with error")
		}

	}
}
