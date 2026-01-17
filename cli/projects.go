package cli

import (
	"encoding/json"
	"fmt"
	"github.com/timhugh/edit_project"
)

type OutputFormat string

const (
	FormatList OutputFormat = "list"
	FormatJSON OutputFormat = "json"
)

func getAllProjects(configPath string) ([]edit_project.Project, error) {
	config, err := edit_project.LoadConfig(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load configuration: %w", err)
	}
	projects, err := edit_project.ListAllProjects(&config)
	if err != nil {
		return nil, fmt.Errorf("failed to list projects: %w", err)
	}
	return projects, nil
}

func ProjectsList(out *Output, configPath string, format OutputFormat) error {
	projects, err := getAllProjects(configPath)
	if err != nil {
		return err
	}

	switch format {
	case FormatJSON:
		jsonOutput, err := json.MarshalIndent(projects, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal projects to JSON: %w", err)
		}
		out.Println(string(jsonOutput))
	default:
		for _, project := range projects {
			out.Println(project.AbsPath)
		}
	}

	return nil
}
