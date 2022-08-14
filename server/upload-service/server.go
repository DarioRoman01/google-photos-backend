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

	filename := ""
	username := ""
	folder := ""
	readed := false

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			location, err := bucket.Upload(buff, filename, username, folder)
			if err != nil {
				return status.Error(codes.Internal, "failed to upload image")
			}

			return stream.SendAndClose(&uploadpb.UploadResponse{
				Location: location,
			})
		}

		if !readed {
			filename = req.Filename
			username = req.Username
			folder = req.Folder
			readed = true
		}

		if err != nil {
			return status.Error(codes.Internal, "failed to process image")
		}

		if _, err := buff.Write(req.Chunk); err != nil {
			return status.Error(codes.Internal, "failed to process image")
		}
	}
}
