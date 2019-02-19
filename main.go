package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Person struct {
	ID        uint64   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
}


func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello", GetPerson).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	person := Person{1, "Hwiesung","Jung"};

	json.NewEncoder(w).Encode(person)
}
