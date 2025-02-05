package cmd

import (
	"encoding/json"

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
	Run:     runList,
}

func runList(cmd *cobra.Command, args []string) {
	todosPath := config.AppInstance.Config.TodosPath
	b, err := files.Load(todosPath)

	if err != nil {
		if err := files.Create(todosPath); err != nil {
			config.ErrorLog.Fatalln("Couldn't create the todos file: ", err)
		}
	}

	todos := []models.Todo{}
	if err := json.Unmarshal(b, &todos); err != nil && len(b) > 0 {
		config.ErrorLog.Fatalln("Couldn't unmarshal the todos file: ", err)
	}

	output.ListTodos(todos, all)
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&all, "all", "a", false, "See all your todos, including the already completed!")
}
