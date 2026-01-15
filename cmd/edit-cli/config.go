package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/timhugh/edit_project"
	"os"
	"strings"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Interact with tool configuration",
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

		cmd.Printf("Workspace Directory [%s]: ", config.WorkspaceDirectory)
		var workspaceInput string
		_, err := fmt.Scanln(&workspaceInput)
		if err != nil {
			if err.Error() != "unexpected newline" {
				panic(err)
			}
		}
		if strings.TrimSpace(workspaceInput) != "" {
			config.WorkspaceDirectory = strings.TrimSpace(workspaceInput)
		}
		
		if config.GitUsername == "" {
			cmd.Printf("Git Username: ")
		} else {
			cmd.Printf("Git Username [%s]: ", config.GitUsername)
		}
		_, err = fmt.Scanln(&config.GitUsername)
		if err != nil {
			if err.Error() != "unexpected newline" {
				panic(err)
			}
		}

		cmd.Printf("Editor [%s]: ", config.Editor)
		_, err = fmt.Scanln(&config.Editor)
		if err != nil {
			if err.Error() != "unexpected newline" {
				panic(err)
			}
		}

		jsonOutput, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			panic(err)
		}
		cmd.Println("Configuration:")
		cmd.Println(string(jsonOutput))
		cmd.Printf("Save to %s? (Y/n): ", configPath)
		var response string
		_, err = fmt.Scanln(&response)
		if err != nil {
			if err.Error() != "unexpected newline" {
				panic(err)
			}
		}
		if response != "Y" && response != "y" && response != "" {
			cmd.Println("Aborting, configuration not saved.")
			return
		}

		err = edit_project.SaveConfig(configPath, config)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configShowCmd)
	configCmd.AddCommand(configCreateCmd)
}
