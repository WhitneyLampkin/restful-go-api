package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoId}", TodoShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}
