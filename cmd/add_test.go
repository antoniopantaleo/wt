package cmd

import (
	"os"
	"testing"

	"wt/internal/config"
	"wt/internal/domain"

	"github.com/spf13/cobra"
)

func newAddSUT(configStore domain.ConfigStore, git domain.Git) (*cobra.Command, domain.ConfigStore) {
	deps := domain.Dependencies{
		ConfigStore: configStore,
		Git: git,
	}
	return NewAddCmd(deps), configStore
}

func TestAddVersion(t *testing.T) {
	sut, _ := newAddSUT(mockConfigStore{}, mockGit{})
	const expectedVersion = "0.1.0"
	if sut.Version != expectedVersion {
		t.Errorf("Expected version %v, got %v instead.", expectedVersion, sut.Version)
	}
}

func TestAddNewPathFailsIfThereIsMoreThenOnePath(t *testing.T) {
	sut, _ := newAddSUT(mockConfigStore{}, mockGit{})
	sut.SetArgs([]string{"/path/1", "/path/2"})
	err := sut.Execute()
	if err == nil {
		t.Errorf("Expected error when providing more than one path, but got no error")
	}
}

func TestAddNewPathFailsIfThereIsNoPath(t *testing.T) {
	sut, _ := newAddSUT(mockConfigStore{}, mockGit{})
	sut.SetArgs([]string{})
	err := sut.Execute()
	if err == nil {
		t.Errorf("Expected error when providing no path, but got no error")
	}
}

func TestAddNewPathDoesNotAddANonGitRepo(t *testing.T) {
	tmpDir := t.TempDir()
	if err := os.WriteFile(tmpDir+"/config.json", []byte("{}"), 0644); err != nil {
		t.Fatalf("failed to create empty config.json: %v", err)
	}
	sut, _ := newAddSUT(
		config.XDGConfigStore{
			Path: tmpDir+"/config.json"}, 
			mockGit{
				isGitRepo: func(path string) bool {
					return false
			},
		},
	)
	sut.SetArgs([]string{"/not/a/git/repo"})
	err := sut.Execute()
	if err == nil {
		t.Errorf("Expected error when providing a non git repo, but got no error")
	}
}

func TestAddNewPathCanAddNewGitRepo(t *testing.T) {
	tmpDir := t.TempDir()
	if err := os.WriteFile(tmpDir+"/config.json", []byte("{}"), 0644); err != nil {
		t.Fatalf("failed to create empty config.json: %v", err)
	}
	sut, store := newAddSUT(
		config.XDGConfigStore{
			Path: tmpDir+"/config.json"}, 
			mockGit{
				isGitRepo: func(path string) bool {
					return true
			},
		},
	)
	sut.SetArgs([]string{"/a/git/repo"})
	err := sut.Execute()
	if err != nil {
		t.Errorf("Expected no error when providing a git repo, but got error %v", err)
	}
	managedPaths, error := store.GetManagedPaths()
	if error != nil {
		t.Errorf("Expected to get managed paths with no error, but got error: %v", error)
	}
	if len(managedPaths) != 1 || managedPaths[0] != "/a/git/repo" {
		t.Errorf("Expected managed paths to contain '/a/git/repo', but got %v instead", managedPaths)
	}
}