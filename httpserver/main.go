package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

func greet(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(rw, "<h1>Hello World! %s</h1>", time.Now())
}

func GetBook(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	book := Book{
		Title:  "The Gunslinger",
		Author: "Stephen King",
		Pages:  304,
	}
	json.NewEncoder(rw).Encode(book)
}

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/book", GetBook)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
