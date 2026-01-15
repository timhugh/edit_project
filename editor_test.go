package edit_project_test

import (
	"testing"

	"github.com/timhugh/edit_project"
)

func TestPathToExecutable(t *testing.T) {
	t.Run("returns path if found", func(t *testing.T) {
		path, err := edit_project.PathToExecutable("go")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if path == "" {
			t.Fatalf("expected a valid path, got empty string")
		}
	})

	t.Run("returns error if not found", func(t *testing.T) {
		_, err := edit_project.PathToExecutable("nonexistent_executable_12345")
		if err == nil {
			t.Fatalf("expected an error, got nil")
		}
	})
}
