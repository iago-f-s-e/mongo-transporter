package domain

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Sender struct {
	dbUri        string
	dbName       string
	dbCollection *mongo.Collection
	dbClient     *mongo.Client
	dbDatabase   *mongo.Database
}

type SenderCofing struct {
	Uri            string `toml:"string-connection"`
	DatabaseName   string `toml:"database-name"`
	CollectionName string `toml:"collection-name"`
}

func NewSender(dbUri string, dbName string, dbCollectName string, dbClient *mongo.Client) Sender {
	database := dbClient.Database(dbName)
	collection := database.Collection(dbCollectName)

	return Sender{
		dbUri:        dbUri,
		dbName:       dbName,
		dbCollection: collection,
		dbClient:     dbClient,
		dbDatabase:   database,
	}
}

func (s Sender) GetCollection(ctx context.Context) ([]interface{}, error) {
	var documents []interface{}

	cursor, err := s.dbCollection.Find(ctx, bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	for cursor.Next(ctx) {
		var document bson.M

		err := cursor.Decode(&document)

		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		documents = append(documents, document)
	}

	return documents, nil
}
