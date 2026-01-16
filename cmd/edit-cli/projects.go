package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/timhugh/edit_project/cli"
)

var projectsCommand = &cobra.Command{
	Use:   "projects",
	Short: "Interact with project directories",
}

var projectsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all projects in configured workspaces",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: add arg for output format
		if err := cli.ProjectsList(stdout, configPath, "JSON"); err != nil {
			stderr.Println("Error listing projects:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(projectsCommand)
	projectsCommand.AddCommand(projectsListCmd)
}
