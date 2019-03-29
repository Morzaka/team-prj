package handlers

import (
	"net/http"
	"team-project/services/database"
)

//GetStartFunc is a handler function for start page
func GetStartFunc(w http.ResponseWriter, r *http.Request) {
	var active = false
	if len(r.Cookies()) <= 0 {
		// If the cookie is not set, return an unauthorized status
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	//else get last logged in user, who hasn;t logged out
	id, arr := len(r.Cookies())-1, r.Cookies()
	cookie := arr[id]
	sessionToken := cookie.Name
	response, err := database.Client.LRange(sessionToken, 0, -1).Result()
	if err != nil {
		// If there is an error fetching from cache, return an internal server error status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	for _, v := range response {
		if v == cookie.Value {
			active = true
		}
	}
	if !active {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// Finally, return the website page to the user
	_, err = w.Write([]byte("Hello golang group-388!"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
