package api

import (
	"encoding/json"
	"net/http"

	"github.com/joaquincamara/alchemist/example/internal/aboutMe"
)

type IAboutMeHandler interface {
	Post(http.ResponseWriter, *http.Request)
	GetAll(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
	Put(http.ResponseWriter, *http.Request)
}

type aboutMeHandler struct {
	aboutMeService aboutMe.Service
}

func NewAboutMeHandler(aboutMeService aboutMe.Service) IAboutMeHandler {
	return &aboutMeHandler{aboutMeService: aboutMeService}
}

func (h *aboutMeHandler) Post(w http.ResponseWriter, r *http.Request) {
	p := &aboutMe.AboutMe{}
	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	err = h.aboutMeService.Add(p)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&p)
}

func (h *aboutMeHandler) Delete(w http.ResponseWriter, r *http.Request) {

	type id struct {
		Id int
	}
	aboutMeId := &id{}
	err := json.NewDecoder(r.Body).Decode(aboutMeId)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	err = h.aboutMeService.Delete(aboutMeId.Id)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(aboutMeId)
}

func (h *aboutMeHandler) GetAll(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	res, err := h.aboutMeService.FindAll()

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&res)
}

func (h *aboutMeHandler) Put(w http.ResponseWriter, r *http.Request) {

	p := &aboutMe.AboutMe{}
	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	err = h.aboutMeService.Update(p)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&p)

}
