package templates

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func CreateModelFile(domainName string, domainValues map[string]string) {

	var domainProperties string

	for k, v := range domainValues {
		domainProperties += fmt.Sprintf("%s %s\n", k, v)

	}

	fmt.Println(domainProperties)
	var modelTemplate = []byte(`package ` + domainName + `

type ` + strings.Title(domainName) + ` struct {
	` + domainProperties + `
}
	`)

	err := ioutil.WriteFile("internal/"+domainName+"/"+"model.go", modelTemplate, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
