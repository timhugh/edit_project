package core

import (
	"fmt"
	"os"
	"syscall"
	"github.com/hairyhenderson/go-which"
)

func PathToExecutable(executableName string) (string, error) {
	path := which.Which(executableName)
	if path == "" {
		return "", fmt.Errorf("executable %s not found in PATH", executableName)
	}
	return path, nil
}

func OpenEditor(config Config, args ...string) error {
	editorPath, err := PathToExecutable(config.Editor)
	if err != nil {
		return fmt.Errorf("failed to find editor executable: %w", err)
	}
	
	command := append([]string{editorPath}, args...)
	env := os.Environ()
	return syscall.Exec(editorPath, command, env)
}
