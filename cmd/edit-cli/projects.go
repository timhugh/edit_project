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

func parseFormat(format string) cli.OutputFormat {
	switch format {
	case "json":
		return cli.FormatJSON
	default:
		return cli.FormatList
	}
}

var projectsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all projects in configured workspaces",
	RunE: func(cmd *cobra.Command, args []string) error {
		formatArg, err := cmd.Flags().GetString("format")
		if err != nil {
			return err
		}

		format := parseFormat(formatArg)

		if err := cli.ProjectsList(stdout, configPath, format); err != nil {
			stderr.Println("Error listing projects:", err)
			os.Exit(1)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(projectsCommand)
	projectsCommand.AddCommand(projectsListCmd)
	projectsListCmd.Flags().String("format", "list", "Output format: list, json")
}
