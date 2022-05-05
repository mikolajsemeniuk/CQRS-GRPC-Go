package application

import (
	"net/http"
)

type Server interface {
	Run() error
}

type server struct{}

func (s server) Run() error {
	return http.ListenAndServe(":8080", nil)
}

func NewServer() Server {
	return &server{}
}
