package repo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/aliaksei-kasiyanik/places-api/models"
)

type (
	PlacesRepo struct {
		session    *mgo.Session
	}
)

func NewPlacesRepo(s *mgo.Session) *PlacesRepo {
	return &PlacesRepo{s}
}

func (repo *PlacesRepo) InsertPlace(place *models.Place) (bson.ObjectId, error) {

	place.Id = bson.NewObjectId()

	session := repo.session.Copy()
	defer session.Close()

	// todo: several attempts
	err := session.DB("places-api").C("places").Insert(&place)

	return place.Id, err
}

func (repo *PlacesRepo) FindAllPlaces() (*[]models.Place, error) {

	var result []models.Place

	session := repo.session.Copy()
	defer session.Close()

	iter := session.DB("places-api").C("places").Find(nil).Limit(100).Iter()
	err := iter.All(&result)
	return &result, err
}

func (repo *PlacesRepo) FindPlaceById(oid *bson.ObjectId) (*models.Place, error) {

	var result models.Place

	session := repo.session.Copy()
	defer session.Close()

	err := session.DB("places-api").C("places").Find(bson.M{"_id": oid}).One(&result)
	return &result, err
}

func (repo *PlacesRepo) RemovePlace(oid *bson.ObjectId) error {

	session := repo.session.Copy()
	defer session.Close()

	return session.DB("places-api").C("places").RemoveId(oid)
}
