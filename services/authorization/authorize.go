package authorization

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
	"team-project/services/authorization/models"
	"team-project/services/authorization/session"
	"team-project/services/database"
	"time"
)

var InMemorySession *session.Session

const (
	CookieName = "sessionId"
)

func init() {
	InMemorySession = session.NewSession()
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil

}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("team-project/services/authorization/frontend/login.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	log.Fatal(tmpl.ExecuteTemplate(w, "login", nil))
}

func SigninFunc(w http.ResponseWriter, r *http.Request) {
	var IsAuthorized bool
	var t time.Time
	var IsRegistered = false
	login := r.FormValue("login")
	password := r.FormValue("password")
	dbpassword := database.GetUser(login)
	if checkPasswordHash(password, dbpassword) {
		IsRegistered = true
	}
	if IsRegistered == true {
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
				log.Println("User is autorized", IsAuthorized)
			}
		}
		http.Redirect(w, r, "/api/v1/startpage", 302)
	} else if IsRegistered == false {
		fmt.Println("Not registered")
		http.Redirect(w, r, "/api/v1/register", 302)
	}
}

func RegisterPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("team-project/services/authorization/frontend/register.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	log.Fatal(tmpl.ExecuteTemplate(w, "register", nil))
}

func SignupFunc(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	surname := r.FormValue("surname")
	role := r.FormValue("role")
	login := r.FormValue("login")
	passwordtmp := r.FormValue("password")
	password, _ := hashPassword(passwordtmp)
	user := models.NewUser(password, name, surname, login, role)
	id := database.AddUser(user)
	log.Println("You are registered with id :", id)
	http.Redirect(w, r, "/api/v1/login", 302)
}
