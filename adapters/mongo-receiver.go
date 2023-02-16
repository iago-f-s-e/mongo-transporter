package adapters

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoReceiver struct {
	dbUri          string
	dbName         string
	collectionName string
	dbCollection   *mongo.Collection
	dbClient       *mongo.Client
	dbDatabase     *mongo.Database
}

func NewMongoReceiver(dbUri string, dbName string, dbCollectName string, dbClient *mongo.Client) MongoReceiver {
	database := dbClient.Database(dbName)
	collection := database.Collection(dbCollectName)

	return MongoReceiver{
		dbUri:          dbUri,
		dbName:         dbName,
		collectionName: dbCollectName,
		dbCollection:   collection,
		dbClient:       dbClient,
		dbDatabase:     database,
	}
}

func (r MongoReceiver) GetCollectionName() string {
	return r.collectionName
}
