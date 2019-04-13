package authorization

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis"
	"github.com/google/uuid"

	"team-project/database"
	"team-project/services/common"
	"team-project/services/data"
	"team-project/services/model"
)

// mockRedis returns client connected to fake Redis server
func mockRedis() *redis.Client {
	mr, err := miniredis.Run()
	if err != nil {
		log.Fatal(err)
	}
	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})
	return client
}

//TestSignin tests function Signin
func TestSignin(t *testing.T) {
	user := data.Signin{Login: "golang", Password: "golang"}
	id := uuid.Must(uuid.Parse("08307904-f18e-4fb8-9d18-29cfad38ffaf"))
	jsonBody, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}
	r, err := http.NewRequest("POST", "/api/v1/login", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	// set back original functions at the end of test
	defer func() {
		RenderJSON = common.RenderJSON
		GetUserPassword = database.GetUserPassword
		CheckPasswordHash = model.CheckPasswordHash
		RedisClient = database.Client
	}()
	for i := 0; i < 3; i++ {
		//mock original functions
		RedisClient = mockRedis()
		SessionInit = func(string) uuid.UUID {
			return id
		}
		RenderJSON = func(w http.ResponseWriter, r *http.Request, status int, response interface{}) {}
		switch i {
		//case when there was no error getting password from db to compare with the password entered
		case 0:
			GetUserPassword = func(login string) (string, error) {
				return login, nil
			}
			CheckPasswordHash = func(password, hash string) bool {
				return true
			}
			http.HandlerFunc(Signin).ServeHTTP(w, r)
		//case when there was an error getting password from db to compare with the password entered
		case 1:
			r.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBody))
			GetUserPassword = func(login string) (string, error) {
				return login, errors.New("Error")
			}
			http.HandlerFunc(Signin).ServeHTTP(w, r)
		//case when password from db and password entered do not match
		case 2:
			r.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBody))
			GetUserPassword = func(login string) (string, error) {
				return login, nil
			}
			CheckPasswordHash = func(password, hash string) bool {
				return false
			}
			http.HandlerFunc(Signin).ServeHTTP(w, r)
		}
	}
}

//TestSignup tests function Signup
func TestSignup(t *testing.T) {
	user := data.User{Signin: data.Signin{Login: "oks", Password: "oks"}, Name: "Oksana", Surname: "Zhykina", Role: "User"}
	jsonBody, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}
	r, err := http.NewRequest("POST", "/api/v1/register", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	defer func() {
		AddUser = database.AddUser
		HashPassword = model.HashPassword
		GenerateID = model.GenerateID
		RenderJSON = common.RenderJSON
	}()
	for i := 0; i < 2; i++ {
		HashPassword = func(password string) (string, error) {
			return "$2a$14$MA.GufeWJj7IryAoAgd8BeuRphle78ubdgqaPFPpjG9GzbxEk7kKu", nil
		}
		GenerateID = func() uuid.UUID {
			return uuid.New()
		}
		RenderJSON = func(w http.ResponseWriter, r *http.Request, status int, response interface{}) {}
		switch i {
		//case when there was no error while adding user to db
		case 0:
			AddUser = func(user data.User) (data.User, error) {
				return user, nil
			}
			http.HandlerFunc(Signup).ServeHTTP(w, r)
		//case when there was an error adding user to db
		case 1:
			//renew request body after being read in case 0
			r.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBody))
			AddUser = func(user data.User) (data.User, error) {
				return user, errors.New("Error")
			}
			http.HandlerFunc(Signup).ServeHTTP(w, r)
		}
	}
}

