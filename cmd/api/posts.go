package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Post struct {
	ID		int
	Title	string
	Content string
}


func AllPosts(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var posts []Post
		rows, err := db.Query("SELECT * FROM Post")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		
		for rows.Next() {
			var post Post
			if err := rows.Scan(&post.ID, &post.Title, &post.Content); err != nil {
				log.Fatal(err)
			}
			posts = append(posts, post)
		}
	
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(posts)
	}
}

func PostByID(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var post Post
		vars := mux.Vars(r)
		id := vars["id"]
		err := db.QueryRow("SELECT * FROM Post WHERE ID = ?", id).Scan(&post.ID, &post.Title, &post.Content)
		if err != nil {
			log.Fatal(err)
		}
	
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(post)
	}
}

func AddPost(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p Post
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			log.Fatal(err)
		}

		res, err := db.Exec("INSERT INTO Post (Title, Content) VALUES (?, ?)", p.Title, p.Content)
		if err != nil {
			http.Error(w, "Failed to insert post", http.StatusInternalServerError)
			log.Printf("insert error: %v\n", err)
			return
		}

		id, err := res.LastInsertId()
		if err != nil {
			http.Error(w, "Failed to retrieve last insert ID", http.StatusInternalServerError)
			return
		}

		p.ID = int(id)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(p)
	}
}

func DelPost(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		_,err := db.Exec("DELETE FROM Post WHERE ID = ?", id)
		if err != nil {
			log.Fatal(err)
		}
	
		json.NewEncoder(w).Encode("User deleted.")
	}
}

func EditPost(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p Post
		vars := mux.Vars(r)
		id := vars["id"]
		json.NewDecoder(r.Body).Decode(&p)

		newID, err := strconv.Atoi(id)
		if err != nil {
			log.Fatal(err)
		}

		_, err = db.Exec("UPDATE Post SET Title = ?, Content = ? WHERE ID = ?", p.Title, p.Content, newID)
		if err != nil {
			log.Fatal(err)
		}
	
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(p)
	}
}