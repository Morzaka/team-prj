package services

import (
	"github.com/go-zoo/bone"
	"team-project/services/authorization"
	"team-project/services/models"
)

//NewRouter creates a router for URL
func NewRouter() *bone.Mux {
	router := bone.New().Prefix("/api")
	subV1 := bone.New()
	router.SubRoute("/v1", subV1)
	// GetFunc, PostFunc etc ... takes http.HandlerFunc
	subV1.GetFunc("/startpage", models.GetStartFunc)
	subV1.PostFunc("/register", authorization.SignupFunc)
	subV1.PostFunc("/login", authorization.SigninFunc)
	subV1.PostFunc("/logout", authorization.LogoutFunc)
	subV1.DeleteFunc("/delete/:id", authorization.DeletePageFunc)
	subV1.PatchFunc("/update/:id", authorization.UpdatePageFunc)
	return router
}
