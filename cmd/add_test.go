package cmd

import (
	"testing"

	"wt/internal/domain"

	"github.com/spf13/cobra"
)

func newAddSUT() *cobra.Command {
	deps := domain.Dependencies{}
	return NewAddCmd(deps)
}

func TestAddVersion(t *testing.T) {
	sut := newAddSUT()
	const expectedVersion = "0.1.0"
	if sut.Version != expectedVersion {
		t.Errorf("Expected version %v, got %v instead.", expectedVersion, sut.Version)
	}
}
