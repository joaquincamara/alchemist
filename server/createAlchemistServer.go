package server

import (
	"io/ioutil"
	"log"
)

func CreateAlchemistServer(appName string) {
	AppServer()

	appServerErr := ioutil.WriteFile(appName+"/server.go", AppServer(), 0644)
	if appServerErr != nil {
		log.Fatal(appServerErr)
	}
}
