package cmd

import (
	"errors"
	"os"
	"testing"
	"wt/internal/config"

	"wt/internal/domain"

	"github.com/spf13/cobra"
)

func newRootSUT() *cobra.Command {
	deps := domain.Dependencies{}
	return NewRootCmd(deps)
}

func TestRootVersion(t *testing.T) {
	sut := newRootSUT()
	const expectedVersion = "0.1.0"
	if sut.Version != expectedVersion {
		t.Errorf("Expected version %v, got %v instead.", expectedVersion, sut.Version)
	}
}

func TestRootPersistentPreRunCreatesConfigFile(t *testing.T) {
	tmpDir := t.TempDir()
	configPath := tmpDir + "/.wt/config.json"

	deps := domain.Dependencies{
		ConfigStore: config.JSONConfigStore{Path: configPath},
		Renderer:    &mockRenderer{},
	}
	sut := NewRootCmd(deps)
	sut.SetArgs([]string{"managed"})

	if err := sut.Execute(); err != nil {
		t.Fatalf("expected command to succeed, got error: %v", err)
	}

	if _, err := os.Stat(configPath); err != nil {
		t.Fatalf("expected config file to be created, got error: %v", err)
	}
}

func TestRootPersistentPreRunReturnsEnsureExistsError(t *testing.T) {
	expectedErr := errors.New("cannot initialize store")
	deps := domain.Dependencies{
		ConfigStore: mockConfigStore{
			ensureExists: func() error { return expectedErr },
		},
	}
	sut := NewRootCmd(deps)
	sut.SetArgs([]string{"managed"})

	err := sut.Execute()
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err.Error() != expectedErr.Error() {
		t.Fatalf("expected error %q, got %q", expectedErr.Error(), err.Error())
	}
}
