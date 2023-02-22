package constants

import "fmt"

const TomlFileError string = "[ERROR] the toml file is not a valid mongo-transport configuration file"

const TomlFileBatchSizeError string = "[ERROR] batch size is not valid, must be an integer"

const TomlFileDbNameError string = "[ERROR] database name is not valid, must not be empty"

const TomlFileTransferCollectionsError string = "[ERROR] transfer collections is not valid, must not be empty"

const TomlFileReceiverUriError string = "[ERROR] receiver connection is not valid, must not be empty"

const TomlFileReceiverAccessKeyError string = "[ERROR] receiver access key id is not valid, must not be empty"

const TomlFileReceiverSecretKeyError string = "[ERROR] receiver secret access key is not valid, must not be empty"

const TomlFileSenderUriError string = "[ERROR] sender connection is not valid, must not be empty"

const TomlFileMappingColletionNameError string = "[ERROR] mapping is not valid, collection name must not be empty"

const TomlFileMappingColletionMapError string = "[ERROR] mapping is not valid, collection map must not be empty"

var TomlFileReceiverTypeError string = fmt.Sprint("[ERROR] receiver type is not valid, must be on the list of compatible receivers: ", ReceiverTypes)
