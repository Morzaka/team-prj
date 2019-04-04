package model

import (
	"errors"
	"net/http"
	"strconv"
	"team-project/configurations"
	"time"

	"github.com/google/uuid"
)

//type Book struct {
//	Isbn   string
//	Title  string
//	Author string
//	Price  float32
//}

type Ticket struct {
	Id             uuid.UUID `json:"id"`
	Train_id       uuid.UUID `json:"train_id"`
	Plane_id       uuid.UUID `json:"plane_id"`
	User_id        uuid.UUID `json:"user_id"`
	Place          int16     `json:"place"`
	Ticket_type    string    `json:"ticket_type"`
	Discount       string    `json:"discount"`
	Price          float32   `json:"price"`
	Total_price    float32   `json:"total_price"`
	Name           string    `json:"name"`
	Surname        string    `json:"surname"`
	From_place     string    `json:"from_place"`
	Departure_date time.Time `json:"departure_date"`
	Departure_time time.Time `json:"departure_time"`
	To_place       string    `json:"to_place"`
	Arrival_date   time.Time `json:"arrival_date"`
	Arrival_time   time.Time `json:"arrival_time"`
}

func AllTickets() ([]Ticket, error) {
	rows, err := configurations.DB.Query("SELECT * FROM tickets")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tkts := make([]Ticket, 0)
	for rows.Next() {
		tk := Ticket{}
		err := rows.Scan(&tk.Id, &tk.Train_id, &tk.Plane_id, &tk.User_id,
			&tk.Place, &tk.Ticket_type, &tk.Discount, &tk.Price, &tk.Total_price,
			&tk.Name, &tk.Surname, &tk.From_place, &tk.Departure_date,
			&tk.Departure_time, &tk.To_place, &tk.Arrival_date,
			&tk.Arrival_time) // order matters
		if err != nil {
			return nil, err
		}
		tkts = append(tkts, tk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return tkts, nil
}

func OneTicket(r *http.Request) (Ticket, error) {
	tk := Ticket{}
	isbn := r.FormValue("isbn")
	if isbn == "" {
		return tk, errors.New("400. Bad Request.")
	}

	row := configurations.DB.QueryRow("SELECT * FROM books WHERE isbn = $1", isbn)

	err := row.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
	if err != nil {
		return bk, err
	}

	return bk, nil
}

func PutBook(r *http.Request) (Book, error) {
	// get form values
	bk := Book{}
	bk.Isbn = r.FormValue("isbn")
	bk.Title = r.FormValue("title")
	bk.Author = r.FormValue("author")
	p := r.FormValue("price")

	// validate form values
	if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || p == "" {
		return bk, errors.New("400. Bad request. All fields must be complete.")
	}

	// convert form values
	f64, err := strconv.ParseFloat(p, 32)
	if err != nil {
		return bk, errors.New("406. Not Acceptable. Price must be a number.")
	}
	bk.Price = float32(f64)

	// insert values
	_, err = configurations.DB.Exec("INSERT INTO books (isbn, title, author, price) VALUES ($1, $2, $3, $4)", bk.Isbn, bk.Title, bk.Author, bk.Price)
	if err != nil {
		return bk, errors.New("500. Internal Server Error." + err.Error())
	}
	return bk, nil
}

func UpdateBook(r *http.Request) (Book, error) {
	// get form values
	bk := Book{}
	bk.Isbn = r.FormValue("isbn")
	bk.Title = r.FormValue("title")
	bk.Author = r.FormValue("author")
	p := r.FormValue("price")

	if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || p == "" {
		return bk, errors.New("400. Bad Request. Fields can't be empty.")
	}

	// convert form values
	f64, err := strconv.ParseFloat(p, 32)
	if err != nil {
		return bk, errors.New("406. Not Acceptable. Enter number for price.")
	}
	bk.Price = float32(f64)

	// insert values
	_, err = configurations.DB.Exec("UPDATE books SET isbn = $1, title=$2, author=$3, price=$4 WHERE isbn=$1;", bk.Isbn, bk.Title, bk.Author, bk.Price)
	if err != nil {
		return bk, err
	}
	return bk, nil
}

func DeleteBook(r *http.Request) error {
	isbn := r.FormValue("isbn")
	if isbn == "" {
		return errors.New("400. Bad Request.")
	}

	_, err := configurations.DB.Exec("DELETE FROM books WHERE isbn=$1;", isbn)
	if err != nil {
		return errors.New("500. Internal Server Error")
	}
	return nil
}
