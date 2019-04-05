package swagger

import (
	"log"
	"net/http"

	"github.com/go-zoo/bone"
)

//GetHello is a handler function for swagger testing
func GetHello(w http.ResponseWriter, r *http.Request) {
	log.Println("Responsing to /hello request")
	log.Println(r.UserAgent())
	// Get the value of the "name" parameters.
	val := bone.GetValue(r, "name")
	_, err := w.Write([]byte("Hello " + val))
	if err != nil {
		panic(err)
	}
}
