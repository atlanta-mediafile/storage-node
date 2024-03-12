package server

import (
	"context"
	pb "mediafile/storage_node/grpc"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FileServiceServer struct {
	pb.UnimplementedFileServiceServer
}

func (s *FileServiceServer) GetSingleFile(context context.Context, request *pb.GetSingleFileRequest) (*pb.GetSingleFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSingleFile not implemented")
}
func (s *FileServiceServer) UploadSingleFile(context context.Context, request *pb.UploadSingleFileRequest) (*pb.UploadSingleFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadSingleFile not implemented")
}
func (s *FileServiceServer) DeleteSingleFile(context context.Context, request *pb.DeleteSingleFileRequest) (*pb.DeleteSingleFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSingleFile not implemented")
}
