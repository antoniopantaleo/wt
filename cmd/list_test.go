package cmd

import (
	"testing"
	"reflect"

	"wt/internal/domain"

	"github.com/spf13/cobra"
)

type mockConfigStore struct {}
func (s mockConfigStore) GetManagedPaths() ([]string, error) {
	return []string{"/Users/antonio/Development/TestProject"}, nil
}

type mockRenderer struct {
	branches []string
}
func (r *mockRenderer) RenderWorktrees(worktrees []domain.Worktree) {
	for _, worktree := range worktrees {
		r.branches = append(r.branches, worktree.Branch)
	}
}

type mockGit struct {}
func (g mockGit) GetWorktreesFromPath(path string) []domain.Worktree {
	return []domain.Worktree{
		{	Path: "/Users/antonio/Development/Worktrees/TestList/develop",
			RepoPath: "/Users/antonio/Development/TestProject",
			Branch: "develop",
			HeadSHA: "4e04b2b0961c494fb643d91c8956813dbfcc799d",
		},
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
	if !reflect.DeepEqual(renderer.branches, []string{"develop"}) {
		t.Errorf("Expected 'develop', received %v instead", renderer.branches)
	}
}
