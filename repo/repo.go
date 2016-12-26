package repo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/aliaksei-kasiyanik/places-api/models"
	"log"
	"time"
)

type (
	PlacesRepo struct {
		session *mgo.Session
	}
)

func NewPlacesRepo(s *mgo.Session) *PlacesRepo {
	ensureGeoIndex(s)
	//mgo doesn't support partial indexes
	//ensureFoursquareIdIndex(s)
	return &PlacesRepo{s}
}

func ensureGeoIndex(s *mgo.Session) {
	session := s.Copy()
	defer session.Close()

	c := session.DB("places-api").C("places")

	index := mgo.Index{
		Key:      []string{"$2dsphere:loc", "name", "_id"},
		Bits:     26, // bits of precision; 26 bits is roughly equivalent to 2 feet or 60 centimeters of precision
		Name:     "GeoIndex",
		Unique:   true,
		DropDups: true,
	}
	log.Print("GeoIndex ensuring...")
	err := c.EnsureIndex(index)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("GeoIndex is created.")
}

func ensureFoursquareIdIndex(s *mgo.Session) {
	session := s.Copy()
	defer session.Close()

	c := session.DB("places-api").C("places")

	index := mgo.Index{
		Key: []string{"fsId"},
		Name: "FsIndex",
		Unique: true,
		DropDups: true,
	}

	log.Print("FoursquareIdIndex ensuring...")
	err := c.EnsureIndex(index)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("FoursquareIdIndex is created.")
}


func (repo *PlacesRepo) InsertPlace(place *models.Place) error {

	place.Id = bson.NewObjectId()

	place.LastModifiedTime = time.Now()

	session := repo.session.Copy()
	defer session.Close()

	return session.DB("places-api").C("places").Insert(&place)
}

func (repo *PlacesRepo) FindAllPlaces(sp *models.SearchParams) (models.Places, error) {

	var results models.Places

	session := repo.session.Copy()
	defer session.Close()

	err := session.DB("places-api").C("places").
		Find(nil).
		Select(bson.M{"id": 1, "name" : 1, "loc" : 1}).
		Skip(sp.Offset).Limit(sp.Limit).All(&results)
	return results, err
}

func (repo *PlacesRepo) FindPlacesByLocation(sp *models.SearchParams) (models.Places, error) {

	var results models.Places

	session := repo.session.Copy()
	defer session.Close()

	err := session.DB("places-api").C("places").
		Find(bson.M{
		"loc": bson.M{
			"$nearSphere": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float64{sp.Lon, sp.Lat},
				},
				"$maxDistance": sp.Rad,
				},
			},
	}).
		Select(bson.M{"id": 1, "name" : 1, "loc" : 1}).
		Skip(sp.Offset).Limit(sp.Limit).All(&results)

	return results, err
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

func (repo *PlacesRepo) UpdatePlace(place *models.Place) error {

	session := repo.session.Copy()
	defer session.Close()

	place.LastModifiedTime = time.Now()

	_, err := session.DB("places-api").C("places").UpsertId(place.Id, bson.M{"$set": place})
	return err
}

func (repo *PlacesRepo) PlaceExist(oid *bson.ObjectId) error {

	session := repo.session.Copy()
	defer session.Close()

	return session.DB("places-api").C("places").RemoveId(oid)
}
