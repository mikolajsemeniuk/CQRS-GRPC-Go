package services

import (
	"context"
	"io"

	"github.com/mikolajsemeniuk/CQRS-GRPC-Go/product-proxy-service/messages"
	read "github.com/mikolajsemeniuk/CQRS-GRPC-Go/product-read-service/proto-read"
	write "github.com/mikolajsemeniuk/CQRS-GRPC-Go/product-write-service/proto-write"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Product interface {
	List() ([]messages.Product, error)
	Read(id string) (messages.Product, error)
	Add(messages.Product) error
	Update(input messages.Product) error
	Remove(id string) error
}

type product struct {
	readProductServiceClient  read.ProductServiceClient
	writeProductServiceClient write.ProductServiceClient
}

func (p *product) List() ([]messages.Product, error) {
	products := []messages.Product{}
	stream, err := p.readProductServiceClient.ListProducts(context.Background(), &emptypb.Empty{})
	if err != nil {
		return products, err
	}

	for {
		product, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return products, err
		}
		products = append(products, messages.Product{
			Id:        product.Id,
			Name:      product.Name,
			Dollars:   product.Dollars,
			Cents:     product.Cents,
			Amount:    product.Amount,
			CreatedAt: product.CreatedAt.AsTime(),
			UpdatedAt: product.UpdatedAt.AsTime(),
		})
	}
	return products, nil
}

func (p product) Read(id string) (messages.Product, error) {
	product, err := p.readProductServiceClient.ReadProduct(context.Background(), &read.ProductId{
		Id: id,
	})
	if err != nil {
		return messages.Product{}, err
	}

	return messages.Product{
		Id:        product.Id,
		Name:      product.Name,
		Dollars:   product.Dollars,
		Cents:     product.Cents,
		Amount:    product.Amount,
		CreatedAt: product.CreatedAt.AsTime(),
		UpdatedAt: product.UpdatedAt.AsTime(),
	}, nil
}

func (p *product) Add(input messages.Product) error {
	_, err := p.writeProductServiceClient.CreateProduct(context.Background(), &write.CreateProductRequest{
		Name:       input.Name,
		Dollars:    input.Dollars,
		Cents:      input.Cents,
		Amount:     input.Amount,
		IsImported: input.IsImported,
	})
	return err
}

func (p *product) Update(input messages.Product) error {
	_, err := p.writeProductServiceClient.UpdateProduct(context.Background(), &write.UpdateProductRequest{
		// Id:         input.Id,
		// Name:       input.Name,
		// Dollars:    input.Dollars,
		// Cents:      input.Cents,
		// Amount:     input.Amount,
		// IsImported: input.IsImported,
	})
	return err
}

func (p *product) Remove(id string) error {
	_, err := p.writeProductServiceClient.RemoveProduct(context.Background(), &write.RemoveProductRequest{
		Id: id,
	})
	return err
}

func NewProduct(writeProductServiceClient write.ProductServiceClient, readProductServiceClient read.ProductServiceClient) Product {
	return &product{
		writeProductServiceClient: writeProductServiceClient,
		readProductServiceClient:  readProductServiceClient,
	}
}
