package config

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"

	"github.com/7ruedzn/todos/internal/files"
	"github.com/7ruedzn/todos/internal/models"
	"github.com/spf13/viper"
)

var AppInstance App

type App struct {
	currentTodos []models.Todo
	config       Config
}

type Config struct {
	todosPath  string
	configPath string
	logsPath   string
}

func createFileIfNotExist(paths []string) error {
	if has := slices.Contains(paths, ""); has {
		return fmt.Errorf("path cannot be empty")
	}

	for _, v := range paths {
		if _, err := os.Stat(v); os.IsNotExist(err) {
			if err := files.Create(v); err != nil {
				return err
			}
			return err
		}
	}

	return nil
}

func LoadConfig() error {
	todosPath := viper.GetString("todos.path")
	configPath := viper.GetString("config.path")
	logsPath := viper.GetString("logs.path")

	if err := createFileIfNotExist([]string{configPath, todosPath, logsPath}); err != nil {
		panic(err)
	}

	b, err := files.Load(todosPath)
	if err != nil {
		return err
	}

	var todos []models.Todo
	if err := json.Unmarshal(b, &todos); err != nil {
		return err
	}

	AppInstance = App{
		currentTodos: todos,
		config: Config{
			todosPath:  todosPath,
			configPath: configPath,
			logsPath:   logsPath,
		},
	}

	return nil
}
