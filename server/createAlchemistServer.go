package server

import (
	"io/ioutil"
	"log"
	"os"
)

var serverTemplate = []byte(`
package main

import (
	"github.com/joaquincamara/silver"
	"github.com/joaquincamara/silver/middleware"
)

func main() {
	router := silver.NewRouter()

	router.Use(middleware.Recovery)

	router.GET("/", silver.AlchemyDoor)
	router.Start("8080", router)
}
`)

func CreateAlchemistServer(appName string) {
	os.Mkdir(appName, 0755)
	appServerErr := ioutil.WriteFile(appName+"/main.go", serverTemplate, 0644)
	if appServerErr != nil {
		log.Fatal(appServerErr)
	}
}
