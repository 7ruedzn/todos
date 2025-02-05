package cmd

import (
	"strconv"

	"github.com/7ruedzn/todos/internal/config"
	"github.com/7ruedzn/todos/internal/models"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"rm", "remove"},
	Args:    cobra.ExactArgs(1), //TODO: accept more than one id to delete todos
	Short:   "Delete a todo",
	Long:    "Delete a todo by it's id.",
	Run:     runDelete,
}

func runDelete(cmd *cobra.Command, args []string) {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		config.WarningLog.Printf("Couldn't delete todo. Provided %s instead of an todo id of type int\n", args[0])
	}

	if err := models.DeleteTodo(id); err != nil {
		config.ErrorLog.Fatalln("Couldn't delete todo: ", err)
	}

	config.InfoLog.Printf("Todo with id %d deleted successfully\n", id)
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
