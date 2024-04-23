package server

import (
	"context"
	"fmt"
	"mediafile/storage_node/config"
	pb "mediafile/storage_node/grpc"
	"mediafile/storage_node/utils"

	"github.com/google/uuid"
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

	nodeIp := fmt.Sprintf("%s:%d", s.node.BindHost, s.node.BindPort)

	return &pb.UploadSingleFileResponse{
		FileId: id.String(),
		NodeIp: nodeIp,
	}, nil
}

func (s *FileServiceServer) DeleteSingleFile(context context.Context, request *pb.DeleteSingleFileRequest) (*pb.DeleteSingleFileResponse, error) {

	err := utils.DeleteFile(s.node.StorageLocation, request.FileId)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteSingleFileResponse{
		Done: true,
	}, nil
}
