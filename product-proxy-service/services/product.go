package services

import (
	"github.com/mikolajsemeniuk/CQRS-GRPC-Go/product-proxy-service/messages"
)

type Product interface {
	List() ([]messages.Product, error)
	Read(id string) (messages.Product, error)
	Add(messages.Product) error
	Update(input messages.Product) error
	Remove(id string) error
}

type product struct {
}

func (p *product) List() ([]messages.Product, error) {
	return []messages.Product{}, nil
}

func (p *product) Read(id string) (messages.Product, error) {
	return messages.Product{}, nil
}

func (p *product) Add(input messages.Product) error {
	return nil
}

func (p *product) Update(input messages.Product) error {
	return nil
}

func (p *product) Remove(id string) error {
	return nil
}

func NewProduct() Product {
	return &product{}
}
