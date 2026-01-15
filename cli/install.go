package cli

import (
	"os"

	"github.com/timhugh/edit_project"
)

func Install(out *Output, shell string) error {
	return edit_project.Install(shell, os.Stdout)
}
