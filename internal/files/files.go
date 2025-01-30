package files

import (
	"fmt"
	"os"
	"path/filepath"
)

// INFO: can be changed if needed
var TODOS_FILE_PATH string = filepath.Join("storage", "todos.json")

func Write(data []byte) error {
	file, err := os.Create(TODOS_FILE_PATH)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	_, wErr := file.Write(data)
	if wErr != nil {
		panic(wErr)
	}

	if sErr := file.Sync(); sErr != nil {
		panic(sErr)
	}

	return nil
}

func Create(defaultData []byte) error {
	file, err := os.Create(TODOS_FILE_PATH)
	if err != nil {
		return err
	}

	defer file.Close()

	_, wErr := file.Write(defaultData)
	if wErr != nil {
		return wErr
	}

	if syncErr := file.Sync(); syncErr != nil {
		return syncErr
	}

	return nil
}

func Load() ([]byte, error) {
	_, statErr := os.Stat(TODOS_FILE_PATH)

	if statErr == nil {
		b, err := os.ReadFile(TODOS_FILE_PATH)

		if err != nil {
			panic(err)
		}

		return b, nil
	}

	return []byte{}, fmt.Errorf("file not found!\n")
}
