package config

import (
	"os"
	"testing"
)

func TestEnsureExistsCreatesMissingFile(t *testing.T) {
	tmpDir := t.TempDir()
	configPath := tmpDir + "/.wt/config.json"

	sut := JSONConfigStore{Path: configPath}
	if err := sut.EnsureExists(); err != nil {
		t.Fatalf("expected EnsureExists to succeed, got error: %v", err)
	}

	if _, err := os.Stat(configPath); err != nil {
		t.Fatalf("expected config file to be created, got error: %v", err)
	}

	paths, err := sut.GetManagedPaths()
	if err != nil {
		t.Fatalf("expected GetManagedPaths to succeed, got error: %v", err)
	}
	if len(paths) != 0 {
		t.Fatalf("expected empty managed paths, got %v", paths)
	}
}

func TestEnsureExistsKeepsExistingFile(t *testing.T) {
	tmpDir := t.TempDir()
	configPath := tmpDir + "/config.json"
	original := "{\"managedRepos\":[\"/a/git/repo\"]}"
	if err := os.WriteFile(configPath, []byte(original), 0644); err != nil {
		t.Fatalf("failed to create config file: %v", err)
	}

	sut := JSONConfigStore{Path: configPath}
	if err := sut.EnsureExists(); err != nil {
		t.Fatalf("expected EnsureExists to succeed, got error: %v", err)
	}

	paths, err := sut.GetManagedPaths()
	if err != nil {
		t.Fatalf("expected GetManagedPaths to succeed, got error: %v", err)
	}
	if len(paths) != 1 || paths[0] != "/a/git/repo" {
		t.Fatalf("expected managed path '/a/git/repo', got %v", paths)
	}
}
