package core

import (
	"flag"
	"log"
	"os"
)

func DecodeFlag() string {

	flagName := flag.String("f", "", "Location of configuration file")
	flag.Parse()

	if *flagName == "" {
		log.Fatal("Error: flag 'f' was not provided")
		os.Exit(1)
	}

	return *flagName
}
