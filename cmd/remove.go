package cmd

import (
	"fmt"
	"wt/internal/debuglog"
	"wt/internal/domain"

	"github.com/spf13/cobra"
)

func NewRemoveCmd(deps domain.Dependencies) *cobra.Command {
	cmd := &cobra.Command{
		Version: "0.1.0",
		Use:     "remove",
		Aliases: []string{"rm"},
		Short:   "Remove a repo from managed",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("please provide exactly one path to remove")
			}
			path := args[0]
			debuglog.Printf("Removing path %v from managed paths", path)
			return deps.ConfigStore.RemoveManagedPath(path)
		},
	}

	return cmd
}
