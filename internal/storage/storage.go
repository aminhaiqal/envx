package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/axelyn/envx/pkg/envx"
)

type Storage struct {
	basePath string
}

// create new storage
func New() (*Storage, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory: %w", err)
	}

	basePath := filepath.Join(home, ".config", "envx")

	if err := os.MkdirAll(basePath, 0755); err != nil {
		return nil, fmt.Errorf("failed ot create config directory: %w", err)
	}

	return &Storage{basePath: basePath}, nil
}

// NewWithBasePath creates storage using a custom base path (for tests)
func NewWithBasePath(basePath string) (*Storage, error) {
	if err := os.MkdirAll(basePath, 0755); err != nil {
		return nil, fmt.Errorf("failed to create base directory: %w", err)
	}

	return &Storage{basePath: basePath}, nil
}


// save project
func(s *Storage) SaveProject(project *envx.Project) error {
	projectPath := filepath.Join(s.basePath, project.Name+".json")

	data, err := json.MarshalIndent(project, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal project: %w", err)
	}

	if err := os.WriteFile(projectPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write project file: %w", err)
	}

	return nil
}

// load project
func (s *Storage) LoadProject(name string) (*envx.Project, error) {
	projectPath := filepath.Join(s.basePath, name+".json")

	data, err := os.ReadFile(projectPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("project '%s' not found", name)
		}
		return nil, fmt.Errorf("failed to read project file %w", err)
	}

	var project envx.Project
	if err := json.Unmarshal(data, &project); err != nil {
		return nil, fmt.Errorf("failed to load unmarshal project: %w", err)
	}
	
	return &project, nil
}

// list project
func (s *Storage) ListProject() ([]string, error) {
	files, err := os.ReadDir(s.basePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}

	var projects []string
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".json" {
			name := file.Name()[:len(file.Name())-5]
			projects = append(projects, name)
		}
	}

	return projects, nil

}

// delete project
func (s *Storage) DeleteProject(name string) error {
	projectPath := filepath.Join(s.basePath, name+".json")

	if err := os.Remove(projectPath); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("project '%s' not found", name)
		}
		return fmt.Errorf("failed to delete project: %w", err)
	}

	return nil
}

// check if project exists
func(s *Storage) ProjectExists(name string) bool {
	projectPath := filepath.Join(s.basePath, name+".json")
	_, err := os.Stat(projectPath)

	return err == nil
}