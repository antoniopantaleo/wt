package cmd

import (
	"testing"

	"wt/internal/domain"

	"github.com/spf13/cobra"
)

type mockConfigStore struct {}
func (s mockConfigStore) GetManagedPaths() ([]string, error) {
	return []string{"mock path 1"}, nil
}

type mockRenderer struct {
	result string
}
func (r *mockRenderer) RenderWorktrees(worktrees []domain.Worktree) {
	r.result = "test"
}

type mockGit struct {}
func (g mockGit) GetWorktreesFromPath(path string) []domain.Worktree {
	return []domain.Worktree{
		{Path: "Example path one"},
	}
}

func newListSUT(renderer mockRenderer) (*cobra.Command, *mockRenderer) {
	deps := domain.Dependencies{
		ConfigStore: mockConfigStore{},
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
	if renderer.result != "test" {
		t.Errorf("Expected 'result', received %v instead", renderer.result)
	}
}
