package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/timhugh/edit_project/cli"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Interact with tool configuration",
}

var configCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Write current configuration to file",
	Run: func(cmd *cobra.Command, args []string) {
		if err := cli.ConfigCreate(stdout, configPath); err != nil {
			stderr.Println("Error creating configuration file:", err)
			os.Exit(1)
		}
	},
}

var configEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit the configuration file in your default editor",
	Run: func(cmd *cobra.Command, args []string) {
		if err := cli.ConfigEdit(stdout, configPath); err != nil {
			stderr.Println("Error editing configuration file:", err)
			os.Exit(1)
		}
	},
}

var configPathCmd = &cobra.Command{
	Use:   "path",
	Short: "Show the path to the configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		if err := cli.ConfigPath(stdout, configPath); err != nil {
			stderr.Println("Error getting configuration file path:", err)
			os.Exit(1)
		}
	},
}

var configResetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset the configuration file to default values",
	Run: func(cmd *cobra.Command, args []string) {
		if err := cli.ConfigReset(stdout, configPath); err != nil {
			stderr.Println("Error resetting configuration file:", err)
			os.Exit(1)
		}
	},
}

var configShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the current configuration",
	Run: func(cmd *cobra.Command, args []string) {
		if err := cli.ConfigShow(stdout, configPath); err != nil {
			stderr.Println("Error showing configuration file:", err)
			os.Exit(1)
		}
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
