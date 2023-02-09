package infra

import (
	"log"

	"github.com/BurntSushi/toml"
)

func TomlFile(path string, config *interface{}) (interface{}, error) {
	if _, err := toml.DecodeFile(path, &config); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return *config, nil
}
