package server

// Creates the standar code for a alchemist RestApi
func ApiServer() []byte {
	appServer := []byte(`
	package main

	import (
		"log"
		"net/http"
		"github.com/joaquincamara/silver"
	)

	func main() {
		router := silver.NewRouter()
        router.GET("/", silver.HomeRoute)
        log.Fatal(http.ListenAndServe(":8080", router))
	}
	`)

	return appServer
}
