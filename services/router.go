package services

import (
	"team-project/services/authorization"
	"team-project/services/booking"
	"team-project/services/model"

	"github.com/go-zoo/bone"
)

//NewRouter creates a router for URL
func NewRouter() *bone.Mux {
	router := bone.New().Prefix("/api")
	subV1 := bone.New()
	router.SubRoute("/v1", subV1)
	// GetFunc, PostFunc etc ... takes http.HandlerFunc
	subV1.GetFunc("/startpage", model.GetStartFunc)
	subV1.PostFunc("/register", authorization.SignupFunc)
	subV1.PostFunc("/login", authorization.SigninFunc)
	subV1.PostFunc("/logout", authorization.LogoutFunc)
	subV1.DeleteFunc("/user/:id", authorization.DeletePageFunc)
	subV1.PatchFunc("/user/:id", authorization.UpdatePageFunc)

	// Tickets routs
	subV1.GetFunc("/tickets", booking.GetAllTickets)
	subV1.GetFunc("/ticket/:id", booking.GetOneTicket)
	subV1.PostFunc("/ticket", booking.CreateTicket)
	subV1.PatchFunc("/ticket/:id", booking.UpdateTicket)
	subV1.DeleteFunc("/ticket/:id", booking.DeleteTicket)

	return router
}
