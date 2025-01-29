package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todos",
	Short: "Todos is a simple CLI tool to manage all your todos in the terminal",
	Long:  "A simple way to manage all your todos in the terminal, written in Go",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// INFO: A nice thing to do would be to customize the help text a lot. This can
// improve the user experience so much to the user. Be as descritive as possibile,
// like listing all resorces availables, the commands, etc.

// INFO: avoid making people to pipe the output or something to your command.
// the final user will be more happy if he can just type:
// $travis pubkey | jq -r .key > mykey.pub
// instead of:
// $travis encrypt MY_SECRET_ENV=super_secret --add env -- THIS IS BETTER
func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// this is a global flag. This can be run in any "sequence" while using the CLI
	// this is a global flag because its being defined at the root cmd.
	// and persistent flags can also be added to other command
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/todos/config.toml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile) // Use config file from the flag.
	} else {
		home, err := os.UserHomeDir() // Find home directory.
		cobra.CheckErr(err)           // check for errors. If found, prints it and exit with code 1

		// Search config in home directory with name ".todos" (without extension).
		viper.AddConfigPath(filepath.Join(home, ".config", "todos"))
		viper.SetConfigType("toml")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
