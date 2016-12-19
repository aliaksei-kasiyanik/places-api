package web

import (
	"github.com/julienschmidt/httprouter"

	"github.com/aliaksei-kasiyanik/places-api/repo"
)

func PlaceApiRouter(repo *repo.PlacesRepo) *httprouter.Router {

	router := httprouter.New()

	pc := NewPlacesController(repo)

	router.GET("/", pc.Index)

	router.GET("/places", pc.SearchPlaces)
	router.POST("/places", pc.CreatePlace)

	router.GET("/places/:id", pc.GetPlaceById)
	router.DELETE("/places/:id", pc.RemovePlace)

	return router
}
