package cmd

import (
	"encoding/json"

	"github.com/7ruedzn/todos/internal/files"
	"github.com/7ruedzn/todos/internal/output"
	"github.com/7ruedzn/todos/models"
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
	b, err := files.Load()
	if err != nil {
		files.Create([]byte("[]"))
		panic(err)
	}

	todos := []models.Todo{}
	if err := json.Unmarshal(b, &todos); err != nil {
		panic(err)
	}

	output.ListTodos(todos, all)
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&all, "all", "a", false, "See all your todos, including the already completed!")
}
