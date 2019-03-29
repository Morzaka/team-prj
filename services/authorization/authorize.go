package authorization

import (
	"encoding/json"
	"fmt"
	"github.com/go-zoo/bone"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"team-project/services/authorization/session"
	"team-project/services/database"
	"team-project/services/models"
	"time"
)

//InMemorySession creates new session in memory
var InMemorySession *session.Session

//init function initializes new session
func init() {
	InMemorySession = session.NewSession()
}

//hashPassword function hashes user's password
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//checkPasswordHash function valides user's password
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil

}

//SigninFunc implements signing in
func SigninFunc(w http.ResponseWriter, r *http.Request) {
	var IsAuthorized bool
	var isRegistered = false
	var user models.Signin
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(data, &user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	dbpassword := database.GetUserPassword(user.Login)
	//if entered password matches the password from database than user is registered
	if checkPasswordHash(user.Password, dbpassword) {
		isRegistered = true
	}
	//if user is registered than write session id for this user to cookie to tack authorized users
	if isRegistered == true {
		sessionID := InMemorySession.Init(user.Login)
		cookie := &http.Cookie{Name: user.Login,
			Value:   sessionID,
			Expires: time.Now().Add(1 * time.Minute),
		}
		if cookie != nil {
			if user.Login == InMemorySession.GetUser(cookie.Value) {
				IsAuthorized = true
				log.Println("User is authorized", IsAuthorized)
				//add cookie to redis db
				_, err := database.Client.LPush(cookie.Name, cookie.Value).Result()
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
			}
		}
		http.SetCookie(w, cookie)
		output, err := json.Marshal(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("content-type", "application/json")
		w.Write(output)
		//else if passwords don't match then redirect user to registration page
	} else if isRegistered == false {
		w.Write([]byte("Something went wrong"))
	}
}

//SignupFunc function implements user's registration
func SignupFunc(w http.ResponseWriter, r *http.Request) {
	var user models.User
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(data, &user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// get entered values from the registration form
	password, _ := hashPassword(user.Signin.Password)
	user.Signin.Password = password
	//add user to database and get his id
	id := database.AddUser(user)
	w.Write([]byte(fmt.Sprintf("You are registered with id %d", id)))
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
	var user models.User
	id, err := strconv.Atoi(bone.GetValue(r, "id"))
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(data, &user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// get entered values from the registration form
	password, _ := hashPassword(user.Signin.Password)
	user.Signin.Password = password
	//add user to database and get his id
	database.UpdateUser(user, id)
	w.Write([]byte(fmt.Sprintf("User with id %d is updated", id)))
}

//DeletePageFunc deletes user's page
func DeletePageFunc(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(bone.GetValue(r, "id"))
	if err != nil {
		log.Fatal(err)
	}
	database.DeleteUser(id)
	w.Write([]byte(fmt.Sprintf("User with id %d is deleted", id)))
}
