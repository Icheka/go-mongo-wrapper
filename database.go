package database

import (
	"log"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
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

func AddIndex(collection *mgm.Collection, config struct {
	field     string
	ascending bool
	unique    bool
}) (string, error) {
	ctx := mgm.Ctx()
	var sort int32 = 1
	if !config.ascending {
		sort = -1
	}
	index := mongo.IndexModel{
		Keys: bsonx.Doc{{
			Key:   config.field,
			Value: bsonx.Int32(sort),
		}},
	}
	if config.unique {
		t := true
		index.Options.Unique = &t
	}

	return collection.Indexes().CreateOne(ctx, index)
}

func AddIndexes(collection *mgm.Collection, config struct {
	fields []struct {
		field     string
		ascending bool
		unique    bool
	}
}) ([]string, error) {
	ctx := mgm.Ctx()

	indexes := []mongo.IndexModel{}
	for _, conf := range config.fields {
		var sort int32 = 1
		if !conf.ascending {
			sort = -1
		}

		index := mongo.IndexModel{
			Keys: bsonx.Doc{{
				Key:   conf.field,
				Value: bsonx.Int32(sort),
			}},
		}
		if conf.unique {
			t := true
			index.Options.Unique = &t
		}

		indexes = append(indexes, index)
	}

	return collection.Indexes().CreateMany(ctx, indexes)
}
