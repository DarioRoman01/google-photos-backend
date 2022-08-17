package main

import (
	"log"
	"net"

	"github.com/DarioRoman01/photos/bucket"
	"github.com/DarioRoman01/photos/uploadpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listener, err := net.Listen("tcp", ":5070")
	if err != nil {
		log.Fatalf("Error listening: %s", err.Error())
	}

	s3Bucket, err := bucket.NewS3BucketRepository()
	if err != nil {
		log.Fatalf("Error creating S3 bucket repository: %s", err.Error())
	}

	bucket.SetBucketRepository(s3Bucket)

	server := NewServer()
	grpcServer := grpc.NewServer()
	uploadpb.RegisterUploadServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Error serving: %s", err.Error())
	}
}
