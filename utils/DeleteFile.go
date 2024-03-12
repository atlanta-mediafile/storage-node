package utils

import (
	"os"
	"path/filepath"
)

func DeleteFile(folder string, file string) error {
	return os.Remove(filepath.Join(folder, file))
}
