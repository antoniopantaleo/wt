package cmd

import (
	"wt/internal/domain"

	"github.com/spf13/cobra"
)

func NewAddCmd(deps domain.Dependencies) *cobra.Command {
	cmd := &cobra.Command{
		Version: "0.1.0",
		Use:     "add",
		Short:   "Add new repo to managed",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd
}
