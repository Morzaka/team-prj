package authorization_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"team-project/database"
	"team-project/services"
	"team-project/services/data"

	"github.com/golang/mock/gomock"
)

var router = services.NewRouter()

type ListAllUsersTestCase struct {
	name        string
	url         string
	want        int
	mockedUsers []data.User
	mockedError error
}

func TestGetUsers(t *testing.T) {
	tests := []ListAllUsersTestCase{
		{
			name:        "Get_Users_200",
			url:         "/api/v1/users",
			want:        http.StatusOK,
			mockedUsers: []data.User{},
			mockedError: nil,
		},
		{
			name:        "Get_Users_404",
			url:         "/api/v1/users",
			want:        http.StatusNoContent,
			mockedUsers: []data.User{},
			mockedError: errors.New("db error"),
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			usersMock := database.NewMockUserCRUD(mockCtrl)

			usersMock.EXPECT().GetAllUsers().Return(tc.mockedUsers, tc.mockedError)

			database.Users = usersMock

			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tc.url, nil)

			router.ServeHTTP(rec, req)

			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
		})
	}
}
