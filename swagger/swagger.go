package swagger

import (
	"log"
	"net/http"

	"github.com/go-zoo/bone"
)

// The purpose of this application is to test go-swagger in a simple GET request.
//
//     Schemes: http
//     Host: localhost:8080
//     Version: 0.0.1
//     Contact: Bohdan<bohdansydor2015@gmail.com>
//
//     Consumes:
//     - text/plain
//
//     Produces:
//     - text/plain
//
// swagger:meta

//GetHello is a handler function for swagger testing
func GetHello(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /hello/{name} hello Hello
	//
	// Returns a simple Hello message
	// ---
	// consumes:
	// - text/plain
	// produces:
	// - text/plain
	// parameters:
	// - name: name
	//   in: path
	//   description: Name to be returned.
	//   required: true
	//   type: string
	// responses:
	//   '200':
	//     description: The hello message
	//     type: string

	log.Println("Responsing to /hello request")
	log.Println(r.UserAgent())

	// Get the value of the "name" parameters.
	val := bone.GetValue(r, "name")

	_, err := w.Write([]byte("Hello " + val))
	if err != nil {
		panic(err)
	}
}
