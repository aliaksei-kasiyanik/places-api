package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"

	"github.com/aliaksei-kasiyanik/places-api/models"
	"github.com/aliaksei-kasiyanik/places-api/repo"
)

type (
	PlacesController struct {
		repo *repo.PlacesRepo
	}
)

func NewPlacesController(r *repo.PlacesRepo) *PlacesController {
	return &PlacesController{r}
}

func (pc PlacesController) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, "Welcome!")
}

func (pc PlacesController) SearchPlaces(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	places, err := pc.repo.FindAllPlaces()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(places); err != nil {
		panic(err)
	}
}

func (pc PlacesController) GetPlaceById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	oid := bson.ObjectIdHex(id)
	place, err := pc.repo.FindPlaceById(&oid)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}


	placeJson, _ := json.Marshal(place)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", placeJson)
}

func (pc PlacesController) CreatePlace(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	place := &models.Place{}

	if err := json.NewDecoder(r.Body).Decode(place); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		return
	}

	oid, err := pc.repo.InsertPlace(place)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(oid)
}

func (pc PlacesController) RemovePlace(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	oid := bson.ObjectIdHex(id)

	if err := pc.repo.RemovePlace(&oid); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
