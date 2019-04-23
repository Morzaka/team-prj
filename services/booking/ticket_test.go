package booking

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"team-project/database"
	"team-project/services/data"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
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

//var router = services.NewRouter()

func TestValidateForm(t *testing.T) {
	for i := 0; i < 8; i++ {
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

type ListAllTicketsTestCase struct {
	name          string
	url           string
	want          int
	mockedTickets []data.Ticket
	mockedError   error
}

func TestGetAllTickets(t *testing.T) {
	tests := []ListAllTicketsTestCase{
		{
			name:          "Get_Tickets_200",
			url:           "/api/v1/tickets",
			want:          http.StatusOK,
			mockedTickets: []data.Ticket{},
			mockedError:   nil,
		},
		{
			name:          "Get_Tickets_500",
			url:           "/api/v1/tickets",
			want:          http.StatusInternalServerError,
			mockedTickets: []data.Ticket{},
			mockedError:   errors.New("db error"),
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ticketMock := database.NewMockTicketRepository(mockCtrl)

			ticketMock.EXPECT().AllTickets().Return(tc.mockedTickets, tc.mockedError)

			database.TicketRepo = ticketMock

			rw := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tc.url, nil)
			//router.ServeHTTP(rw, req)
			http.HandlerFunc(GetAllTickets).ServeHTTP(rw, req)

			if rw.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rw.Code)
			}
		})
	}
}

type ListTicketTestCase struct {
	name          string
	id            uuid.UUID
	url           string
	want          int
	mockedTickets data.Ticket
	mockedError   error
}

func TestGetOneTicket(t *testing.T) {
	tests := []ListTicketTestCase{
		{
			name:          "Get_Ticket_200",
			id:            uuid.Must(uuid.Parse("fcb33af4-40a3-4c82-afb1-218731052309")),
			url:           "/api/v1/ticket?id=fcb33af4-40a3-4c82-afb1" +
				"-218731052309",
			want:          http.StatusOK,
			mockedTickets: testData,
			mockedError:   nil,
		},
		{
			name:          "Get_Tickets_500",
			id:            uuid.Must(uuid.Parse("0e3763c6-a7ed-4532-afd7-420c5a480000")),
			url:           "/api/v1/ticket?id=0e3763c6-a7ed-4532-afd7" +
				"-420c5a480000",
			want:          http.StatusInternalServerError,
			mockedTickets: data.Ticket{},
			mockedError:   errors.New("db error"),
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ticketMock := database.NewMockTicketRepository(mockCtrl)

			ticketMock.EXPECT().GetTicket(tc.id).Return(tc.mockedTickets,
				tc.mockedError)

			database.TicketRepo = ticketMock

			rw := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tc.url, nil)

			//router.ServeHTTP(rw, req)
			http.HandlerFunc(GetOneTicket).ServeHTTP(rw, req)

			if rw.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rw.Code)
			}
		})
	}
}
