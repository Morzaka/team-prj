package services

import(
	"log"
        "net/http"
	"gitlab.com/golang-lv-388/team-project/services/handlers"
        "github.com/go-zoo/bone"
)

//NewRouter creates a router for URL
func NewRouter()*bone.Mux {
        router := bone.New().Prefix("/api")
        subV1 := bone.New()
        router.SubRoute("/v1", subV1)
         // GetFunc, PostFunc etc ... takes http.HandlerFunc
        subV1.GetFunc("/startpage", handlers.GetStartFunc)
	return router
}

