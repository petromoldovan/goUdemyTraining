package main

import (
	"fmt"
	"net/http"
	"database/sql"
	_ "github.com/lib/pq"
	"encoding/json"
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
	http.HandleFunc("/users", userShowByID)
	http.ListenAndServe(":8080", nil)
}

//
// get all users
//
func usersShow(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
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

	marshaledUsers, err := json.Marshal(users)

	//set cors
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Server", "User Service")
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(marshaledUsers)
}

//
// get specific user by id
//
func userShowByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(511), http.StatusMethodNotAllowed)
		return
	}

	//reading query param
	id := r.FormValue("id")
	if id == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	row := db.QueryRow("SELECT * FROM users WHERE id=$1", id)

	user := User{}

	err := row.Scan(&user.ID, &user.FirstName, &user.Surname, &user.PhoneNumber, &user.Email, &user.isActive)
	switch {
		case err == sql.ErrNoRows:
			http.NotFound(w, r)
			return
		case err != nil:
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
	}

	json.NewEncoder(w).Encode(user)
}
