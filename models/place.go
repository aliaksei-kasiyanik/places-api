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
		Id          bson.ObjectId `json:"-" bson:"_id"`
		Location    GeoJson       `json:"location" bson:"loc"`
		Name        string        `json:"name" bson:"name"`
		Description string        `json:"description" bson:"desc,omitempty"`
		Categories  []string      `json:"categories" bson:"cat,omitempty"`
		Image       string        `json:"image" bson:"img,omitempty"`
	}

	Places []*Place

	PlaceMeta struct {
		Self string `json:"self"`
	}

	PlaceWrapper struct {
		Item *Place     `json:"item"`
		Meta *PlaceMeta `json:"_meta"`
	}

	PlacesMeta struct {
		Self string `json:"self"`
		Next string `json:"next,omitempty"`
		Prev string `json:"prev,omitempty"`
	}

	PlacesWrapper struct {
		Items []*PlaceWrapper `json:"items"`
		Meta  *PlacesMeta     `json:"_meta"`
	}
)

func (places Places) WrapPlaces(self string) *PlacesWrapper {
	var items []*PlaceWrapper
	for _, p := range places {
		items = append(items, p.WrapPlace())
	}
	pw := &PlacesWrapper{
		Items: items,
		Meta:  &PlacesMeta{Self: self},
	}
	return pw
}

func (p *Place) WrapPlace() *PlaceWrapper {
	return &PlaceWrapper{
		Item: p,
		Meta: &PlaceMeta{"/places/" + p.Id.Hex()},
	}
}
