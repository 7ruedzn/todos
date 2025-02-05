package cmd

import (
	"encoding/json"

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
	Run:     runAdd,
}

func runAdd(cmd *cobra.Command, args []string) {
	todos, err := models.GetTodos()
	if err != nil {
		config.ErrorLog.Fatalln("Couldn't get todos: ", err)
	}

	todosPath := config.AppInstance.Config.TodosPath
	newTodos, todo := models.AddTodo(todos, args[0])
	b, err := json.Marshal(newTodos)

	if err != nil {
		config.ErrorLog.Fatalf("Couldn't marshal new todos %+v: %v\n", newTodos, err)
	}

	if err := files.Write(b, todosPath); err != nil {
		config.ErrorLog.Fatalf("Couldn't write a new todo with %s: %v: ", string(b), err)
	}

	config.InfoLog.Printf("Todo with id %d added successfully\n", todo.Id)
	output.ListAddedTodo(todo)
}

func init() {
	rootCmd.AddCommand(addCmd)
}
