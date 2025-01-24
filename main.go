package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

type Contact struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type ContactService struct {
	Contacts map[int]Contact
}

func main() {
	service := ContactService{Contacts: make(map[int]Contact)}
	mux := http.NewServeMux()
	mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
