package main

import (
	"github.com/spf13/cobra"
	"github.com/timhugh/edit_project/internal/cli"
)

var installCmd = &cobra.Command{
	Use:   "install [shell]",
	Short: "Install the edit CLI tool",
	Long: `Outputs the shell code to install the edit_project and open_project commands for the specified shell.

Example usage:

	edit-cli install bash >> ~/.bashrc
	edit-cli install zsh >> ~/.zshrc`,

	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		shell := args[0]
		return cli.Install(stdout, shell)
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
