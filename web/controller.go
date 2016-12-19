package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
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

func (pc PlacesController) SearchPlaces(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	places, err := pc.repo.FindAllPlaces()

	if err != nil {
		ErrorResponse(w, "Database Error", http.StatusInternalServerError)
		return
	}

	ResponseOK(w, places)
}

func (pc PlacesController) GetPlaceById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		ErrorResponse(w, "id is corrupted", http.StatusBadRequest)
		return
	}

	oid := bson.ObjectIdHex(id)
	place, err := pc.repo.FindPlaceById(&oid)
	if err != nil {
		switch err {
		default:
			ErrorResponse(w, "Database error", http.StatusInternalServerError)
			return
		case mgo.ErrNotFound:
			ErrorResponse(w, "Place not found", http.StatusNotFound)
			return
		}
	}

	ResponseOK(w, place)
}

func (pc PlacesController) CreatePlace(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	place := &models.Place{}
	if err := json.NewDecoder(r.Body).Decode(place); err != nil {
		ErrorResponse(w, "Place entity is corrupted", http.StatusBadRequest)
	}

	err := pc.repo.InsertPlace(place)
	if err != nil {
		ErrorResponse(w, "Database Error", http.StatusInternalServerError)
		return
	}

	Response(w, place, http.StatusCreated)
}

func (pc PlacesController) RemovePlace(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		ErrorResponse(w, "id is corrupted", http.StatusBadRequest)
		return
	}

	oid := bson.ObjectIdHex(id)
	if err := pc.repo.RemovePlace(&oid); err != nil {

		switch err {
		default:
			ErrorResponse(w, "Database error", http.StatusInternalServerError)
			return
		case mgo.ErrNotFound:
			ErrorResponse(w, "Place not found", http.StatusNotFound)
			return
		}

	}

	w.WriteHeader(http.StatusNoContent)
}

func Response(w http.ResponseWriter, v interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		ErrorResponse(w, "JSON Encode Error", http.StatusInternalServerError)
	}
}

func ResponseOK(w http.ResponseWriter, v interface{}) {
	Response(w, v, http.StatusOK)
}

func ErrorResponse(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	fmt.Fprintf(w, `{ "code": "%d", "message": %q}`, code, message)
}
