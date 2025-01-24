package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
)

type Contact struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type ContactService struct {
	Contacts map[int]Contact
}

func (c *ContactService) Create(w http.ResponseWriter, r *http.Request) {
	var contact Contact
	err := json.NewDecoder(r.Body).Decode(&contact)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := len(c.Contacts) + 1
	contact.Id = id

	c.Contacts[id] = contact

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contact)
	w.WriteHeader(http.StatusCreated)

}

func main() {
	service := &ContactService{Contacts: make(map[int]Contact)}
	mux := http.NewServeMux()

	mux.HandleFunc("/contacts", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		case http.MethodPost:
			service.Create(w, r)
		default:
			http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
