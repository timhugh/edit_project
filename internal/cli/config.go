package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/timhugh/edit_project/internal/core"
)

func loadConfigOrDefault(configPath string) (core.Config, error) {
	config, err := core.LoadConfig(configPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return core.DefaultConfig(), nil
		}
		return core.Config{}, fmt.Errorf("failed to load configuration: %w", err)
	}
	return config, nil
}

func saveConfig(out *Output, configPath string, config *core.Config, confirm bool) error {
	if confirm {
		jsonOutput, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal configuration: %w", err)
		}
		out.Printf("Configuration to be written to %s:\n", configPath)
		out.Println(string(jsonOutput))
		shouldContinue, err := confirmPrompt(out, "Continue?", false)
		if err != nil {
			return err
		}
		if !shouldContinue {
			out.Println("Aborting.")
			return nil
		}
	}
	if err := core.SaveConfig(configPath, config); err != nil {
		return fmt.Errorf("failed to write configuration to file %s: %w", configPath, err)
	}
	jsonOutput, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal configuration: %w", err)
	}
	out.Println("Configuration written to", configPath)
	out.Println(string(jsonOutput))
	return nil
}

func ConfigCreate(out *Output, configPath string) error {
	config, err := loadConfigOrDefault(configPath)
	if err != nil {
		return err
	}

	return saveConfig(out, configPath, &config, true)
}

func ConfigEdit(out *Output, configPath string) error {
	config, err := core.LoadConfig(configPath)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("failed to load configuration: %w", err)
		}
		out.Println("Configuration file does not exist; creating with default values.")
		if err := saveConfig(out, configPath, &config, false); err != nil {
			return fmt.Errorf("failed to write default configuration: %w", err)
		}
	}
	return core.OpenEditor(config, configPath)
}

func ConfigEditorPath(out *Output, configPath string) error {
	config, err := core.LoadConfig(configPath)
	if err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}
	editorPath, err := config.EditorFullPath()
	if err != nil {
		return fmt.Errorf("failed to get editor path: %w", err)
	}
	out.Println(editorPath)
	return nil
}

func ConfigPath(out *Output, configPath string) error {
	out.Println(configPath)
	return nil
}

func ConfigReset(out *Output, configPath string) error {
	defaultConfig := core.DefaultConfig()
	return saveConfig(out, configPath, &defaultConfig, true)
}

func ConfigShow(out *Output, configPath string) error {
	config, err := core.LoadConfig(configPath)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("failed to load configuration: %w", err)
	}
	jsonOutput, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal configuration: %w", err)
	}
	out.Println(string(jsonOutput))
	return nil
}
