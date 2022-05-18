package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/mikolajsemeniuk/CQRS-GRPC-Go/product-proxy-service/messages"
	"github.com/mikolajsemeniuk/CQRS-GRPC-Go/product-proxy-service/services"
)

type Product interface {
	List(c *gin.Context)
	Read(c *gin.Context)
	Add(c *gin.Context)
	Update(c *gin.Context)
	Remove(c *gin.Context)
}

type product struct {
	productService services.Product
}

func (p product) List(c *gin.Context) {
	products, err := p.productService.List()
	if err == nil {
		err = errors.New("")
	}
	c.JSON(200, gin.H{
		"data":   products,
		"errors": []string{err.Error()},
	})
}

func (p product) Read(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Read",
	})
}

func (p product) Add(c *gin.Context) {
	err := p.productService.Add(messages.Product{
		Name:       "dsadsa",
		Dollars:    14,
		Cents:      13,
		Amount:     10,
		IsImported: true,
	})
	if err == nil {
		err = errors.New("")
	}
	c.JSON(200, gin.H{
		"message": "Add",
		"errors":  []string{err.Error()},
	})
}

func (p product) Update(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Update",
	})
}

func (p product) Remove(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Remove",
	})
}

func NewProduct(productService services.Product) Product {
	return &product{
		productService: productService,
	}
}
