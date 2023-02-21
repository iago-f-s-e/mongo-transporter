package core

import (
	"log"
	"mongo_transporter/constants"
	"mongo_transporter/domain"
	"mongo_transporter/infra"
	"mongo_transporter/utils"
)

func decodeReceiver(config interface{}) domain.ReceiverConfig {
	receiverconfig := domain.ReceiverConfig{}
	receiver, ok := config.(map[string]interface{})["receiver"]

	if !ok {
		return receiverconfig
	}

	connection, ok := receiver.(map[string]interface{})["connection"]

	if !ok {
		return receiverconfig
	}

	receiverconfig.Uri = string(connection.(string))
	receiverconfig.Type = decodeReceiverType(receiver)
	receiverconfig.Region = decodeReceiverRegion(receiver)
	receiverconfig.AccessKeyId = decodeReceiverAccessKeyId(receiver)
	receiverconfig.SecretAccessKey = decodeReceiverSecretAccessKey(receiver)
	receiverconfig.SessionToken = decodeReceiverSessionToken(receiver)

	return receiverconfig
}

func decodeSender(config interface{}) domain.SenderCofing {
	senderconfig := domain.SenderCofing{}
	receiver, ok := config.(map[string]interface{})["sender"]

	if !ok {
		return senderconfig
	}

	connection, ok := receiver.(map[string]interface{})["connection"]

	if !ok {
		return senderconfig
	}

	senderconfig.Uri = string(connection.(string))

	return senderconfig
}

func decodeMapping(config interface{}) domain.MappingConfig {
	mappingConfig := domain.MappingConfig{}
	mappingCollectionConfig := make(map[string]string)

	mapping, ok := config.(map[string]interface{})["mapping"]

	if !ok {
		return mappingConfig
	}

	mappingList := mapping.([]map[string]interface{})

	for _, values := range mappingList {
		collectionName, okName := values["collection-name"]
		mapCollectionTo, okMap := values["collection-map"]

		if okName && okMap {
			mappingCollectionConfig[collectionName.(string)] = string(mapCollectionTo.(string))
		}
	}

	mappingConfig.Collection = mappingCollectionConfig

	return mappingConfig
}

func decodeBatchSize(config interface{}) int64 {
	batchSizeConfig, ok := config.(map[string]interface{})["batch-size"]

	if !ok {
		return 1000
	}

	batchSize, err := utils.ConvertInterfaceToInt64(batchSizeConfig)

	if err != nil {
		return -1
	}

	if batchSize < 1000 {
		batchSize = 1000
	}

	return batchSize
}

func decodeDbName(config interface{}) string {
	dbName, ok := config.(map[string]interface{})["database-name"]

	if !ok {
		return ""
	}

	return string(dbName.(string))
}

func decodeReceiverRegion(receiver interface{}) string {
	receiverRegion, ok := receiver.(map[string]interface{})["region"]

	if !ok {
		return constants.ReceiverDefaultLocally

	}

	return string(receiverRegion.(string))
}

func decodeReceiverAccessKeyId(receiver interface{}) string {
	accessKeyId, ok := receiver.(map[string]interface{})["access-key-id"]

	if !ok {
		return ""

	}

	return string(accessKeyId.(string))
}

func decodeReceiverSecretAccessKey(receiver interface{}) string {
	secretAccessKey, ok := receiver.(map[string]interface{})["secret-access-key"]

	if !ok {
		return ""

	}

	return string(secretAccessKey.(string))
}

func decodeReceiverSessionToken(receiver interface{}) string {
	sessionToken, ok := receiver.(map[string]interface{})["session-token"]

	if !ok {
		return ""

	}

	return string(sessionToken.(string))
}

func decodeReceiverType(receiver interface{}) string {

	receiverType, ok := receiver.(map[string]interface{})["type"]

	if !ok {
		return "mongodb"

	}
	if !utils.IsReceiverType(receiverType.(string)) {
		return ""
	}

	return string(receiverType.(string))
}

func decodeTransferCollections(config interface{}) []string {
	transferCollections, ok := config.(map[string]interface{})["transfer-collections"]
	var collections []string

	if !ok {
		return collections
	}

	for _, collection := range transferCollections.([]interface{}) {
		collections = append(collections, string(collection.(string)))
	}

	return collections
}

func decodeWatchCollections(config interface{}) []string {
	watchCollections, ok := config.(map[string]interface{})["watch-collections"]
	var collections []string

	if !ok {
		return collections
	}

	for _, collection := range watchCollections.([]interface{}) {
		collections = append(collections, string(collection.(string)))
	}

	return collections
}

func DecodeConfig(path string) (domain.Config, error) {
	config := domain.Config{}

	var configInterface interface{} = &domain.Config{}

	decodedConfig, err := infra.TomlFile(path, &configInterface)

	if err != nil {
		log.Fatal(err)

		return config, err
	}

	batchSize := decodeBatchSize(decodedConfig)
	dbName := decodeDbName(decodedConfig)
	transferCollections := decodeTransferCollections(decodedConfig)
	watchCollections := decodeWatchCollections(decodedConfig)

	config.BatchSize = batchSize
	config.DatabaseName = dbName
	config.TransferCollections = transferCollections
	config.WatchCollections = watchCollections
	config.Receiver = decodeReceiver(decodedConfig)
	config.Sender = decodeSender(decodedConfig)
	config.Mapping = decodeMapping(decodedConfig)

	err = config.Error()

	if err != nil {
		return config, err
	}

	return config, nil
}
