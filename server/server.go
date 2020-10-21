package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//
func homeRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome fellow Alchemist to the ApiRest transmutation")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homeRoute)

	log.Fatal(http.ListenAndServe(":8080", router))
}
