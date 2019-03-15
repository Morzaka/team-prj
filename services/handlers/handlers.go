package handlers

import (
	"net/http"
)

//GetStartFunc is a handler function for start page
func GetStartFunc(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello golang group-388"))
	if err != nil {
		panic(err)
	}
}
