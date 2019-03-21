package authorization

import (
	"golang.org/x/crypto/bcrypt"
	"fmt"
	"log"
	"net/http"
	"team-project/services/authorization/models"
	"team-project/services/authorization/session"
	"time"
	"team-project/services/database"
	"io/ioutil"
	"encoding/json"
)

var InMemorySession *session.Session

const (
	CookieName = "sessionId"
)
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
	var t time.Time
	var isRegistered = false
	var user models.Signin
	data,err:=ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err!=nil{
		http.Error(w, err.Error(), 500)
		return
	}
	err=json.Unmarshal(data,&user)
	if err!=nil{
		http.Error(w, err.Error(), 500)
		return
	}
	dbpassword := database.GetUserPassword(user.Login)
	//if entered password matches the password from database than user is registered
	if checkPasswordHash(user.Password, dbpassword) {
		isRegistered = true
	}
	//if user is registered than write session id for this user to cookie to tack authorized users
	if isRegistered == true {
		t = time.Now().Add(1 * time.Minute)
		sessionId := InMemorySession.Init(user.Login)
		cookie := &http.Cookie{Name: CookieName,
			Value:   sessionId,
			Expires: t,
		}
		http.SetCookie(w, cookie)
		if cookie != nil {
			if user.Login == InMemorySession.GetUser(cookie.Value) {
				IsAuthorized = true
				log.Println("User is authorized", IsAuthorized)
			}
		}
		output,err:=json.Marshal(user)
		if err!=nil{
			http.Error(w, err.Error(), 500)
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
	if err!=nil{
		http.Error(w, err.Error(), 500)
                return
	}
        err=json.Unmarshal(data, &user)
	if err!=nil{
		http.Error(w, err.Error(), 500)
                return
	}
	// get entered values from the registration form
	password, _ := hashPassword(user.Signin.Password)
	user.Signin.Password=password
	//add user to database and get his id
	id := database.AddUser(user)
	w.Write([]byte(fmt.Sprintf("You are registered with id %d", id)))
	//redirect registered user to log in page
}

