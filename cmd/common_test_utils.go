package cmd

import "wt/internal/domain"

// Mock config store

type mockConfigStore struct{}

func (s mockConfigStore) Exists() bool { return true }
func (s mockConfigStore) GetManagedPaths() ([]string, error) {
	return []string{"/Users/antonio/Development/TestProject"}, nil
}
func (s mockConfigStore) AddManagedPath(path string) error {
	return nil
}

// Mock renderer

type mockRenderer struct {
	branches []string
}

func (r *mockRenderer) RenderWorktrees(worktrees []domain.Worktree) {
	for _, worktree := range worktrees {
		r.branches = append(r.branches, worktree.Branch)
	}
}
func (r mockRenderer) RenderManagedPaths(paths []string) {
	// No-op for this test
}

// Mock Git

type mockGit struct{}

func (g mockGit) GetWorktreesFromPath(path string) []domain.Worktree {
	return []domain.Worktree{
		{Path: "/Users/antonio/Development/Worktrees/TestList/develop",
			RepoPath: "/Users/antonio/Development/TestProject",
			Branch:   "develop",
			HeadSHA:  "4e04b2b0961c494fb643d91c8956813dbfcc799d",
		},
	}
}
