package edit_project

import (
	_ "embed"
	"fmt"
	"io"
)

//go:embed install.bash
var bashInstall string

//go:embed install.zsh
var zshInstall string

func Install(shell string, w io.Writer) error {
	switch shell {
	case "bash":
		if _, err := fmt.Fprintln(w, bashInstall); err != nil {
			return err
		}
		return nil
	case "zsh":
		if _, err := fmt.Fprintln(w, zshInstall); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("unsupported shell type: %s", shell)
	}
}
