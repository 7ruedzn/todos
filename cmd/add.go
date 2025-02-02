package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/7ruedzn/todos/internal/files"
	"github.com/7ruedzn/todos/internal/output"
	"github.com/7ruedzn/todos/models"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add",                                                  // this is the command name
	Aliases: []string{"new", "include"},                             // if any of these are wrote, the commmand will still run
	Args:    cobra.ExactArgs(1),                                     // you need to provide at least one new todo description
	Short:   "Add a new todo",                                       // short description of the command
	Long:    "Add a new todo you're planning or already working on", // long description of the command
	Run: func(cmd *cobra.Command, args []string) {
		todos := models.GetTodos()
		newTodos, todo := models.AddTodo(todos, args[0])
		fmt.Println("new todos", newTodos)
		b, err := json.Marshal(newTodos)
		cobra.CheckErr(err)
		files.Write(b)
		output.ListAddedTodo(todo)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	// addCmd.PersistentFlags().String("output", "", "outputs the content of the todos in a specific file format")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
