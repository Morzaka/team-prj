package database

import (
	"testing"
	"fmt"

	"github.com/google/uuid"
	"github.com/DATA-DOG/go-sqlmock"

	"team-project/services/data"
)

func TestAddUser(t *testing.T){
	s := "08307904-f18e-4fb8-9d18-29cfad38ffaf"
	id, err := uuid.Parse(s)
	if err != nil {
		fmt.Printf("Error while parsing string to uuid, %s \n", err)
		return
	}
	user:=data.User{ID: id,
		Signin: data.Signin{
			Login:    "whythat",
			Password: "whythat",
		},
		Name:    "Yuri",
		Surname: "Zhykin",
		Role:    "User"}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	Db=db
	defer db.Close()
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO public.user").WithArgs(user).WillReturnResult(sqlmock.NewResult(user))
	mock.ExpectCommit()

	// now we execute our method
	if user, err = AddUser(user); err != nil {
		t.Errorf("error was not expected while adding user: %s", err)
	}
	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

/*
func MockDatabase() error {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		"localhost", "5432", "postgres", "postgres", "travel_test")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	Db = db
	return nil
}

func TestAddUser(t *testing.T) {
	err := MockDatabase()
	defer Db.Close()
	if err != nil {
		fmt.Printf("Error while mocking connection database, %s \n", err)
		return
	}
	s := "08307904-f18e-4fb8-9d18-29cfad38ffaf"
	id, err := uuid.Parse(s)
	if err != nil {
		fmt.Printf("Error while parsing string to uuid, %s \n", err)
		return
	}
	testData := data.User{
		ID: id,
		Signin: data.Signin{
			Login:    "whythat",
			Password: "whythat",
		},
		Name:    "Yuri",
		Surname: "Zhykin",
		Role:    "User",
	}
	_, err = AddUser(testData)
	if err != nil {
		t.Errorf("Failed adding user")
	}
}

func TestGetUserPassword(t *testing.T) {
	err := MockDatabase()
	defer Db.Close()
	if err != nil {
		fmt.Printf("Error while mocking connection database, %s \n", err)
		return
	}
	testData := []struct {
		login        string
		expectedpswd string
		err          error
	}{
		{"romich", "romich", nil},
		{"whythat", "whythat", nil},
		{"golang", "", errors.New("sql: no rows in result set")},
	}
	for _, testCase := range testData {
		dbpswd, dberr := GetUserPassword(testCase.login)
		if testCase.expectedpswd != dbpswd && testCase.err == dberr {
			t.Errorf("Failed getting password")
		}
	}
}

func TestUpdateUser(t *testing.T) {
	err := MockDatabase()
	if err != nil {
		fmt.Printf("Error while mocking connection database, %s \n", err)
		return
	}
	str := "08307904-f18e-4fb8-9d18-29cfad38ffaf"
	id, err := uuid.Parse(str)
	if err != nil {
		fmt.Printf("Error while parsing string to uuid, %s \n", err)
		return
	}
	user := data.User{
		ID: id,
		Signin: data.Signin{
			Login:    "whyso",
			Password: "whyso",
		},
		Name:    "Yurko",
		Surname: "Zhykin",
		Role:    "User",
	}
	err = UpdateUser(user, id)
	if err != nil {
		t.Errorf("Failed updating user")
	}
}

func TestDeleteUser(t *testing.T) {
	err := MockDatabase()
	if err != nil {
		fmt.Printf("Error while mocking connection database, %s \n", err)
		return
	}
	strOK := "08307904-f18e-4fb8-9d18-29cfad38ffaf"
	idOK, err := uuid.Parse(strOK)
	if err != nil {
		fmt.Printf("Error while parsing string to uuid, %s \n", err)
		return
	}
	if dberr := DeleteUser(idOK); dberr != nil {
		t.Errorf("Failed deleting user")
	}
}
*/