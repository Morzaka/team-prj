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

	"team-project/database"
	"team-project/mocks"
	"team-project/services/data"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
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

//TestSignin tests Signin function
func TestSignin(t *testing.T) {
	id := uuid.Must(uuid.Parse("08307904-f18e-4fb8-9d18-29cfad38ffaf"))
	user := data.Signin{Login: "golang", Password: "golang"}
	jsonBody, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}
	r, err := http.NewRequest("POST", "/api/v1/login", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	mockCtrlUser := gomock.NewController(t)
	mockCtrlModel := gomock.NewController(t)
	mockCtrlCommon := gomock.NewController(t)
	defer mockCtrlUser.Finish()
	defer mockCtrlModel.Finish()
	defer mockCtrlCommon.Finish()
	mockUserCRUD := mocks.NewMockUserCRUD(mockCtrlUser)
	mockModel := mocks.NewMockModel(mockCtrlModel)
	mockCommon := mocks.NewMockCommon(mockCtrlCommon)
	UserCrud = mockUserCRUD
	CommonVar = mockCommon
	ModelVar = mockModel
	defer func() {
		RedisClient = database.Client
	}()
	for i := 0; i < 4; i++ {
		RedisClient = mockRedis()
		switch i {
		//case when there was no error getting password from db to compare with the password entered
		case 0:
			mockModel.EXPECT().GenerateID().Return(id)
			mockUserCRUD.EXPECT().GetUserPassword(user.Login).Return(user.Password, nil)
			mockModel.EXPECT().CheckPasswordHash(user.Password, user.Password).Return(true)
			mockCommon.EXPECT().RenderJSON(w, r, http.StatusOK, user)
			http.HandlerFunc(Signin).ServeHTTP(w, r)
		//case when there was an error getting password from db to compare with the password entered
		case 1:

			r.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBody))
			mockUserCRUD.EXPECT().GetUserPassword(user.Login).Return(user.Password, errors.New("Errors"))
			mockCommon.EXPECT().RenderJSON(w, r, http.StatusUnauthorized, emptyResponse)
			http.HandlerFunc(Signin).ServeHTTP(w, r)
		//case when password from db and password entered do not match
		case 2:
			r.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBody))
			mockUserCRUD.EXPECT().GetUserPassword(user.Login).Return(user.Password, nil)
			mockModel.EXPECT().CheckPasswordHash(user.Password, user.Password).Return(false)
			mockCommon.EXPECT().RenderJSON(w, r, http.StatusUnauthorized, emptyResponse)
			http.HandlerFunc(Signin).ServeHTTP(w, r)
		//case when error with redis db occured and user couldn't log in
		case 3:
			r.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBody))
			mockUserCRUD.EXPECT().GetUserPassword(user.Login).Return(user.Password, nil)
			mockModel.EXPECT().CheckPasswordHash(user.Password, user.Password).Return(true)
			mockCommon.EXPECT().RenderJSON(w, r, http.StatusInternalServerError, emptyResponse)
			mockModel.EXPECT().GenerateID().Return(id)
			RedisClient.Close()
			http.HandlerFunc(Signin).ServeHTTP(w, r)
		}
	}
}

