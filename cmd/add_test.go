package cmd

import (
	"fmt"
	"os"
	"testing"

	"wt/internal/config"
	"wt/internal/domain"

	"github.com/spf13/cobra"
)

func newAddSUT(configStore domain.ConfigStore) (*cobra.Command, domain.ConfigStore) {
	deps := domain.Dependencies{
		ConfigStore: configStore,
	}
	return NewAddCmd(deps), configStore
}

func TestAddVersion(t *testing.T) {
	sut, _ := newAddSUT(mockConfigStore{})
	const expectedVersion = "0.1.0"
	if sut.Version != expectedVersion {
		t.Errorf("Expected version %v, got %v instead.", expectedVersion, sut.Version)
	}
}

func TestAddNewPathFailsIfThereIsMoreThenOnePath(t *testing.T) {
	sut, _ := newAddSUT(mockConfigStore{})
	sut.SetArgs([]string{"/path/1", "/path/2"})
	err := sut.Execute()
	if err == nil {
		t.Errorf("Expected error when providing more than one path, but got no error")
	}
}

func TestAddNewPathFailsIfThereIsNoPath(t *testing.T) {
	sut, _ := newAddSUT(mockConfigStore{})
	sut.SetArgs([]string{})
	err := sut.Execute()
	if err == nil {
		t.Errorf("Expected error when providing no path, but got no error")
	}
}

func TestAddNewPath(t *testing.T) {
	tmpDir := t.TempDir()
	if err := os.WriteFile(tmpDir+"/config.json", []byte("{}"), 0644); err != nil {
		t.Fatalf("failed to create empty config.json: %v", err)
	}
	fmt.Println(tmpDir)
	sut, store := newAddSUT(config.XDGConfigStore{Path: tmpDir+"/config.json"})
	sut.SetArgs([]string{"/path/1"})
	err := sut.Execute()
	if err != nil {
		t.Errorf("Expected no error when providing exactly one path, but got error: %v", err)
	}
	managedPaths, error := store.GetManagedPaths()
	if error != nil {
		t.Errorf("Expected to get managed paths with no error, but got error: %v", error)
	}
	if len(managedPaths) != 1 || managedPaths[0] != "/path/1" {
		t.Errorf("Expected managed paths to contain '/path/1', but got %v instead", managedPaths)
	}
}