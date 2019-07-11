package authorization

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"team-project/database"
	"team-project/services/data"
	"team-project/services/model"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis"
	"github.com/go-zoo/bone"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	_ "github.com/google/uuid"
)

//testRouter returns router
func testRouter() *bone.Mux {
	router := bone.New().Prefix("/api")
	subV1 := bone.New()
	router.SubRoute("/v1", subV1)
	subV1.PostFunc("/register", Signup)
	subV1.PostFunc("/login", Signin)
	subV1.PostFunc("/logout", Logout)
	subV1.DeleteFunc("/user/:id", DeleteUserPage)
	subV1.PatchFunc("/user/:id", UpdateUserPage)
	subV1.GetFunc("/users", ListAllUsers)
	subV1.GetFunc("/user/:id", GetOneUser)
	return router
}

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

//TestCase contains data for testing
type TestCase struct {
	name                string
	url                 string
	statusWant          int
	mockedUserList      []data.User
	mockedUser          data.User
	mockedError         error
	mockedRole          string
	mockedAffected      int64
	mockedLogicalResult bool
	messageWant         string
	cookieSet           bool
	openRedis           bool
	resultWant          bool
	loggedIn            bool
}

//TestSignin function tests Signin function
func TestSignin(t *testing.T) {
	tests := []TestCase{
		{
			name:        "Fail_Not_Registered",
			url:         "/api/v1/login",
			mockedUser:  data.User{},
			statusWant:  http.StatusUnauthorized,
			mockedError: errors.New("mocked error"),
		},
		{
			name:                "Fail_Redis_DBerror",
			url:                 "/api/v1/login",
			mockedUser:          data.User{},
			statusWant:          http.StatusInternalServerError,
			mockedError:         nil,
			mockedLogicalResult: true,
		},
		{
			name:                "Fail_Invalid_Password",
			url:                 "/api/v1/login",
			mockedUser:          data.User{},
			statusWant:          http.StatusUnauthorized,
			mockedError:         nil,
			mockedLogicalResult: false,
		},
		{
			name:                "Success_Logged_In",
			url:                 "/api/v1/login",
			mockedUser:          data.User{},
			statusWant:          http.StatusOK,
			mockedError:         nil,
			mockedLogicalResult: true,
			openRedis:           true,
		},
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	for _, tc := range tests {
		RedisClient = mockRedis()
		if !tc.openRedis {
			RedisClient.Close()
		}
		rec := httptest.NewRecorder()
		jsonBody, err := json.Marshal(tc.mockedUser)
		if err != nil {
			t.Fatal(err)
		}
		req, _ := http.NewRequest(http.MethodPost, tc.url, bytes.NewBuffer(jsonBody))
		usersMock := database.NewMockUserCRUD(mockCtrl)
		modelMock := model.NewMockModel(mockCtrl)
		model.HelperModel = modelMock
		database.Users = usersMock
		usersMock.EXPECT().GetUserPassword(tc.mockedUser.Login).Return(tc.mockedUser.Password, tc.mockedError)
		if tc.name != "Fail_Not_Registered" {
			modelMock.EXPECT().CheckPasswordHash(tc.mockedUser.Password, tc.mockedUser.Password).Return(tc.mockedLogicalResult)
		}
		testRouter().ServeHTTP(rec, req)
		if rec.Code != tc.statusWant {
			t.Errorf("Expected %d, got %d", rec.Code, tc.statusWant)
		}
	}
	RedisClient = database.Client
}

//TestSignup tests Signup function
func TestSignup(t *testing.T) {
	id, err := uuid.Parse("08307904-f18e-4fb8-9d18-29cfad38ffaf")
	if err != nil {
		t.Fatal(err)
	}
	tests := []TestCase{
		{
			name:           "Failure_LoginNotAllowed",
			url:            "/api/v1/register",
			statusWant:     http.StatusNotAcceptable,
			mockedUser:     data.User{ID: id},
			mockedError:    nil,
			mockedUserList: []data.User{{}},
		},
		{
			name:           "Failure_DBerror",
			url:            "/api/v1/register",
			statusWant:     http.StatusInternalServerError,
			mockedUser:     data.User{ID: id},
			mockedError:    errors.New("mocked error"),
			mockedUserList: []data.User{},
		},
		{
			name:           "Success_Registered",
			url:            "/api/v1/register",
			statusWant:     http.StatusOK,
			mockedUser:     data.User{ID: id},
			mockedError:    nil,
			mockedUserList: []data.User{},
		},
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	for _, tc := range tests {
		Validate = func(user data.User) (bool, string) {
			return true, ""
		}
		rec := httptest.NewRecorder()
		jsonBody, err := json.Marshal(tc.mockedUser)
		if err != nil {
			t.Fatal(err)
		}
		req, _ := http.NewRequest(http.MethodPost, tc.url, bytes.NewBuffer(jsonBody))

		usersMock := database.NewMockUserCRUD(mockCtrl)
		modelMock := model.NewMockModel(mockCtrl)
		model.HelperModel = modelMock
		database.Users = usersMock
		usersMock.EXPECT().GetAllUsers().Return(tc.mockedUserList, tc.mockedError)
		if tc.name == "Success_Registered" {
			modelMock.EXPECT().HashPassword(tc.mockedUser.Password).Return(tc.mockedUser.Password, nil)
			usersMock.EXPECT().AddUser(tc.mockedUser).Return(tc.mockedUser, tc.mockedError)
		}
		testRouter().ServeHTTP(rec, req)
		if rec.Code != tc.statusWant {
			t.Errorf("Expected %d, got %d", tc.statusWant, rec.Code)
		}
	}
	Validate = Validation
}

//TestLogout tests Logout function
func TestLogout(t *testing.T) {
	tests := []TestCase{
		{
			name:       "Success_LoggedOut",
			url:        "/api/v1/logout",
			statusWant: http.StatusNoContent,
			openRedis:  true,
		},
		{
			name:       "Failure_DBerror",
			url:        "/api/v1/logout",
			statusWant: http.StatusInternalServerError,
			openRedis:  false,
		},
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	for _, tc := range tests {
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, tc.url, nil)
		cookie := http.Cookie{Name: "Cookie",
			Value: "expected",
		}
		http.SetCookie(rec, &cookie)
		req.AddCookie(&cookie)
		RedisClient = mockRedis()
		if !tc.openRedis {
			_ = RedisClient.Close()
		} else {
			_, err := RedisClient.LPush(cookie.Name, cookie.Value).Result()
			if err != nil {
				t.Fatal(err)
			}
		}
		testRouter().ServeHTTP(rec, req)
		if rec.Code != tc.statusWant {
			t.Errorf("Expected %d, got %d", tc.statusWant, rec.Code)
		}
	}
	RedisClient = database.Client
}

//TestListAllUsers tests ListAllUsers function
func TestListAllUsers(t *testing.T) {
	tests := []TestCase{
		{
			name:           "Get_Users_200",
			url:            "/api/v1/users",
			statusWant:     http.StatusOK,
			mockedUserList: []data.User{},
			mockedError:    nil,
		},
		{
			name:           "Get_Users_404",
			url:            "/api/v1/users",
			statusWant:     http.StatusNoContent,
			mockedUserList: []data.User{},
			mockedError:    errors.New("db error"),
		},
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	for _, tc := range tests {
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, tc.url, nil)

		usersMock := database.NewMockUserCRUD(mockCtrl)
		database.Users = usersMock

		usersMock.EXPECT().GetAllUsers().Return(tc.mockedUserList, tc.mockedError)

		testRouter().ServeHTTP(rec, req)
		if rec.Code != tc.statusWant {
			t.Errorf("Expected: %d , got %d", tc.statusWant, rec.Code)
		}

	}
}

//TestUpdateUserPage tests UpdateUserPage function
func TestUpdateUserPage(t *testing.T) {
	id := uuid.Must(uuid.Parse("08307904-f18e-4fb8-9d18-29cfad38ffaf"))
	tests := []TestCase{
		{
			name:           "Success_Update_User",
			url:            "/api/v1/user/08307904-f18e-4fb8-9d18-29cfad38ffaf",
			mockedUser:     data.User{ID: id},
			mockedAffected: 1,
			mockedError:    nil,
			statusWant:     http.StatusOK,
		},
		{
			name:           "Fail_Update_User_DBerror",
			url:            "/api/v1/user/08307904-f18e-4fb8-9d18-29cfad38ffaf",
			mockedUser:     data.User{ID: id},
			mockedAffected: 0,
			mockedError:    errors.New("mocked error"),
			statusWant:     http.StatusInternalServerError,
		},
		{
			name:           "Fail_Update_User_No_content",
			url:            "/api/v1/user/08307904-f18e-4fb8-9d18-29cfad38ffaf",
			mockedUser:     data.User{ID: id},
			mockedAffected: 0,
			mockedError:    nil,
			statusWant:     http.StatusNotFound,
		},
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	for _, tc := range tests {
		rec := httptest.NewRecorder()
		jsonBody, err := json.Marshal(tc.mockedUser)
		if err != nil {
			t.Fatal(err)
		}
		req, _ := http.NewRequest(http.MethodPatch, tc.url, bytes.NewBuffer(jsonBody))
		usersMock := database.NewMockUserCRUD(mockCtrl)
		modelMock := model.NewMockModel(mockCtrl)
		model.HelperModel = modelMock
		database.Users = usersMock
		modelMock.EXPECT().HashPassword(tc.mockedUser.Password).Return(tc.mockedUser.Password, nil)
		usersMock.EXPECT().UpdateUser(tc.mockedUser, id).Return(tc.mockedAffected, tc.mockedError)
		testRouter().ServeHTTP(rec, req)
		//http.HandlerFunc(UpdateUserPage).ServeHTTP(rec, req)
		if rec.Code != tc.statusWant {
			t.Errorf("Expected %d, got %d", rec.Code, tc.statusWant)
		}
	}
}

//TestDeleteUserPage tests DeleteUserPage function
func TestDeleteUserPage(t *testing.T) {
	id, err := uuid.Parse("08307904-f18e-4fb8-9d18-29cfad38ffaf")
	if err != nil {
		t.Fatal(err)
	}
	tests := []TestCase{
		{
			name:           "Success_Delete_User",
			url:            "/api/v1/user/08307904-f18e-4fb8-9d18-29cfad38ffaf",
			mockedAffected: 1,
			mockedError:    nil,
			statusWant:     http.StatusOK,
		},
		{
			name:           "Fail_Delete_User_DBerror",
			url:            "/api/v1/user/08307904-f18e-4fb8-9d18-29cfad38ffaf",
			mockedAffected: 0,
			mockedError:    errors.New("mocked error"),
			statusWant:     http.StatusInternalServerError,
		},
		{
			name:           "Fail_Delete_User_No_content",
			url:            "/api/v1/user/08307904-f18e-4fb8-9d18-29cfad38ffaf",
			mockedAffected: 0,
			mockedError:    nil,
			statusWant:     http.StatusNotFound,
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	for _, tc := range tests {
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodDelete, tc.url, nil)
		usersMock := database.NewMockUserCRUD(mockCtrl)
		database.Users = usersMock
		usersMock.EXPECT().DeleteUser(id).Return(tc.mockedAffected, tc.mockedError)
		testRouter().ServeHTTP(rec, req)
		if rec.Code != tc.statusWant {
			t.Errorf("Expected %d, got %d", rec.Code, tc.statusWant)
		}

	}
}

//TestCheckAccess tests CheckAccess fucntion
func TestCheckAccess(t *testing.T) {
	tests := []TestCase{
		{
			name:       "Access_Denied_No_Cookie",
			url:        "/api/v1/trains",
			resultWant: false,
			statusWant: http.StatusUnauthorized,
			cookieSet:  false,
			openRedis:  true,
		},
		{
			name:       "Access_Denied_Redis_Error",
			url:        "/api/v1/trains",
			resultWant: false,
			statusWant: http.StatusInternalServerError,
			cookieSet:  true,
			openRedis:  false,
		},
		{
			name:       "Access_Granted",
			url:        "/api/v1/trains",
			resultWant: true,
			statusWant: http.StatusOK,
			cookieSet:  true,
			openRedis:  true,
		},
	}
	for _, tc := range tests {
		var cookie http.Cookie
		RedisClient = mockRedis()
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, tc.url, nil)
		if tc.cookieSet {
			cookie = http.Cookie{Name: "Cookie",
				Value: "expected",
			}
			http.SetCookie(rec, &cookie)
			req.AddCookie(&cookie)
		}
		if !tc.openRedis {
			RedisClient.Close()
		} else {
			_, err := RedisClient.LPush(cookie.Name, cookie.Value).Result()
			if err != nil {
				t.Fatal(err)
			}
		}
		if tc.resultWant != CheckAccess(rec, req) {
			t.Errorf("Expected: %t , got %t", tc.resultWant, CheckAccess(rec, req))
		}
		if tc.statusWant != rec.Code {
			t.Errorf("Expected: %d , got %d", tc.statusWant, rec.Code)
		}
	}
	RedisClient = database.Client
}

func TestGetUserByLogin(t *testing.T) {
	id, err := uuid.Parse("08307904-f18e-4fb8-9d18-29cfad38ffaf")
	if err != nil {
		t.Fatal(err)
	}
	tests := []TestCase{
		{
			name:        "Get_Users_200",
			url:         "/api/v1/user/08307904-f18e-4fb8-9d18-29cfad38ffaf",
			statusWant:  http.StatusOK,
			mockedUser:  data.User{ID: id},
			mockedError: nil,
		},
		{
			name:        "Get_Users_404",
			url:         "/api/v1/user/08307904-f18e-4fb8-9d18-29cfad38ffaf",
			statusWant:  http.StatusNoContent,
			mockedUser:  data.User{ID: id},
			mockedError: errors.New("db error"),
		},
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	for _, tc := range tests {
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, tc.url, nil)

		usersMock := database.NewMockUserCRUD(mockCtrl)
		database.Users = usersMock

		usersMock.EXPECT().GetUser(tc.mockedUser.ID).Return(tc.mockedUser, tc.mockedError)
		testRouter().ServeHTTP(rec, req)
		if rec.Code != tc.statusWant {
			t.Errorf("Expected: %d , got %d", tc.statusWant, rec.Code)
		}

	}
}

//TestCheckAdmin tests CheckAdmin function
func TestCheckAdmin(t *testing.T) {
	tests := []TestCase{
		{
			name:        "Success_Admin_role",
			url:         "/api/v1/trains",
			loggedIn:    true,
			mockedRole:  "Admin",
			mockedError: nil,
			resultWant:  true,
		},
		{
			name:        "Success_User_role",
			url:         "/api/v1/trains",
			loggedIn:    true,
			mockedRole:  "User",
			mockedError: nil,
			resultWant:  false,
		},
		{
			name:        "Failure_database_error",
			url:         "/api/v1/trains",
			loggedIn:    true,
			mockedError: errors.New("mocked error"),
			resultWant:  false,
		},
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	cookie := http.Cookie{Name: "Cookie",
		Value: "expected",
	}
	for _, tc := range tests {
		LoggedIn = func(http.ResponseWriter, *http.Request) bool {
			return tc.loggedIn
		}
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, tc.url, nil)
		http.SetCookie(rec, &cookie)
		req.AddCookie(&cookie)
		usersMock := database.NewMockUserCRUD(mockCtrl)
		database.Users = usersMock
		usersMock.EXPECT().GetUserRole(cookie.Name).Return(tc.mockedRole, tc.mockedError)
		res := CheckAdmin(rec, req)
		if res != tc.resultWant {
			t.Errorf("Expected: %t , got %t", tc.resultWant, res)
		}
	}
	LoggedIn = CheckAccess
}

func TestValidation(t *testing.T) {
	tests := []TestCase{
		{
			mockedUser: data.User{
				Name: "Oksana", Surname: "Zhykina", Login: "oks_zh", Password: "oksana88zh", Email: "oks88zh@gmail.com",
			},
			messageWant: "",
			resultWant:  true,
		},
		{
			mockedUser: data.User{
				Name: "Oksana", Surname: "Zhykina", Login: "oks_zh", Password: "ok", Email: "oks88zh@gmail.com",
			},
			messageWant: "Invalid Password",
			resultWant:  false,
		},
		{
			mockedUser: data.User{
				Name: "356oks", Surname: "Zhykina", Login: "oks_zh", Password: "oks_zh888", Email: "oks88zh@gmail.com",
			},
			messageWant: " Invalid Name",
			resultWant:  false,
		},
		{
			mockedUser: data.User{
				Name: "Oksana", Surname: "999zhyk", Login: "oks_zh", Password: "oks_zh888", Email: "oks88zh@gmail.com",
			},
			messageWant: "Invalid Surname",
			resultWant:  false,
		},
	}
	for _, tc := range tests {
		if valid, msg := Validation(tc.mockedUser); valid != tc.resultWant && msg != tc.messageWant {
			t.Errorf("Expected %t, %s, got %t, %s", tc.resultWant, tc.messageWant, valid, msg)
		}
	}
}
