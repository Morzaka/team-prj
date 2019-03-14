package authorization

import (
	"html/template"
	"net/http"
	"time"
	"golang.org/x/crypto/bcrypt"
	"github.com/satori/go.uuid"
	"gitlab.com/golang-lv-388/team-project/services/authorization/models"
	"gitlab.com/golang-lv-388/team-project/services/authorization/session"
)

var users []*models.User
var InMemorySession *session.Session
var IsAuthorized bool = false
var t time.Time

const (
	COOKIE_NAME = "sessionId"
)

func init(){
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
	tmpl, err := template.ParseFiles("gitlab.com/golang-lv-388/team-project/services/authorization/frontend/login.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	tmpl.ExecuteTemplate(w, "login", nil)
}

func SigninFunc(w http.ResponseWriter, r *http.Request) {
	var IsRegistered bool = false
	login := r.FormValue("login")
	password := r.FormValue("password")
	for _, value := range users {
		if value.Login == login {
			if checkPasswordHash(password, value.Password) {
				IsRegistered = true
				break
			}
		}
	}
	if IsRegistered == true{
		t = time.Now().Add(1 * time.Minute)
		sessionId := InMemorySession.Init(login)
		cookie := &http.Cookie{Name: COOKIE_NAME,
			Value:   sessionId,
			Expires: t,
		}
		http.SetCookie(w, cookie)
		if cookie != nil {
			if login == InMemorySession.GetUser(cookie.Value) {
				IsAuthorized = true
			}
		}
		http.Redirect(w, r, "/api/v1/startpage", 302)
	} else if IsRegistered == false {
		fmt.Println("Not registered")
		http.Redirect(w, r, "/api/v1/register", 302)
	}
}

func RegisterPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("gitlab.com/golang-lv-388/team-project/services/authorization/frontend/register.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	tmpl.ExecuteTemplate(w, "register", nil)
}

func SignupFunc(w http.ResponseWriter, r *http.Request) {
	id, _ :=uuid.NewV4()
	name:=r.FormValue("name")
	surname:=r.FormValue("surname")
	role:=r.FormValue("role")
	login := r.FormValue("login")
	passwordtmp := r.FormValue("password")
	password, _ := hashPassword(passwordtmp)
	user := models.NewUser(id, password, name, surname,login,role)
	users=append(users,user)
	http.Redirect(w, r, "/api/v1/login", 302)
}


