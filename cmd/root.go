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

	addCommand := func(
		root *cobra.Command, 
		command *cobra.Command, 
		groupID string,
	) *cobra.Command {
		root.AddCommand(command)
		command.GroupID = groupID
		return command
	}

	cmd.AddGroup(&cobra.Group{ID: "management", Title: "Management"})

	addCommand(cmd, NewListCmd(deps), "management")

	return cmd
}
