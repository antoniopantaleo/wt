package cmd

import (
	"reflect"
	"testing"

	"wt/internal/domain"

	"github.com/spf13/cobra"
)

func newListSUT(renderer mockRenderer) (*cobra.Command, *mockRenderer) {
	deps := domain.Dependencies{
		ConfigStore: mockConfigStore{
			getManagedPaths: func() ([]string, error) {
				return []string{"/Users/antonio/Development/TestProject"}, nil
			},
		},
		Renderer: &renderer,
		Git: mockGit{},
	}
	return NewListCmd(deps), &renderer
}

func TestListVersion(t *testing.T) {
	sut, _ := newListSUT(mockRenderer{})
	const expectedVersion = "0.1.0"
	if sut.Version != expectedVersion {
		t.Errorf("Expected version %v, got %v instead.", expectedVersion, sut.Version)
	}
}

func TestListCommand(t *testing.T) {
	sut, renderer := newListSUT(mockRenderer{})
	err := sut.Execute()
	if err != nil {
		t.Errorf("Expected to pass with no error, got %v insted", err)
	}
	if !reflect.DeepEqual(renderer.branches, []string{"develop"}) {
		t.Errorf("Expected 'develop', received %v instead", renderer.branches)
	}
}
