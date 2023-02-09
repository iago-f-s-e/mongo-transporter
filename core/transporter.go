package core

import (
	"context"
	"mongo_transporter/domain"
	"mongo_transporter/infra"
)

func Sender(ctx context.Context, dbUri string, dbName string, dbCollection string) domain.Sender {
	client := infra.MongoConnection(dbUri)

	sender := domain.NewSender(dbUri, dbName, dbCollection, client)

	return sender
}

func Receiver(ctx context.Context, dbUri string, dbName string, dbCollection string) domain.Receiver {
	client := infra.MongoConnection(dbUri)

	recevier := domain.NewReceiver(dbUri, dbName, dbCollection, client)

	return recevier
}
