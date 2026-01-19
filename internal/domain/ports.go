package domain

type Renderer interface {
	RenderWorktrees(worktrees []Worktree)
}

type ConfigStore interface {
	Exists() bool
	GetManagedPaths() ([]string, error)
}

type Git interface {
	GetWorktreesFromPath(path string) []Worktree
}
