package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hello world.")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", helloWorld).Methods("GET")
	myRouter.HandleFunc("/posts", AllPosts).Methods("GET")
	myRouter.HandleFunc("/posts/{id}", PostByID).Methods("GET")
	myRouter.HandleFunc("/posts", AddPost).Methods("POST")
	myRouter.HandleFunc("/posts/{id}", DelPost).Methods("DELETE")
	myRouter.HandleFunc("/posts/{id}", EditPost).Methods("PATCH")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

var db *sql.DB

func main() {
	fmt.Println("Go MySQL Test")
	db, err := sql.Open("mysql", "new_user:password@tcp(127.0.0.1:3306)/testdb")

	defer db.Close()

	if err != nil {
		panic(err)
	}

	handleRequests()

	fmt.Println("DB Successful connection.")
}