package edit_project

import (
	"encoding/json"
	"os"
	"path/filepath"
)

var DefaultConfigPath = os.ExpandEnv("$HOME/.config/edit_project/config.json")

type Config struct {
	Workspaces []string `json:"workspaces"`
	GitUsers   []string `json:"git_users"`
	Editor     string   `json:"editor"`
}

func DefaultConfig() Config {
	return Config{
		Workspaces: []string{"~/git"},
		GitUsers:   []string{},
		Editor:     "nvim",
	}
}

func LoadConfig(path string) (Config, error) {
	config := DefaultConfig()
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

func SaveConfig(path string, config *Config) error {
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
