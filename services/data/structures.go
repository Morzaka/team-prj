package data

import "github.com/google/uuid"

//User struct contains user data
type User struct {
	ID       uuid.UUID `json:"id"`
	Login    string    `json:"login"`
	Password string    `json:"password"`
	Name     string    `json:"name"`
	Surname  string    `json:"surname"`
	Email    string    `json:"email"`
	Role     string    `json:"role"`
}

//Signin contains data for logging in
type Signin struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

//Trip database structure
type Trip struct {
	TripID             uuid.UUID `json:"Trip_id"`
	TripName           string    `json:"Trip_name"`
	TripTicketID       uuid.UUID `json:"Trip_ticket_id"`
	TripReturnTicketID uuid.UUID `json:"Trip_return_ticket_id"`
	TotalTripPrice     float32   `json:"Total_trip_price"`
}

//Train is Model
type Train struct {
	ID            uuid.UUID `json:"id"`
	DepartureCity string    `json:"DepartureCity"`
	ArrivalCity   string    `json:"ArrivalCity"`
	DepartureTime string    `json:"DepartureTime"`
	DepartureDate string    `json:"DepartureDate"`
	ArrivalTime   string    `json:"ArrivalTime"`
	ArrivalDate   string    `json:"ArrivalDate"`
}

// Plane struct contains plane data
type Plane struct {
	ID            uuid.UUID `json:"id"`
	DepartureCity string    `json:"departureCity"`
	ArrivalCity   string    `json:"arrivalCity"`
}

//Routes contains result of algorithm
type Routes struct {
	RouteID  string   `json:"route"`
	Stations Stations `json:"stations"`
}

//Stations contains departure and arrival stations
type Stations struct {
	StartRoute string `json:"start"`
	EndRoute   string `json:"end"`
}

//RouteStruct is a struct of Routes
type RouteStruct struct {
	Index  string   `json:"index"`
	Routes []string `json:"routes"`
}

// RouteStructs is a slice of Struct of Routes
type RouteStructs []RouteStruct

//Ticket datastructure
type Ticket struct {
	ID         uuid.UUID `json:"id"`
	TrainID    uuid.UUID `json:"train_id"`
	PlaneID    uuid.UUID `json:"plane_id"`
	UserID     uuid.UUID `json:"user_id"`
	Place      int       `json:"place"`
	TicketType string    `json:"ticket_type"`
	Discount   string    `json:"discount"`
	Price      float32   `json:"price"`
	TotalPrice float32   `json:"total_price"`
	Name       string    `json:"name"`
	Surname    string    `json:"surname"`
	//From_place     string    `json:"from_place"`
	//Departure_date time.Time `json:"departure_date"`
	//Departure_time time.Time `json:"departure_time"`
	//To_place       string    `json:"to_place"`
	//Arrival_date   time.Time `json:"arrival_date"`
	//Arrival_time   time.Time `json:"arrival_time"`
}
