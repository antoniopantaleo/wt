package cmd

import (
	"wt/internal/domain"

	"github.com/spf13/cobra"
)

func NewListCmd(deps domain.Dependencies) *cobra.Command {
	cmd := &cobra.Command{
		Version: "0.1.0",
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List all currently managed git worktrees",
		RunE:    func(cmd *cobra.Command, args []string) error {
			managedPaths, err := deps.ConfigStore.GetManagedPaths()
			if err != nil { return err }
			var managedWorktrees []domain.Worktree
			for _, path := range managedPaths {
				worktrees := deps.Git.GetWorktreesFromPath(path)
				managedWorktrees = append(managedWorktrees, worktrees...)
			}
			deps.Renderer.RenderWorktrees(managedWorktrees)
			return nil
		},
	}

	return cmd
}
