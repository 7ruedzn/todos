package cmd

import (
	"encoding/json"
	"strconv"

	"github.com/7ruedzn/todos/internal/files"
	"github.com/7ruedzn/todos/internal/models"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var completeCmd = &cobra.Command{
	Use:     "complete",
	Aliases: []string{"finish", "done"},
	Short:   "Set a todo as finished!",
	Long:    "Set a todo you've created as done. This way you can keep track off the complete todos and the still on progress",
	Run:     runComplete,
}

func runComplete(cmd *cobra.Command, args []string) {
	todosPath := viper.GetString("todos.path")
	todos := models.GetTodos()
	id, err := strconv.Atoi(args[0])
	if err != nil {
		panic(err)
	}

	todo, err := models.GetTodo(id, todos)
	if err != nil {
		panic(err)
	}

	updatedTodos, err := todo.UpdateTodos()
	if err != nil {
		panic(err)
	}

	b, err := json.Marshal(&updatedTodos)
	if err != nil {
		panic(err)
	}

	if err := files.Write(b, todosPath); err != nil {
		panic(err)
	}
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
