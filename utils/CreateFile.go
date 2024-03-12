package utils

import (
	"os"
	"path/filepath"
)

func CreateFile(folder string, fileName string, file *[]byte) error {

	buf, err := os.Create(filepath.Join(folder, fileName))

	if err != nil {
		buf.Close()
		return err
	}

	_, err = buf.Write(*file)

	if err != nil {
		buf.Close()
		return err
	}

	err = buf.Sync()

	if err != nil {
		buf.Close()
		return err
	}

	buf.Close()

	return nil
}
