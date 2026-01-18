package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/timhugh/edit_project/internal/cli"
)

var projectsCommand = &cobra.Command{
	Use:   "projects",
	Short: "Interact with project directories",
}

func parseFormat(cmd *cobra.Command) cli.OutputFormat {
	format, _ := cmd.Flags().GetString("format")
	switch format {
	case "json":
		return cli.FormatJSON
	default:
		return cli.FormatList
	}
}

func parsePathOutput(cmd *cobra.Command) cli.PathOutput {
	relative, _ := cmd.Flags().GetBool("relative")
	if relative {
		return cli.RelativePathOutput
	}
	return cli.AbsolutePathOutput
}

var projectsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all projects in configured workspaces",
	RunE: func(cmd *cobra.Command, args []string) error {
		format := parseFormat(cmd)
		pathOutput := parsePathOutput(cmd)

		if err := cli.ProjectsList(stdout, configPath, format, pathOutput); err != nil {
			stderr.Println("Error listing projects:", err)
			os.Exit(1)
		}
		return nil
	},
}

var projectsSearchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for projects",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var query string
		if len(args) > 0 {
			query = args[0]
		}

		err := cli.ProjectsSearch(stdout, configPath, query)
		if err != nil {
			stderr.Println("Error searching projects:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(projectsCommand)
	projectsCommand.AddCommand(projectsListCmd)
	projectsListCmd.Flags().String("format", "list", "Output format: list, json")
	projectsListCmd.Flags().Bool("relative", false, "Show relative paths (absolute paths are shown by default)")
	projectsCommand.AddCommand(projectsSearchCmd)
}
