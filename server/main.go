package server

import (
	"fmt"
	"log"
	"mediafile/storage_node/config"
	"net"

	pb "mediafile/storage_node/grpc"

	"google.golang.org/grpc"
)

func StartGrpcServer(_node *config.Node) (*grpc.Server, error) {

	address := fmt.Sprintf("%s:%d", _node.BindHost, _node.BindPort)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}
	srv := grpc.NewServer()

	pb.RegisterFileServiceServer(srv, &FileServiceServer{
		node: *_node,
	})

	go func() {
		if err := srv.Serve(lis); err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
	}()

	return srv, nil
}
