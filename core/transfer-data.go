package core

import (
	"context"
	"log"
	"mongo_transporter/domain"
	"mongo_transporter/utils"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func transferData(ctx context.Context, batchSize int64, receiver domain.Receiver, sender *domain.Sender, wg *sync.WaitGroup) {
	collectionName := receiver.GetCollectionName()

	var lastPosition int64 = 0
	count := 1
	receiver.SetupCollection(ctx)

	for {
		documents, newLastPosition, err := sender.GetCollectionWithPagination(ctx, batchSize, lastPosition)

		if err != nil {
			os.Exit(1)
		}

		if len(documents) == 0 {
			break
		}

		wg.Add(1)
		go func(documents []interface{}, batchNumber int) {
			defer wg.Done()

			utils.PrintWithCollection(collectionName, "[BATCH NUMBER]", batchNumber, "[BATCH SIZE]", len(documents))

			receiver.InsertOnCollection(ctx, documents)
		}(documents, count)

		lastPosition = newLastPosition
		count++
	}
}

func transferDataOnInsertEvent(ctx context.Context, event primitive.M, receiver domain.Receiver) {
	collectionName := receiver.GetCollectionName()

	fullDocument := utils.MakeMongoMap(event["fullDocument"])

	utils.PrintWithCollection(collectionName, "[INSERT]", fullDocument)

	receiver.ReflectWatchOnInsert(ctx, fullDocument)
}

func transferDataOnDeleteEvent(ctx context.Context, event primitive.M, receiver domain.Receiver) {
	collectionName := receiver.GetCollectionName()

	_id := event["documentKey"].(primitive.D).Map()["_id"].(primitive.ObjectID)

	utils.PrintWithCollection(collectionName, "[DELETE]", _id)

	receiver.ReflectWatchOnDelete(ctx, _id)
}

func transferDataOnUpdateEvent(ctx context.Context, event primitive.M, receiver domain.Receiver) {
	collectionName := receiver.GetCollectionName()

	_id := event["documentKey"].(primitive.D).Map()["_id"].(primitive.ObjectID)
	updateDescription := event["updateDescription"].(primitive.D).Map()

	updatedFields := updateDescription["updatedFields"].(primitive.D)
	removedFields := updateDescription["removedFields"].(primitive.A)

	mappedRemovedFields := bson.M{}

	for _, field := range removedFields {
		mappedRemovedFields[field.(string)] = 1
	}

	utils.PrintWithCollection(collectionName, "[UPDATE]", _id)
	utils.PrintWithCollection(collectionName, "[UPDATE FIELDS]", updatedFields)
	utils.PrintWithCollection(collectionName, "[REMOVE FIELDS]", mappedRemovedFields)

	receiver.ReflectWatchOnUpdate(ctx, _id, updatedFields, mappedRemovedFields)
}

func transferDataOnWatch(ctx context.Context, watcher *mongo.ChangeStream, receiver domain.Receiver, wg *sync.WaitGroup) {
	var mutex sync.Mutex

	for watcher.Next(ctx) {
		var event primitive.D

		err := watcher.Decode(&event)

		if err != nil {
			log.Fatal(err)
		}

		mappedEvent := event.Map()

		operationType, ok := mappedEvent["operationType"].(string)

		if !ok {
			log.Fatal(err)
		}

		wg.Add(1)
		mutex.Lock()
		go func(operation string, event primitive.M) {
			defer mutex.Unlock()
			defer wg.Done()

			switch operation {
			case "insert":
				transferDataOnInsertEvent(ctx, event, receiver)

			case "delete":
				transferDataOnDeleteEvent(ctx, event, receiver)

			case "update":
				transferDataOnUpdateEvent(ctx, event, receiver)

			default:
				log.Fatal("Invalid event type")
				os.Exit(1)
			}

		}(operationType, mappedEvent)
	}
}
