package main

import (
	"github.com/urfave/negroni"
	"gopkg.in/mgo.v2"

	"github.com/aliaksei-kasiyanik/places-api/repo"
	"github.com/aliaksei-kasiyanik/places-api/web"
	"github.com/aliaksei-kasiyanik/places-api/utils"
)

func main() {

	config := utils.NewConfiguration("config.json")

	session, err := mgo.DialWithInfo(config.GetMongoDialInfo())
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	placesRepo := repo.NewPlacesRepo(session)
	router := web.PlaceApiRouter(placesRepo)

	n := negroni.New(negroni.NewRecovery(), web.NewLogger())
	n.UseHandler(router)

	n.Run(config.AppAddr)
}
