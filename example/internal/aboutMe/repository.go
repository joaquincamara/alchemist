
package main

import (
	"github.com/joaquincamara/silver"
	"github.com/joaquincamara/silver/middleware"
)

func main() {
	router := silver.NewRouter()

	router.Use(middleware.Logger)
	router.GET("/", silver.AlchemyDoor)
	router.Start("8080", router)
}
