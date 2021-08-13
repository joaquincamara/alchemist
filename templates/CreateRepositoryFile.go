package templates

import (
	"io/ioutil"
	"log"
	"strings"
)

func CreateRespositoryFile(domainName string) {

	var repositoryTemplate = []byte(`package ` + domainName + `

type Repository interface {
	Add(` + domainName + ` *` + strings.Title(domainName) + `) error
	FindAll() ([]*` + strings.Title(domainName) + `, error)
	Update(` + domainName + ` *` + strings.Title(domainName) + `) error
	Delete(id int) error
}
	`)

	err := ioutil.WriteFile("internal/"+domainName+"/"+"repository.go", repositoryTemplate, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
