package main

import (
	"errors"
	"log"
	"os"
)

var (
	PathEmptyStringError = errors.New("path to json file with data cannot be empty string")
)

func main() {
	data := os.Getenv("ES_DATA_PATH")
	if data == "" {
		log.Println(PathEmptyStringError)
		os.Exit(1)
	}
}
