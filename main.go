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
		wgOnStart.Add(1)

		go core.Start(ctx, transferCollection, &config, &wgOnStart)
	}

	wgOnStart.Wait()

	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	var wgOnWatch sync.WaitGroup

	for _, watchCollection := range config.WatchCollections {
		wgOnWatch.Add(1)

		go core.Watch(ctx, watchCollection, &config, &wgOnWatch)
	}

	wgOnWatch.Wait()
}
