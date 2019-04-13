package session

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
)

//TestNewSession tests function NewSession
func TestNewSession(t *testing.T) {
	if session := NewSession(); session == nil {
		t.Error("Session was not created")
	}
}

//TestInit tests method Init
func TestInit(t *testing.T) {
	session := NewSession()
	uuidType := "uuid.UUID"
	value := session.Init("golang")
	if reflect.TypeOf(value).String() != uuidType {
		t.Error("Returned value is not uuid type")
	}
}

//TestGetUser tests method TestUser
func TestGetUser(t *testing.T) {
	session := NewSession()
	s := "08307904-f18e-4fb8-9d18-29cfad38ffaf"
	id, err := uuid.Parse(s)
	if err != nil {
		t.Fatal(err)
	}
	testData := []struct {
		id           uuid.UUID
		expectedUser string
	}{
		{session.Init("golang"), "golang"},
		{id, ""},
	}
	for _, testCase := range testData {
		realUser := session.GetUser(testCase.id)
		if realUser != testCase.expectedUser {
			t.Errorf("Error: expected %s, got %s", testCase.expectedUser, realUser)
		}
	}
}
