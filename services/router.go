package services

import (
	"team-project/services/authorization"
	"team-project/services/booking"
	"team-project/services/plane"
	"team-project/services/routing"
	"team-project/services/train"
	"team-project/services/trip"
	"team-project/swagger"

	"github.com/go-zoo/bone"
)

//NewRouter creates a router for URL
func NewRouter() *bone.Mux {
	router := bone.New().Prefix("/api")
	subV1 := bone.New()
	router.SubRoute("/v1", subV1)
	subV1.PostFunc("/register", authorization.Signup)
	subV1.PostFunc("/login", authorization.Signin)
	subV1.PostFunc("/logout", authorization.Logout)
	subV1.DeleteFunc("/user/:id", authorization.DeleteUserPage)
	subV1.PatchFunc("/user/:id", authorization.UpdateUserPage)
	subV1.GetFunc("/user/:id", authorization.GetOneUser)
	subV1.GetFunc("/users", authorization.ListAllUsers)
	subV1.GetFunc("/hello/:name", swagger.GetHello)

	// Tickets routs
	subV1.GetFunc("/tickets", booking.GetAllTickets)
	subV1.GetFunc("/ticket/:id", booking.GetOneTicket)
	subV1.PostFunc("/ticket", booking.CreateTicket)
	subV1.PatchFunc("/ticket/:id", booking.UpdateTicket)
	subV1.DeleteFunc("/ticket/:id", booking.DeleteTicket)

	// Train routes
	subV1.GetFunc("/trains", train.GetTrains)
	subV1.GetFunc("/train/:id", train.GetSingleTrain)
	subV1.PostFunc("/train", train.CreateTrain)
	subV1.PatchFunc("/train/:id", train.UpdateTrain)
	subV1.DeleteFunc("/train/:id", train.DeleteTrain)

	// Plane routes
	subV1.GetFunc("/planes", plane.GetPlanes)
	subV1.GetFunc("/plane/:id", plane.GetPlane)
	subV1.PostFunc("/plane", plane.CreatePlane)
	subV1.PatchFunc("/plane/:id", plane.UpdatePlane)
	subV1.DeleteFunc("/plane/:id", plane.DeletePlane)

	//Get Routes
	subV1.PostFunc("/routes", routing.FindPath)

	// Trip routes
	subV1.GetFunc("/trips", trip.GetTrips)
	subV1.GetFunc("/trip/:id", trip.GetTrip)
	subV1.PostFunc("/trip", trip.CreateTrip)
	subV1.PatchFunc("/trip/:id", trip.UpdateTrip)
	subV1.DeleteFunc("/trip/:id", trip.DeleteTrip)

	return router
}
