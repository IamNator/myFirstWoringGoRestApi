package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

type Book struct {
	ID     string  `json: "id"`
	Isbn   string  `json: "isbn"`
	Title  string  `json:"title"`
	Author *Author `json: "author"`
}

func main() {

	mybook1 := Book{
		ID:    "1",
		Isbn:  "23533342",
		Title: "Success",
		Author: &Author{
			FirstName: "Nator",
			LastName:  "Verinumbe",
		},
	}

	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/api/books", mybook1.getBooks).Methods("GET")

	fmt.Println("Server running on localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func (book *Book) getBooks(res http.ResponseWriter, req *http.Request) {
	_ = json.NewEncoder(res).Encode(book)
}
