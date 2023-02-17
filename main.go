package main

import (
	"context"
	"log"
	"mongo_transporter/core"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	config, err := core.DecodeConfig(core.DecodeFlag())

	if err != nil {
		log.Fatal(err)
	}

	var wgOnStart sync.WaitGroup

	for _, transferCollection := range config.TransferCollections {
		collectionMap := config.GetCollectionMap(transferCollection)

		wgOnStart.Add(1)

		go core.Start(ctx, transferCollection, collectionMap, &config, &wgOnStart)
	}

	wgOnStart.Wait()

	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	var wgOnWatch sync.WaitGroup

	for _, watchCollection := range config.WatchCollections {
		collectionMap := config.GetCollectionMap(watchCollection)

		wgOnWatch.Add(1)

		go core.Watch(ctx, watchCollection, collectionMap, &config, &wgOnWatch)
	}

	wgOnWatch.Wait()
}
