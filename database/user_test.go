package database

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"

	"team-project/services/data"
)

func TestAddUser(t *testing.T) {
	s := "08307904-f18e-4fb8-9d18-29cfad38ffaf"
	id, err := uuid.Parse(s)
	if err != nil {
		fmt.Printf("Error while parsing string to uuid, %s \n", err)
		return
	}
	user := data.User{ID: id,
		Signin: data.Signin{
			Login:    "whythat",
			Password: "whythat",
		},
		Name:    "Yuri",
		Surname: "Zhykin",
		Role:    "User",
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	Db = db
	defer db.Close()
	mock.ExpectExec("INSERT INTO public.user").WithArgs(user.ID, user.Name, user.Surname, user.Signin.Login, user.Signin.Password, user.Role).WillReturnResult(sqlmock.NewResult(1, 1))
	// now we execute our method
	if user, err = AddUser(user); err != nil {
		t.Errorf("error was not expected while adding user: %s", err)
	}
	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	Db = db
	defer db.Close()
	strOK := "08307904-f18e-4fb8-9d18-29cfad38ffaf"
	idOK, err := uuid.Parse(strOK)
	if err != nil {
		fmt.Printf("Error while parsing string to uuid, %s \n", err)
		return
	}
	mock.ExpectExec("DELETE").WithArgs(idOK).WillReturnResult(sqlmock.NewResult(0, 1))
	// now we execute our method
	if err = DeleteUser(idOK); err != nil {
		t.Errorf("error was not expected while deleting user: %s", err)
	}
	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	Db = db
	defer db.Close()
	str := "08307904-f18e-4fb8-9d18-29cfad38ffaf"
	id, err := uuid.Parse(str)
	if err != nil {
		fmt.Printf("Error while parsing string to uuid, %s \n", err)
		return
	}
	user := data.User{
		ID: id,
		Signin: data.Signin{
			Login:    "whythat",
			Password: "whythat",
		},
		Name:    "Jakob",
		Surname: "Spalding",
		Role:    "User",
	}
	mock.ExpectExec("UPDATE public.user").WithArgs(id, user.Name, user.Surname, user.Signin.Login, user.Signin.Password, user.Role).WillReturnResult(sqlmock.NewResult(0, 1))
	// now we execute our method
	if err = UpdateUser(user, id); err != nil {
		t.Errorf("error was not expected while deleting user: %s", err)
	}
	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

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
	if _, err = GetUserPassword(loginOK); err != nil {
		t.Errorf("error was not expected while getting user: %s", err)
	}
	if _, err = GetUserPassword(loginERR); err == nil {
		t.Errorf("error was not expected while getting user: %s", err)
	}
	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
