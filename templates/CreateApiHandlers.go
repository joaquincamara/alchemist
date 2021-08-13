package templates

import (
	"io/ioutil"
	"log"
	"strings"
)

func CreateApiHandlers(domainName string) {

	var apiHandlerTemplate = []byte(`package api

	import (
		"encoding/json"
		"net/http"
	)
	
	type I` + strings.Title(domainName) + `Handler interface {
		Post(http.ResponseWriter, *http.Request)
		GetAll(http.ResponseWriter, *http.Request)
		Delete(http.ResponseWriter, *http.Request)
		Put(http.ResponseWriter, *http.Request)
	}
	
	type ` + domainName + `Handler struct {
		` + domainName + `Service ` + domainName + `.Service
	}
	
	func New` + strings.Title(domainName) + `Handler(` + domainName + `Service ` + domainName + `.Service) I` + strings.Title(domainName) + `Handler {
		return &` + domainName + `Handler{` + domainName + `Service: ` + domainName + `Service}
	}
	
	func (h *` + domainName + `Handler) Post(w http.ResponseWriter, r *http.Request) {
		p := &` + domainName + `.` + strings.Title(domainName) + `{}
		err := json.NewDecoder(r.Body).Decode(&p)
	
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()
	
		err = h.` + domainName + `Service.Add(p)
	
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(&p)
	}
	
	func (h *` + domainName + `Handler) Delete(w http.ResponseWriter, r *http.Request) {
	
		type id struct {
			Id int
		}
		` + domainName + `Id := &id{}
		err := json.NewDecoder(r.Body).Decode(` + domainName + `Id)
	
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	
		defer r.Body.Close()
	
		err = h.` + domainName + `Service.Delete(` + domainName + `Id.Id)
	
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(` + domainName + `Id)
	}
	
	func (h *` + domainName + `Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
	
		res, err := h.` + domainName + `Service.FindAll()
	
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(&res)
	}
	
	func (h *` + domainName + `Handler) Put(w http.ResponseWriter, r *http.Request) {
	
		p := &` + domainName + `.` + strings.Title(domainName) + `{}
		err := json.NewDecoder(r.Body).Decode(&p)
	
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()
	
		err = h.` + domainName + `Service.Update(p)
	
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(&p)
	
	}
	
	`)

	err := ioutil.WriteFile("api/"+domainName+"Handler.go", apiHandlerTemplate, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
