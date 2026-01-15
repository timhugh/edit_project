package main

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/spf13/cobra"
	"github.com/timhugh/edit_project"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Interact with tool configuration",
}

var configPathCmd = &cobra.Command{
	Use:   "path",
	Short: "Show the path to the configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println(configPath)
	},
}

var configShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the current configuration",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := edit_project.LoadConfig(configPath)
		if err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				panic(err)
			}
		}
		jsonOutput, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			panic(err)
		}
		cmd.Println(string(jsonOutput))
	},
}

var configCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a configuration file (will prompt for values)",
	Run: func(cmd *cobra.Command, args []string) {
		// pulling current config values to use as defaults,
		// but we don't care if it fails
		config, _ := edit_project.LoadConfig(configPath)

		if err := edit_project.SaveConfig(configPath, &config); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configPathCmd)
	configCmd.AddCommand(configShowCmd)
	configCmd.AddCommand(configCreateCmd)
}
