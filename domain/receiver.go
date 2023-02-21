package domain

import (
	"context"
	"errors"
	"mongo_transporter/constants"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReceiverConfig struct {
	Uri         string `toml:"connection"`
	DisablleSSL bool   `toml:"disable-ssl,omitempty"`
	Region      string `toml:"region,omitempty"`
	Type        string `toml:"type,omitempty"`
}

func (r ReceiverConfig) Error() error {
	if r.Uri == "" {
		return errors.New(constants.TomlFileReceiverUriError)
	}

	if r.Type == "" {
		return errors.New(constants.TomlFileReceiverTypeError)
	}

	return nil
}

type Receiver interface {
	GetCollectionName() string
	SetupCollection(ctx context.Context)
	InsertOnCollection(ctx context.Context, documents []interface{})
	ReflectWatchOnInsert(ctx context.Context, fullDocument primitive.M)
	ReflectWatchOnDelete(ctx context.Context, id primitive.ObjectID)
	ReflectWatchOnUpdate(ctx context.Context, id primitive.ObjectID, updatedFields primitive.D, removedFields primitive.M)
}
