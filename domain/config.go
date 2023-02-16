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
}

type Flags struct {
	ConfigFile string
}

func (c Config) Error() error {
	err := c.yourselfError()

	if err != nil {
		return err
	}

	err = c.receiverConfigError()

	if err != nil {
		return err
	}

	err = c.senderConfigError()

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

func (c Config) receiverConfigError() error {
	if c.Receiver.Uri == "" {
		return errors.New(constants.TomlFileReceiverUriError)
	}

	if c.Receiver.Type == "" {
		return errors.New(constants.TomlFileReceiverTypeError)
	}

	return nil
}

func (c Config) senderConfigError() error {
	if c.Sender.Uri == "" {
		return errors.New(constants.TomlFileSenderUriError)
	}

	return nil
}
