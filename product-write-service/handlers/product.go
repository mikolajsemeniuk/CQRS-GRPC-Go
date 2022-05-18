package handlers

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/mikolajsemeniuk/CQRS-GRPC-Go/product-write-service/messages"
	proto "github.com/mikolajsemeniuk/CQRS-GRPC-Go/product-write-service/proto-write"
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

func (p *product) CreateProduct(c context.Context, input *proto.CreateProductRequest) (*emptypb.Empty, error) {
	name := input.GetName()
	if len(name) < 3 || len(name) > 255 {
		return nil, InvalidNameError
	}

	dollars := input.GetDollars()
	if dollars <= 0 {
		return nil, InvalidDollarsError
	}

	cents := input.GetCents()
	if cents <= 0 {
		return nil, InvalidCentsError
	}

	amount := input.GetAmount()
	if amount <= 0 {
		return nil, InvalidAmountError
	}

	imported := input.GetIsImported()
	created := time.Now()

	product := messages.Product{
		Id:         uuid.New().String(),
		Name:       &name,
		Dollars:    &dollars,
		Cents:      &cents,
		Amount:     &amount,
		IsImported: &imported,
		CreatedAt:  &created,
		UpdatedAt:  &time.Time{},
	}

	err := p.productService.Create(product)
	if err != nil {
		return nil, err
	}

	event := messages.Event{
		Method:    "CREATE",
		Data:      product,
		Timestamp: time.Now(),
	}

	bytes, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}

	// TODO: move to configuration
	err = p.sender.Send("product-import-queue", string(bytes))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, err
}

func (p *product) UpdateProduct(c context.Context, input *proto.UpdateProductRequest) (*emptypb.Empty, error) {
	id, err := uuid.Parse(input.GetId())
	if err != nil {
		return nil, InvalidIdError
	}
	product := messages.Product{
		Id: id.String(),
	}

	name := input.GetName()
	if name != nil && len(name.Value) < 3 || name != nil && len(name.Value) > 255 {
		return nil, InvalidNameError
	}
	if name != nil {
		product.Name = &name.Value
	}

	dollars := input.GetDollars()
	if dollars != nil && dollars.Value <= 0 {
		return nil, InvalidDollarsError
	}
	if dollars != nil {
		product.Dollars = &dollars.Value
	}

	cents := input.GetCents()
	if cents != nil && cents.Value <= 0 {
		return nil, InvalidCentsError
	}
	if cents != nil {
		product.Cents = &cents.Value
	}

	amount := input.GetAmount()
	if amount != nil && amount.Value <= 0 {
		return nil, InvalidAmountError
	}
	if amount != nil {
		product.Amount = &amount.Value
	}

	isImported := input.GetIsImported()
	if isImported != nil {
		product.IsImported = &isImported.Value
	}

	updated := time.Now()
	product.UpdatedAt = &updated

	err = p.productService.Update(product)
	if err != nil {
		return nil, err
	}

	event := messages.Event{
		Method:    "UPDATE",
		Data:      product,
		Timestamp: time.Now(),
	}

	bytes, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}

	// TODO: move to configuration
	err = p.sender.Send("product-import-queue", string(bytes))
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

	err = p.productService.Remove(id.String())
	if err != nil {
		return nil, err
	}

	event := messages.Event{
		Method:    "REMOVE",
		Data:      messages.Product{Id: id.String()},
		Timestamp: time.Now(),
	}

	bytes, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}

	// TODO: move to configuration
	err = p.sender.Send("product-import-queue", string(bytes))
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
