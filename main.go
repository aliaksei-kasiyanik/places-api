package main

import (
	"gopkg.in/mgo.v2"
	"github.com/codegangsta/negroni"

	"github.com/aliaksei-kasiyanik/places-api/repo"
	"github.com/aliaksei-kasiyanik/places-api/web"
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

	n := negroni.New(negroni.NewRecovery(), negroni.NewLogger())
	n.UseHandler(router)
	n.Run(":8080")
}
