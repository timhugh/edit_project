package core

import(
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func expandTildePath(path string) (string, error) {
	if strings.HasPrefix(path, "~/") {
		userHomeDir, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("failed to get user home directory: %w", err)
		}
		return filepath.Join(userHomeDir, path[2:]), nil
	}
	return path, nil
}
