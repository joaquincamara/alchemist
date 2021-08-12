package templates

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func CreateModelFile(domainName string, domainValues map[string]string) {
	fmt.Println(domainValues)
	var modelTemplate = []byte(`package ` + domainName + `

type ` + strings.Title(domainName) + ` struct {

}
	`)

	err := ioutil.WriteFile("internal/"+domainName+"/"+"model.go", modelTemplate, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
