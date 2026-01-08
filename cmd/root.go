// Package cmd implements the command line interface for the application.
package cmd

import (
	"wt/internal/app"

	"github.com/spf13/cobra"
)

func NewRootCmd(deps app.Dependencies) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "wt",
		Short: "A manager for git worktrees",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	}

	cmd.AddCommand(NewListCmd(deps))

	return cmd
}
