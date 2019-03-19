package authorization

import (
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
	"team-project/services/authorization/models"
	"team-project/services/authorization/session"
	"time"
	"team-project/services/database"
	"os"
	"path/filepath"
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
//LoginPage function loads html form for logging in
func LoginPage(w http.ResponseWriter, r *http.Request) {
	cwd, err := os.Getwd()
        if err != nil {
                http.Error(w, err.Error(), 400)
                return
        }
	tmpl, err := template.ParseFiles(filepath.Join(cwd, "/services/authorization/frontend/login.html"))
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	err=tmpl.ExecuteTemplate(w, "login", nil)
	if err != nil {
                http.Error(w, err.Error(), 400)
                return
        }
}

//SigninFunc implements signing in
func SigninFunc(w http.ResponseWriter, r *http.Request) {
	var IsAuthorized bool
	var t time.Time
	var isRegistered = false
	login := r.FormValue("login")
	password := r.FormValue("password")
	dbpassword := database.GetUserPassword(login)
	//if entered password matches the password from database than user is registered
	if checkPasswordHash(password, dbpassword) {
		isRegistered = true
	}
	//if user is registered than write session id for this user to cookie to tack authorized users
	if isRegistered == true {
		t = time.Now().Add(1 * time.Minute)
		sessionId := InMemorySession.Init(login)
		cookie := &http.Cookie{Name: CookieName,
			Value:   sessionId,
			Expires: t,
		}
		http.SetCookie(w, cookie)
		if cookie != nil {
			if login == InMemorySession.GetUser(cookie.Value) {
				IsAuthorized = true
				log.Println("User is authorized", IsAuthorized)
			}
		}
		log.Println("cookie: ", cookie)
		http.Redirect(w, r, "/api/v1/startpage", 302)
	//else if passwords don't match then redirect user to registration page
	} else if isRegistered == false {
		log.Println("Not registered")
		http.Redirect(w, r, "/api/v1/register", 302)
	}
}
//RegisterPage function loads html registration form
func RegisterPage(w http.ResponseWriter, r *http.Request) {
	cwd, err := os.Getwd()
	if err != nil {
                http.Error(w, err.Error(), 400)
                return
        }
	tmpl, err := template.ParseFiles(filepath.Join(cwd, "/services/authorization/frontend/register.html"))
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	err=tmpl.ExecuteTemplate(w, "register", nil)
	if err != nil {
                http.Error(w, err.Error(), 400)
                return
        }
}
//SignupFunc function implements user's registration
func SignupFunc(w http.ResponseWriter, r *http.Request) {
	// get entered values from the registration form 
	name := r.FormValue("name")
	surname := r.FormValue("surname")
	role := r.FormValue("role")
	login := r.FormValue("login")
	passwordtmp := r.FormValue("password")
	password, _ := hashPassword(passwordtmp)
	//create user with received data 
	user := models.NewUser(password, name, surname, login, role)
	//add user to database and get his id
	id := database.AddUser(user)
	log.Println("You are registered with id :",id)
	//redirect registered user to log in page
	http.Redirect(w, r, "/api/v1/login", 302)
}

