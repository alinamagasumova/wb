package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"database/sql"
	"github.com/lib/pq"
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
func handleReq() {
	http.HandleFunc("/", home_page)
	http.ListenAndServe(":3009", nil)
}

func main() {
	handleReq()
	connStr := "user=postgres dbname=orders password=123 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connecting to bd successful")
	defer db.Close()
	  
}