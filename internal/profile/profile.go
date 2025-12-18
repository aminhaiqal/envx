package profile

import (
	"fmt"
	"time"

	"github.com/axelyn/envx/internal/storage"
	"github.com/axelyn/envx/pkg/envx"
)

type Manager struct {
	storage *storage.Storage
}

func New(storage *storage.Storage) *Manager {
	return &Manager{storage: storage}
}

func (m *Manager) InitProject(name, description, env string) error {
	if m.storage.ProjectExists(name) {
		return fmt.Errorf("project '%s' already exists", name)
	}

	project := &envx.Project{
		Base: envx.Base{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name: name,
		Description: description,
		Environments: map[string]envx.Environment{
			env: {
				Name: env,
				Variables: make(map[string]envx.Variable),
			},
		},
		DefaultEnv: env,
	}

	return m.storage.SaveProject(project)
}

func (m *Manager) SetVariable(projectName, env, key, value, description string, isSecret bool) error {
	project, err := m.storage.LoadProject(projectName)
	if err != nil {
		return err
	}

	if _, exists := project.Environments[env]; !exists {
		project.Environments[env] = envx.Environment{
			Name: 		env,
			Variables: 	make(map[string]envx.Variable),
		}
	}

	environment := project.Environments[env]
	environment.Variables[key] = envx.Variable{
		Base: envx.Base{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Key: key,
		Value: value,
		Description: description,
		IsSecret: isSecret,		
	}

	project.Environments[env] = environment
	project.Base.UpdatedAt = time.Now()

	return m.storage.SaveProject(project)
}

func (m *Manager) GetVariable(projectName, env, key string) (*envx.Variable, error) {
	project, err := m.storage.LoadProject(projectName)
	if err != nil {
		return nil, err
	}

	environment, exists := project.Environments[env]
	if !exists {
		return nil, fmt.Errorf("environment '%s' not found", env)
	}

	variable, exists := environment.Variables[key]
	if !exists {
		return nil, fmt.Errorf("variable '%s' not found", key)
	}

	return &variable, nil
}

func (m *Manager) ListVariable(projectName, env string) (map[string]envx.Variable, error) {
	project, err := m.storage.LoadProject(projectName)
	if err != nil {
		return nil, err
	}

	environment, exists := project.Environments[env]
	if !exists {
		return nil, fmt.Errorf("environment '%s' not found", env)
	}

	return environment.Variables, nil
}

func (m *Manager) DeleteVariable(projectName, env, key string) error {
	project, err := m.storage.LoadProject(projectName)
	if err != nil {
		return err
	}

	environment, exists := project.Environments[env]
	if !exists {
		return fmt.Errorf("environment '%s' not found", env)
	}

	if _, exists := environment.Variables[key]; !exists {
		return  fmt.Errorf("variable '%s' not found", key)
	}

	delete(environment.Variables, key)
	project.Environments[key] = environment
	project.Base.UpdatedAt = time.Now()

	return m.storage.SaveProject(project)
}