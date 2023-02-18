package core

import (
	"context"
	"fmt"
	"mongo_transporter/adapters"
	"mongo_transporter/constants"
	"mongo_transporter/domain"
	"mongo_transporter/infra"
	"os"
	"sync"
)

func sender(ctx context.Context, dbUri string, dbName string, dbCollection string) domain.Sender {
	client := infra.MongoConnection(ctx, dbUri)

	sender := domain.NewSender(dbUri, dbName, dbCollection, client)

	return sender
}

func receiver(ctx context.Context, collection string, config *domain.Config) domain.Receiver {
	switch config.Receiver.Type {

	case constants.ReceiverTypeDynamoDb: // WIP
		infra.DynamoConnection(config.Receiver.Uri, config.Receiver.Region, config.Receiver.DisablleSSL)

		os.Exit(1)

		var wip domain.Receiver

		return wip

	default:
		client := infra.MongoConnection(ctx, config.Receiver.Uri)

		recevier := adapters.NewMongoReceiver(config.Receiver.Uri, config.DatabaseName, collection, client)

		return recevier
	}
}

func Start(ctx context.Context, dbCollection string, mapCollection string, config *domain.Config, wgOnStart *sync.WaitGroup) {
	defer wgOnStart.Done()

	fmt.Println("Start collection: ", dbCollection)

	sender := sender(ctx, config.Sender.Uri, config.DatabaseName, dbCollection)
	receiver := receiver(ctx, mapCollection, config)

	var wg sync.WaitGroup

	transferData(ctx, config.BatchSize, receiver, &sender, &wg)

	wg.Wait()

	fmt.Println("End collection: ", dbCollection)
}

func Watch(ctx context.Context, dbCollection string, mapCollection string, config *domain.Config, wgOnWatch *sync.WaitGroup) {
	defer wgOnWatch.Done()

	fmt.Println("Watch collection: ", dbCollection)

	sender := sender(ctx, config.Sender.Uri, config.DatabaseName, dbCollection)
	receiver := receiver(ctx, mapCollection, config)

	watcher := sender.WatchCollection(ctx)

	defer watcher.Close(ctx)

	var wg sync.WaitGroup

	transferDataOnWatch(ctx, watcher, receiver, &wg)

	wg.Wait()
}
