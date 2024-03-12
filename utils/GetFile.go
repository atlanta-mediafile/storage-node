package utils

import (
	"os"
	"path/filepath"
)

func GetFile(folder string, file string) ([]byte, error) {
	return os.ReadFile(filepath.Join(folder, file))
}
