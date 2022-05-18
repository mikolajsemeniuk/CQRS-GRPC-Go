package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

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
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": products})
}

func (p product) Read(c *gin.Context) {
	id := c.Param("id")

	product, err := p.productService.Read(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	c.JSON(200, gin.H{"message": product})
}

func (p product) Add(c *gin.Context) {
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	product := messages.Product{}
	if err := json.Unmarshal(bytes, &product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = p.productService.Add(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Success"})
}

func (p product) Update(c *gin.Context) {
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id := c.Param("id")
	product := messages.Update{
		Id: id,
	}
	if err := json.Unmarshal(bytes, &product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = p.productService.Update(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Success"})
}

func (p product) Remove(c *gin.Context) {
	id := c.Param("id")

	err := p.productService.Remove(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Success"})
}

func NewProduct(productService services.Product) Product {
	return &product{
		productService: productService,
	}
}
