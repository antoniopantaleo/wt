package cmd

import (
	"fmt"
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
			err := deps.ConfigStore.AddManagedPath(path)
			return err
		},
	}
	return cmd
}
