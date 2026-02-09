package cmd

import (
	"sync"
	"wt/internal/domain"

	"github.com/spf13/cobra"
)

func NewListCmd(deps domain.Dependencies) *cobra.Command {
	cmd := &cobra.Command{
		Version: "0.1.0",
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List all currently managed git worktrees",
		RunE: func(cmd *cobra.Command, args []string) error {
			managedPaths, err := deps.ConfigStore.GetManagedPaths()
			if err != nil {
				return err
			}

			worktreesByPath := make([][]domain.Worktree, len(managedPaths))
			var wg sync.WaitGroup
			wg.Add(len(managedPaths))

			for i, path := range managedPaths {
				go func(index int, managedPath string) {
					defer wg.Done()
					worktreesByPath[index] = deps.Git.GetWorktreesFromPath(managedPath)
				}(i, path)
			}

			wg.Wait()

			var managedWorktrees []domain.Worktree
			for _, worktrees := range worktreesByPath {
				managedWorktrees = append(managedWorktrees, worktrees...)
			}

			deps.Renderer.RenderWorktrees(managedWorktrees)
			return nil
		},
	}

	return cmd
}
