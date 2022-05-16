package handlers

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/mikolajsemeniuk/CQRS-GRPC-Go/product-write-service/messages"
	"github.com/mikolajsemeniuk/CQRS-GRPC-Go/product-write-service/proto"
	"github.com/mikolajsemeniuk/CQRS-GRPC-Go/product-write-service/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	InvalidIdError      = status.Errorf(codes.InvalidArgument, "id is not a type of uuid")
	InvalidNameError    = status.Errorf(codes.InvalidArgument, "name has to be between 3 and 255")
	InvalidAmountError  = status.Errorf(codes.InvalidArgument, "amount has to be between 3 and 255")
	InvalidDollarsError = status.Errorf(codes.InvalidArgument, "dollars has to be greater than 0")
	InvalidCentsError   = status.Errorf(codes.InvalidArgument, "cents has to be greater than 0")
)

type Product interface {
	CreateProduct(c context.Context, in *proto.CreateProductRequest) (*emptypb.Empty, error)
	UpdateProduct(c context.Context, in *proto.UpdateProductRequest) (*emptypb.Empty, error)
	RemoveProduct(c context.Context, in *proto.RemoveProductRequest) (*emptypb.Empty, error)
}

type product struct {
	productService services.Product
	sender         services.Sender
}

func (p *product) CreateProduct(c context.Context, product *proto.CreateProductRequest) (*emptypb.Empty, error) {
	name := product.GetName()
	if len(name) < 3 || len(name) > 255 {
		return nil, InvalidNameError
	}

	dollars := product.GetDollars()
	if dollars <= 0 {
		return nil, InvalidDollarsError
	}

	cents := product.GetCents()
	if cents <= 0 {
		return nil, InvalidCentsError
	}

	amount := product.GetAmount()
	if amount <= 0 {
		return nil, InvalidAmountError
	}

	imported := product.GetIsImported()
	created := time.Now()

	message := messages.Product{
		Id:         uuid.New().String(),
		Name:       &name,
		Dollars:    &dollars,
		Cents:      &cents,
		Amount:     &amount,
		IsImported: &imported,
		CreatedAt:  &created,
	}

	err := p.productService.Create(message)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, err
}

func (p *product) UpdateProduct(c context.Context, product *proto.UpdateProductRequest) (*emptypb.Empty, error) {
	id, err := uuid.Parse(product.GetId())
	if err != nil {
		return nil, InvalidIdError
	}
	message := messages.Product{
		Id: id.String(),
	}

	name := product.GetName()
	if name != nil && len(name.Value) < 3 || name != nil && len(name.Value) > 255 {
		return nil, InvalidNameError
	}
	if name != nil {
		message.Name = &name.Value
	}

	dollars := product.GetDollars()
	if dollars != nil && dollars.Value <= 0 {
		return nil, InvalidDollarsError
	}
	if dollars != nil {
		message.Dollars = &dollars.Value
	}

	cents := product.GetCents()
	if cents != nil && cents.Value <= 0 {
		return nil, InvalidCentsError
	}
	if cents != nil {
		message.Cents = &cents.Value
	}

	amount := product.GetAmount()
	if amount != nil && amount.Value <= 0 {
		return nil, InvalidAmountError
	}
	if amount != nil {
		message.Amount = &amount.Value
	}

	isImported := product.GetIsImported()
	if isImported != nil {
		message.IsImported = &isImported.Value
	}

	updated := time.Now()
	message.UpdatedAt = &updated

	err = p.productService.Update(message)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, err
}

func (p *product) RemoveProduct(c context.Context, product *proto.RemoveProductRequest) (*emptypb.Empty, error) {
	id, err := uuid.Parse(product.GetId())
	if err != nil {
		return nil, InvalidIdError
	}

	err = p.sender.Send("product-import-queue", "hello there")
	if err != nil {
		return nil, err
	}

	err = p.productService.Remove(id.String())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, err
}

func NewProduct(productService services.Product, sender services.Sender) Product {
	return &product{
		productService: productService,
		sender:         sender,
	}
}
