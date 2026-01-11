package cmd

import (
	"testing"

	"wt/internal/domain"

	"github.com/spf13/cobra"
)

func newManagedSUT() *cobra.Command {
	deps := domain.Dependencies{}
	return NewManagedCmd(deps)
}

func TestManagedVersion(t *testing.T) {
	sut := newManagedSUT()
	const expectedVersion = "0.1.0"
	if sut.Version != expectedVersion {
		t.Errorf("Expected version %v, got %v instead.", expectedVersion, sut.Version)
	}
}
