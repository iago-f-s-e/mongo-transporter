package utils

import "fmt"

func PrintWithCollection(collection string, messages ...any) {
	fmt.Printf("<%s> ", collection)
	fmt.Println(messages...)
}
