package database

import (
	"log"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var initialisationTriesCount = 1
var maxInitialisationTries = 6

func InitialiseDatabase(dbName string, mongoURI string) {
	err := mgm.SetDefaultConfig(
		nil,
		dbName,
		options.Client().ApplyURI(mongoURI),
	)
	if err != nil && initialisationTriesCount <= maxInitialisationTries {
		initialisationTriesCount++
		InitialiseDatabase(dbName, mongoURI)
		return
	} else if err != nil {
		log.Println("MongoDB connection timed out!")
		return
	}
	log.Println("MongoDB connection instituted!")
}

func Collection(collection mgm.Model) *mgm.Collection {
	return mgm.Coll(collection)
}

func InsertOne(collection *mgm.Collection, object mgm.Model) error {
	return collection.Create(object)
}

func Find(collection *mgm.Collection, result interface{}, filter bson.M) error {
	return collection.SimpleFind(result, filter)
}

func First(collection *mgm.Collection, object mgm.Model, filter bson.M) error {
	return collection.First(filter, object)
}

func Replace(collection *mgm.Collection, object mgm.Model) error {
	return collection.Update(object)
}

func Delete(collection *mgm.Collection, object mgm.Model) error {
	return collection.Delete(object)
}

func FindByID(id string, collection *mgm.Collection, obj mgm.Model) error {
	return collection.FindByID(id, obj)
}
