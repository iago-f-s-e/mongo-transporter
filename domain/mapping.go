package domain

import (
	"errors"
	"mongo_transporter/constants"
)

type MappingConfig struct {
	Collection map[string]string
}

func (m MappingConfig) Error() error {
	for key, mapCollectionTo := range m.Collection {
		if key == "" {
			return errors.New(constants.TomlFileMappingColletionNameError)
		}

		if mapCollectionTo == "" {
			return errors.New(constants.TomlFileMappingColletionMapError)
		}
	}

	return nil
}
