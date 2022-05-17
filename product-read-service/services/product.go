package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/mikolajsemeniuk/CQRS-GRPC-Go/product-read-service/messages"
)

var (
	ProductNotFoundError = errors.New("Product not found")
)

type Product interface {
	List() ([]messages.Product, error)
	Read(id string) (messages.Product, error)
	Create(product messages.Product) error
	Update(product messages.Product) error
	Remove(id string) error
}

type product struct {
	index  string
	client *elasticsearch.Client
}

func (product) List() ([]messages.Product, error) {
	return []messages.Product{}, nil
}

func (p product) Read(id string) (messages.Product, error) {
	request := esapi.GetRequest{
		Index:      p.index,
		DocumentID: id,
	}

	response, err := request.Do(context.Background(), p.client)
	if err != nil {
		return messages.Product{}, err
	}

	if response.StatusCode == 404 {
		return messages.Product{}, ProductNotFoundError
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return messages.Product{}, err
	}

	var result struct {
		Source messages.Product `json:"_source"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return messages.Product{}, err
	}

	defer func() {
		err = response.Body.Close()
	}()

	return result.Source, err
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
