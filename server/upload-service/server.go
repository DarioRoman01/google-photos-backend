package main

import (
	"bytes"
	"io"

	"github.com/DarioRoman01/photos/bucket"
	"github.com/DarioRoman01/photos/uploadpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	uploadpb.UnimplementedUploadServiceServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Upload(stream uploadpb.UploadService_UploadServer) error {
	buff := bytes.NewBuffer(nil)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			location, err := bucket.Upload(buff, req.Filename, req.Username, req.Folder)
			if err != nil {
				return status.Error(codes.Internal, "failed to upload image")
			}

			return stream.SendAndClose(&uploadpb.UploadResponse{
				Location: location,
			})
		}

		if err != nil {
			return status.Error(codes.Internal, "failed to process image")
		}

		if _, err := buff.Write(req.Chunk); err != nil {
			return status.Error(codes.Internal, "failed to process image")
		}
	}
}