//TestSignup tests Signup function
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
	mockCtrlUser := gomock.NewController(t)
	mockCtrlModel := gomock.NewController(t)
	mockCtrlCommon := gomock.NewController(t)
	defer mockCtrlUser.Finish()
	defer mockCtrlModel.Finish()
	defer mockCtrlCommon.Finish()
	mockUserCRUD := mocks.NewMockUserCRUD(mockCtrlUser)
	mockModel := mocks.NewMockModel(mockCtrlModel)
	mockCommon := mocks.NewMockCommon(mockCtrlCommon)
	UserCrud = mockUserCRUD
	CommonVar = mockCommon
	ModelVar = mockModel
	id := uuid.Must(uuid.Parse("08307904-f18e-4fb8-9d18-29cfad38ffaf"))
	for i := 0; i < 2; i++ {
		switch i {
		//case when there was no error
		case 0:
			mockModel.EXPECT().GenerateID().Return(id)
			mockModel.EXPECT().HashPassword(user.Signin.Password).Return(user.Signin.Password, nil)
			mockUserCRUD.EXPECT().AddUser(user).Return(user, nil)
			mockCommon.EXPECT().RenderJSON(w, r, http.StatusOK, user)
			http.HandlerFunc(Signup).ServeHTTP(w, r)
			//case when there was error while adding user to db
		case 1:
			r.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBody))
			mockModel.EXPECT().GenerateID().Return(id)
			mockModel.EXPECT().HashPassword(user.Signin.Password).Return(user.Signin.Password, nil)
			mockUserCRUD.EXPECT().AddUser(user).Return(user, errors.New("error"))
			mockCommon.EXPECT().RenderJSON(w, r, http.StatusInternalServerError, emptyResponse)
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
	mockCtrlCommon := gomock.NewController(t)
	defer mockCtrlCommon.Finish()
	mockCommon := mocks.NewMockCommon(mockCtrlCommon)
	CommonVar = mockCommon
	defer func() {
		RedisClient = database.Client

	}()
	RedisClient = mockRedis()
	for i := 0; i < 2; i++ {
		switch i {
		//case when there was no error while accessing redis db, cookie was deleted
		case 0:
			_, err = RedisClient.LPush(cookie.Name, cookie.Value).Result()
			if err != nil {
				t.Fatal(err)
			}
			mockCommon.EXPECT().RenderJSON(w, r, http.StatusNoContent, emptyResponse)
			http.HandlerFunc(Logout).ServeHTTP(w, r)
			//check whether cookie was deleted successfully
			if cookie, err := r.Cookie(cookie.Name); err != nil {
				if cookie.MaxAge >= 0 {
					t.Error("Users hasn't logged out successfully")
				}
				t.Fatal(err)
			}
		//case when there was an error while accessing redis db
		case 1:
			RedisClient.Close()
			mockCommon.EXPECT().RenderJSON(w, r, http.StatusInternalServerError, emptyResponse)
			http.HandlerFunc(Logout).ServeHTTP(w, r)
		}
	}
}

//TestUpdateUserPage tests UpdateUserPage function
func TestUpdateUserPage(t *testing.T) {
	id := uuid.Must(uuid.Parse("08307904-f18e-4fb8-9d18-29cfad38ffaf"))
	user := data.User{ID: id, Signin: data.Signin{Login: "oks", Password: "oks"}, Name: "Oksana", Surname: "Zhykina", Role: "User"}
	jsonBody, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}
	r, err := http.NewRequest("PATCH", "/api/v1/user", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}
	q := r.URL.Query()
	q.Add("id", id.String())
	r.URL.RawQuery = q.Encode()
	w := httptest.NewRecorder()
	mockCtrlUser := gomock.NewController(t)
	mockCtrlModel := gomock.NewController(t)
	mockCtrlCommon := gomock.NewController(t)
	defer mockCtrlUser.Finish()
	defer mockCtrlModel.Finish()
	defer mockCtrlCommon.Finish()
	mockUserCRUD := mocks.NewMockUserCRUD(mockCtrlUser)
	mockModel := mocks.NewMockModel(mockCtrlModel)
	mockCommon := mocks.NewMockCommon(mockCtrlCommon)
	UserCrud = mockUserCRUD
	CommonVar = mockCommon
	ModelVar = mockModel
	for i := 0; i < 2; i++ {
		switch i {
		//case where there was no error
		case 0:
			mockModel.EXPECT().HashPassword(user.Signin.Password).Return(user.Signin.Password, nil)
			mockUserCRUD.EXPECT().UpdateUser(user, id).Return(nil)
			mockCommon.EXPECT().RenderJSON(w, r, http.StatusOK, user)
			http.HandlerFunc(UpdateUserPage).ServeHTTP(w, r)
			//case when there was error while updating user in db
		case 1:
			r.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBody))
			mockModel.EXPECT().HashPassword(user.Signin.Password).Return(user.Signin.Password, nil)
			mockUserCRUD.EXPECT().UpdateUser(user, id).Return(errors.New("error"))
			mockCommon.EXPECT().RenderJSON(w, r, http.StatusInternalServerError, emptyResponse)
			http.HandlerFunc(UpdateUserPage).ServeHTTP(w, r)
		}
	}
}

