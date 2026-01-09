package cmd

import (
	"testing"

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
