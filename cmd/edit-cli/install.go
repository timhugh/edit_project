package main

import (
	"github.com/spf13/cobra"
	"github.com/timhugh/edit_project"
	"os"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install the edit CLI tool",
	Run: func(cmd *cobra.Command, args []string) {
		shell, err := cmd.Flags().GetString("shell")
		if err != nil {
			panic(err)
		}
		err = edit_project.Install(shell, os.Stdout)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
	installCmd.Flags().StringP("shell", "s", "bash", "Specify the shell type (e.g., bash, zsh)")
}
