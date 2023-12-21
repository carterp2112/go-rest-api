package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	ID		int
	Title	string
	Content string
}

func AllPosts(w http.ResponseWriter, r *http.Request) {
	var posts []Post

	rows, err := db.Query("SELECT id, title, content FROM Post")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var p Post
		if err := rows.Scan(&p.ID, &p.Title, &p.Content); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		posts = append(posts, p)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func PostByID(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Single post endpoint hit.")
}

func AddPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Add Post endpoint hit.")
}

func DelPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Delete Post endpoint hit.")
}

func EditPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Edit Post endpoint hit.")
}