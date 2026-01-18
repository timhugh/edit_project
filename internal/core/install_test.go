package core_test

import (
	_ "embed"
	"github.com/timhugh/edit_project/internal/core"
	"strings"
	"testing"
)

//go:embed install.bash
var bashInstall string

//go:embed install.zsh
var zshInstall string

func TestInstall(t *testing.T) {
	tests := []struct {
		Shell          string
		ExpectedOutput string
		ExpectError    bool
	}{
		{"bash", bashInstall + "\n", false},
		{"zsh", zshInstall + "\n", false},
		{"unsupported shell", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.Shell, func(t *testing.T) {
			w := &strings.Builder{}
			err := core.Install(tt.Shell, w)
			if tt.ExpectError {
				if err == nil {
					t.Fatalf("expected error for unsupported shell, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			output := w.String()
			if output != tt.ExpectedOutput {
				t.Fatalf("expected:\n%q\ngot:\n%q", tt.ExpectedOutput, output)
			}
		})
	}
}
