package database

import (
	"team-project/services/data"

	"github.com/google/uuid"
)

//go:generate mockgen -destination=./mock_userCRUD.go -package=database team-project/database UserCRUD

//UserCRUD for mocking
type UserCRUD interface {
	AddUser(user data.User) (data.User, error)
	GetUserPassword(login string) (string, error)
	GetUserRole(login string) (string, error)
	UpdateUser(user data.User, id uuid.UUID) (int64, error)
	DeleteUser(id uuid.UUID) (int64, error)
	GetAllUsers() ([]data.User, error)
	GetUser(id uuid.UUID) (data.User, error)
}

//IUser structure contains interface UserCRUD
type IUser struct {
	UserMethods UserCRUD
}

// Users is an instance presented IUser
var Users UserCRUD = &IUser{}

var (
	insertUser = `INSERT INTO public.user (id,name,surname,login,password,email)
	VALUES ($1, $2, $3, $4, $5, $6);`
	selectUserPassword = `SELECT password FROM public.user WHERE login=$1;`
	selectUserRole     = `SELECT role FROM public.user WHERE login=$1;`
	selectAllUsers     = `SELECT * from public.user;`
	selectUser         = `SELECT * from public.user WHERE id=$1;`
	updateUser         = `UPDATE public.user SET name = $2, surname = $3, login=$4, password=$5, email=$6 WHERE id = $1;`
	deleteUser         = `DELETE FROM public.user WHERE id = $1;`
)

//AddUser adds info about new user to the database
func (*IUser) AddUser(user data.User) (data.User, error) {
	//insert values to the database
	_, err := Db.Exec(insertUser, user.ID, user.Name, user.Surname, user.Login, user.Password, user.Email)
	if err != nil {
		return data.User{}, err
	}
	return user, nil
}

//GetUser gets user by login
func (*IUser) GetUser(id uuid.UUID) (data.User, error) {
	var user data.User
	err := Db.QueryRow(selectUser, id).Scan(&user.ID, &user.Name, &user.Surname, &user.Login, &user.Password, &user.Role, &user.Email)
	//if there's no matches for login return empty value
	if err != nil {
		return user, err
	}
	return user, nil
}

//GetUserPassword gets user's password and returns password
func (*IUser) GetUserPassword(login string) (string, error) {
	var password string
	//get user's password for given login
	err := Db.QueryRow(selectUserPassword, login).Scan(&password)
	//if there's no matches for login return empty value
	if err != nil {
		return password, err
	}
	//else return password
	return password, nil
}

//GetUserRole get's user's role and returns it with nil error, otherwise returns error
func (*IUser) GetUserRole(login string) (string, error) {
	var role string
	err := Db.QueryRow(selectUserRole, login).Scan(&role)
	if err != nil {
		return role, err
	}
	return role, nil
}

//UpdateUser updates user's personal information
func (*IUser) UpdateUser(user data.User, id uuid.UUID) (int64, error) {
	var count int64
	res, err := Db.Exec(updateUser, id, user.Name, user.Surname, user.Login, user.Password, user.Email)
	if err != nil {
		return count, err
	}
	count, err = res.RowsAffected()
	if err != nil {
		return count, err
	}
	return count, nil
}

//DeleteUser deletes user's page from db
func (*IUser) DeleteUser(id uuid.UUID) (int64, error) {
	var count int64
	res, err := Db.Exec(deleteUser, id)
	if err != nil {
		return count, err
	}
	count, err = res.RowsAffected()
	if err != nil {
		return count, err
	}
	return count, nil
}

//GetAllUsers returns slice with all users in db with possible error
func (*IUser) GetAllUsers() ([]data.User, error) {
	var users []data.User
	rows, err := Db.Query(selectAllUsers)
	if err != nil {
		return users, err
	}
	defer rows.Close()
	for rows.Next() {
		var user data.User
		err = rows.Scan(&user.ID, &user.Name, &user.Surname, &user.Login, &user.Password, &user.Email, &user.Role)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		return users, err
	}
	return users, nil
}
