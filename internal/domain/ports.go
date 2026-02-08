package domain

type Renderer interface {
	RenderWorktrees(worktrees []Worktree)
	RenderManagedPaths(paths []string)
}

type ConfigStore interface {
	Exists() bool
	GetManagedPaths() ([]string, error)
	AddManagedPath(path string) error
}

type Git interface {
	GetWorktreesFromPath(path string) []Worktree
}