//TestLogout tests Logout function
func TestLogout(t *testing.T) {
	r, err := http.NewRequest("POST", "/api/v1/logout", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	cookie := &http.Cookie{Name: "Cookie",
		Value: "expected",
	}
	http.SetCookie(w, cookie)
	// Copy the Cookie over to a new Request
	r = &http.Request{Header: http.Header{"Cookie": w.HeaderMap["Set-Cookie"]}}
	RedisClient = mockRedis()
	_, err = RedisClient.LPush(cookie.Name, cookie.Value).Result()
	if err != nil {
		t.Fatal(err)
	}
	RenderJSON = func(w http.ResponseWriter, r *http.Request, status int, response interface{}) {}
	defer func() {
		RedisClient = database.Client
		RenderJSON = common.RenderJSON
	}()
	http.HandlerFunc(Logout).ServeHTTP(w, r)
	//check whether cookie was deleted successfully
	if cookie, err := r.Cookie(cookie.Name); err != nil {
		if cookie.MaxAge >= 0 {
			t.Error("Users hasn't logged out successfully")
		}
		t.Fatal(err)
	}
}

//TestUpdateUserPage tests UpdateUserPage function
func TestUpdateUserPage(t *testing.T) {
	user := data.User{Signin: data.Signin{Login: "oks", Password: "oks"}, Name: "Oksana", Surname: "Zhykina", Role: "User"}
	jsonBody, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}
	r, err := http.NewRequest("PATCH", "/api/v1/user/:id=", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}
	q := r.URL.Query()
	q.Add("id", "61c364d9-591a-4879-a9fb-79ae67945d38")
	r.URL.RawQuery = q.Encode()
	w := httptest.NewRecorder()
	defer func() {
		UpdateUser = database.UpdateUser
		HashPassword = model.HashPassword
		RenderJSON = common.RenderJSON
	}()
	for i := 0; i < 2; i++ {
		HashPassword = func(pswd string) (string, error) {
			return "$2a$14$MA.GufeWJj7IryAoAgd8BeuRphle78ubdgqaPFPpjG9GzbxEk7kKu", nil
		}
		RenderJSON = func(w http.ResponseWriter, r *http.Request, status int, response interface{}) {}
		switch i {
		//case when there was no error while updating user in db
		case 0:
			UpdateUser = func(user data.User, id uuid.UUID) error {
				return nil
			}
			http.HandlerFunc(UpdateUserPage).ServeHTTP(w, r)
		//case when there was an error while updating user in db
		case 1:
			r.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBody))
			UpdateUser = func(user data.User, id uuid.UUID) error {
				return errors.New("Error")
			}
			http.HandlerFunc(UpdateUserPage).ServeHTTP(w, r)
		}
	}
}

//TestDeleteUser tests DeleteUser function
func TestDeleteUser(t *testing.T) {
	r, err := http.NewRequest("DELETE", "/api/v1/user/:id=", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := r.URL.Query()
	q.Add("id", "61c364d9-591a-4879-a9fb-79ae67945d38")
	r.URL.RawQuery = q.Encode()
	w := httptest.NewRecorder()
	defer func() {
		DeleteUser = database.DeleteUser
		RenderJSON = common.RenderJSON
	}()
	for i := 0; i < 2; i++ {
		RenderJSON = func(w http.ResponseWriter, r *http.Request, status int, response interface{}) {}
		switch i {
		//case when there was no error while deleting user from db
		case 0:
			DeleteUser = func(id uuid.UUID) error {
				return nil
			}
			http.HandlerFunc(DeleteUserPage).ServeHTTP(w, r)
		//case when there was an error while deleting user from db
		case 1:
			DeleteUser = func(id uuid.UUID) error {
				return errors.New("Error")
			}
			http.HandlerFunc(DeleteUserPage).ServeHTTP(w, r)
		}
	}
}

//TestListAllUsers tests ListAllUsers function
func TestListAllUsers(t *testing.T) {
	r, err := http.NewRequest("GET", "/api/v1/users", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	defer func() {
		GetAllUsers = database.GetAllUsers
		RenderJSON = common.RenderJSON
	}()
	for i := 0; i < 2; i++ {
		RenderJSON = func(w http.ResponseWriter, r *http.Request, status int, response interface{}) {}
		switch i {
		//case when there was no error while getting all users from db
		case 0:
			GetAllUsers = func() ([]data.User, error) {
				return []data.User{}, nil
			}
			http.HandlerFunc(ListAllUsers).ServeHTTP(w, r)
		//case when there was an error while getting all users from db
		case 1:
			GetAllUsers = func() ([]data.User, error) {
				return []data.User{}, errors.New("Error")
			}
			http.HandlerFunc(ListAllUsers).ServeHTTP(w, r)
		}
	}
}

//TestCheckAccess tests CheckAccess function
func TestCheckAccess(t *testing.T) {
	r, err := http.NewRequest("GET", "/api/v1/trains", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	for i := 0; i < 2; i++ {
		switch i {
		//case when user was not logged in, thus there was no cookie in request
		case 0:
			if CheckAccess(w, r) != false {
				t.Error("Function should return false with no cookies in request")
			}
		//case when user was logged in, thus there was cookie in request
		case 1:
			cookie := &http.Cookie{Name: "Cookie",
				Value: "expected",
			}
			http.SetCookie(w, cookie)
			// Copy the Cookie over to a new Request
			r = &http.Request{Header: http.Header{"Cookie": w.HeaderMap["Set-Cookie"]}}
			defer func() {
				RedisClient = database.Client
			}()
			RedisClient = mockRedis()
			_, err = RedisClient.LPush(cookie.Name, cookie.Value).Result()
			if err != nil {
				t.Fatal(err)
			}
			if CheckAccess(w, r) != true {
				t.Error("Function doesn't work properly")
			}
		}
	}
}
