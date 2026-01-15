package main

import (
	"github.com/spf13/cobra"
	"github.com/timhugh/edit_project/cli"
	"os"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install the edit CLI tool",
	Run: func(cmd *cobra.Command, args []string) {
		shell, err := cmd.Flags().GetString("shell")
		if err != nil {
			stderr.Println("Error reading shell flag:", err)
			os.Exit(1)
		}
		if err := cli.Install(stdout, shell); err != nil {
			stderr.Println("Failed to install edit-cli tool:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
	installCmd.Flags().StringP("shell", "s", "bash", "Specify the shell type (e.g., bash, zsh)")
}
