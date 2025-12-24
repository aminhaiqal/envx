package unit

import (
	"testing"

	"github.com/axelyn/envx/internal/profile"
	"github.com/axelyn/envx/internal/storage"
)

func TestInitProjectCreatesDefaultEnvironment(t *testing.T) {
	tmpDir := t.TempDir()

	store, err := storage.NewWithBasePath(tmpDir)
	if err != nil {
		t.Fatalf("failed to create storage: %v", err)
	}

	manager := profile.New(store)

	err = manager.InitProject(
		"myapp",
		"My test app",
		"development",
	)
	if err != nil {
		t.Fatalf("InitProject failed: %v", err)
	}

	project, err := store.LoadProject("myapp")
	if err != nil {
		t.Fatalf("project not found after init")
	}

	if project.DefaultEnv != "development" {
		t.Errorf("expected default env development, got %s", project.DefaultEnv)
	}

	env, ok := project.Environments["development"]
	if !ok {
		t.Fatalf("development environment not created")
	}

	if env.Name != "development" {
		t.Errorf("expected env name development, got %s", env.Name)
	}
}
