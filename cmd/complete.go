package cmd

import (
	"encoding/json"
	"log/slog"
	"strconv"

	"github.com/7ruedzn/todos/internal/config"
	"github.com/7ruedzn/todos/internal/files"
	"github.com/7ruedzn/todos/internal/models"
	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:     "complete",
	Aliases: []string{"finish", "done"},
	Short:   "Set a todo as finished!",
	Long:    "Set a todo you've created as done. This way you can keep track off the complete todos and the still on progress",
	Args:    cobra.ExactArgs(1),
	RunE:    runComplete,
	PostRun: completeLog,
}

func runComplete(cmd *cobra.Command, args []string) error {
	todosPath := config.AppInstance.Config.TodosPath
	todos, err := models.GetTodos()
	if err != nil {
		slog.Error("Couldn't get todos", "cmd", cmd.Name(), "error", err, "path", todosPath)
		return err
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		slog.Error("Couldn't parse todo id to int", "cmd", cmd.Name(), "error", err, "args", args[0])
		return err
	}

	todo, err := models.GetTodo(id, todos)
	if err != nil {
		slog.Error("Couldn't get todo with id", "cmd", cmd.Name(), "error", err, "args", args[0], "id", id)
		return err
	}

	updatedTodos, err := todo.UpdateTodos()
	if err != nil {
		slog.Error("Couldn't update todos", "cmd", cmd.Name(), "error", err, "args", args[0])
		return err
	}

	b, err := json.Marshal(&updatedTodos)
	if err != nil {
		slog.Error("Couldn't marshal updated todos", "cmd", cmd.Name(), "error", err, "args", args[0])
		return err
	}

	if err := files.Write(b, todosPath); err != nil {
		slog.Error("Couldn't write the updated todo into file", "cmd", cmd.Name(), "bytes", string(b), "path", todosPath, "error", err)
		return err
	}

	return nil
}

func completeLog(cmd *cobra.Command, args []string) {
	slog.Info("Todo was completed successfully", "cmd", cmd.Name(), "args", args)
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
