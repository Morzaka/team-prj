package authorization

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/google/uuid"

	"team-project/database"
	"team-project/logger"
	"team-project/services/authorization/session"
	"team-project/services/common"
	"team-project/services/data"
	"team-project/services/model"
)

var (
	//InMemorySession creates new session in memory
	InMemorySession *session.Session
	emptyResponse   interface{}
	//LoggedIn variable holds value of CheckAccess function
	LoggedIn = CheckAccess
	//SessionID variable holds the value of InMemorySession.Init function
	SessionID uuid.UUID
	//AddUser variable holds the value of database.AddUser function
	AddUser = database.AddUser
	//UpdateUser variable holds the value of database.UpdateUser function
	UpdateUser = database.UpdateUser
	//DeleteUser variable holds the value of Database.DeleteUser function
	DeleteUser = database.DeleteUser
	//GetAllUsers variable holds the value of database.GetAllUsers function
	GetAllUsers = database.GetAllUsers
	//GenerateID variable holds the value of model.GenerateID function
	GenerateID = model.GenerateID
	//HashPassword variable holds the value of model.HashPassword function
	HashPassword = model.HashPassword
	//RenderJSON variable holds the value of common.RenderJSON function
	RenderJSON = common.RenderJSON
	//GetUserPassword variable holds the value of database.GetUserPassword function
	GetUserPassword = database.GetUserPassword
	//GetUserRole variable holds the value of database.GetUserRole function
	GetUserRole = database.GetUserRole
	//CheckPasswordHash variable holds the value ofmodel.CheckPasswordHash function
	CheckPasswordHash = model.CheckPasswordHash
	//RedisClient holds the value of redis client
	RedisClient = database.Client
)

//init function initializes new session
func init() {
	InMemorySession = session.NewSession()
}

//Signin implements signing in
func Signin(w http.ResponseWriter, r *http.Request) {
	if RedisClient == nil {
		RedisClient = database.Client
	}
	var user data.Signin
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		logger.Logger.Errorf("Error, %s", err)
	}
	dbpassword, err := GetUserPassword(user.Login)
	if err != nil {
		common.RenderJSON(w, r, http.StatusUnauthorized, emptyResponse)
		return
	}
	//if entered password matches the password from database than user is registered
	if CheckPasswordHash(user.Password, dbpassword) {
		//if user is registered than write session id for this user to cookie to tack authorized users
		SessionID := InMemorySession.Init(user.Login)
		cookie := &http.Cookie{Name: user.Login,
			Value:   SessionID.String(),
			Expires: time.Now().Add(15 * time.Minute),
		}
		//add cookie to redis db
		_, err := RedisClient.LPush(cookie.Name, cookie.Value).Result()
		if err != nil {
			RenderJSON(w, r, http.StatusInternalServerError, emptyResponse)
			return
		}
		//delele this session value from redis in 15 minutes
		go func() {
			time.Sleep(15 * time.Minute)
			_, err := RedisClient.LRem(cookie.Name, 0, cookie.Value).Result()
			if err != nil {
				RenderJSON(w, r, http.StatusInternalServerError, emptyResponse)
				return
			}
		}()
		http.SetCookie(w, cookie)
		RenderJSON(w, r, http.StatusOK, user)
		//else if passwords don't match then render status unauthorized
	} else {
		RenderJSON(w, r, http.StatusUnauthorized, emptyResponse)
		return
	}
}

//Signup function implements user's registration
func Signup(w http.ResponseWriter, r *http.Request) {
	var user data.User
	user.ID = GenerateID()
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		logger.Logger.Errorf("Error, %s", err)
	}
	// get entered values from the registration form
	password, err := HashPassword(user.Signin.Password)
	if err != nil {
		logger.Logger.Errorf("Error, %s", err)
	}
	user.Signin.Password = password
	//add user to database and get his id
	user, err = AddUser(user)
	if err != nil {
		RenderJSON(w, r, http.StatusInternalServerError, emptyResponse)
		return
	}
	RenderJSON(w, r, http.StatusOK, user)
}

//Logout implements logging out - deletes cookie from db
func Logout(w http.ResponseWriter, r *http.Request) {
	id, arr := len(r.Cookies())-1, r.Cookies()
	cookie := arr[id]
	sessionToken := cookie.Name
	_, err := RedisClient.LRem(sessionToken, 0, cookie.Value).Result()
	if err != nil {
		RenderJSON(w, r, http.StatusInternalServerError, emptyResponse)
		return
	}
	cookie = &http.Cookie{
		Name:   sessionToken,
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
	RenderJSON(w, r, http.StatusNoContent, emptyResponse)
}

//UpdateUserPage deletes user
func UpdateUserPage(w http.ResponseWriter, r *http.Request) {
	var user data.User
	u, err := url.Parse(r.URL.String())
	if err != nil {
		RenderJSON(w, r, http.StatusInternalServerError, emptyResponse)
		return
	}
	id, err := uuid.Parse(u.Query().Get("id"))
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
	password, err := HashPassword(user.Signin.Password)
	if err != nil {
		logger.Logger.Errorf("Error, %s", err)
	}
	user.Signin.Password = password
	//add user to database and get his id
	err = UpdateUser(user, id)
	if err != nil {
		RenderJSON(w, r, http.StatusInternalServerError, emptyResponse)
		return
	}
	user.ID = id
	RenderJSON(w, r, http.StatusOK, user)
}

//DeleteUserPage deletes user's page
func DeleteUserPage(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		RenderJSON(w, r, http.StatusInternalServerError, emptyResponse)
		return
	}
	id, err := uuid.Parse(u.Query().Get("id"))
	if err != nil {
		logger.Logger.Errorf("Error, %s", err)
	}
	err = DeleteUser(id)
	if err != nil {
		RenderJSON(w, r, http.StatusInternalServerError, emptyResponse)
		return
	}
	RenderJSON(w, r, http.StatusNotFound, "User was deleted successfully!")
}

//ListAllUsers makes a request to db to get all users
func ListAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := GetAllUsers()
	if err != nil {
		RenderJSON(w, r, http.StatusNoContent, users)
		return
	}
	RenderJSON(w, r, http.StatusOK, users)
}

//CheckAccess checks whether user is logged in to give him access to services
func CheckAccess(w http.ResponseWriter, r *http.Request) bool {
	var active = false
	if len(r.Cookies()) <= 0 {
		// If the cookie is not set, return an unauthorized status
		w.WriteHeader(http.StatusUnauthorized)
		return false
	}
	id, arr := len(r.Cookies())-1, r.Cookies()
	cookie := arr[id]
	sessionToken := cookie.Name
	response, err := RedisClient.LRange(sessionToken, 0, -1).Result()
	if err != nil {
		// If there is an error fetching from cache, return an internal server error status
		w.WriteHeader(http.StatusInternalServerError)
		return false
	}
	for _, v := range response {
		if v == cookie.Value {
			active = true
		}
	}
	return active
}

//CheckAdmin function checks where user is admin to give him special access to some services
func CheckAdmin(w http.ResponseWriter, r *http.Request) bool {
	in := LoggedIn(w, r)
	if in {
		id, arr := len(r.Cookies())-1, r.Cookies()
		cookie := arr[id]
		role, err := GetUserRole(cookie.Name)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return false
		}
		if role != "Admin" {
			return false
		}
		return true
	}
	return false
}
