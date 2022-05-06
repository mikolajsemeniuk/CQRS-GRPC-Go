package controllers

import (
	"net/http"

	"github.com/mikolajsemeniuk/CQRS-GRPC-Go/product-proxy-service/services"
	"github.com/mikolajsemeniuk/CQRS-GRPC-Go/product-proxy-service/settings"
)

type Product interface {
	Index(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	Read(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Remove(w http.ResponseWriter, r *http.Request)
}

type product struct {
	configuration settings.Configuration
	product       services.Product
}

func (p *product) Index(w http.ResponseWriter, r *http.Request) {

}

func (p *product) List(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (p *product) Read(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (p *product) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (p *product) Remove(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func NewProduct(configuration settings.Configuration, service services.Product) Product {
	return &product{
		configuration: configuration,
		product:       service,
	}
}
