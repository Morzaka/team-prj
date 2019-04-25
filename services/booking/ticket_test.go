package booking_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"team-project/database"
	"team-project/services"
	"team-project/services/authorization"
	"team-project/services/booking"
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

var router = services.NewRouter()

type ListTicketTestCase struct {
	name            string
	id              uuid.UUID
	url             string
	want            int
	mockedTicket    data.Ticket
	mockedTickets   []data.Ticket
	mockedError     error
	mockedAuthorize bool
}

func TestValidateForm(t *testing.T) {
	for i := 0; i < 8; i++ {
		switch i {
		case 0:
			v := booking.ValidateForm(testData)
			if v != nil {
				t.Error("Expected nil, got ", v)
			}
		case 1:
			testData.Place = 0
			v := booking.ValidateForm(testData)
			if v == nil {
				t.Error("Expected nil, got ", v)
			}
		case 2:
			testData.Place = 21
			testData.TicketType = ""
			v := booking.ValidateForm(testData)
			if v == nil {
				t.Error("Expected nil, got ", v)
			}
		case 3:
			testData.TicketType = "Train"
			testData.Discount = ""
			v := booking.ValidateForm(testData)
			if v == nil {
				t.Error("Expected nil, got ", v)
			}
		case 4:
			testData.Discount = "-10%"
			testData.Price = -2
			v := booking.ValidateForm(testData)
			if v == nil {
				t.Error("Expected nil, got ", v)
			}
		case 5:
			testData.Price = 23.32
			testData.TotalPrice = 0
			v := booking.ValidateForm(testData)
			if v == nil {
				t.Error("Expected nil, got ", v)
			}
		case 6:
			testData.TotalPrice = 23.32
			testData.Name = ""
			v := booking.ValidateForm(testData)
			if v == nil {
				t.Error("Expected nil, got ", v)
			}
		case 7:
			testData.Name = "Pavlo"
			testData.Surname = ""
			v := booking.ValidateForm(testData)
			if v == nil {
				t.Error("Expected nil, got ", v)
			}
		}
	}
}

func TestGetAllTickets(t *testing.T) {
	tests := []ListTicketTestCase{
		{
			name:            "Get_Tickets_200",
			url:             "/api/v1/tickets",
			want:            http.StatusOK,
			mockedTickets:   []data.Ticket{},
			mockedError:     nil,
			mockedAuthorize: true,
		},
		{
			name:            "Get_Tickets_500",
			url:             "/api/v1/tickets",
			want:            http.StatusInternalServerError,
			mockedTickets:   []data.Ticket{},
			mockedError:     errors.New("db error"),
			mockedAuthorize: true,
		},
		{
			name:            "Get_Tickets_403",
			url:             "/api/v1/tickets",
			want:            http.StatusForbidden,
			mockedAuthorize: false,
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	defer func() { authorization.AdminRole = authorization.CheckAccess }()

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			authorization.AdminRole = func(w http.ResponseWriter, r *http.Request) bool {
				return tc.mockedAuthorize
			}
			if tc.mockedAuthorize {
				ticketMock := database.NewMockTicketRepository(mockCtrl)
				ticketMock.EXPECT().AllTickets().Return(tc.mockedTickets, tc.mockedError)
				database.TicketRepo = ticketMock
			}
			rw := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tc.url, nil)
			router.ServeHTTP(rw, req)

			if rw.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rw.Code)
			}
		})
	}
}

func TestGetOneTicket(t *testing.T) {
	tests := []ListTicketTestCase{
		{
			name:            "Get_Ticket_200",
			id:              uuid.Must(uuid.Parse("fcb33af4-40a3-4c82-afb1-218731052309")),
			url:             "/api/v1/ticket/fcb33af4-40a3-4c82-afb1-218731052309",
			want:            http.StatusOK,
			mockedTicket:    testData,
			mockedError:     nil,
			mockedAuthorize: true,
		},
		{
			name:            "Get_Tickets_500",
			id:              uuid.Must(uuid.Parse("0e3763c6-a7ed-4532-afd7-420c5a480000")),
			url:             "/api/v1/ticket/0e3763c6-a7ed-4532-afd7-420c5a480000",
			want:            http.StatusInternalServerError,
			mockedTicket:    data.Ticket{},
			mockedError:     errors.New("db error"),
			mockedAuthorize: true,
		},
		{
			name:            "Get_Tickets_403",
			url:             "/api/v1/ticket/fcb33af4-40a3-4c82-afb1-218731052309",
			want:            http.StatusForbidden,
			mockedAuthorize: false,
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	defer func() { authorization.LoggedIn = authorization.CheckAccess }()

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			authorization.LoggedIn = func(w http.ResponseWriter, r *http.Request) bool {
				return tc.mockedAuthorize
			}
			if tc.mockedAuthorize {
				ticketMock := database.NewMockTicketRepository(mockCtrl)
				ticketMock.EXPECT().GetTicket(tc.id).Return(tc.mockedTicket,
					tc.mockedError)
				database.TicketRepo = ticketMock
			}

			rw := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tc.url, nil)
			router.ServeHTTP(rw, req)
			if rw.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rw.Code)
			}
		})
	}
}

