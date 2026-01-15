package edit_project_test

import (
	"testing"

	"github.com/timhugh/edit_project"
)

func TestDefaultConfig(t *testing.T) {
	config := edit_project.DefaultConfig()
	if config.WorkspaceDirectory != "~/git" {
		t.Errorf("expected WorkspaceDir to be '~/git', got '%s'", config.WorkspaceDirectory)
	}
	if config.GitUsername != "" {
		t.Errorf("expected Username to be '', got '%s'", config.GitUsername)
	}
	if config.Editor != "nvim" {
		t.Errorf("expected Editor to be 'nvim', got '%s'", config.Editor)
	}
}

func TestLoadConfig(t *testing.T) {
	t.Run("valid config file", func(t *testing.T) {
		config, err := edit_project.LoadConfig("testdata/full_config.json")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if config.WorkspaceDirectory != "~/my_workspace" {
			t.Errorf("expected WorkspaceDir to be '~/my_workspace', got '%s'", config.WorkspaceDirectory)
		}
		if config.GitUsername != "my_username" {
			t.Errorf("expected Username to be 'my_username', got '%s'", config.GitUsername)
		}
		if config.Editor != "emacs" {
			t.Errorf("expected Editor to be 'emacs', got '%s'", config.Editor)
		}
	})

	t.Run("partial config file", func(t *testing.T) {
		defaultConfig := edit_project.DefaultConfig()
		config, err := edit_project.LoadConfig("testdata/partial_config.json")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if config.WorkspaceDirectory != defaultConfig.WorkspaceDirectory {
			t.Errorf("expected WorkspaceDir to be '%s', got '%s'", defaultConfig.WorkspaceDirectory, config.WorkspaceDirectory)
		}
		if config.GitUsername != "git_user" {
			t.Errorf("expected Username to be 'git_user', got '%s'", config.GitUsername)
		}
		if config.Editor != defaultConfig.Editor {
			t.Errorf("expected Editor to be '%s', got '%s'", defaultConfig.Editor, config.Editor)
		}
	})

	t.Run("non-existent config file", func(t *testing.T) {
		_, err := edit_project.LoadConfig("testdata/non_existent.json")
		if err == nil {
			t.Fatalf("expected error for non-existent file, got nil")
		}
	})
}
