package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/timhugh/edit_project"
	"github.com/timhugh/edit_project/cli"
)

var configPath string
var stdout = cli.NewOutput(os.Stdout)
var stderr = cli.NewOutput(os.Stderr)

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "edit-cli",
	Short: "Quickly open projects in your editor",
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the current version of edit-cli",
	Run: func(cmd *cobra.Command, args []string) {
		stdout.Printf("edit-cli version %s", edit_project.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.PersistentFlags().StringVar(&configPath, "config", edit_project.DefaultConfigPath, "Path to configuration file")
}
