package cmd

import (
	"encoding/json"

	"github.com/7ruedzn/todos/internal/files"
	"github.com/7ruedzn/todos/internal/models"
	"github.com/7ruedzn/todos/internal/output"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	todosPath := viper.GetString("todos.path")
	b, err := files.Load(todosPath)

	if err != nil {
		if err := files.Create(todosPath); err != nil {
			panic(err)
		}
	}

	todos := []models.Todo{}
	if err := json.Unmarshal(b, &todos); err != nil && len(b) > 0 {
		panic(err)
	}

	output.ListTodos(todos, all)
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&all, "all", "a", false, "See all your todos, including the already completed!")
}
