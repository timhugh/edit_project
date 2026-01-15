package edit_project_test

import (
	"encoding/json"
	"errors"
	"os"
	"testing"

	"github.com/go-test/deep"
	"github.com/timhugh/edit_project"
)

func TestDefaultConfig(t *testing.T) {
	config := edit_project.DefaultConfig()
	if diff := deep.Equal(config, edit_project.Config{
		Workspaces: []string{"~/git"},
		GitUsers:   []string{},
		Editor:     "nvim",
	}); diff != nil {
		t.Errorf("default config does not match expected: %v", diff)
	}
}

func TestLoadConfig(t *testing.T) {
	t.Run("valid config file populates config", func(t *testing.T) {
		config, err := edit_project.LoadConfig("testdata/full_config.json")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if diff := deep.Equal(config, edit_project.Config{
			Workspaces: []string{"~/projects", "~/work"},
			GitUsers:   []string{"my_git_user", "my_work_org"},
			Editor:     "emacs",
		}); diff != nil {
			t.Errorf("loaded config does not match expected: %v", diff)
		}
	})

	t.Run("partial config file uses defaults for missing values", func(t *testing.T) {
		defaultConfig := edit_project.DefaultConfig()
		config, err := edit_project.LoadConfig("testdata/partial_config.json")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if diff := deep.Equal(config, edit_project.Config{
			Workspaces: defaultConfig.Workspaces,
			GitUsers:   []string{"my_git_user"},
			Editor:     defaultConfig.Editor,
		}); diff != nil {
			t.Errorf("loaded config does not match expected: %v", diff)
		}
	})

	t.Run("non-existent config file returns expected error", func(t *testing.T) {
		_, err := edit_project.LoadConfig("testdata/non_existent.json")
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

		writtenConfig := &edit_project.Config{
			Workspaces: []string{"~/my_projects"},
			GitUsers:   []string{"user1", "user2"},
			Editor:     "code",
		}
		err = edit_project.SaveConfig(file.Name(), writtenConfig)
		if err != nil {
			t.Fatalf("unexpected error saving config: %v", err)
		}

		readConfigRaw, err := os.ReadFile(file.Name())
		if err != nil {
			t.Fatalf("unexpected error reading config file: %v", err)
		}
		var readConfig edit_project.Config
		err = json.Unmarshal(readConfigRaw, &readConfig)
		if err != nil {
			t.Fatalf("unexpected error unmarshaling config: %v", err)
		}

		if diff := deep.Equal(readConfig, *writtenConfig); diff != nil {
			t.Errorf("saved config does not match written config: %v", diff)
		}
	})
}
