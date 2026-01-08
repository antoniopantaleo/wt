package cmd

import (
	"fmt"

	"wt/internal/app"

	"github.com/spf13/cobra"
)

func NewListCmd(deps app.Dependencies) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all current git worktrees",
		Long:  "List all current git worktrees in the repository.",
		Run:   runList,
	}

	return cmd
}

func runList(cmd *cobra.Command, args []string) {
	fmt.Println("Mock response")
}
