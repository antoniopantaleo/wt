// Package cmd implements the command line interface for the application.
package cmd

import (
	"wt/internal/domain"

	"github.com/spf13/cobra"
)

func NewRootCmd(deps domain.Dependencies) *cobra.Command {
	cmd := &cobra.Command{
		Version: "0.1.0",
		Use:     "wt",
		Short:   "A manager for git worktrees",
	}

	list := NewListCmd(deps)
	cmd.AddCommand(list)

	cmd.AddGroup(&cobra.Group{ID: "daily", Title: "Daily usage"})
	list.GroupID = "daily"

	return cmd
}
