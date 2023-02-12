package core

import (
	"context"
	"fmt"
	"mongo_transporter/domain"
	"mongo_transporter/infra"
	"sync"
)

func sender(ctx context.Context, dbUri string, dbName string, dbCollection string) domain.Sender {
	client := infra.MongoConnection(ctx, dbUri)

	sender := domain.NewSender(dbUri, dbName, dbCollection, client)

	return sender
}

func receiver(ctx context.Context, dbUri string, dbName string, dbCollection string) domain.Receiver {
	client := infra.MongoConnection(ctx, dbUri)

	recevier := domain.NewReceiver(dbUri, dbName, dbCollection, client)

	return recevier
}

func Start(ctx context.Context, dbCollection string, config *domain.Config, wgOnStart *sync.WaitGroup) {
	defer wgOnStart.Done()

	fmt.Println("Start collection: ", dbCollection)

	receiver := receiver(ctx, config.Receiver.Uri, config.DatabaseName, dbCollection)
	sender := sender(ctx, config.Sender.Uri, config.DatabaseName, dbCollection)

	var wg sync.WaitGroup

	transferData(ctx, config.BatchSize, &receiver, &sender, &wg)

	wg.Wait()

	fmt.Println("End collection: ", dbCollection)
}

func Watch(ctx context.Context, dbCollection string, config *domain.Config, wgOnWatch *sync.WaitGroup) {
	defer wgOnWatch.Done()

	fmt.Println("Watch collection: ", dbCollection)

	sender := sender(ctx, config.Sender.Uri, config.DatabaseName, dbCollection)
	receiver := receiver(ctx, config.Receiver.Uri, config.DatabaseName, dbCollection)

	watcher := sender.WatchCollection(ctx)

	defer watcher.Close(ctx)

	var wg sync.WaitGroup

	transferDataOnWatch(ctx, watcher, &receiver, &wg)

	wg.Wait()
}
