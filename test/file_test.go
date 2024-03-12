package test

import (
	"mediafile/storage_node/utils"
	"os"
	"path"
	"path/filepath"
	"testing"
)

func TestFile(t *testing.T) {

	// Get current dir
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		t.Fatalf("Cannot get current folder: %v", err)
	}

	// Create folder
	dir = path.Join(dir, "hello")
	err = utils.CreateFolder(dir)
	if err != nil {
		t.Fatalf("Cannot create folder: %v", err)
	}

	// Set file
	file_content := "Hello, world!"
	file := []byte(file_content)

	// Save file
	err = utils.CreateFile(dir, "helloworld.txt", &file)

	if err != nil {
		t.Fatalf("Cannot save the file: %v", err)
	}

	// Get file
	file2, err := utils.GetFile(dir, "helloworld.txt")

	if err != nil {
		t.Fatalf("Cannot get the file: %v", err)
	}

	file2_content := string(file2)

	if file_content != file2_content {
		t.Fatal("File content are not equal")
	}

}
