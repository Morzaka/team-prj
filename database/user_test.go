package database

import (
	"fmt"
	"testing"

	"team-project/services/data"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
)

var userTest *IUser

//TestAddUser tests function AddUser
func TestAddUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	Db = db
	defer db.Close()
	id, err := uuid.Parse("08307904-f18e-4fb8-9d18-29cfad38ffaf")
	if err != nil {
		t.Fatal(err)
	}
	user := data.User{ID: id,
		Login:    "whythat",
		Password: "whythat",
		Name:     "Yuri",
		Surname:  "Zhykin",
		Email:    "yuri77@gmail.com",
	}
	mock.ExpectExec("INSERT INTO public.user").WithArgs(user.ID, user.Name, user.Surname, user.Login, user.Password, user.Email).WillReturnResult(sqlmock.NewResult(1, 1))
	// now we execute our method
	if user, err = userTest.AddUser(user); err != nil {
		t.Errorf("error was not expected while adding user: %s", err)
	}
	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

//TestDeleteUser tests function DeleteUser
func TestDeleteUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	Db = db
	defer db.Close()
	id, err := uuid.Parse("08307904-f18e-4fb8-9d18-29cfad38ffaf")
	if err != nil {
		t.Fatal(err)
	}
	mock.ExpectExec("DELETE").WithArgs(id).WillReturnResult(sqlmock.NewResult(0, 1))
	// now we execute our method
	if _, err = userTest.DeleteUser(id); err != nil {
		t.Errorf("error was not expected while deleting user: %s", err)
	}
}

//TestUpdateUser tests function UpdateUser
func TestUpdateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	Db = db
	defer db.Close()
	id, err := uuid.Parse("08307904-f18e-4fb8-9d18-29cfad38ffaf")
	if err != nil {
		t.Fatal(err)
	}
	user := data.User{
		ID:       id,
		Login:    "whythat",
		Password: "whythat",
		Name:     "Jakob",
		Surname:  "Spalding",
		Email:    "jake11spal@gmail.com",
	}
	mock.ExpectExec("UPDATE public.user").WithArgs(id, user.Name, user.Surname, user.Login, user.Password, user.Email).WillReturnResult(sqlmock.NewResult(0, 1))
	// now we execute our method
	if _, err = userTest.UpdateUser(user, id); err != nil {
		t.Errorf("error was not expected while deleting user: %s", err)
	}
}

//TestGetUserPassword tests function TestGetUserPassword
func TestGetUserPassword(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	Db = db
	defer db.Close()
	loginOK := "golang"
	loginERR := "java"
	rows := sqlmock.NewRows([]string{"password"}).AddRow("golang")
	mock.ExpectQuery("SELECT").WithArgs(loginOK).WillReturnRows(rows)
	mock.ExpectQuery("SELECT").WithArgs(loginERR).WillReturnError(fmt.Errorf("no rows found"))
	if _, err = userTest.GetUserPassword(loginOK); err != nil {
		t.Errorf("error was not expected while getting user: %s", err)
	}
	if _, err = userTest.GetUserPassword(loginERR); err == nil {
		t.Errorf("error was not expected while getting user: %s", err)
	}
	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

//TestGetUserRole tests function TestGetUserRole
func TestGetUserRole(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	Db = db
	defer db.Close()
	loginOK := "oksana"
	loginERR := "yuri"
	rows := sqlmock.NewRows([]string{"role"}).AddRow("Admin")
	mock.ExpectQuery("SELECT").WithArgs(loginOK).WillReturnRows(rows)
	mock.ExpectQuery("SELECT").WithArgs(loginERR).WillReturnError(fmt.Errorf("no rows found"))
	if _, err = userTest.GetUserRole(loginOK); err != nil {
		t.Errorf("error was not expected while getting user: %s", err)
	}
	if _, err = userTest.GetUserRole(loginERR); err == nil {
		t.Errorf("error was not expected while getting user: %s", err)
	}
	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

//TestGetAllUsers tests function GetAllUsers
func TestGetAllUsers(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	Db = db
	defer db.Close()
	id, err := uuid.Parse("08307904-f18e-4fb8-9d18-29cfad38ffaf")
	if err != nil {
		t.Fatal(err)
	}
	rows := sqlmock.NewRows([]string{"id", "name", "surname", "login", "password", "role", "email"}).AddRow(id, "Oksana", "Zhykina", "litleskew", "littleskew", "User", "oks8@gmail.com")
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	if _, err = userTest.GetAllUsers(); err != nil {
		t.Errorf("error was not expected while getting user: %s", err)
	}
	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

//TestGetUser tests function GetUser
func TestGetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	Db = db
	defer db.Close()
	id, err := uuid.Parse("08307904-f18e-4fb8-9d18-29cfad38ffaf")
	if err != nil {
		t.Fatal(err)
	}
	row := sqlmock.NewRows([]string{"id", "name", "surname", "login", "password", "role", "email"}).AddRow(id, "Oksana", "Zhykina", "litleskew", "littleskew", "User", "oks8@gmail.com")
	mock.ExpectQuery("SELECT").WillReturnRows(row)
	if _, err = userTest.GetUser(id); err != nil {
		t.Errorf("error was not expected while getting user: %s", err)
	}
	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
