package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/elastic/go-elasticsearch/esapi"
	elasticsearch "github.com/elastic/go-elasticsearch/v8"
	"github.com/mikolajsemeniuk/CQRS-GRPC-Go/product-write-service/messages"
)

var (
	ProductNotFoundError = errors.New("Product not found")
)

type Product interface {
	Create(product messages.Product) error
	Update(product messages.Product) error
	Remove(id string) error
}

type product struct {
	index  string
	client *elasticsearch.Client
}

func (p product) Create(product messages.Product) error {
	bytes, err := json.Marshal(product)
	if err != nil {
		return err
	}

	request := esapi.IndexRequest{
		Index:      p.index,
		DocumentID: product.Id,
		Body:       strings.NewReader(string(bytes)),
	}

	response, err := request.Do(context.Background(), p.client)
	if err != nil {
		return err
	}

	if response.IsError() {
		return err
	}

	defer func() {
		err = response.Body.Close()
	}()

	return err
}

func (p product) Update(product messages.Product) error {
	bytes, err := json.Marshal(product)
	if err != nil {
		return err
	}

	request := esapi.UpdateRequest{
		Index:      p.index,
		DocumentID: product.Id,
		Body:       strings.NewReader(fmt.Sprintf(`{"doc":%s}`, bytes)),
	}

	response, err := request.Do(context.Background(), p.client)
	if err != nil {
		return err
	}

	if response.StatusCode == 404 {
		return ProductNotFoundError
	}

	defer func() {
		err = response.Body.Close()
	}()

	return err
}

func (p product) Remove(id string) error {
	request := esapi.DeleteRequest{
		Index:      p.index,
		DocumentID: id,
	}

	response, err := request.Do(context.Background(), p.client)
	if err != nil {
		return err
	}

	if response.StatusCode == 404 {
		return ProductNotFoundError
	}

	defer func() {
		err = response.Body.Close()
	}()

	return err
}

func NewProduct(client *elasticsearch.Client) Product {
	return product{
		index:  "products",
		client: client,
	}
}
