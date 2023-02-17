package core

import (
	"context"
	"fmt"
	"mongo_transporter/adapters"
	"mongo_transporter/domain"
	"mongo_transporter/infra"
	"sync"
)

func sender(ctx context.Context, dbUri string, dbName string, dbCollection string) domain.Sender {
	client := infra.MongoConnection(ctx, dbUri)

	sender := domain.NewSender(dbUri, dbName, dbCollection, client)

	return sender
}

func receiver(ctx context.Context, dbUri string, dbName string, dbCollection string, receiverType string) domain.Receiver {
	switch receiverType {

	default:
		client := infra.MongoConnection(ctx, dbUri)

		recevier := adapters.NewMongoReceiver(dbUri, dbName, dbCollection, client)

		return recevier
	}
}

func Start(ctx context.Context, dbCollection string, mapCollection string, config *domain.Config, wgOnStart *sync.WaitGroup) {
	defer wgOnStart.Done()

	fmt.Println("Start collection: ", dbCollection)

	sender := sender(ctx, config.Sender.Uri, config.DatabaseName, dbCollection)
	receiver := receiver(ctx, config.Receiver.Uri, config.DatabaseName, mapCollection, config.Receiver.Type)

	var wg sync.WaitGroup

	transferData(ctx, config.BatchSize, receiver, &sender, &wg)

	wg.Wait()

	fmt.Println("End collection: ", dbCollection)
}

func Watch(ctx context.Context, dbCollection string, mapCollection string, config *domain.Config, wgOnWatch *sync.WaitGroup) {
	defer wgOnWatch.Done()

	fmt.Println("Watch collection: ", dbCollection)

	sender := sender(ctx, config.Sender.Uri, config.DatabaseName, dbCollection)
	receiver := receiver(ctx, config.Receiver.Uri, config.DatabaseName, mapCollection, config.Receiver.Type)

	watcher := sender.WatchCollection(ctx)

	defer watcher.Close(ctx)

	var wg sync.WaitGroup

	transferDataOnWatch(ctx, watcher, receiver, &wg)

	wg.Wait()
}
