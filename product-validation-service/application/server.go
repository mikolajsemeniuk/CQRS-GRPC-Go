package application

import (
	"net"
)

type Server interface {
	Run() error
}

type server struct{}

func (s server) Run() error {
	_, err := net.Listen("tcp", "0.0.0.0:50051")
	return err
}

func NewServer() Server {
	return &server{}
}
