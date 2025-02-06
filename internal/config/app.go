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
	CurrentTodos []models.Todo
	Config       Config
}

type Config struct {
	TodosPath  string
	ConfigPath string
	LogsPath   string
}

func createFileIfNotExist(paths []string) error {
	if has := slices.Contains(paths, ""); has {
		return fmt.Errorf("path cannot be empty")
	}

	for _, v := range paths {
		if _, err := os.Stat(v); os.IsNotExist(err) {
			fmt.Println("filee doesnt exist", v)
			if err := files.Create(v); err != nil {
				fmt.Println("error creating file path", v)
				fmt.Println("error", err)
				return err
			}
			return nil
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
		CurrentTodos: todos,
		Config: Config{
			TodosPath:  todosPath,
			ConfigPath: configPath,
			LogsPath:   logsPath,
		},
	}

	return nil
}
