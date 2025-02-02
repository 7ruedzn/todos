package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string // provided from the --output flag
)

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

func init() {
	cobra.OnInitialize(initConfig)
	// this is a global flag. This can be run in any "sequence" while using the CLI
	// this is a global flag because its being defined at the root cmd.
	// and persistent flags can also be added to other command
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "provide a custom config file path to load your configurations. The default is $HOME/.config/todos/config.toml")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	fmt.Println("cfg file: ", cfgFile)
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile) // Use config file from the flag --config flag
	} else {
		home, err := os.UserHomeDir() // Find home directory.
		cobra.CheckErr(err)           // check for errors. If found, prints it and exit with code 1

		// Search config in $HOME/.config/todos/
		defaultPath := filepath.Join(home, ".config", "todos")
		viper.SetDefault("todos_path", filepath.Join(defaultPath, "todos.json"))
		viper.SetDefault("config_path", filepath.Join(defaultPath, "config.toml"))
		viper.AddConfigPath(defaultPath)
		viper.AddConfigPath(".") // look into current dir
		viper.SetConfigType("toml")
		viper.SetConfigName("config") // find for config.toml file
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
