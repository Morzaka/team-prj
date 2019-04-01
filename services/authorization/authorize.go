package authorization

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-zoo/bone"
	"github.com/google/uuid"

	"team-project/database"
	"team-project/logger"
	"team-project/services/authorization/session"
	"team-project/services/data"
	"team-project/services/model"
)

//InMemorySession creates new session in memory
var InMemorySession *session.Session

//init function initializes new session
func init() {
	InMemorySession = session.NewSession()
}

//SigninFunc implements signing in
func SigninFunc(w http.ResponseWriter, r *http.Request) {
	var isRegistered = false
	var user data.Signin
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("Decoding")
		logger.Logger.Errorf("Error, %s", err)
	}
	dbpassword := database.GetUserPassword(user.Login)
	//if entered password matches the password from database than user is registered
	if model.CheckPasswordHash(user.Password, dbpassword) {
		isRegistered = true
	}
	//if user is registered than write session id for this user to cookie to tack authorized users
	if isRegistered == true {
		sessionID := InMemorySession.Init(user.Login)
		cookie := &http.Cookie{Name: user.Login,
			Value:   sessionID.String(),
			Expires: time.Now().Add(1 * time.Minute),
		}
		if cookie != nil {
			//add cookie to redis db
			_, err := database.Client.LPush(cookie.Name, cookie.Value).Result()
			if err != nil {
				fmt.Println("pushing redis")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
		http.SetCookie(w, cookie)
		w.Header().Set("content-type", "application/json")
		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			fmt.Println("Encoding")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		//else if passwords don't match then redirect user to registration page
	} else if isRegistered == false {
		w.Header().Set("content-type", "application/json")
	}
}

//SignupFunc function implements user's registration
func SignupFunc(w http.ResponseWriter, r *http.Request) {
	var user data.User
	user.ID = model.GenerateID()
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		logger.Logger.Errorf("Error, %s", err)
	}
	// get entered values from the registration form
	password, _ := model.HashPassword(user.Signin.Password)
	user.Signin.Password = password
	//add user to database and get his id
	database.AddUser(user)
	w.Header().Set("content-type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

//LogoutFunc implements logging out - deletes cookie from db
func LogoutFunc(w http.ResponseWriter, r *http.Request) {
	id, arr := len(r.Cookies())-1, r.Cookies()
	cookie := arr[id]
	sessionToken := cookie.Name
	_, err := database.Client.LRem(sessionToken, 0, cookie.Value).Result()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	cookie = &http.Cookie{
		Name:   sessionToken,
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
	w.Write([]byte("You're logged out!\n"))
}

//UpdatePageFunc deletes user
func UpdatePageFunc(w http.ResponseWriter, r *http.Request) {
	var user data.User
	id, err := uuid.Parse(bone.GetValue(r, "id"))
	if err != nil {
		logger.Logger.Errorf("Error, %s", err)
	}
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		logger.Logger.Errorf("Error, %s", err)
	}
	err = json.Unmarshal(data, &user)
	if err != nil {
		logger.Logger.Errorf("Error, %s", err)
	}
	// get entered values from the registration form
	password, _ := model.HashPassword(user.Signin.Password)
	user.Signin.Password = password
	//add user to database and get his id
	database.UpdateUser(user, id)
	w.Header().Set("content-type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

//DeletePageFunc deletes user's page
func DeletePageFunc(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(bone.GetValue(r, "id"))
	fmt.Println(id)
	if err != nil {
		logger.Logger.Errorf("Error, %s", err)
	}
	database.DeleteUser(id)
	w.Write([]byte(fmt.Sprintf("User with id %s is deleted", id)))
}
