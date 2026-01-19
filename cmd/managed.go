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
			paths, error := deps.ConfigStore.GetManagedPaths()
			if error != nil { return error }
			deps.Renderer.RenderManagedPaths(paths)	
			return nil
		},
	}

	return cmd
}
