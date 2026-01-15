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

func init() {
	rootCmd.PersistentFlags().StringVar(&configPath, "config", edit_project.DefaultConfigPath, "Path to configuration file")
}
