package cmd

import (
	"encoding/json"

	"github.com/7ruedzn/todos/internal/files"
	"github.com/7ruedzn/todos/internal/models"
	"github.com/7ruedzn/todos/internal/output"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addCmd = &cobra.Command{
	Use:     "add",                                                  // this is the command name
	Aliases: []string{"new", "include"},                             // if any of these are wrote, the commmand will still run
	Args:    cobra.ExactArgs(1),                                     //TODO: accept more than one descrition to add multiple todos at once
	Short:   "Add a new todo",                                       // short description of the command
	Long:    "Add a new todo you're planning or already working on", // long description of the command
	Run:     runAdd,
}

func runAdd(cmd *cobra.Command, args []string) {
	todos := models.GetTodos()
	todosPath := viper.GetString("todos.path")
	newTodos, todo := models.AddTodo(todos, args[0])
	b, err := json.Marshal(newTodos)
	cobra.CheckErr(err)

	if err := files.Write(b, todosPath); err != nil {
		cobra.CheckErr(err)
	}

	output.ListAddedTodo(todo)
}

func init() {
	rootCmd.AddCommand(addCmd)
}
