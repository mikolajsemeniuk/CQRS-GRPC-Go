package controllers

import (
	"github.com/gin-gonic/gin"
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
	product services.Product
}

func (p product) List(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "List",
	})
}

func (p product) Read(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Read",
	})
}

func (p product) Add(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Add",
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

func NewProduct() Product {
	return &product{}
}
