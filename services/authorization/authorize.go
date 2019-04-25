package authorization

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"

	"team-project/database"
	"team-project/logger"
	"team-project/services/common"
	"team-project/services/data"
	"team-project/services/model"

	"github.com/go-redis/redis"
	"github.com/go-zoo/bone"
	"github.com/google/uuid"
)

var (
	emptyResponse interface{}
	//RedisClient variable to refer to redis db
	RedisClient *redis.Client
	//LoggedIn variable holds the value of CheckAccess function
	LoggedIn = CheckAccess
	//AdminRole variable to refer to function CheckAdmin
	AdminRole = CheckAdmin
	//Validate variable holds the value of Validation function
	Validate = Validation
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
	dbpassword, err := database.Users.GetUserPassword(user.Login)
	if err != nil {
		common.RenderJSON(w, r, http.StatusUnauthorized, emptyResponse)
		return
	}
	//if entered password matches the password from database than user is registered
	if model.HelperModel.CheckPasswordHash(user.Password, dbpassword) {
		//if user is registered than write session id for this user to cookie to tack authorized users
		sessionID := uuid.New()
		cookie := &http.Cookie{Name: user.Login,
			Value:   sessionID.String(),
			Expires: time.Now().Add(15 * time.Minute),
		}
		//add cookie to redis db
		_, err := RedisClient.LPush(cookie.Name, cookie.Value).Result()
		if err != nil {
			common.RenderJSON(w, r, http.StatusInternalServerError, emptyResponse)
			return
		}
		//delele this session value from redis in 15 minutes
		go func() {
			time.Sleep(15 * time.Minute)
			_, err := RedisClient.LRem(cookie.Name, 0, cookie.Value).Result()
			if err != nil {
				common.RenderJSON(w, r, http.StatusInternalServerError, emptyResponse)
				return
			}
		}()
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
	users, err := database.Users.GetAllUsers()
	if err != nil {
		common.RenderJSON(w, r, http.StatusInternalServerError, emptyResponse)
		return
	}
	for _, each := range users {
		if each.Login == user.Login {
			common.RenderJSON(w, r, http.StatusNotAcceptable, "Login is not allowed")
			return
		}
	}
	// get entered values from the registration form
	password, err := model.HelperModel.HashPassword(user.Password)
	if err != nil {
		logger.Logger.Errorf("Error, %s", err)
	}
	user.Password = password
	//add user to database and get his id
	user, err = database.Users.AddUser(user)
	if err != nil {
		common.RenderJSON(w, r, http.StatusInternalServerError, emptyResponse)
		return
	}
	user.Role = "User"
	common.RenderJSON(w, r, http.StatusOK, user)
}

//Logout implements logging out - deletes cookie from db
func Logout(w http.ResponseWriter, r *http.Request) {
	id, arr := len(r.Cookies())-1, r.Cookies()
	cookie := arr[id]
	sessionToken := cookie.Name
	_, err := RedisClient.LRem(sessionToken, 0, cookie.Value).Result()
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
		common.RenderJSON(w, r, http.StatusInternalServerError, emptyResponse)
		return
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
	password, err := model.HelperModel.HashPassword(user.Password)
	if err != nil {
		logger.Logger.Errorf("Error, %s", err)
	}
	user.Password = password
	//add user to database and get his id
	affected, err := database.Users.UpdateUser(user, id)
	if err != nil {
		common.RenderJSON(w, r, http.StatusInternalServerError, emptyResponse)
		return
	}
	user.ID = id
	if affected > 0 {
		common.RenderJSON(w, r, http.StatusOK, "User was updated successfully!")
	} else {
		common.RenderJSON(w, r, http.StatusNotFound, "No rows affected!")
	}
}

//DeleteUserPage deletes user's page
func DeleteUserPage(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(bone.GetValue(r, "id"))
	if err != nil {
		common.RenderJSON(w, r, http.StatusInternalServerError, emptyResponse)
		return
	}
	affected, err := database.Users.DeleteUser(id)
	if err != nil {
		common.RenderJSON(w, r, http.StatusInternalServerError, emptyResponse)
		return
	}
	if affected > 0 {
		common.RenderJSON(w, r, http.StatusOK, "User was deleted successfully!")
	} else {
		common.RenderJSON(w, r, http.StatusNotFound, "No rows affected!")
	}
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

//GetOneUser gets user from db by id
func GetOneUser(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(bone.GetValue(r, "id"))
	if err != nil {
		common.RenderJSON(w, r, http.StatusInternalServerError, emptyResponse)
		return
	}
	user, err := database.Users.GetUser(id)
	if err != nil {
		common.RenderJSON(w, r, http.StatusNoContent, emptyResponse)
		return
	}
	common.RenderJSON(w, r, http.StatusOK, user)
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
	w.WriteHeader(http.StatusOK)
	return active
}

//CheckAdmin function checks where user is admin to give him special access to some services
func CheckAdmin(w http.ResponseWriter, r *http.Request) bool {
	if LoggedIn(w, r) {
		id, arr := len(r.Cookies())-1, r.Cookies()
		cookie := arr[id]
		role, err := database.Users.GetUserRole(cookie.Name)
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

//Validation function checks whether user password login name and surname are valid
//and are between 0 and 40 characters
func Validation(user data.User) (bool, string) {
	errMessage := ""
	var checkPass = regexp.MustCompile(`^[[:graph:]]*$`)
	var checkName = regexp.MustCompile(`^[A-Z]{1}[a-z]+$`)
	var checkLogin = regexp.MustCompile(`^[[:graph:]]*$`)
	var validPass, validName, validSurname, validLogin bool
	if len(user.Password) >= 8 && checkPass.MatchString(user.Password) {
		validPass = true
	} else {
		errMessage += "Invalid Password"
	}
	if checkName.MatchString(user.Name) && len(user.Name) < 40 {
		validName = true
	} else {
		errMessage += " Invalid Name"
	}
	if checkName.MatchString(user.Surname) && len(user.Name) < 40 {
		validSurname = true
	} else {
		errMessage += "Invalid Surname"
	}
	if checkLogin.MatchString(user.Login) && len(user.Login) < 40 {
		validLogin = true
	} else {
		errMessage += " Invalid Login"
	}
	return validName && validLogin && validPass && validSurname, errMessage
}
