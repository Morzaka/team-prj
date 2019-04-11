package authorization

import(
	"testing"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"bytes"

	"github.com/google/uuid"
	"github.com/go-redis/redis"
	
	"team-project/services/data"
	"team-project/database"
	"team-project/services/model"
	"team-project/services/common"
)
//TestSignin tests function Signin
func TestSignin(t *testing.T){
	user:=data.Signin{Login:"golang", Password:"golang"}
	jsonBody, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}
	r, err := http.NewRequest("POST", "/api/v1/login", bytes.NewBuffer(jsonBody))
    if err != nil {
        t.Fatal(err)
	}
	w :=httptest.NewRecorder()
	RenderJSON=func(w http.ResponseWriter, r *http.Request, status int, response interface{}){}
	GetUserPassword=func(login string)(string,error){
		return login,nil
	}
	CheckPasswordHash=func(password,hash string) bool{
		return true
	}
	RedisPush=func(string, ...interface {}) *redis.IntCmd{
		var res *redis.IntCmd
		return res
	}
	RedisRem=func(string, int64, interface {}) *redis.IntCmd {
		var res *redis.IntCmd
		return res
	}
	SessionID=uuid.New()
	defer func(){
		RenderJSON=common.RenderJSON
		GetUserPassword=database.GetUserPassword
		CheckPasswordHash=model.CheckPasswordHash
		RedisPush=database.Client.LPush
		RedisRem=database.Client.LRem
	}()
	http.HandlerFunc(Signup).ServeHTTP(w, r)
}



func TestSignup(t *testing.T){	
	user:=data.User{Signin:data.Signin{Login:"oks", Password:"oks"}, Name:"Oksana", Surname:"Zhykina", Role:"User"}
	jsonBody, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}
	r, err := http.NewRequest("POST", "/api/v1/register", bytes.NewBuffer(jsonBody))
    if err != nil {
        t.Fatal(err)
	}
	w :=httptest.NewRecorder()
	AddUser=func(user data.User)(data.User,error){
		return user,nil
	}
	HashPassword=func(password string)(string,error){
		return "$2a$14$MA.GufeWJj7IryAoAgd8BeuRphle78ubdgqaPFPpjG9GzbxEk7kKu",nil
	}
	GenerateID=func()uuid.UUID{
		return uuid.New()
	}
	RenderJSON=func(w http.ResponseWriter, r *http.Request, status int, response interface{}){}
	defer func(){
		AddUser = database.AddUser // set back original func at end of test
		HashPassword=model.HashPassword
		GenerateID=model.GenerateID
		RenderJSON=common.RenderJSON
	}()
	http.HandlerFunc(Signup).ServeHTTP(w, r)
}