package services

import (
	"github.com/go-zoo/bone"
	"team-project/services/authorization"
	"team-project/services/model"
	"team-project/swagger"
)

//NewRouter creates a router for URL
func NewRouter() *bone.Mux {
	router := bone.New().Prefix("/api")
	subV1 := bone.New()
	router.SubRoute("/v1", subV1)
	// GetFunc, PostFunc etc ... takes http.HandlerFunc
	subV1.GetFunc("/startpage", model.GetStart)
	subV1.PostFunc("/register", authorization.Signup)
	subV1.PostFunc("/login", authorization.Signin)
	subV1.PostFunc("/logout", authorization.Logout)
	subV1.DeleteFunc("/user/:id", authorization.DeleteUserPage)
	subV1.PatchFunc("/user/:id", authorization.UpdateUserPage)
	subV1.GetFunc("/hello/:name", swagger.GetHello)
	return router
}
