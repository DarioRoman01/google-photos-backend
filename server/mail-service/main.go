package main

import (
	"log"
	"net"

	"github.com/DarioRoman01/photos/mailpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listener, err := net.Listen("tcp", ":5060")
	if err != nil {
		log.Fatalf("Error listening: %s", err.Error())
	}

	server := NewServer()
	grpcServer := grpc.NewServer()
	mailpb.RegisterMailServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Error serving: %s", err.Error())
	}
}
