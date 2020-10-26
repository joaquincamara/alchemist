package server

import (
	"io/ioutil"
	"log"
	"os"
)

func CreateAlchemistServer(appName string) {
	var appServer = ApiServer()
	os.Mkdir(appName, 0755)
	appServerErr := ioutil.WriteFile(appName+"/server.go", appServer, 0644)
	if appServerErr != nil {
		log.Fatal(appServerErr)
	}
}
