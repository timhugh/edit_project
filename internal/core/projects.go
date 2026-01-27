package core

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/timhugh/edit_project/internal/config"
	"github.com/timhugh/edit_project/internal/util"
)

type Project struct {
	Workspace string
	AbsPath   string
	RelPath   string
}

func ListProjectsInWorkspace(workspace string, includeUserPrefix bool) ([]Project, error) {
	fullDir, err := util.ExpandTildePath(workspace)
	if err != nil {
		return nil, err
	}
	workspaceName := filepath.Base(workspace)

	var dirs []string
	if includeUserPrefix {
		dirs, err = filepath.Glob(filepath.Join(fullDir, "*", "*"))
	} else {
		dirs, err = filepath.Glob(filepath.Join(fullDir, "*"))
	}
	if err != nil {
		return nil, err
	}

	var projects []Project
	for _, dir := range dirs {
		info, err := os.Stat(dir)
		if err != nil {
			return nil, err
		}

		// ignore non-directories and hidden directories
		// TODO: the dot checks are pretty naive and the second one could pick up false positives
		if !info.IsDir() ||
			strings.HasPrefix(filepath.Base(dir), ".") ||
			strings.Contains(dir, "/.") {
			continue
		}

		project := Project{
			Workspace: workspace,
			AbsPath:   dir,
		}
		relPath, err := filepath.Rel(fullDir, dir)
		if err != nil {
			return nil, err
		}
		if includeUserPrefix {
			project.RelPath = relPath
		} else {
			project.RelPath = filepath.Join(workspaceName, relPath)
		}

		projects = append(projects, project)
	}

	return projects, nil
}

func ListAllProjects(cfg *config.Config) ([]Project, error) {
	var projects []Project
	for _, workspace := range cfg.Workspaces {
		workspaceProjects, err := ListProjectsInWorkspace(workspace.Path, workspace.UserPrefixes)
		if err != nil {
			return nil, err
		}
		projects = append(projects, workspaceProjects...)
	}
	return projects, nil
}
