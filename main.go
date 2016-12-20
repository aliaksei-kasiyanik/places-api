package main

import (
	"os"

	"github.com/urfave/negroni"
	"gopkg.in/mgo.v2"

	"github.com/aliaksei-kasiyanik/places-api/repo"
	"github.com/aliaksei-kasiyanik/places-api/utils"
	"github.com/aliaksei-kasiyanik/places-api/web"
)

const (
	DEFAULT_CONFIG_PATH = "config.json"
)

func main() {
	config := utils.NewConfiguration(getConfigPath())

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

func getConfigPath() string {
	if len(os.Args) == 2 {
		return os.Args[1]
	} else {
		return DEFAULT_CONFIG_PATH
	}
}
