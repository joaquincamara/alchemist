package templates

import (
	"io/ioutil"
	"log"
)

var readmeTemplate = []byte(`
## Alchemist README
`)

func CreateAlchemistReadme() {
	err := ioutil.WriteFile("README.md", readmeTemplate, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
