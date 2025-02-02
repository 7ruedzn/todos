package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/7ruedzn/todos/models"
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
		fmt.Fprintf(os.Stderr, "you provided %s instead of an integer. Please provide an integer to delete an existing todo or for help see usage with %q", args[0], "help")
		cobra.CheckErr(err)
	}

	err = models.DeleteTodo(id)

	if err != nil {
		cobra.CheckErr(err)
	}
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
