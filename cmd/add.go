/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
// WARN: The problem with this approache is that every single new command i need
// to add, i have to create another cobra command and so on.
var addCmd = &cobra.Command{
	Use:        "add", // this is the command name
	SuggestFor: []string{"include", "new"},
	Args:       cobra.MinimumNArgs(1),                                             // you need to provide at least one new todo description
	Short:      "Add one or more new todos",                                       // short description of the command
	Long:       "Add one or more new todos you're planning or already working on", // long description of the command
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("called add cmd")
		fmt.Println("args on add: ", args)
		fmt.Println("args in zero", args[0])
		fmt.Println("args in 1", args[1])

		outputFlag, _ := cmd.Flags().GetString("output")
		if outputFlag != "" {
			fmt.Println("output flag is :", outputFlag)
		} else {
			fmt.Println("no output flag provided")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.PersistentFlags().String("output", "", "outputs the content of the todos in a specific file format")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
