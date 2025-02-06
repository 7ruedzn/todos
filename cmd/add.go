package cmd

import (
	"encoding/json"
	"log/slog"

	"github.com/7ruedzn/todos/internal/config"
	"github.com/7ruedzn/todos/internal/files"
	"github.com/7ruedzn/todos/internal/models"
	"github.com/7ruedzn/todos/internal/output"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add",                                                  // this is the command name
	Aliases: []string{"new", "include"},                             // if any of these are wrote, the commmand will still run
	Args:    cobra.ExactArgs(1),                                     //TODO: accept more than one descrition to add multiple todos at once
	Short:   "Add a new todo",                                       // short description of the command
	Long:    "Add a new todo you're planning or already working on", // long description of the command
	RunE:    runAdd,
	PostRun: logAdd,
}

func logAdd(cmd *cobra.Command, args []string) {
	slog.Info("Todo with  added successfully", "cmd", cmd.Name(), "args", args)
}

func runAdd(cmd *cobra.Command, args []string) error {
	todos, err := models.GetTodos()
	if err != nil {
		slog.Error("Couldn't get todos", "cmd", cmd.Name(), "error", err)
		return err
	}

	todosPath := config.AppInstance.Config.TodosPath
	newTodos, todo := models.AddTodo(todos, args[0])
	b, err := json.Marshal(newTodos)

	if err != nil {
		slog.Error("Couldn't marshal new todos", "cmd", cmd.Name(), "newTodos", newTodos, "error", err)
		return err
	}

	if err := files.Write(b, todosPath); err != nil {
		slog.Error("Couldn't write the new todo into file", "cmd", cmd.Name(), "bytes", string(b), "path", todosPath, "error", err)
		return err
	}

	output.ListAddedTodo(todo)
	return nil
}

func init() {
	rootCmd.AddCommand(addCmd)
}
