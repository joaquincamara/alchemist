package templates

import (
	"io/ioutil"
	"log"
	"strings"
)

func CreateServiceFile(domainName string) {

	var serviceTemplate = []byte(`package ` + domainName + `

type Service interface {
	Add(` + domainName + ` *` + strings.Title(domainName) + `) error
	FindAll() ([]*` + strings.Title(domainName) + `, error)
	Update(` + domainName + ` *` + strings.Title(domainName) + `) error
	Delete(id int) error
}

type service struct {
	` + domainName + `Repo Repository
}

func New` + strings.Title(domainName) + `Service(` + domainName + `Repo Repository) Service {
	return &service{` + domainName + `Repo: ` + domainName + `Repo}
}
	
func (s *service) Add(` + domainName + ` *` + strings.Title(domainName) + `) error {
	return s.` + domainName + `Repo.Add(` + domainName + `)
}
	
func (s *service) FindAll() ([]*` + strings.Title(domainName) + `, error) {
	return s.` + domainName + `Repo.FindAll()
}
	
func (s *service) Delete(id int) error {
	return s.` + domainName + `Repo.Delete(id)
}
	
func (s *service) Update(` + domainName + ` *` + strings.Title(domainName) + `) error {
	return s.` + domainName + `Repo.Update(` + domainName + `)
}
`)

	err := ioutil.WriteFile("internal/"+domainName+"/"+"service.go", serviceTemplate, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