func TestCreateTicket(t *testing.T) {
	ticket := data.Ticket{
		Place:      29,
		TicketType: "Bus",
		Discount:   "-2%",
		Price:      52.60,
		TotalPrice: 51.15,
		Name:       "Ivo",
		Surname:    "Bobul",
	}
	tests := []ListTicketTestCase{
		{
			name:            "Post_Ticket_201",
			url:             "/api/v1/ticket",
			want:            http.StatusCreated,
			mockedTicket:    ticket,
			mockedError:     nil,
			mockedAuthorize: true,
		},
		{
			name:            "Post_Tickets_500",
			url:             "/api/v1/ticket",
			want:            http.StatusInternalServerError,
			mockedTicket:    ticket,
			mockedError:     errors.New("db error"),
			mockedAuthorize: true,
		},
		{
			name:            "Post_Tickets_406",
			url:             "/api/v1/ticket",
			want:            http.StatusNotAcceptable,
			mockedTicket:    data.Ticket{},
			mockedError:     errors.New("validation failure"),
			mockedAuthorize: true,
		},
		{
			name:            "Get_Tickets_403",
			url:             "/api/v1/ticket",
			want:            http.StatusForbidden,
			mockedAuthorize: false,
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	defer func() { authorization.AdminRole = authorization.CheckAccess }()

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			authorization.AdminRole = func(w http.ResponseWriter, r *http.Request) bool {
				return tc.mockedAuthorize
			}
			if tc.name != "Post_Tickets_406" && tc.mockedAuthorize {
				ticketMock := database.NewMockTicketRepository(mockCtrl)
				ticketMock.EXPECT().CreateTicket(tc.mockedTicket).Return(tc.mockedError)
				database.TicketRepo = ticketMock
			}
			b, err := json.Marshal(tc.mockedTicket)
			if err != nil {
				t.Error(err)
			}

			rw := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, tc.url, bytes.NewBuffer(b))
			router.ServeHTTP(rw, req)
			if rw.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rw.Code)
			}
		})
	}
}

func TestUpdateTicket(t *testing.T) {
	ticket := data.Ticket{
		ID:         uuid.Must(uuid.Parse("fcb33af4-40a3-4c82-afb1-218731052309")),
		Place:      29,
		TicketType: "Bus",
		Discount:   "-2%",
		Price:      52.60,
		TotalPrice: 51.15,
		Name:       "Ivo",
		Surname:    "Bobul",
	}
	tests := []ListTicketTestCase{
		{
			name:            "Patch_Ticket_200",
			url:             "/api/v1/ticket/fcb33af4-40a3-4c82-afb1-218731052309",
			want:            http.StatusOK,
			mockedTicket:    ticket,
			mockedError:     nil,
			mockedAuthorize: true,
		},
		{
			name:            "Patch_Tickets_500",
			url:             "/api/v1/ticket/fcb33af4-40a3-4c82-afb1-218731052309",
			want:            http.StatusInternalServerError,
			mockedTicket:    ticket,
			mockedError:     errors.New("db error"),
			mockedAuthorize: true,
		},
		{
			name:            "Patch_Tickets_406",
			url:             "/api/v1/ticket/fcb33af4-40a3-4c82-afb1-218731052309",
			want:            http.StatusNotAcceptable,
			mockedTicket:    data.Ticket{},
			mockedError:     errors.New("validation failure"),
			mockedAuthorize: true,
		},
		{
			name:            "Get_Tickets_403",
			url:             "/api/v1/ticket/fcb33af4-40a3-4c82-afb1-218731052309",
			want:            http.StatusForbidden,
			mockedAuthorize: false,
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	defer func() { authorization.AdminRole = authorization.CheckAccess }()

	for _, tc := range tests {
		authorization.AdminRole = func(w http.ResponseWriter, r *http.Request) bool {
			return tc.mockedAuthorize
		}
		t.Run(tc.name, func(t *testing.T) {
			if tc.name != "Patch_Tickets_406" && tc.mockedAuthorize {
				ticketMock := database.NewMockTicketRepository(mockCtrl)
				ticketMock.EXPECT().UpdateTicket(tc.mockedTicket).Return(tc.mockedError)
				database.TicketRepo = ticketMock
			}
			b, err := json.Marshal(tc.mockedTicket)
			if err != nil {
				t.Error(err)
			}

			rw := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPatch, tc.url, bytes.NewBuffer(b))
			router.ServeHTTP(rw, req)
			if rw.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rw.Code)
			}
		})
	}
}

func TestDeleteTicket(t *testing.T) {
	tests := []ListTicketTestCase{
		{
			name:            "Delete_Ticket_200",
			id:              uuid.Must(uuid.Parse("fcb33af4-40a3-4c82-afb1-218731052309")),
			url:             "/api/v1/ticket/fcb33af4-40a3-4c82-afb1-218731052309",
			want:            http.StatusOK,
			mockedError:     nil,
			mockedAuthorize: true,
		},
		{
			name:            "Delete_Tickets_500",
			id:              uuid.Must(uuid.Parse("fcb33af4-40a3-4c82-afb1-218731052309")),
			url:             "/api/v1/ticket/fcb33af4-40a3-4c82-afb1-218731052309",
			want:            http.StatusInternalServerError,
			mockedError:     errors.New("db error"),
			mockedAuthorize: true,
		},
		{
			name:            "Get_Tickets_403",
			url:             "/api/v1/ticket/fcb33af4-40a3-4c82-afb1-218731052309",
			want:            http.StatusForbidden,
			mockedAuthorize: false,
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	defer func() { authorization.AdminRole = authorization.CheckAccess }()

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			authorization.AdminRole = func(w http.ResponseWriter, r *http.Request) bool {
				return tc.mockedAuthorize
			}
			if tc.mockedAuthorize {
				ticketMock := database.NewMockTicketRepository(mockCtrl)
				ticketMock.EXPECT().DeleteTicket(tc.id).Return(tc.mockedError)
				database.TicketRepo = ticketMock
			}
			rw := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodDelete, tc.url, nil)
			router.ServeHTTP(rw, req)
			if rw.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rw.Code)
			}
		})
	}
}
