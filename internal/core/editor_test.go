package core_test

import (
	"testing"

	"github.com/timhugh/edit_project/internal/util"
)

func TestPathToExecutable(t *testing.T) {
	t.Run("returns path if found", func(t *testing.T) {
		path, err := util.PathToExecutable("go")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if path == "" {
			t.Fatalf("expected a valid path, got empty string")
		}
	})

	t.Run("returns error if not found", func(t *testing.T) {
		_, err := util.PathToExecutable("nonexistent_executable_12345")
		if err == nil {
			t.Fatalf("expected an error, got nil")
		}
	})
}
