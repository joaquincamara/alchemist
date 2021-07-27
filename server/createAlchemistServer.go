package server

import (
	"io/ioutil"
	"log"
	"os"
)

var serverTemplate = []byte(`
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

func CreateAlchemistServer(appName string) {
	os.Mkdir(appName, 0755)
	appServerErr := ioutil.WriteFile(appName+"/main.go", serverTemplate, 0644)
	if appServerErr != nil {
		log.Fatal(appServerErr)
	}
}
