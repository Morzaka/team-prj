package model

import (
	"testing"
)

var modelTest *IModel

//TestHashPassword tests function HashPassword
func TestHashPassword(t *testing.T) {
	pswd := "golang"
	pswdHash, err := modelTest.HashPassword(pswd)
	if err != nil {
		t.Fatal(err)
	}
	if len(pswdHash) == 0 {
		t.Error("Password length shouldn't be 0")
	}
}

//TestCheckPasswordHash tests function CheckPasswordHash
func TestCheckPasswordHash(t *testing.T) {
	testData := []struct {
		pswd     string
		pswdHash string
		expected bool
	}{
		{"littleskew", "$2a$14$MA.GufeWJj7IryAoAgd8BeuRphle78ubdgqaPFPpjG9GzbxEk7kKu", true},
		{"whythat", "$2a$14$MA.GufeWJj7IryAoAgd8BeuRphle78ubdgqaPFPpjG9GzbxEk7kKu", false},
	}
	for _, testCase := range testData {
		compare := modelTest.CheckPasswordHash(testCase.pswd, testCase.pswdHash)
		if compare != testCase.expected {
			t.Errorf("Expected comparing result %t, got %t", testCase.expected, compare)
		}
	}
}
