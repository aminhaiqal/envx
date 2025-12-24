package unit

import (
	"testing"

	"github.com/axelyn/envx/internal/storage"
	"github.com/axelyn/envx/pkg/envx"
)

func TestStorage_SaveAndLoadProject(t *testing.T) {
	tmpDir := t.TempDir()

	store, err := storage.NewWithBasePath(tmpDir)
	if err != nil {
		t.Fatalf("failed to create storage: %v", err)
	}

	project := envx.Project{
		Name:        "myapp",
		Description: "test project",
		DefaultEnv:  "development",
		Environments: map[string]envx.Environment{
			"development": {
				Name: "development",
				Variables: map[string]envx.Variable{
					"PORT": {
						Key:   "PORT",
						Value: "3000",
					},
				},
			},
		},
	}

	// Save requires pointer
	if err := store.SaveProject(&project); err != nil {
		t.Fatalf("save failed: %v", err)
	}

	loaded, err := store.LoadProject("myapp")
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}

	// Project-level assertions
	if loaded.Name != "myapp" {
		t.Errorf("expected project name myapp, got %s", loaded.Name)
	}

	if loaded.DefaultEnv != "development" {
		t.Errorf("expected default env development, got %s", loaded.DefaultEnv)
	}

	// Environment-level assertions
	env, ok := loaded.Environments["development"]
	if !ok {
		t.Fatalf("development environment not found")
	}

	if env.Name != "development" {
		t.Errorf("expected env name development, got %s", env.Name)
	}

	// Variable-level assertions
	portVar, ok := env.Variables["PORT"]
	if !ok {
		t.Fatalf("PORT variable not found")
	}

	if portVar.Value != "3000" {
		t.Errorf("expected PORT=3000, got %s", portVar.Value)
	}
}
