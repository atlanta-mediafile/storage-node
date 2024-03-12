package server

import (
	"context"
	"mediafile/storage_node/config"
	pb "mediafile/storage_node/grpc"
	"mediafile/storage_node/utils"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FileServiceServer struct {
	node config.Node
	pb.UnimplementedFileServiceServer
}

func (s *FileServiceServer) GetSingleFile(context context.Context, request *pb.GetSingleFileRequest) (*pb.GetSingleFileResponse, error) {

	file, err := utils.GetFile(s.node.StorageLocation, request.FileId)
	if err != nil {
		return nil, err
	}

	return &pb.GetSingleFileResponse{
		File: file,
	}, nil
}

func (s *FileServiceServer) UploadSingleFile(context context.Context, request *pb.UploadSingleFileRequest) (*pb.UploadSingleFileResponse, error) {
	id := uuid.New()

	err := utils.CreateFile(s.node.StorageLocation, id.String(), &request.File)
	if err != nil {
		return nil, err
	}

	return &pb.UploadSingleFileResponse{
		FileId: id.String(),
	}, nil
}

func (s *FileServiceServer) DeleteSingleFile(context context.Context, request *pb.DeleteSingleFileRequest) (*pb.DeleteSingleFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSingleFile not implemented")
}
