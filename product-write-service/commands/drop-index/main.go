package main

import (
	"errors"
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
)

var (
	IndexNameEmptyStringError = errors.New("index name cannot be empty string")
	IndexNotExistsError       = errors.New("elasticsearch index does not exist")
)

func main() {
	index := os.Getenv("ES_INDEX")
	if index == "" {
		log.Println(IndexNameEmptyStringError)
		os.Exit(1)
	}

	elastic, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
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

	if response.StatusCode != 200 {
		log.Println(IndexNotExistsError)
		os.Exit(1)
	}

	response, err = elastic.Indices.Delete([]string{index})
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	if response.IsError() {
		log.Println(err)
		os.Exit(1)
	}
}
