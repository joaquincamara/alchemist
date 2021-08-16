package templates

import (
	"io/ioutil"
	"log"
)

var mainTemplate = []byte(`
package main

import (
	"net/http"

	"github.com/joaquincamara/silver"
	"github.com/joaquincamara/silver/middleware"
)

func main() {
	router := silver.NewRouter()

	router.Use(middleware.Logger)
	router.GET("/", openAlchemyDoor)
	router.Start("8080", router)
}

func openAlchemyDoor(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../public/index.html")
}

`)

func CreateAlchemistMain() {
	err := ioutil.WriteFile("cmd/main.go", mainTemplate, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
