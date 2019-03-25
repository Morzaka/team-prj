package handlers

import (
	"team-project/services/database"
	"net/http"
)

//GetStartFunc is a handler function for start page
func GetStartFunc(w http.ResponseWriter, r *http.Request) {
	if len(r.Cookies())<=0{
                // If the cookie is not set, return an unauthorized status
                w.WriteHeader(http.StatusUnauthorized)
                return
	} else {
		//else get last logged in user, who hasn;t logged out
		id,arr:=len(r.Cookies())-1, r.Cookies()
		cookie := arr[id]
		sessionToken := cookie.Name
		response, err := database.GetRedisValue(sessionToken)
		if err != nil {
		// If there is an error fetching from cache, return an internal server error status
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if response != cookie.Value {
		// If the session token is not present in db, return an unauthorized error
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	}
	// Finally, return the website page to the user
	_, err := w.Write([]byte("Hello golang group-388"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
