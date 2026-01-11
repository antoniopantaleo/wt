package cmd

import (
	"testing"

	"wt/internal/domain"

	"github.com/spf13/cobra"
)

func newRemoveSUT() *cobra.Command {
	deps := domain.Dependencies{}
	return NewRemoveCmd(deps)
}

func TestRemoveVersion(t *testing.T) {
	sut := newRemoveSUT()
	const expectedVersion = "0.1.0"
	if sut.Version != expectedVersion {
		t.Errorf("Expected version %v, got %v instead.", expectedVersion, sut.Version)
	}
}
