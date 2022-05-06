package services

import (
	"github.com/mikolajsemeniuk/CQRS-GRPC-Go/product-proxy-service/settings"
	"github.com/mikolajsemeniuk/CQRS-GRPC-Go/product-validation-service/inputs"
	"github.com/mikolajsemeniuk/CQRS-GRPC-Go/product-validation-service/payloads"
)

type Product interface {
	List() ([]payloads.Product, error)
	Read(id string) (*payloads.Product, error)
	Update(input inputs.Product) error
	Remove(id string) error
}

type product struct {
	configuration settings.Configuration
}

func (p *product) List() ([]payloads.Product, error) {
	return []payloads.Product{}, nil
}

func (p *product) Read(id string) (*payloads.Product, error) {
	return nil, nil
}

func (p *product) Update(input inputs.Product) error {
	return nil
}

func (p *product) Remove(id string) error {
	return nil
}

func NewProduct(configuration settings.Configuration) Product {
	return &product{
		configuration: configuration,
	}
}
