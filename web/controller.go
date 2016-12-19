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
		w.WriteHeader(404)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(places); err != nil {
		panic(err)
	}
}

func (pc PlacesController) GetPlaceById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)
	place, err := pc.repo.FindPlaceById(&oid)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	w.WriteHeader(200)

	placeJson, _ := json.Marshal(place)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", placeJson)
}

func (pc PlacesController) CreatePlace(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	place := &models.Place{}

	if err := json.NewDecoder(r.Body).Decode(place); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	oid, err := pc.repo.InsertPlace(place)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(oid)
}

func (pc PlacesController) RemovePlace(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)

	if err := pc.repo.RemovePlace(&oid); err != nil {
		w.WriteHeader(404)
		return
	}

	w.WriteHeader(200)
}
