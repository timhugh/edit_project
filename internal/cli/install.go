package cli

import (
	"os"

	"github.com/timhugh/edit_project/internal/core"
)

func Install(out *Output, shell string) error {
	return core.Install(shell, os.Stdout)
}
