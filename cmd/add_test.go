package cmd

import (
	"testing"

	"wt/internal/domain"

	"github.com/spf13/cobra"
)

func newAddSUT() *cobra.Command {
	deps := domain.Dependencies{
		ConfigStore: mockConfigStore{},
	}
	return NewAddCmd(deps)
}

func TestAddVersion(t *testing.T) {
	sut := newAddSUT()
	const expectedVersion = "0.1.0"
	if sut.Version != expectedVersion {
		t.Errorf("Expected version %v, got %v instead.", expectedVersion, sut.Version)
	}
}

func TestAddNewPathFailsIfThereIsMoreThenOnePath(t *testing.T) {
	sut := newAddSUT()
	sut.SetArgs([]string{"/path/1", "/path/2"})
	err := sut.Execute()
	if err == nil {
		t.Errorf("Expected error when providing more than one path, but got no error")
	}
}