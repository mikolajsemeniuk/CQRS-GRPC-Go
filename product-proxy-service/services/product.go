package services

import (
	"context"
	"io"

	"github.com/mikolajsemeniuk/CQRS-GRPC-Go/product-proxy-service/messages"
	read "github.com/mikolajsemeniuk/CQRS-GRPC-Go/product-read-service/proto-read"
	write "github.com/mikolajsemeniuk/CQRS-GRPC-Go/product-write-service/proto-write"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Product interface {
	List() ([]messages.Product, error)
	Read(id string) (messages.Product, error)
	Add(messages.Product) error
	Update(input messages.Update) error
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
			Id:         product.Id,
			Name:       product.Name,
			Dollars:    product.Dollars,
			Cents:      product.Cents,
			Amount:     product.Amount,
			IsImported: product.IsImported,
			CreatedAt:  product.CreatedAt.AsTime(),
			UpdatedAt:  product.UpdatedAt.AsTime(),
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
		Id:         product.Id,
		Name:       product.Name,
		Dollars:    product.Dollars,
		Cents:      product.Cents,
		Amount:     product.Amount,
		IsImported: product.IsImported,
		CreatedAt:  product.CreatedAt.AsTime(),
		UpdatedAt:  product.UpdatedAt.AsTime(),
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

func (p *product) Update(input messages.Update) error {
	payload := &write.UpdateProductRequest{Id: input.Id}

	if input.Name != nil {
		payload.Name = &wrapperspb.StringValue{Value: *input.Name}
	}

	if input.Dollars != nil {
		payload.Dollars = &wrapperspb.UInt64Value{Value: *input.Dollars}
	}

	if input.Cents != nil {
		payload.Cents = &wrapperspb.UInt32Value{Value: *input.Cents}
	}

	if input.Amount != nil {
		payload.Amount = &wrapperspb.UInt32Value{Value: *input.Amount}
	}

	if input.IsImported != nil {
		payload.IsImported = &wrapperspb.BoolValue{Value: *input.IsImported}
	}

	_, err := p.writeProductServiceClient.UpdateProduct(context.Background(), payload)
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
