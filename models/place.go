package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
	"net/http"
)

type (
	GeoJson struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	}

	BasePlace struct {
		Id       bson.ObjectId `json:"-" bson:"_id"`
		Location GeoJson       `json:"location" bson:"loc"`
		Name     string        `json:"name" bson:"name"`
	}

	Place struct {
		Id               bson.ObjectId `json:"-" bson:"_id"`
		Location         GeoJson       `json:"location" bson:"loc"`
		Name             string        `json:"name" bson:"name"`
		Description      string        `json:"description" bson:"desc,omitempty"`
		Categories       []string      `json:"categories" bson:"cat,omitempty"`
		Image            string        `json:"image" bson:"img,omitempty"`
		LastModifiedTime time.Time     `json:"-" bson:"lastModified"`
	}

	Places []*BasePlace

	PlaceMeta struct {
		Self string `json:"self"`
	}

	PlaceWrapper struct {
		Item *Place     `json:"item"`
		Meta *PlaceMeta `json:"_meta"`
	}

	BasePlaceWrapper struct {
		Item *BasePlace     `json:"item"`
		Meta *PlaceMeta `json:"_meta"`
	}

	PlacesMeta struct {
		Self string `json:"self"`
		Next string `json:"next,omitempty"`
		Prev string `json:"prev,omitempty"`
	}

	PlacesWrapper struct {
		Items []*BasePlaceWrapper `json:"items"`
		Meta  *PlacesMeta     `json:"_meta"`
	}
)

func (p *Place) Validate() string {
	if p.Name == "" {
		return "Name is not provided."
	}
	if p.Location.Coordinates == nil || len(p.Location.Coordinates) != 2 || p.Location.Type != "Point" {
		return "Location is corrupted."
	}
	return ""
}

func (places Places) Wrap(r *http.Request, searchParams *SearchParams, resultCount int) *PlacesWrapper {
	items := make([]*BasePlaceWrapper, 0)
	for _, p := range places {
		items = append(items, p.wrapPlace())
	}
	pw := &PlacesWrapper{
		Items: items,
		Meta: createMeta(r, searchParams, resultCount),
	}
	return pw
}

func (p *Place) Wrap(r *http.Request) *PlaceWrapper {
	return &PlaceWrapper{
		Item: p,
		Meta: &PlaceMeta{Self: r.RequestURI},
	}
}

func (p *BasePlace) wrapPlace() *BasePlaceWrapper {
	return &BasePlaceWrapper{
		Item: p,
		Meta: &PlaceMeta{"/places/" + p.Id.Hex()},
	}
}

func createMeta(r *http.Request, searchParams *SearchParams, resultCount int) *PlacesMeta {
	prev := ""
	prevParams := searchParams.getPrev()
	if prevParams != "" {
		prev = r.URL.Path + prevParams
	}
	next := ""
	if resultCount == searchParams.Limit {
		nextParams := searchParams.getNext()
		next = r.URL.Path + nextParams
	}
	meta := &PlacesMeta{Self: r.RequestURI, Next: next, Prev: prev}
	return meta
}
