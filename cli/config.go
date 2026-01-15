package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"syscall"

	"github.com/hairyhenderson/go-which"
	"github.com/timhugh/edit_project"
)

func ConfigCreate(out *Output, configPath string) error {
	// We don't care if this fails because there probably isn't a config file yet
	// and we'll at least get the default values
	config, _ := edit_project.LoadConfig(configPath)

	if err := edit_project.SaveConfig(configPath, &config); err != nil {
		return err
	}

	jsonOutput, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	out.Println("Configuration written to", configPath)
	out.Println(string(jsonOutput))
	return nil
}

func ConfigEdit(out *Output, configPath string) error {
	config, err := edit_project.LoadConfig(configPath)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("failed to load configuration: %w", err)
		}
		out.Println("Configuration file does not exist; creating with default values.")
		if err := edit_project.SaveConfig(configPath, &config); err != nil {
			return fmt.Errorf("failed to write default configuration: %w", err)
		}
	}
	pathToEditor := which.Which(config.Editor)
	command := []string{config.Editor, configPath}
	env := os.Environ()
	out.Println("Opening configuration file in editor:", pathToEditor)
	if err := syscall.Exec(pathToEditor, command, env); err != nil {
		return fmt.Errorf("failed to open editor: %w", err)
	}
	return nil
}

func ConfigPath(out *Output, configPath string) error {
	out.Println(configPath)
	return nil
}

func ConfigReset(out *Output, configPath string) error {
	defaultConfig := edit_project.DefaultConfig()
	jsonOutput, err := json.MarshalIndent(defaultConfig, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal default configuration: %w", err)
	}
	out.Printf("Configuration file %s will be reset to default values\n:", configPath)
	out.Println(string(jsonOutput))
	out.Printf("Continue? (y/N): ")
	var response string
	_, err = fmt.Scanln(&response)
	if err != nil || (response != "y" && response != "Y") {
		out.Println("Aborting.")
		return nil
	}
	if err := edit_project.SaveConfig(configPath, &defaultConfig); err != nil {
		return fmt.Errorf("failed to write configuration to file %s: %w", configPath, err)
	}
	out.Println("Configuration reset")
	return nil
}

func ConfigShow(out *Output, configPath string) error {
	config, err := edit_project.LoadConfig(configPath)
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
