package domain

import (
	"errors"
	"mongo_transporter/constants"
)

type Config struct {
	BatchSize           int64    `toml:"batch-size,omitempty"`
	DatabaseName        string   `toml:"database-name"`
	TransferCollections []string `toml:"transfer-collections"`
	WatchCollections    []string `toml:"watch-collections,omitempty"`
	Receiver            ReceiverConfig
	Sender              SenderCofing
	Mapping             MappingConfig
}

type Flags struct {
	ConfigFile string
}

func (c Config) Error() error {
	err := c.yourselfError()

	if err != nil {
		return err
	}

	err = c.Mapping.Error()

	if err != nil {
		return err
	}

	err = c.Receiver.Error()

	if err != nil {
		return err
	}

	err = c.Sender.Error()

	if err != nil {
		return err
	}

	return nil
}

func (c Config) yourselfError() error {
	if c.BatchSize < 0 {
		return errors.New(constants.TomlFileBatchSizeError)
	}

	if c.DatabaseName == "" {
		return errors.New(constants.TomlFileDbNameError)
	}

	if len(c.TransferCollections) == 0 {
		return errors.New(constants.TomlFileTransferCollectionsError)
	}

	return nil
}

func (c Config) GetCollectionMap(name string) string {
	collection, ok := c.Mapping.Collection[name]

	if !ok {
		return name
	}

	return collection
}
