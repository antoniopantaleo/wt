package cmd

import (
	"fmt"
	"path/filepath"
	"wt/internal/debuglog"
	"wt/internal/domain"

	"github.com/spf13/cobra"
)

func NewAddCmd(deps domain.Dependencies) *cobra.Command {
	cmd := &cobra.Command{
		Version: "0.1.0",
		Use:     "add",
		Short:   "Add new repo to managed",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			path, err := filepath.Abs(args[0])
			if err != nil {
				return fmt.Errorf("resolving path: %w", err)
			}
			debuglog.Printf("Adding path %v to managed paths", path)
			deps.Git.IsGitRepo(path)
			if !deps.Git.IsGitRepo(path) {
				return fmt.Errorf("provided path is not a git repository")
			}
			return deps.ConfigStore.AddManagedPath(path)
		},
	}
	return cmd
}
