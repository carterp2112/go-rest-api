package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	_ "github.com/go-sql-driver/mysql"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hello world.")
}

var db *sql.DB

func main() {
	fmt.Println("Go MySQL Test")
	db, err := sql.Open("mysql", "new_user:password@tcp(127.0.0.1:3306)/testdb")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", helloWorld).Methods("GET")
	myRouter.HandleFunc("/posts", AllPosts(db)).Methods("GET")
	myRouter.HandleFunc("/posts/{id}", PostByID(db)).Methods("GET")
	myRouter.HandleFunc("/posts", AddPost(db)).Methods("POST")
	myRouter.HandleFunc("/posts/{id}", DelPost(db)).Methods("DELETE")
	myRouter.HandleFunc("/posts/{id}", EditPost(db)).Methods("PATCH")
	myRouter.HandleFunc("/users", AllUsers(db)).Methods("GET")
	myRouter.HandleFunc("/users/{id}", UserByID(db)).Methods("GET")
	myRouter.HandleFunc("/users", AddUser(db)).Methods("POST")
	myRouter.HandleFunc("/users/{id}", DelUser(db)).Methods("DELETE")
	myRouter.HandleFunc("/users/{id}", EditUser(db)).Methods("PATCH")

	c := cors.New(cors.Options{
        AllowedOrigins: []string{"http://example.com"},
        // AllowedOrigins: []string{"*"},
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders: []string{"Content-Type", "X-Requested-With"},
        AllowCredentials: true,
        MaxAge: 86400,
    })

    handler := c.Handler(myRouter)
	log.Fatal(http.ListenAndServe(":8081", handler))

	fmt.Println("DB Successful connection.")
}