package cli

import (
	"encoding/json"
	"fmt"

	fzf "github.com/junegunn/fzf/src"
	"github.com/timhugh/edit_project/internal/config"
	"github.com/timhugh/edit_project/internal/core"
)

type PathOutput int

const (
	AbsolutePathOutput PathOutput = iota
	RelativePathOutput PathOutput = iota
)

type OutputFormat int

const (
	FormatList OutputFormat = iota
	FormatJSON OutputFormat = iota
)

func getAllProjects(configPath string) ([]core.Project, error) {
	cfg, err := config.Load(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load configuration: %w", err)
	}
	projects, err := core.ListAllProjects(&cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to list projects: %w", err)
	}
	return projects, nil
}

func ProjectsList(out *Output, configPath string, format OutputFormat, pathOutput PathOutput) error {
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
			if pathOutput == RelativePathOutput {
				out.Println(project.RelPath)
			} else {
				out.Println(project.AbsPath)
			}
		}
	}

	return nil
}

func ProjectsSearch(out *Output, configPath string, query string) error {
	projects, err := getAllProjects(configPath)
	if err != nil {
		return err
	}

	inputChan := make(chan string)
	outputChan := make(chan string)

	// not handling errors since we're not parsing anything
	options, _ := fzf.ParseOptions(true, []string{})
	options.Input = inputChan
	options.Output = outputChan
	options.Query = query
	options.Exit0 = true
	options.Select1 = true

	complete := make(chan bool)
	defer close(complete)

	go func() {
		for _, project := range projects {
			inputChan <- project.RelPath
		}
		close(inputChan)
	}()
	go func() {
		for s := range outputChan {
			for _, project := range projects {
				if project.RelPath == s {
					out.Println(project.AbsPath)
				}
			}
			complete <- true
		}
	}()

	if _, err := fzf.Run(options); err != nil {
		return fmt.Errorf("fzf search failed: %w", err)
	}

	// with options.Select1, fzf.Run() seems to return before outputChan gets a chance to be read.
	// this ensures we wait for that to complete.
	<-complete
	return nil
}
