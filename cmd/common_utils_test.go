package cmd

import "wt/internal/domain"

// Mock config store

type mockConfigStore struct {
	getManagedPaths   func() ([]string, error)
	exists            func() bool
	ensureExists      func() error
	addManagedPath    func(path string) error
	removeManagedPath func(path string) error
}

func (s mockConfigStore) Exists() bool {
	if s.exists == nil {
		return true
	}
	return s.exists()
}
func (s mockConfigStore) EnsureExists() error {
	if s.ensureExists == nil {
		return nil
	}
	return s.ensureExists()
}
func (s mockConfigStore) GetManagedPaths() ([]string, error) {
	if s.getManagedPaths == nil {
		return []string{}, nil
	}
	return s.getManagedPaths()
}
func (s mockConfigStore) AddManagedPath(path string) error {
	if s.addManagedPath == nil {
		return nil
	}
	return s.addManagedPath(path)
}
func (s mockConfigStore) RemoveManagedPath(path string) error {
	if s.removeManagedPath == nil {
		return nil
	}
	return s.removeManagedPath(path)
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

type mockGit struct {
	isGitRepo func(path string) bool
}

func (g mockGit) GetWorktreesFromPath(path string) []domain.Worktree {
	return []domain.Worktree{
		{Path: "/Users/antonio/Development/Worktrees/TestList/develop",
			RepoPath: "/Users/antonio/Development/TestProject",
			Branch:   "develop",
			HeadSHA:  "4e04b2b0961c494fb643d91c8956813dbfcc799d",
		},
	}
}

func (g mockGit) IsGitRepo(path string) bool {
	return g.isGitRepo(path)
}
