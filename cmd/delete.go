package cmd

import (
	"log/slog"
	"strconv"

	"github.com/7ruedzn/todos/internal/models"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"rm", "remove"},
	Args:    cobra.ExactArgs(1), //TODO: accept more than one id to delete todos
	Short:   "Delete a todo",
	Long:    "Delete a todo by it's id.",
	RunE:    runDelete,
	PostRun: logDelete,
}

func logDelete(cmd *cobra.Command, args []string) {
	slog.Info("Delete completed successfully", "cmd", cmd.Name(), "args", args)
}

func runDelete(cmd *cobra.Command, args []string) error {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		slog.Error("Couldn't parse todo id to int", "cmd", cmd.Name(), "error", err, "args", args[0])
		return err
	}

	if err := models.DeleteTodo(id); err != nil {
		slog.Error("Couldn't delete the todo", "cmd", cmd.Name(), "error", err, "id", id)
		return err
	}

	return nil
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
