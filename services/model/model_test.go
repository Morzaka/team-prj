package model

import(
	"testing"
	"reflect"
	"net/http"
)

//TestGenerateID tests function GenerateID
func TestGenerateID(t *testing.T){
	uuidType:="uuid.UUID"
	value:=GenerateID()
	if reflect.TypeOf(value).String()!=uuidType{
		t.Error("Returned value is not uuid type")
	}
}

//TestHashPassword tests function HashPassword
func TestHashPassword(t *testing.T){
	pswd:="golang"
	pswdHash,err:=HashPassword(pswd)
	if err!=nil{
		t.Fatal(err)
	}
	if len(pswdHash)==0{
		t.Error("Password length shouldn't be 0")
	}
}

//TestCheckPasswordHash tests function CheckPasswordHash
func TestCheckPasswordHash(t *testing.T){
	testData:=[]struct{
		pswd		string
		pswdHash 	string
		expected	bool
	}{
		{"littleskew", "$2a$14$MA.GufeWJj7IryAoAgd8BeuRphle78ubdgqaPFPpjG9GzbxEk7kKu", true},
		{"whythat", "$2a$14$MA.GufeWJj7IryAoAgd8BeuRphle78ubdgqaPFPpjG9GzbxEk7kKu", false},
	}
	for _,testCase:=range testData{
		compare:=CheckPasswordHash(testCase.pswd, testCase.pswdHash)
		if compare!=testCase.expected{
			t.Errorf("Expected comparing result %t, got %t", testCase.expected, compare)
		}
	}
}

//TestGetID tests function GetID
func TestGetID(t *testing.T){
	rOK, err := http.NewRequest("GET", "/api/v1/:id=",nil)
	if err != nil {
        t.Fatal(err)
	}
	q := rOK.URL.Query()
    q.Add("id","61c364d9-591a-4879-a9fb-79ae67945d38")
	rOK.URL.RawQuery = q.Encode()
	_,err=GetID(rOK)
	if err!=nil{
		t.Fatal(err)
	}
	rErr,err:=http.NewRequest("GET", "/api/v1",nil)
	_,err=GetID(rErr)
	if err==nil{
		t.Fatal(err)
	}
}
