package cmd

import (
	"testing"

	"wt/internal/domain"

	"github.com/spf13/cobra"
)

func newListSUT() *cobra.Command {
	deps := domain.Dependencies{}
	return NewListCmd(deps)
}

func TestListVersion(t *testing.T) {
	sut := newListSUT()
	const expectedVersion = "0.1.0"
	if sut.Version != expectedVersion {
		t.Errorf("Expected version %v, got %v instead.", expectedVersion, sut.Version)
	}
}

func TestListCommand(t *testing.T) {
	sut := newListSUT()
	err := sut.Execute()
	if err != nil {
		t.Errorf("Expected to pass with no error, got %v insted", err)
	}
}
