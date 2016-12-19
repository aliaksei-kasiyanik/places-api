package repo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/aliaksei-kasiyanik/places-api/models"
)

type (
	PlacesRepo struct {
		session    *mgo.Session
		db         *mgo.Database
		collection *mgo.Collection
	}
)

func NewPlacesRepo(s *mgo.Session) *PlacesRepo {
	repo := &PlacesRepo{}
	repo.session = s
	repo.db = s.DB("places-api")
	repo.collection = s.DB("places-api").C("places")
	return repo
}

func (repo *PlacesRepo) InsertPlace(place *models.Place) (bson.ObjectId, error) {
	place.Id = bson.NewObjectId()

	// todo: several attempts
	err := repo.collection.Insert(&place)

	return place.Id, err
}

func (repo *PlacesRepo) FindAllPlaces() (*[]models.Place, error) {
	var result []models.Place
	iter := repo.collection.Find(nil).Limit(100).Iter()
	err := iter.All(&result)
	return &result, err
}

func (repo *PlacesRepo) FindPlaceById(oid *bson.ObjectId) (*models.Place, error) {
	var result models.Place
	err := repo.collection.Find(bson.M{"_id": oid}).One(&result)
	return &result, err
}

func (repo *PlacesRepo) RemovePlace(oid *bson.ObjectId) error {
	return repo.collection.RemoveId(oid)
}

