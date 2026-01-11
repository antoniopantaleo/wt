package cmd

import (
	"wt/internal/domain"

	"github.com/spf13/cobra"
)

func NewManagedCmd(deps domain.Dependencies) *cobra.Command {
	cmd := &cobra.Command{
		Version: "0.1.0",
		Use:     "managed",
		Short:   "Show all managed repos",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd
}
