package authorization

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-zoo/bone"
	"github.com/google/uuid"

	"team-project/database"
	"team-project/logger"
	"team-project/services/authorization/session"
	"team-project/services/common"
	"team-project/services/data"
	"team-project/services/model"
)

//InMemorySession creates new session in memory
var InMemorySession *session.Session
var emptyResponse interface{}

//init function initializes new session
func init() {
	InMemorySession = session.NewSession()
}

//Signin implements signing in
func Signin(w http.ResponseWriter, r *http.Request) {
	var user data.Signin
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		logger.Logger.Errorf("Error, %s", err)
	}
	dbpassword, err := database.GetUserPassword(user.Login)
	if err != nil {
		common.RenderJSON(w, r, http.StatusUnauthorized, emptyResponse)
		return
	}
	//if entered password matches the password from database than user is registered
	if model.CheckPasswordHash(user.Password, dbpassword) {
		//if user is registered than write session id for this user to cookie to tack authorized users
		sessionID := InMemorySession.Init(user.Login)
		cookie := &http.Cookie{Name: user.Login,
			Value:   sessionID.String(),
			Expires: time.Now().Add(15 * time.Minute),
		}
		if cookie != nil {
			//add cookie to redis db
			_, err := database.Client.LPush(cookie.Name, cookie.Value).Result()
			if err != nil {
				common.RenderJSON(w, r, http.StatusInternalServerError, emptyResponse)
				return
			}
			//delele this session value from redis in 15 minutes
			go func() {
				time.Sleep(15 * time.Minute)
				_, err := database.Client.LRem(cookie.Name, 0, cookie.Value).Result()
				if err != nil {
					logger.Logger.Errorf("Error, %s", err)
				}
			}()
		}
		http.SetCookie(w, cookie)
		common.RenderJSON(w, r, http.StatusOK, user)
		//else if passwords don't match then render status unauthorized
	} else {
		common.RenderJSON(w, r, http.StatusUnauthorized, emptyResponse)
		return
	}
}

//Signup function implements user's registration
func Signup(w http.ResponseWriter, r *http.Request) {
	var user data.User
	user.ID = uuid.New()
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		logger.Logger.Errorf("Error, %s", err)
	}
	// get entered values from the registration form
	password, err := model.HashPassword(user.Signin.Password)
	if err != nil {
		logger.Logger.Errorf("Error, %s", err)
	}
	user.Signin.Password = password
	//add user to database and get his id
	user, err = database.AddUser(user)
	if err != nil {
		common.RenderJSON(w, r, http.StatusInternalServerError, emptyResponse)
		return
	}
	common.RenderJSON(w, r, http.StatusOK, user)
}

//Logout implements logging out - deletes cookie from db
func Logout(w http.ResponseWriter, r *http.Request) {
	id, arr := len(r.Cookies())-1, r.Cookies()
	cookie := arr[id]
	sessionToken := cookie.Name
	_, err := database.Client.LRem(sessionToken, 0, cookie.Value).Result()
	if err != nil {
		common.RenderJSON(w, r, http.StatusInternalServerError, emptyResponse)
		return
	}
	cookie = &http.Cookie{
		Name:   sessionToken,
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
	common.RenderJSON(w, r, http.StatusNoContent, emptyResponse)
}

//UpdateUserPage deletes user
func UpdateUserPage(w http.ResponseWriter, r *http.Request) {
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
	password, err := model.HashPassword(user.Signin.Password)
	if err != nil {
		logger.Logger.Errorf("Error, %s", err)
	}
	user.Signin.Password = password
	//add user to database and get his id
	err = database.UpdateUser(user, id)
	if err != nil {
		common.RenderJSON(w, r, http.StatusInternalServerError, emptyResponse)
		return
	}
	user.ID=id
	common.RenderJSON(w, r, http.StatusOK, user)
}

//DeleteUserPage deletes user's page
func DeleteUserPage(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(bone.GetValue(r, "id"))
	if err != nil {
		logger.Logger.Errorf("Error, %s", err)
	}
	err = database.DeleteUser(id)
	if err != nil {
		common.RenderJSON(w, r, http.StatusInternalServerError, emptyResponse)
		return
	}
	common.RenderJSON(w, r, http.StatusNotFound, "User was deleted successfully!")
}

//GetAllUsers makes a request to db to get all users
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := database.GetAllUsers()
	if err != nil {
		common.RenderJSON(w, r, http.StatusNoContent, users)
		return
	}
	common.RenderJSON(w, r, http.StatusOK, users)
}
