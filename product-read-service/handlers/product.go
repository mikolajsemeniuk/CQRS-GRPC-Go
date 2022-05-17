package handlers

import (
	"context"

	"github.com/google/uuid"
	"github.com/mikolajsemeniuk/CQRS-GRPC-Go/product-read-service/proto"
	"github.com/mikolajsemeniuk/CQRS-GRPC-Go/product-read-service/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	InvalidIdError = status.Errorf(codes.InvalidArgument, "id is not a type of uuid")
)

type Product interface {
	ListProducts(in *emptypb.Empty, stream proto.ProductService_ListProductsServer) error
	ReadProduct(c context.Context, input *proto.ProductId) (*proto.Product, error)
}

type product struct {
	productService services.Product
}

func (p product) ListProducts(in *emptypb.Empty, stream proto.ProductService_ListProductsServer) error {
	products, err := p.productService.List()
	if err != nil {
		return err
	}

	for _, product := range products {
		stream.Send(&proto.Product{
			Id:        product.Id,
			Name:      *product.Name,
			Dollars:   *product.Dollars,
			Cents:     *product.Cents,
			Amount:    *product.Amount,
			CreatedAt: timestamppb.New(*product.CreatedAt),
			UpdatedAt: timestamppb.New(*product.UpdatedAt),
		})
	}
	return nil
}

func (p product) ReadProduct(c context.Context, input *proto.ProductId) (*proto.Product, error) {
	_, err := uuid.Parse(input.GetId())
	if err != nil {
		return nil, InvalidIdError
	}

	result, err := p.productService.Read(input.GetId())
	if err != nil {
		return nil, err
	}

	product := &proto.Product{
		Id:         result.Id,
		Name:       *result.Name,
		Dollars:    *result.Dollars,
		Cents:      *result.Cents,
		Amount:     *result.Amount,
		IsImported: *result.IsImported,
		CreatedAt:  timestamppb.New(*result.CreatedAt),
		UpdatedAt:  timestamppb.New(*result.UpdatedAt),
	}

	return product, nil
}

func NewProduct(productService services.Product) Product {
	return product{
		productService: productService,
	}
}
