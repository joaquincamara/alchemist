package templates

import (
	"fmt"
	"io/ioutil"
	"log"
)

func CreateModelFile(domainName string, domainValues map[string]string) {
	fmt.Println(domainValues)
	var modelTemplate = []byte(`
		package aboutMe

		type` + domainName + `struct {

		}
	`)

	err := ioutil.WriteFile("internal/"+domainName+"/"+"model.go", modelTemplate, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
