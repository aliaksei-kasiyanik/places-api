package models

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	GeoJson struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	}

	Place struct {
		Id          bson.ObjectId `json:"id" bson:"_id"`
		Location    GeoJson       `json:"location" bson:"loc"`
		Name        string        `json:"name" bson:"name"`
		Description string        `json:"description" bson:"desc,omitempty"`
		Categories  []string      `json:"categories" bson:"cat,omitempty"`
		Image       string        `json:"image" bson:"img,omitempty"`
	}

	PlaceMeta struct {
		Self string `json:"self"`
	}

	PlaceWrapper struct {
		Item *Place    `json:"item"`
		Meta PlaceMeta `json:"_meta"`
	}

	PlacesMeta struct {
		Self string `json:"self"`
		Next string `json:"next,omitempty"`
		Prev string `json:"prev,omitempty"`
	}

	PlacesWrapper struct {
		Items []PlacesWrapper `json:"items"`
		Meta  PlacesMeta      `json:"_meta"`
	}
)
