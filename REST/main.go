package main

import (
	"fmt"
	"net/http"
	"database/sql"
	_ "github.com/lib/pq"
	"encoding/json"
	"os"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:postgres@localhost/bonddb?sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("connected")
}

type User struct {
	ID 					int
	FirstName 	string
	Surname			string
	PhoneNumber	string
	Email				string
	isActive		bool
}

func main() {
	http.HandleFunc("/users/show", usersShow)
	http.ListenAndServe(":8080", nil)
}

func usersShow(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(511), http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.Query("SELECT * FROM users")

	fmt.Println(rows)

	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.ID, &user.FirstName, &user.Surname, &user.PhoneNumber, &user.Email, &user.isActive)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	json.NewEncoder(os.Stdout).Encode(users)
}



