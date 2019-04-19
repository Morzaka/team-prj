package booking

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"net/http"
	"net/http/httptest"
	"team-project/services"
	"team-project/services/data"
	"testing"
)


var testData = data.Ticket{
	ID:         uuid.Must(uuid.Parse("fcb33af4-40a3-4c82-afb1-218731052309")),
	TrainID:    uuid.Must(uuid.Parse("a521d12f-148a-4689-a0ff-e05ec1a40699")),
	PlaneID:    uuid.Must(uuid.Parse("b0ffec41-eb5f-41a4-adab-4d6944a748ad")),
	UserID:     uuid.Must(uuid.Parse("0e3763c6-a7ed-4532-afd7-420c5a48cea9")),
	Place:      61,
	TicketType: "Plane",
	Discount:   "-2%",
	Price:      52.60,
	TotalPrice: 51.15,
	Name:       "Oleh",
	Surname:    "Vynnyk",
}

var router = services.NewRouter()

func TestValidateForm(t *testing.T){
	for i := 0 ; i < 8; i++ {
		switch i {
		case 0:
			v := validateForm(testData)
			if v != nil {
				t.Error("Expected nil, got ", v)
			}
		case 1:
			testData.Place = 0
			v := validateForm(testData)
			if v == nil {
				t.Error("Expected nil, got ", v)
			}
		case 2:
			testData.Place = 21
			testData.TicketType = ""
			v := validateForm(testData)
			if v == nil {
				t.Error("Expected nil, got ", v)
			}
		case 3:
			testData.TicketType = "Train"
			testData.Discount = ""
			v := validateForm(testData)
			if v == nil {
				t.Error("Expected nil, got ", v)
			}
		case 4:
			testData.Discount = "-10%"
			testData.Price = -2
			v := validateForm(testData)
			if v == nil {
				t.Error("Expected nil, got ", v)
			}
		case 5:
			testData.Price = 23.32
			testData.TotalPrice = 0
			v := validateForm(testData)
			if v == nil {
				t.Error("Expected nil, got ", v)
			}
		case 6:
			testData.TotalPrice = 23.32
			testData.Name = ""
			v := validateForm(testData)
			if v == nil {
				t.Error("Expected nil, got ", v)
			}
		case 7:
			testData.Name = "Pavlo"
			testData.Surname = ""
			v := validateForm(testData)
			if v == nil {
				t.Error("Expected nil, got ", v)
			}
		}
	}
}

func GetAllTickets(t *testing.T) {
	//tests := []ListAllUsersTestCase{
	//	{
	//		name:        "Get_Users_200",
	//		url:         "/api/v1/users",
	//		want:        http.StatusOK,
	//		mockedUsers: []data.User{},
	//		mockedError: nil,
	//	},
	//	{
	//		name:        "Get_Users_404",
	//		url:         "/api/v1/users",
	//		want:        http.StatusNoContent,
	//		mockedUsers: []data.User{},
	//		mockedError: errors.New("db error"),
	//	},
	//}

	ts := []*data.Ticket{
		&data.Ticket{
			ID:         uuid.Must(uuid.Parse("fcb33af4-40a3-4c82-afb1-218731052309")),
			TrainID:    uuid.Must(uuid.Parse("a521d12f-148a-4689-a0ff-e05ec1a40699")),
			PlaneID:    uuid.Must(uuid.Parse("b0ffec41-eb5f-41a4-adab-4d6944a748ad")),
			UserID:     uuid.Must(uuid.Parse("0e3763c6-a7ed-4532-afd7-420c5a48cea9")),
		},
		&data.Ticket{
			ID:         uuid.Must(uuid.Parse("fcb33af4-40a3-4c82-afb1-218731052309")),
			TrainID:    uuid.Must(uuid.Parse("a521d12f-148a-4689-a0ff-e05ec1a40699")),
			PlaneID:    uuid.Must(uuid.Parse("b0ffec41-eb5f-41a4-adab-4d6944a748ad")),
			UserID:     uuid.Must(uuid.Parse("0e3763c6-a7ed-4532-afd7-420c5a48cea9")),
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
