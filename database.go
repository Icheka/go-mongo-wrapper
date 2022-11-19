package database

import (
	"log"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var initialisationTriesCount = 1
var maxInitialisationTries = 6

func InitialiseDatabase(params struct {
	dbName      string
	mongoConfig *mgm.Config
	mongoURI    string
}) {
	var config *mgm.Config = nil
	if params.mongoConfig != nil {
		config = params.mongoConfig
	}

	err := mgm.SetDefaultConfig(
		config,
		"development",
		options.Client().ApplyURI(params.mongoURI),
	)
	if err != nil && initialisationTriesCount <= maxInitialisationTries {
		initialisationTriesCount++
		InitialiseDatabase(params)
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

func Update(collection *mgm.Collection, object mgm.Model) error {
	return collection.Update(object)
}

func Delete(collection *mgm.Collection, object mgm.Model) error {
	return collection.Delete(object)
}

func FindByID(id string, collection *mgm.Collection, obj mgm.Model) error {
	return collection.FindByID(id, obj)
}
