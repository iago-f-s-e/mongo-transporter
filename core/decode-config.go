package core

import (
	"errors"
	"log"
	"mongo_transporter/constants"
	"mongo_transporter/domain"
	"mongo_transporter/infra"
)

func decodeReceiver(config interface{}) (domain.ReceiverConfig, error) {
	receiverconfig := domain.ReceiverConfig{}
	receiver, ok := config.(map[string]interface{})["receiver"]

	if !ok {
		return receiverconfig, errors.New(constants.TomlFileError)
	}

	collection, ok := receiver.(map[string]interface{})["collection-name"]

	if !ok {
		return receiverconfig, errors.New(constants.TomlFileError)
	}

	database, ok := receiver.(map[string]interface{})["database-name"]

	if !ok {
		return receiverconfig, errors.New(constants.TomlFileError)
	}

	connection, ok := receiver.(map[string]interface{})["string-connection"]

	if !ok {
		return receiverconfig, errors.New(constants.TomlFileError)
	}

	receiverconfig.CollectionName = string(collection.(string))
	receiverconfig.DatabaseName = string(database.(string))
	receiverconfig.Uri = string(connection.(string))

	return receiverconfig, nil
}

func decodeSender(config interface{}) (domain.SenderCofing, error) {
	senderconfig := domain.SenderCofing{}
	receiver, ok := config.(map[string]interface{})["sender"]

	if !ok {
		return senderconfig, errors.New(constants.TomlFileError)
	}

	collection, ok := receiver.(map[string]interface{})["collection-name"]

	if !ok {
		return senderconfig, errors.New(constants.TomlFileError)
	}

	database, ok := receiver.(map[string]interface{})["database-name"]

	if !ok {
		return senderconfig, errors.New(constants.TomlFileError)
	}

	connection, ok := receiver.(map[string]interface{})["string-connection"]

	if !ok {
		return senderconfig, errors.New(constants.TomlFileError)
	}

	senderconfig.CollectionName = string(collection.(string))
	senderconfig.DatabaseName = string(database.(string))
	senderconfig.Uri = string(connection.(string))

	return senderconfig, nil
}

func DecodeConfig(path string) (domain.Config, error) {
	config := domain.Config{}

	var configInterface interface{} = &domain.Config{}

	decodedConfig, err := infra.TomlFile(path, &configInterface)

	if err != nil {
		log.Fatal(err)

		return config, err
	}

	receiverConfig, err := decodeReceiver(decodedConfig)

	if err != nil {
		log.Fatal(err)

		return config, err
	}

	senderConfig, err := decodeSender(decodedConfig)

	if err != nil {
		log.Fatal(err)

		return config, err
	}

	config.Receiver = receiverConfig
	config.Sender = senderConfig

	if !config.IsValid() {
		return config, errors.New(constants.TomlFileError)
	}

	return config, nil
}
