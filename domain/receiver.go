package domain

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Receiver struct {
	dbUri        string
	dbName       string
	dbCollection *mongo.Collection
	dbClient     *mongo.Client
	dbDatabase   *mongo.Database
}

type ReceiverConfig struct {
	Uri            string `toml:"string-connection"`
	DatabaseName   string `toml:"database-name"`
	CollectionName string `toml:"collection-name"`
}

func NewReceiver(dbUri string, dbName string, dbCollectName string, dbClient *mongo.Client) Receiver {
	database := dbClient.Database(dbName)
	collection := database.Collection(dbCollectName)

	return Receiver{
		dbUri:        dbUri,
		dbName:       dbName,
		dbCollection: collection,
		dbClient:     dbClient,
		dbDatabase:   database,
	}
}
func (r Receiver) CreateCollection(ctx context.Context, name string) {
	err := r.dbDatabase.CreateCollection(ctx, name)

	if err != nil {
		log.Fatal(err)
	}
}

func (r Receiver) DeleteAllCollection(ctx context.Context) {
	_, err := r.dbCollection.DeleteMany(ctx, bson.M{})

	if err != nil {
		log.Fatal(err)
	}
}

func (r Receiver) InsertOnCollection(ctx context.Context, documents []interface{}) {
	r.DeleteAllCollection(ctx)

	_, err := r.dbCollection.InsertMany(ctx, documents)

	if err != nil {
		log.Fatal(err)
	}
}
