package services

<<<<<<< HEAD
import(
	"log"
        "net/http"
	"./services/handlers"
        "./services/authorization"
=======
import (
	"./authorization"
	"./handlers"
>>>>>>> 1aa5d329f6abd061e71950a90bdb990a20b080c3
	"github.com/go-zoo/bone"
)

//NewRouter creates a router for URL
func NewRouter() *bone.Mux {
	router := bone.New().Prefix("/api")
	subV1 := bone.New()
	router.SubRoute("/v1", subV1)
	// GetFunc, PostFunc etc ... takes http.HandlerFunc
	subV1.GetFunc("/startpage", handlers.GetStartFunc)
	subV1.PostFunc("/signup", authorization.SignupFunc)
	subV1.PostFunc("/signin", authorization.SigninFunc)
	subV1.GetFunc("/login", authorization.LoginPage)
	subV1.GetFunc("/register", authorization.RegisterPage)
	return router
}
