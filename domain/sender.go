package domain

import (
	"context"
	"errors"
	"log"
	"mongo_transporter/constants"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SenderCofing struct {
	Uri string `toml:"connection"`
}

func (s SenderCofing) Error() error {
	if s.Uri == "" {
		return errors.New(constants.TomlFileSenderUriError)
	}

	return nil
}

type Sender struct {
	dbUri        string
	dbName       string
	dbCollection *mongo.Collection
	dbClient     *mongo.Client
	dbDatabase   *mongo.Database
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

func (s Sender) GetCollectionWithPagination(ctx context.Context, batchSize int64, lastPosition int64) ([]interface{}, int64, error) {
	var documents []interface{}

	findOptions := options.Find().SetLimit(int64(batchSize)).SetSkip(int64(lastPosition))

	cursor, err := s.dbCollection.Find(ctx, bson.D{}, findOptions)

	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var document bson.M

		err := cursor.Decode(&document)

		if err != nil {
			log.Fatal(err)
			return nil, -1, err
		}

		documents = append(documents, document)
	}

	return documents, lastPosition + batchSize, nil
}

func (s Sender) WatchCollection(ctx context.Context) *mongo.ChangeStream {
	pipeline := mongo.Pipeline{}

	watcher, err := s.dbCollection.Watch(ctx, pipeline)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return watcher
}
