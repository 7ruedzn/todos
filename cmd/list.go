package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/7ruedzn/todos/internal/files"
	"github.com/7ruedzn/todos/internal/output"
	"github.com/7ruedzn/todos/models"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "get"},
	Short:   "List your todos",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("args: ", args)
		b, err := files.Load()
		if err != nil {
			files.Create([]byte("[]"))
			panic(err)
		}

		todos := []models.Todo{}
		if err := json.Unmarshal(b, &todos); err != nil {
			panic(err)
		}

		output.ListTodos(todos, true) //TODO: create a flag to list if the todo is done or not
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP("all", "a", false, "See all your todos, including the already completed!")
}
