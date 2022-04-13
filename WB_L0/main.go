package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"database/sql"
	_"github.com/lib/pq"
)

func home_page(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./public/first.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func main() {
	// connect to db
	connStr := "user=postgres dbname=orders password=123 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connecting to bd successful")

	// creating http
	http.HandleFunc("/", home_page)
	http.ListenAndServe(":3009", nil)
}