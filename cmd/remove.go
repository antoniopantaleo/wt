package cmd

import (
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
			return nil
		},
	}

	return cmd
}
