package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"SearchPlaces",
		"GET",
		"/places",
		SearchPlaces,
	},
	Route{
		"GetPlaceById",
		"GET",
		"/places/{todoId}",
		GetPlaceById,
	},
	Route{
		"CreatePlace",
		"POST",
		"/places",
		CreatePlace,
	},
}
