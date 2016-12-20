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

	Place struct {
		Id               bson.ObjectId `json:"-" bson:"_id"`
		Location         GeoJson       `json:"location" bson:"loc"`
		Name             string        `json:"name" bson:"name"`
		Description      string        `json:"description" bson:"desc,omitempty"`
		Categories       []string      `json:"categories" bson:"cat,omitempty"`
		Image            string        `json:"image" bson:"img,omitempty"`
		LastModifiedTime time.Time     `json:"-" bson:"lastModified"`
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

func (*Place) Validate() error {
	//TODO
	return nil
}

func (*Places) Validate() error {
	//TODO
	return nil
}

func (places Places) Wrap(r *http.Request) *PlacesWrapper {
	//var items []*PlaceWrapper
	items := make([]*PlaceWrapper, 0)
	for _, p := range places {
		items = append(items, p.wrapPlace())
	}
	pw := &PlacesWrapper{
		Items: items,
		Meta:  &PlacesMeta{Self: r.RequestURI},
	}
	return pw
}

func (p *Place) Wrap(r *http.Request) *PlaceWrapper {
	return &PlaceWrapper{
		Item: p,
		Meta: &PlaceMeta{Self: r.RequestURI},
	}
}

func (p *Place) wrapPlace() *PlaceWrapper {
	return &PlaceWrapper{
		Item: p,
		Meta: &PlaceMeta{"/places/" + p.Id.Hex()},
	}
}
