package edit_project

import (
	"fmt"
	"io"
	_ "embed"
)

//go:embed install.bash
var bashInstall string

//go:embed install.zsh
var zshInstall string

func Install(shell string, w io.Writer) error {
	switch shell {
	case "bash":
		fmt.Fprintln(w, bashInstall)
		return nil
	case "zsh":
		fmt.Fprintln(w, zshInstall)
		return nil
	default:
		return fmt.Errorf("unsupported shell type: %s", shell)
	}
}
