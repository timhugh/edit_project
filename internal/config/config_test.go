package config_test

import (
	"encoding/json"
	"errors"
	"os"
	"testing"

	"github.com/go-test/deep"
	"github.com/timhugh/edit_project/internal/config"
)

func TestDefaultConfig(t *testing.T) {
	cfg := config.Default()
	if diff := deep.Equal(cfg, config.Config{
		Workspaces: []config.WorkspaceConfig{{Path: "~/git", UserPrefixes: true}},
		GitUsers:   []string{},
		Editor:     "nvim",
	}); diff != nil {
		t.Errorf("default config does not match expected: %v", diff)
	}
}

func TestLoadConfig(t *testing.T) {
	t.Run("valid config file populates config", func(t *testing.T) {
		cfg, err := config.Load("testdata/full_config.json")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if diff := deep.Equal(cfg, config.Config{
			Workspaces: []config.WorkspaceConfig{
				{Path: "~/projects", UserPrefixes: true},
				{Path: "~/work", UserPrefixes: false},
			},
			GitUsers: []string{"my_git_user", "my_work_org"},
			Editor:   "emacs",
		}); diff != nil {
			t.Errorf("loaded config does not match expected: %v", diff)
		}
	})

	t.Run("partial config file uses defaults for missing values", func(t *testing.T) {
		defaultConfig := config.Default()
		cfg, err := config.Load("testdata/partial_config.json")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if diff := deep.Equal(cfg, config.Config{
			Workspaces: defaultConfig.Workspaces,
			GitUsers:   []string{"my_git_user"},
			Editor:     defaultConfig.Editor,
		}); diff != nil {
			t.Errorf("loaded config does not match expected: %v", diff)
		}
	})

	t.Run("non-existent config file returns expected error", func(t *testing.T) {
		_, err := config.Load("testdata/non_existent.json")
		if err == nil {
			t.Fatalf("expected error for non-existent file, got nil")
		}
		if !errors.Is(err, os.ErrNotExist) {
			t.Fatalf("expected os.ErrNotExist, got %v", err)
		}
	})
}

func TestSaveConfig(t *testing.T) {
	t.Run("saves config to file correctly", func(t *testing.T) {
		file, err := os.CreateTemp("", "config_*.json")
		if err != nil {
			t.Fatalf("unexpected error creating temp file: %v", err)
		}

		writtenConfig := &config.Config{
			Workspaces: []config.WorkspaceConfig{
				{Path: "~/my_projects", UserPrefixes: true},
			},
			GitUsers: []string{"user1", "user2"},
			Editor:   "code",
		}
		err = config.Save(file.Name(), writtenConfig)
		if err != nil {
			t.Fatalf("unexpected error saving config: %v", err)
		}

		readConfigRaw, err := os.ReadFile(file.Name())
		if err != nil {
			t.Fatalf("unexpected error reading config file: %v", err)
		}
		var readConfig config.Config
		err = json.Unmarshal(readConfigRaw, &readConfig)
		if err != nil {
			t.Fatalf("unexpected error unmarshaling config: %v", err)
		}

		if diff := deep.Equal(readConfig, *writtenConfig); diff != nil {
			t.Errorf("saved config does not match written config: %v", diff)
		}
	})
}
