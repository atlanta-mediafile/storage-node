package utils

import "os"

func CreateFolder(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0777)
		if err != nil {
			return err
		}
	}
	return nil
}
