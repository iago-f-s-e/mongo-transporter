package main

import (
	"context"
	"fmt"
	"log"
	"mongo_transporter/core"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	config, err := core.DecodeConfig(core.DecodeFlag())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Starting sender instance...")

	sender :=
		core.Sender(ctx, config.Sender.Uri, config.Sender.DatabaseName, config.Sender.CollectionName)

	fmt.Println("Starting receiver instance...")

	receiver :=
		core.Receiver(ctx, config.Receiver.Uri, config.Receiver.DatabaseName, config.Receiver.CollectionName)

	fmt.Println("Getting the sender collection...")

	documents, _ := sender.GetCollection(ctx)

	fmt.Println("Successful get sender collection")

	fmt.Println("Inserting documents into the receiver collection...")

	receiver.InsertOnCollection(ctx, documents)

	fmt.Println("Successful insert all documents into the receiver collection")
}
