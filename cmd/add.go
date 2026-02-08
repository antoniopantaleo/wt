package cmd

import (
	"fmt"
	"wt/internal/debuglog"
	"wt/internal/domain"

	"github.com/spf13/cobra"
)

func NewAddCmd(deps domain.Dependencies) *cobra.Command {
	cmd := &cobra.Command{
		Version: "0.1.0",
		Use:     "add",
		Short:   "Add new repo to managed",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("please provide exactly one path to add")
			}
			path := args[0]
			debuglog.Printf("Adding path %v to managed paths", path)
			deps.Git.IsGitRepo(path)
			if !deps.Git.IsGitRepo(path) {
				return fmt.Errorf("provided path is not a git repository")
			}
			err := deps.ConfigStore.AddManagedPath(path)
			return err
		},
	}
	return cmd
}
