package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	"grpctesting/internal/handlers"
	"grpctesting/pkg/orderservice"
)

func main() {
	fmt.Println("Starting gRPC server...")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	server := &handlers.Server{}

	orderservice.RegisterOrdersAPIServer(s, server)

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
