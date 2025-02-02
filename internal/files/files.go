package files

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// INFO: can be changed if needed
func Write(data []byte) error {
	path := viper.GetString("todos_path")
	if path == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		path = filepath.Join(home, ".config", "todos", "todos.json")
	}

	file, err := os.Create(path)
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

func Create(b []byte) error {
	home, err := os.UserHomeDir()

	if err != nil {
		return err
	}

	file, err := os.Create(filepath.Join(home, viper.GetString("todos.path")))
	if err != nil {
		fmt.Println("err creating file: ", err)
		return err
	}

	defer file.Close()

	_, wErr := file.Write(b)
	if wErr != nil {
		return wErr
	}

	if syncErr := file.Sync(); syncErr != nil {
		return syncErr
	}

	return nil
}

func Load() ([]byte, error) {
	home, err := os.UserHomeDir()

	if err != nil {
		return nil, err
	}

	path := filepath.Join(home, viper.GetString("todos.path"))
	_, statErr := os.Stat(path)

	if statErr != nil {
		err = Create([]byte("[]"))
		if err != nil {
			panic(err)
		}
	}

	b, err := os.ReadFile(path)

	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return b, nil
}
