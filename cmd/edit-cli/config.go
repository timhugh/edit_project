package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"syscall"

	"github.com/hairyhenderson/go-which"
	"github.com/spf13/cobra"
	"github.com/timhugh/edit_project"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Interact with tool configuration",
}

var configCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Write current configuration to file",
	Run: func(cmd *cobra.Command, args []string) {
		// We don't care if this fails because there probably isn't a config file yet
		// and we'll at least get the default values
		config, _ := edit_project.LoadConfig(configPath)

		if err := edit_project.SaveConfig(configPath, &config); err != nil {
			panic(err)
		}

		jsonOutput, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			panic(err)
		}
		cmd.Println("Configuration written to", configPath)
		cmd.Println(string(jsonOutput))
	},
}

var configEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit the configuration file in your default editor",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := edit_project.LoadConfig(configPath)
		if err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				panic(err)
			}
			cmd.Println("Configuration file does not exist; creating with default values.")
			if err := edit_project.SaveConfig(configPath, &config); err != nil {
				panic(err)
			}
		}
		pathToEditor := which.Which(config.Editor)
		command := []string{config.Editor, configPath}
		env := os.Environ()
		cmd.Println("Opening configuration file in editor:", pathToEditor)
		if err := syscall.Exec(pathToEditor, command, env); err != nil {
			panic(err)
		}
	},
}

var configPathCmd = &cobra.Command{
	Use:   "path",
	Short: "Show the path to the configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println(configPath)
	},
}

var configResetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset the configuration file to default values",
	Run: func(cmd *cobra.Command, args []string) {
		defaultConfig := edit_project.DefaultConfig()
		jsonOutput, err := json.MarshalIndent(defaultConfig, "", "  ")
		if err != nil {
			panic(err)
		}
		cmd.Printf("Configuration file %s will be reset to default values\n:", configPath)
		cmd.Println(string(jsonOutput))
		cmd.Printf("Continue? (y/N): ")
		var response string
		_, err = fmt.Scanln(&response)
		if err != nil || (response != "y" && response != "Y") {
			cmd.Println("Aborting.")
			return
		}
		if err := edit_project.SaveConfig(configPath, &defaultConfig); err != nil {
			panic(err)
		}
		cmd.Println("Configuration reset")
	},
}

var configShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the current configuration",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := edit_project.LoadConfig(configPath)
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			panic(err)
		}
		jsonOutput, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			panic(err)
		}
		cmd.Println(string(jsonOutput))
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configCreateCmd)
	configCmd.AddCommand(configEditCmd)
	configCmd.AddCommand(configPathCmd)
	configCmd.AddCommand(configShowCmd)
	configCmd.AddCommand(configResetCmd)
}
