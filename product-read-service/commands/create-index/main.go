package main

import (
	"errors"
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
)

var (
	IndexNameEmptyStringError = errors.New("index name cannot be empty string")
	IndexAlreadyExistsError   = errors.New("elasticsearch index already exists")
)

func main() {
	index := os.Getenv("ES_INDEX")
	if index == "" {
		log.Println(IndexNameEmptyStringError)
		os.Exit(1)
	}

	elastic, err := elasticsearch.NewClient(elasticsearch.Config{
		// TODO: move to configuration
		Addresses: []string{"http://localhost:9201"},
	})

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	response, err := elastic.Indices.Exists([]string{index})
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	if response.StatusCode != 404 {
		log.Println(IndexAlreadyExistsError)
		os.Exit(1)
	}

	response, err = elastic.Indices.Create(index)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	if response.IsError() {
		log.Println(err)
		os.Exit(1)
	}
}
