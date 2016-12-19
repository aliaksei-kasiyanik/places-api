package main

import (
	"log"
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/aliaksei-kasiyanik/places-api/web"
	"github.com/aliaksei-kasiyanik/places-api/repo"
)

func main() {
	mongoSession, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	defer mongoSession.Close()
	//session.SetMode(mgo.Monotonic, true)

	placesRepo := repo.NewPlacesRepo(mongoSession)

	router := web.PlaceApiRouter(placesRepo)
	log.Fatal(http.ListenAndServe(":8080", router))
}
