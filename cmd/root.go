package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/7ruedzn/todos/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	app     config.App
	cfgFile string // provided from the --config flag
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
	home, err := os.UserHomeDir()
	if err != nil {
		config.ErrorLog.Fatalln("Couldn't find home dir: ", err)
	}

	defaultPath := filepath.Join(home, ".config", "todos")
	configPath := filepath.Join(defaultPath, "config.toml")

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile) // Use config file provided by the --config flag
	} else {
		viper.SetDefault("config.path", filepath.Join(defaultPath, "config.toml"))
		viper.SetDefault("todos.path", filepath.Join(defaultPath, "todos.json"))
		viper.SetDefault("logs.path", filepath.Join(defaultPath, "logs.txt"))
		viper.AddConfigPath(defaultPath) // look into $HOME/.config/todos/
		viper.AddConfigPath(".")         // look into current dir
		viper.SetConfigType("toml")      // set the config type to look into
		viper.SetConfigName("config")    // set the config name to look into
	}

	viper.AutomaticEnv() // read in environment variables that match

	// if the config file is found, reads it
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			config.WarningLog.Println("Config file was not found. Using the default params.")
			if err := viper.SafeWriteConfigAs(configPath); err != nil {
				config.ErrorLog.Fatalln("Couldn't create the config file: ", err)
			}
		} else {
			config.ErrorLog.Fatalln("Couldn't read the config file: ", err)
		}
	}

	if err := config.LoadConfig(); err != nil {
		config.ErrorLog.Fatalln("Couldn't load config: ", err)
	}

	if err := config.SetupLogs(); err != nil {
		config.ErrorLog.Fatalln("Couldn't setup logs: ", err)
	}

	defer config.LogFile.Close()
	config.InfoLog.Println("Setup completed successfully")
}
