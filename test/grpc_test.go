package test

import (
	"context"
	"fmt"
	"regexp"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"mediafile/storage_node/config"
	pb "mediafile/storage_node/grpc"
	"mediafile/storage_node/server"
)

func TestDownloadFile(t *testing.T) {

	node, err := config.GetConfig()
	if err != nil {
		t.Fatalf("Cannot get config: %v", err)
	}

	// START SERVER
	srv, err := server.StartGrpcServer(&node)
	if err != nil {
		t.Fatalf("Failed to start grpc server: %v", err)
	}
	defer srv.Stop()

	// START CLIENT
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	address := fmt.Sprintf("%s:%d", node.BindHost, node.BindPort)

	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		t.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := pb.NewFileServiceClient(conn)

	file_content := "Hello, world!"
	file := []byte(file_content)

	// UPLOAD FILE
	response, err := client.UploadSingleFile(context.Background(), &pb.UploadSingleFileRequest{
		File: file,
	})

	if err != nil {
		t.Fatalf("Failed to call UploadSingleFile: %v", err)
	}

	uuidRegex := regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)

	validUUID := uuidRegex.MatchString(response.FileId)

	if !validUUID {
		t.Fatal("Invalid UUID format.")
	}

	// DOWNLOAD FILE
	getResponse, err := client.GetSingleFile(context.Background(), &pb.GetSingleFileRequest{
		FileId: response.FileId,
	})

	if err != nil {
		t.Fatalf("Failed to call GetSingleFile: %v", err)
	}

	content := string(getResponse.File)

	if content != file_content {
		t.Fatal("File content are not equal")
	}

	// DELETE FILE
	deleteResponse, err := client.DeleteSingleFile(context.Background(), &pb.DeleteSingleFileRequest{
		FileId: response.FileId,
	})

	if err != nil {
		t.Fatalf("Failed to call DeleteSingleFile: %v", err)
	}

	if !deleteResponse.Done {
		t.Fatalf("The file was not deleted")
	}

	// TRY GET THE FILE
	file2, _ := client.GetSingleFile(context.Background(), &pb.GetSingleFileRequest{
		FileId: response.FileId,
	})

	if file2 != nil {
		t.Fatalf("File found, not deleted")
	}

}
