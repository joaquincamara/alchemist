package api

import (
	"encoding/json"
	"net/http"

	"github.com/joaquincamara/alchemist/example/internal/coolFeatures"
)

type ICoolFeaturesHandler interface {
	Post(http.ResponseWriter, *http.Request)
	GetAll(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
	Put(http.ResponseWriter, *http.Request)
}

type coolFeaturesHandler struct {
	coolFeaturesService coolFeatures.Service
}

func NewCoolFeaturesHandler(coolFeaturesService coolFeatures.Service) ICoolFeaturesHandler {
	return &coolFeaturesHandler{coolFeaturesService: coolFeaturesService}
}

func (h *coolFeaturesHandler) Post(w http.ResponseWriter, r *http.Request) {
	p := &coolFeatures.CoolFeatures{}
	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	err = h.coolFeaturesService.Add(p)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&p)
}

func (h *coolFeaturesHandler) Delete(w http.ResponseWriter, r *http.Request) {

	type id struct {
		Id int
	}
	coolFeaturesId := &id{}
	err := json.NewDecoder(r.Body).Decode(coolFeaturesId)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	err = h.coolFeaturesService.Delete(coolFeaturesId.Id)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(coolFeaturesId)
}

func (h *coolFeaturesHandler) GetAll(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	res, err := h.coolFeaturesService.FindAll()

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&res)
}

func (h *coolFeaturesHandler) Put(w http.ResponseWriter, r *http.Request) {

	p := &coolFeatures.CoolFeatures{}
	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	err = h.coolFeaturesService.Update(p)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&p)

}
