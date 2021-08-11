package templates

import (
	"io/ioutil"
	"log"
)

func CreateServiceFile(domainName string) {
	err := ioutil.WriteFile("internal/"+domainName+"/"+"service.go", mainTemplate, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
