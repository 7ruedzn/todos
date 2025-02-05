package files

import (
	"fmt"
	"os"
)

func Write(data []byte, path string) error {
	//if the file doesnt exists, create it in the given path.
	//O_TRUNC truncates de files to zero len before writing to it.
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.Write(data); err != nil {
		return err
	}

	return file.Sync()
}

func Create(path string) error {
	if path == "" {
		return fmt.Errorf("path cannot be empty")
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}

func Load(path string) ([]byte, error) {
	if path == "" {
		return nil, fmt.Errorf("Path cannot be empty")
	}

	b, err := os.ReadFile(path)

	if err != nil {
		if os.IsNotExist(err) {
			if err := Write([]byte(""), path); err != nil {
				return nil, err
			}
			return []byte(""), nil
		}
		return nil, err
	}

	return b, nil
}
