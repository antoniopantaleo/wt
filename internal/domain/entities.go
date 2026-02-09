// Package domain serves as a dependency manager
package domain

type Dependencies struct {
	Git         Git
	Renderer    Renderer
	ConfigStore ConfigStore
}

type Worktree struct {
	RepoPath string
	Path     string
	HeadSHA  string
	// e.g. "main" or "(detached)"
	Branch string
}
