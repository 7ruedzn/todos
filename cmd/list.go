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

var all bool

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "get"},
	Short:   "List your todos",
	Long:    "List all your todos. You can specify with the -a or --all flag to also list already completed todos",
	RunE:    runList,
	PostRun: logList,
}

func logList(cmd *cobra.Command, args []string) {
	slog.Info("List completed successfully", "cmd", cmd.Name(), "args", args)
}

func runList(cmd *cobra.Command, args []string) error {
	todosPath := config.AppInstance.Config.TodosPath
	b, err := files.Load(todosPath)

	if err != nil {
		if err := files.Create(todosPath); err != nil {
			slog.Error("Couldn't create the todos file", "cmd", cmd.Name(), "err", err, "path", todosPath)
			return err
		}
		return err
	}

	todos := []models.Todo{}
	if err := json.Unmarshal(b, &todos); err != nil && len(b) > 0 {
		slog.Error("Couldn't unmarshal the todos file", "cmd", cmd.Name(), "bytes", string(b), "err", err, "path", todosPath)
		return err
	}

	output.ListTodos(todos, all)
	return nil
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&all, "all", "a", false, "See all your todos, including the already completed!")
}
