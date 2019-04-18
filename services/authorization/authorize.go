package authorization

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"team-project/database"
	"team-project/logger"
	"team-project/services/common"
	"team-project/services/data"
	"team-project/services/model"

	"github.com/go-redis/redis"
	"github.com/google/uuid"
)

var (
	emptyResponse interface{}
	//UserCrud is an interface to call methods from database package
	UserCrud database.UserCRUD = &database.IUser{}
	//ModelVar is an interface to call methods from model package
	ModelVar model.Model = &model.IModel{}
	//CommonVar is an interface to call methods from common package
	CommonVar common.Common = &common.ICommon{}
	//RedisClient variable to refer to redis db
	RedisClient *redis.Client
	//LoggedIn variable holds the value of CheckAccess function
	LoggedIn = CheckAccess
)

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
	dbpassword, err := UserCrud.GetUserPassword(user.Login)
	if err != nil {
		CommonVar.RenderJSON(w, r, http.StatusUnauthorized, emptyResponse)
		return
	}
	//if entered password matches the password from database than user is registered
	if ModelVar.CheckPasswordHash(user.Password, dbpassword) {
		//if user is registered than write session id for this user to cookie to tack authorized users
		sessionID := ModelVar.GenerateID()
		cookie := &http.Cookie{Name: user.Login,
			Value:   sessionID.String(),
			Expires: time.Now().Add(15 * time.Minute),
		}
		//add cookie to redis db
		_, err := RedisClient.LPush(cookie.Name, cookie.Value).Result()
		if err != nil {
			CommonVar.RenderJSON(w, r, http.StatusInternalServerError, emptyResponse)
			return
		}
		//delele this session value from redis in 15 minutes
		go func() {
			time.Sleep(15 * time.Minute)
			_, err := RedisClient.LRem(cookie.Name, 0, cookie.Value).Result()
			if err != nil {
				CommonVar.RenderJSON(w, r, http.StatusInternalServerError, emptyResponse)
				return
			}
		}()
		http.SetCookie(w, cookie)
		CommonVar.RenderJSON(w, r, http.StatusOK, user)
		//else if passwords don't match then render status unauthorized
	} else {
		CommonVar.RenderJSON(w, r, http.StatusUnauthorized, emptyResponse)
		return
	}

}

//Signup function implements user's registration
func Signup(w http.ResponseWriter, r *http.Request) {
	var user data.User
	user.ID = ModelVar.GenerateID()
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		logger.Logger.Errorf("Error, %s", err)
	}
	// get entered values from the registration form
	password, err := ModelVar.HashPassword(user.Signin.Password)
	if err != nil {
		logger.Logger.Errorf("Error, %s", err)
	}
	user.Signin.Password = password
	//add user to database and get his id
	user, err = UserCrud.AddUser(user)
	if err != nil {
		CommonVar.RenderJSON(w, r, http.StatusInternalServerError, emptyResponse)
		return
	}
	CommonVar.RenderJSON(w, r, http.StatusOK, user)
}

//Logout implements logging out - deletes cookie from db
func Logout(w http.ResponseWriter, r *http.Request) {
	id, arr := len(r.Cookies())-1, r.Cookies()
	cookie := arr[id]
	sessionToken := cookie.Name
	_, err := RedisClient.LRem(sessionToken, 0, cookie.Value).Result()
	if err != nil {
		CommonVar.RenderJSON(w, r, http.StatusInternalServerError, emptyResponse)
		return
	}
	cookie = &http.Cookie{
		Name:   sessionToken,
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
	CommonVar.RenderJSON(w, r, http.StatusNoContent, emptyResponse)
}

//UpdateUserPage deletes user
func UpdateUserPage(w http.ResponseWriter, r *http.Request) {
	var user data.User
	u, err := url.Parse(r.URL.String())
	if err != nil {
		CommonVar.RenderJSON(w, r, http.StatusInternalServerError, emptyResponse)
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
	password, err := ModelVar.HashPassword(user.Signin.Password)
	if err != nil {
		logger.Logger.Errorf("Error, %s", err)
	}
	user.Signin.Password = password
	//add user to database and get his id
	err = UserCrud.UpdateUser(user, id)
	if err != nil {
		CommonVar.RenderJSON(w, r, http.StatusInternalServerError, emptyResponse)
		return
	}
	user.ID = id
	CommonVar.RenderJSON(w, r, http.StatusOK, user)
}

//DeleteUserPage deletes user's page
func DeleteUserPage(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		CommonVar.RenderJSON(w, r, http.StatusInternalServerError, emptyResponse)
		return
	}
	id, err := uuid.Parse(u.Query().Get("id"))
	if err != nil {
		logger.Logger.Errorf("Error, %s", err)
	}
	err = UserCrud.DeleteUser(id)
	if err != nil {
		CommonVar.RenderJSON(w, r, http.StatusInternalServerError, emptyResponse)
		return
	}
	CommonVar.RenderJSON(w, r, http.StatusNotFound, "User was deleted successfully!")
}

//ListAllUsers makes a request to db to get all users
func ListAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := database.Users.GetAllUsers()
	if err != nil {
		common.RenderJSON(w, r, http.StatusNoContent, emptyResponse)
		return
	}
	common.RenderJSON(w, r, http.StatusOK, users)
}

//CheckAccess checks whether user is logged in to give him access to services
func CheckAccess(w http.ResponseWriter, r *http.Request) bool {
	var active bool
	if len(r.Cookies()) <= 0 {
		// If the cookie is not set, return an unauthorized status
		w.WriteHeader(http.StatusUnauthorized)
		return active
	}
	id, arr := len(r.Cookies())-1, r.Cookies()
	cookie := arr[id]
	sessionToken := cookie.Name
	response, err := RedisClient.LRange(sessionToken, 0, -1).Result()
	if err != nil {
		// If there is an error fetching from cache, return an internal server error status
		w.WriteHeader(http.StatusInternalServerError)
		return active
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
	if LoggedIn(w, r) {
		id, arr := len(r.Cookies())-1, r.Cookies()
		cookie := arr[id]
		role, err := UserCrud.GetUserRole(cookie.Name)
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
