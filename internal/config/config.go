package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/timhugh/edit_project/internal/util"
)

var DefaultPath = os.ExpandEnv("$HOME/.config/edit_project/config.json")

type Config struct {
	Workspaces []WorkspaceConfig `json:"workspaces"`
	GitUsers   []string          `json:"git_users"`
	Editor     string            `json:"editor"`
}

type WorkspaceConfig struct {
	Path         string `json:"path"`
	UserPrefixes bool   `json:"user_prefixes"`
}

func (c *Config) EditorFullPath() (string, error) {
	return util.PathToExecutable(c.Editor)
}

func Default() Config {
	return Config{
		Workspaces: []WorkspaceConfig{
			{
				Path:         "~/git",
				UserPrefixes: true,
			},
		},
		GitUsers: []string{},
		Editor:   "nvim",
	}
}

func Load(path string) (Config, error) {
	config := Default()
	file, err := os.Open(path)
	if err != nil {
		return config, err
	}
	defer func() { _ = file.Close() }()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return config, err
	}

	return config, nil
}

func Save(path string, config *Config) error {
	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return err
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(&config); err != nil {
		return err
	}

	return nil
}
