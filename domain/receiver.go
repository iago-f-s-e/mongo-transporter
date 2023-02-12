package domain

import (
	"context"
	"log"
	"mongo_transporter/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Receiver struct {
	dbUri          string
	dbName         string
	CollectionName string
	dbCollection   *mongo.Collection
	dbClient       *mongo.Client
	dbDatabase     *mongo.Database
}

type ReceiverConfig struct {
	Uri string `toml:"connection"`
}

func NewReceiver(dbUri string, dbName string, dbCollectName string, dbClient *mongo.Client) Receiver {
	database := dbClient.Database(dbName)
	collection := database.Collection(dbCollectName)

	return Receiver{
		dbUri:          dbUri,
		dbName:         dbName,
		CollectionName: dbCollectName,
		dbCollection:   collection,
		dbClient:       dbClient,
		dbDatabase:     database,
	}
}

func (r Receiver) DeleteAllCollection(ctx context.Context) {
	_, err := r.dbCollection.DeleteMany(ctx, bson.M{})

	if err != nil {
		log.Fatal(err)
	}
}

func (r Receiver) InsertOnCollection(ctx context.Context, documents []interface{}) {
	_, err := r.dbCollection.InsertMany(ctx, documents)

	if err != nil {
		log.Fatal(err)
	}
}

func (r Receiver) ReflectWatchOnInsert(ctx context.Context, fullDocument primitive.M) {
	_, err := r.dbCollection.InsertOne(ctx, fullDocument)

	if err != nil {
		log.Fatal(err)
	}
}

func (r Receiver) ReflectWatchOnDelete(ctx context.Context, id primitive.ObjectID) {
	res, err := r.dbCollection.DeleteOne(ctx, bson.M{"_id": id})

	if err != nil {
		log.Fatal(err)
	}

	utils.PrintWithCollection(r.CollectionName, "[DELETE COUNT]", res.DeletedCount)
}

func (r Receiver) ReflectWatchOnUpdate(ctx context.Context, id primitive.ObjectID, updatedFields primitive.D, removedFields primitive.M) {
	res, err := r.dbCollection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": updatedFields, "$unset": removedFields})

	if err != nil {
		log.Fatal(err)
	}

	utils.PrintWithCollection(r.CollectionName, "[UPDATE COUNT]", res.ModifiedCount)
}
