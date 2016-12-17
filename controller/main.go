package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"encoding/json"
)

type Place struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

type Places []Place

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/index", Index)
	router.HandleFunc("/places", SearchPlaces)
	router.HandleFunc("/places/{id}", GetPlaceById)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func SearchPlaces(w http.ResponseWriter, r *http.Request) {
	places := Places{
		Place{Id: 1, Name: "Mir"},
		Place{Id: 2, Name: "Niasvizh"},
	}

	json.NewEncoder(w).Encode(places)
}

func GetPlaceById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintln(w, "place id:", id)
}
