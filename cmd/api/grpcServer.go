package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	api "petProject/pkg/api/grpc"
	tp "petProject/transport/grpc"
)

func main() {
	// Введение в gRPC: пишем сервер на Go
	s := grpc.NewServer()
	srv := &tp.GRPCServer{}
	api.RegisterUserServiceServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
