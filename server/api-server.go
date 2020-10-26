package server

// Creates the standar code for a alchemist RestApi
func ApiServer() []byte {
	appServer := []byte(`
	package main

	import (
		alchemy "github.com/joaquincamara/alchemist/server"
	)

	func main() {
		router := alchemy.NewRouter()
        router.GET("/", alchemy.HomeRoute)
        log.Fatal(http.ListenAndServe(":8080", router))
	}
	`)

	return appServer
}
