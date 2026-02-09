package cmd

import (
	"os"
	"testing"

	"wt/internal/config"
	"wt/internal/domain"

	"github.com/spf13/cobra"
)

func newRemoveSUT(configStore domain.ConfigStore) (*cobra.Command, domain.ConfigStore) {
	deps := domain.Dependencies{ConfigStore: configStore}
	return NewRemoveCmd(deps), configStore
}

func TestRemoveVersion(t *testing.T) {
	sut, _ := newRemoveSUT(mockConfigStore{})
	const expectedVersion = "0.1.0"
	if sut.Version != expectedVersion {
		t.Errorf("Expected version %v, got %v instead.", expectedVersion, sut.Version)
	}
}

func TestRemovePathFailsIfThereIsMoreThanOnePath(t *testing.T) {
	sut, _ := newRemoveSUT(mockConfigStore{})
	sut.SetArgs([]string{"/path/1", "/path/2"})
	err := sut.Execute()
	if err == nil {
		t.Errorf("Expected error when providing more than one path, but got no error")
	}
}

func TestRemovePathFailsIfThereIsNoPath(t *testing.T) {
	sut, _ := newRemoveSUT(mockConfigStore{})
	sut.SetArgs([]string{})
	err := sut.Execute()
	if err == nil {
		t.Errorf("Expected error when providing no path, but got no error")
	}
}

func TestRemovePathCanRemoveManagedRepo(t *testing.T) {
	tmpDir := t.TempDir()
	configFilePath := tmpDir + "/config.json"
	configJSON := `{"managedRepos":["/a/git/repo","/b/git/repo"]}`
	if err := os.WriteFile(configFilePath, []byte(configJSON), 0644); err != nil {
		t.Fatalf("failed to create config.json: %v", err)
	}

	sut, store := newRemoveSUT(config.JSONConfigStore{Path: configFilePath})
	sut.SetArgs([]string{"/a/git/repo"})
	err := sut.Execute()
	if err != nil {
		t.Errorf("Expected no error when removing a managed repo, but got error %v", err)
	}

	managedPaths, getErr := store.GetManagedPaths()
	if getErr != nil {
		t.Errorf("Expected to get managed paths with no error, but got error: %v", getErr)
	}
	if len(managedPaths) != 1 || managedPaths[0] != "/b/git/repo" {
		t.Errorf("Expected managed paths to contain '/b/git/repo', but got %v instead", managedPaths)
	}
}

func TestRemovePathFailsWhenPathIsNotManaged(t *testing.T) {
	tmpDir := t.TempDir()
	configFilePath := tmpDir + "/config.json"
	configJSON := `{"managedRepos":["/a/git/repo"]}`
	if err := os.WriteFile(configFilePath, []byte(configJSON), 0644); err != nil {
		t.Fatalf("failed to create config.json: %v", err)
	}

	sut, _ := newRemoveSUT(config.JSONConfigStore{Path: configFilePath})
	sut.SetArgs([]string{"/not/managed"})
	err := sut.Execute()
	if err == nil {
		t.Errorf("Expected error when removing a path that is not managed, but got no error")
	}
}
