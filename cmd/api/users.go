package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	ID			int
	First_name	string
	Last_name 	string
	Password 	string
	Username 	string
	isAdmin		bool
}


func AllUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var users []User
		rows, err := db.Query("SELECT * FROM User")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		
		for rows.Next() {
			var user User
			if err := rows.Scan(&user.ID, &user.First_name, &user.Last_name, &user.Password, &user.Username, &user.isAdmin); err != nil {
				log.Fatal(err)
			}
			users = append(users, user)
		}
	
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}

func UserByID(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user User
		vars := mux.Vars(r)
		id := vars["id"]
		err := db.QueryRow("SELECT * FROM User WHERE ID = ?", id).Scan(&user.ID, &user.First_name, &user.Last_name, &user.Password, &user.Username, &user.isAdmin)
		if err != nil {
			log.Fatal(err)
		}
	
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}

func AddUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			log.Fatal(err)
		}

		res, err := db.Exec("INSERT INTO User (First_name, Last_name, Username, Password, Is_admin) VALUES (?, ?, ?, ?, ?)", u.First_name, u.Last_name, u.Username, u.Password, false)
		if err != nil {
			http.Error(w, "Failed to insert user", http.StatusInternalServerError)
			log.Printf("insert error: %v\n", err)
			return
		}

		id, err := res.LastInsertId()
		if err != nil {
			http.Error(w, "Failed to retrieve last insert ID", http.StatusInternalServerError)
			return
		}

		u.ID = int(id)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(u)
	}
}

func DelUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		_,err := db.Exec("DELETE FROM User WHERE ID = ?", id)
		if err != nil {
			log.Fatal(err)
		}
	
		json.NewEncoder(w).Encode("User deleted.")
	}
}

func EditUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u User
		vars := mux.Vars(r)
		id := vars["id"]
		json.NewDecoder(r.Body).Decode(&u)

		newID, err := strconv.Atoi(id)
		if err != nil {
			log.Fatal(err)
		}

		_, err = db.Exec("UPDATE User SET First_name = ?, Last_name = ?, Password = ?, Username = ? WHERE ID = ?", u.First_name, u.Last_name, u.Password, u.Username, newID)
		if err != nil {
			log.Fatal(err)
		}
	
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(u)
	}
}