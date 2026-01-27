package core

import (
	"fmt"
	"os"
	"syscall"

	"github.com/timhugh/edit_project/internal/config"
	"github.com/timhugh/edit_project/internal/util"
)

func OpenEditor(cfg config.Config, args ...string) error {
	editorPath, err := util.PathToExecutable(cfg.Editor)
	if err != nil {
		return fmt.Errorf("failed to find editor executable: %w", err)
	}

	command := append([]string{editorPath}, args...)
	env := os.Environ()
	return syscall.Exec(editorPath, command, env)
}
