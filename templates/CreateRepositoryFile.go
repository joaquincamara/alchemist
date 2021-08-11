package templates

import (
	"io/ioutil"
	"log"
)

func CreateRespositoryFile(domainName string) {
	err := ioutil.WriteFile("internal/"+domainName+"/"+"repository.go", mainTemplate, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
