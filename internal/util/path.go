package util

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/hairyhenderson/go-which"
)

func PathToExecutable(executableName string) (string, error) {
	path := which.Which(executableName)
	if path == "" {
		return "", fmt.Errorf("executable %s not found in PATH", executableName)
	}
	return path, nil
}

func ExpandTildePath(path string) (string, error) {
	if strings.HasPrefix(path, "~/") {
		userHomeDir, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("failed to get user home directory: %w", err)
		}
		return filepath.Join(userHomeDir, path[2:]), nil
	}
	return path, nil
}