//TestDeleteUser tests DeleteUser function
func TestDeleteUser(t *testing.T) {
	id := uuid.Must(uuid.Parse("08307904-f18e-4fb8-9d18-29cfad38ffaf"))
	r, err := http.NewRequest("DELETE", "/api/v1/user", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := r.URL.Query()
	q.Add("id", id.String())
	r.URL.RawQuery = q.Encode()
	w := httptest.NewRecorder()
	mockCtrlUser := gomock.NewController(t)
	mockCtrlCommon := gomock.NewController(t)
	defer mockCtrlUser.Finish()
	defer mockCtrlCommon.Finish()
	mockUserCRUD := mocks.NewMockUserCRUD(mockCtrlUser)
	mockCommon := mocks.NewMockCommon(mockCtrlCommon)
	UserCrud = mockUserCRUD
	CommonVar = mockCommon
	for i := 0; i < 2; i++ {
		switch i {
		//case when there was no error while deleting user from db
		case 0:
			mockUserCRUD.EXPECT().DeleteUser(id).Return(nil)
			mockCommon.EXPECT().RenderJSON(w, r, http.StatusNotFound, "User was deleted successfully!")
			http.HandlerFunc(DeleteUserPage).ServeHTTP(w, r)
		//case when there was an error while deleting user from db
		case 1:
			mockUserCRUD.EXPECT().DeleteUser(id).Return(errors.New("error"))
			mockCommon.EXPECT().RenderJSON(w, r, http.StatusInternalServerError, emptyResponse)
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
	mockCtrlUser := gomock.NewController(t)
	mockCtrlCommon := gomock.NewController(t)
	defer mockCtrlUser.Finish()
	defer mockCtrlCommon.Finish()
	mockUserCRUD := mocks.NewMockUserCRUD(mockCtrlUser)
	mockCommon := mocks.NewMockCommon(mockCtrlCommon)
	UserCrud = mockUserCRUD
	CommonVar = mockCommon
	for i := 0; i < 2; i++ {
		switch i {
		//case when there was no error while getting all users from db
		case 0:
			mockUserCRUD.EXPECT().GetAllUsers().Return([]data.User{}, nil)
			mockCommon.EXPECT().RenderJSON(w, r, http.StatusOK, []data.User{})
			http.HandlerFunc(ListAllUsers).ServeHTTP(w, r)
		//case when there was an error while getting all users from db
		case 1:
			mockUserCRUD.EXPECT().GetAllUsers().Return([]data.User{}, errors.New("error"))
			mockCommon.EXPECT().RenderJSON(w, r, http.StatusNoContent, emptyResponse)
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
	defer func() {
		RedisClient = database.Client
	}()
	for i := 0; i < 3; i++ {
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
		//case when there was an error while accessing redis db
		case 2:
			defer func() {
				RedisClient = database.Client
			}()
			RedisClient = mockRedis()
			RedisClient.Close()
			if CheckAccess(w, r) != false {
				t.Error("Function doesn't work properly")
			}
		}
	}
}

//TestCheckAdmin tests function CheckAdmin
func TestCheckAdmin(t *testing.T) {
	r, err := http.NewRequest("GET", "/api/v1/trains", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	mockCtrlUser := gomock.NewController(t)
	defer mockCtrlUser.Finish()
	mockUserCRUD := mocks.NewMockUserCRUD(mockCtrlUser)
	UserCrud = mockUserCRUD
	defer func() {
		LoggedIn = CheckAccess
	}()
	for i := 0; i < 4; i++ {
		cookie := &http.Cookie{Name: "Cookie",
			Value: "expected",
		}
		http.SetCookie(w, cookie)
		r = &http.Request{Header: http.Header{"Cookie": w.HeaderMap["Set-Cookie"]}}
		switch i {
		//case where user wasn't logged in at all
		case 0:
			LoggedIn = func(http.ResponseWriter, *http.Request) bool {
				return false
			}
			if CheckAdmin(w, r) != false {
				t.Error("Function doesn't work properly")
			}
		//case when user was logged in with role Admin
		case 1:
			LoggedIn = func(http.ResponseWriter, *http.Request) bool {
				return true
			}
			mockUserCRUD.EXPECT().GetUserRole(cookie.Name).Return("Admin", nil)
			if CheckAdmin(w, r) != true {
				t.Error("Function doesn't work properly")
			}
		//case when user was logged in with role User
		case 2:
			LoggedIn = func(http.ResponseWriter, *http.Request) bool {
				return true
			}
			mockUserCRUD.EXPECT().GetUserRole(cookie.Name).Return("User", nil)
			if CheckAdmin(w, r) != false {
				t.Error("Function doesn't work properly")
			}
		//case when there was error while getting user role from db
		case 3:
			LoggedIn = func(http.ResponseWriter, *http.Request) bool {
				return true
			}
			mockUserCRUD.EXPECT().GetUserRole(cookie.Name).Return("", errors.New("Error"))
			if CheckAdmin(w, r) != false {
				t.Error("Function doesn't work properly")
			}
		}
	}
}
