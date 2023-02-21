package adapters

import (
	"context"
	"log"
	"mongo_transporter/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r MongoReceiver) SetupCollection(ctx context.Context) {
	_, err := r.dbCollection.DeleteMany(ctx, bson.M{})

	if err != nil {
		log.Fatal(err)
	}
}

func (r MongoReceiver) InsertOnCollection(ctx context.Context, documents []interface{}) {
	options := options.InsertMany().SetOrdered(false)

	_, err := r.dbCollection.InsertMany(ctx, documents, options)

	if err != nil {
		log.Fatal(err)
	}
}

func (r MongoReceiver) ReflectWatchOnInsert(ctx context.Context, fullDocument primitive.M) {
	_, err := r.dbCollection.InsertOne(ctx, fullDocument)

	if err != nil {
		log.Fatal(err)
	}
}

func (r MongoReceiver) ReflectWatchOnDelete(ctx context.Context, id primitive.ObjectID) {
	res, err := r.dbCollection.DeleteOne(ctx, bson.M{"_id": id})

	if err != nil {
		log.Fatal(err)
	}

	utils.PrintWithCollection(r.collectionName, "[DELETE COUNT]", res.DeletedCount)
}

func (r MongoReceiver) ReflectWatchOnUpdate(ctx context.Context, id primitive.ObjectID, updatedFields primitive.D, removedFields primitive.M) {
	res, err := r.dbCollection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": updatedFields, "$unset": removedFields})

	if err != nil {
		log.Fatal(err)
	}

	utils.PrintWithCollection(r.collectionName, "[UPDATE COUNT]", res.ModifiedCount)
}
