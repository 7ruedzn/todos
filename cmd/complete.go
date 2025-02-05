package cmd

import (
	"encoding/json"
	"os"
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
	Run:     runComplete,
}

func runComplete(cmd *cobra.Command, args []string) {
	todosPath := config.AppInstance.Config.TodosPath
	todos, err := models.GetTodos()
	if err != nil {
		config.ErrorLog.Fatalln("Couldn't get todos: ", err)
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		config.ErrorLog.Fatalf("Couldn't parse %s into type of int: %v\n", os.Args[0], err)
	}

	todo, err := models.GetTodo(id, todos)
	if err != nil {
		config.ErrorLog.Fatalf("Couldn't get todo by the id %d: %v\n", id, err)
	}

	updatedTodos, err := todo.UpdateTodos()
	if err != nil {
		config.ErrorLog.Fatalln("Couldn't update the todos: ", err)
	}

	b, err := json.Marshal(&updatedTodos)
	if err != nil {
		config.ErrorLog.Fatalf("Couldn't marshal %+v: %v\n", *updatedTodos, err)
	}

	if err := files.Write(b, todosPath); err != nil {
		config.ErrorLog.Fatalf("Couldn't write %s into the todos file: %v\n", string(b), err)
	}

	config.InfoLog.Printf("Todo with id %d completed successfully\n", todo.Id)
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
